// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesDomainResultDataSourceEnvelope struct {
	Result PagesDomainDataSourceModel `json:"result,computed"`
}

type PagesDomainResultListDataSourceEnvelope struct {
	Result *[]*PagesDomainDataSourceModel `json:"result,computed"`
}

type PagesDomainDataSourceModel struct {
	AccountID   types.String                         `tfsdk:"account_id" path:"account_id"`
	DomainName  types.String                         `tfsdk:"domain_name" path:"domain_name"`
	ProjectName types.String                         `tfsdk:"project_name" path:"project_name"`
	Filter      *PagesDomainFindOneByDataSourceModel `tfsdk:"filter"`
}

type PagesDomainFindOneByDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	ProjectName types.String `tfsdk:"project_name" path:"project_name"`
}
