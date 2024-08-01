// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StaticRouteResultEnvelope struct {
	Result StaticRouteModel `json:"result,computed"`
}

type StaticRouteModel struct {
	AccountID     types.String               `tfsdk:"account_id" path:"account_id"`
	RouteID       types.String               `tfsdk:"route_id" path:"route_id"`
	Description   types.String               `tfsdk:"description" json:"description"`
	Nexthop       types.String               `tfsdk:"nexthop" json:"nexthop"`
	Prefix        types.String               `tfsdk:"prefix" json:"prefix"`
	Priority      types.Int64                `tfsdk:"priority" json:"priority"`
	Weight        types.Int64                `tfsdk:"weight" json:"weight"`
	Scope         *StaticRouteScopeModel     `tfsdk:"scope" json:"scope"`
	Deleted       types.Bool                 `tfsdk:"deleted" json:"deleted,computed"`
	Modified      types.Bool                 `tfsdk:"modified" json:"modified,computed"`
	Routes        *[]*StaticRouteRoutesModel `tfsdk:"routes" json:"routes,computed"`
	DeletedRoute  jsontypes.Normalized       `tfsdk:"deleted_route" json:"deleted_route,computed"`
	ModifiedRoute jsontypes.Normalized       `tfsdk:"modified_route" json:"modified_route,computed"`
	Route         jsontypes.Normalized       `tfsdk:"route" json:"route,computed"`
}

type StaticRouteScopeModel struct {
	ColoNames   *[]types.String `tfsdk:"colo_names" json:"colo_names"`
	ColoRegions *[]types.String `tfsdk:"colo_regions" json:"colo_regions"`
}

type StaticRouteRoutesModel struct {
	Nexthop     types.String                 `tfsdk:"nexthop" json:"nexthop"`
	Prefix      types.String                 `tfsdk:"prefix" json:"prefix"`
	Priority    types.Int64                  `tfsdk:"priority" json:"priority"`
	ID          types.String                 `tfsdk:"id" json:"id"`
	CreatedOn   timetypes.RFC3339            `tfsdk:"created_on" json:"created_on,computed"`
	Description types.String                 `tfsdk:"description" json:"description"`
	ModifiedOn  timetypes.RFC3339            `tfsdk:"modified_on" json:"modified_on,computed"`
	Scope       *StaticRouteRoutesScopeModel `tfsdk:"scope" json:"scope"`
	Weight      types.Int64                  `tfsdk:"weight" json:"weight"`
}

type StaticRouteRoutesScopeModel struct {
	ColoNames   *[]types.String `tfsdk:"colo_names" json:"colo_names"`
	ColoRegions *[]types.String `tfsdk:"colo_regions" json:"colo_regions"`
}
