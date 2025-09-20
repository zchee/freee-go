// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/zchee/freee-go/internal/apijson"
	"github.com/zchee/freee-go/internal/apiquery"
	"github.com/zchee/freee-go/internal/requestconfig"
	"github.com/zchee/freee-go/option"
	"github.com/zchee/freee-go/packages/pagination"
	"github.com/zchee/freee-go/packages/param"
	"github.com/zchee/freee-go/packages/respjson"
)

// PartnerService contains methods and other services that help with interacting
// with the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPartnerService] method instead.
type PartnerService struct {
	Options []option.RequestOption
}

// NewPartnerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPartnerService(opts ...option.RequestOption) (r PartnerService) {
	r = PartnerService{}
	r.Options = opts
	return
}

// 登録されている取引先の一覧を返します。
func (r *PartnerService) List(ctx context.Context, query PartnerListParams, opts ...option.RequestOption) (res *pagination.PartnersOffset[PartnerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "partners"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// 登録されている取引先の一覧を返します。
func (r *PartnerService) ListAutoPaging(ctx context.Context, query PartnerListParams, opts ...option.RequestOption) *pagination.PartnersOffsetAutoPager[PartnerListResponse] {
	return pagination.NewPartnersOffsetAutoPager(r.List(ctx, query, opts...))
}

// 取引先情報
type PartnerListResponse struct {
	// 取引先 ID
	ID int64 `json:"id"`
	// 取引先 code
	Code string `json:"code,nullable"`
	// 取引先名前
	Name string `json:"name"`
	// 発注先として登録されているプロジェクト一覧
	ProjectsAsContractor []PartnerListResponseProjectsAsContractor `json:"projects_as_contractor"`
	// 発注元として登録されているプロジェクト一覧
	ProjectsAsOrderer []PartnerListResponseProjectsAsOrderer `json:"projects_as_orderer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Code                 respjson.Field
		Name                 respjson.Field
		ProjectsAsContractor respjson.Field
		ProjectsAsOrderer    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartnerListResponse) RawJSON() string { return r.JSON.raw }
func (r *PartnerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartnerListResponseProjectsAsContractor struct {
	// プロジェクト ID
	ProjectID int64 `json:"project_id"`
	// プロジェクト名
	ProjectName string `json:"project_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProjectID   respjson.Field
		ProjectName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartnerListResponseProjectsAsContractor) RawJSON() string { return r.JSON.raw }
func (r *PartnerListResponseProjectsAsContractor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartnerListResponseProjectsAsOrderer struct {
	// プロジェクト ID
	ProjectID int64 `json:"project_id"`
	// プロジェクト名
	ProjectName string `json:"project_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProjectID   respjson.Field
		ProjectName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartnerListResponseProjectsAsOrderer) RawJSON() string { return r.JSON.raw }
func (r *PartnerListResponseProjectsAsOrderer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartnerListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PartnerListParams]'s query parameters as `url.Values`.
func (r PartnerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
