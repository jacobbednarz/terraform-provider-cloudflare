// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type CloudConnectorRulesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*CloudConnectorRulesDataSource)(nil)

func NewCloudConnectorRulesDataSource() datasource.DataSource {
	return &CloudConnectorRulesDataSource{}
}

func (d *CloudConnectorRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_connector_rules"
}

func (d *CloudConnectorRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *CloudConnectorRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CloudConnectorRulesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := CloudConnectorRulesResultListDataSourceEnvelope{}
	page, err := d.client.CloudConnector.Rules.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes := []byte(page.JSON.RawJSON())
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
		return
	}

	if count := len(env.Result.Elements()); count != 1 {
		resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
		return
	}
	ts, diags := env.Result.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	data = &ts[0]

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
