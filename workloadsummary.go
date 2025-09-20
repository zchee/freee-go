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

// WorkloadSummaryService contains methods and other services that help with
// interacting with the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkloadSummaryService] method instead.
type WorkloadSummaryService struct {
	Options []option.RequestOption
}

// NewWorkloadSummaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkloadSummaryService(opts ...option.RequestOption) (r WorkloadSummaryService) {
	r = WorkloadSummaryService{}
	r.Options = opts
	return
}

// 取得対象の従業員の工数実績のサマリを返します。 取得対象従業員と年月の取得範囲で
// 絞り込みできます。
func (r *WorkloadSummaryService) List(ctx context.Context, query WorkloadSummaryListParams, opts ...option.RequestOption) (res *pagination.WorkloadSummariesOffset[WorkloadSummaryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "workload_summaries"
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

// 取得対象の従業員の工数実績のサマリを返します。 取得対象従業員と年月の取得範囲で
// 絞り込みできます。
func (r *WorkloadSummaryService) ListAutoPaging(ctx context.Context, query WorkloadSummaryListParams, opts ...option.RequestOption) *pagination.WorkloadSummariesOffsetAutoPager[WorkloadSummaryListResponse] {
	return pagination.NewWorkloadSummariesOffsetAutoPager(r.List(ctx, query, opts...))
}

// 工数実績サマリ
type WorkloadSummaryListResponse struct {
	// 工数登録期間 from
	FromDate string `json:"from_date"`
	// 工数実績（分）
	Minutes int64 `json:"minutes"`
	// 対象従業員のユーザー id
	PersonID int64 `json:"person_id"`
	// 対象従業員氏名
	PersonName string `json:"person_name"`
	// 生産時間実績（分）
	ProductiveMinutes int64 `json:"productive_minutes"`
	// 工数登録期間 to
	ThruDate string `json:"thru_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromDate          respjson.Field
		Minutes           respjson.Field
		PersonID          respjson.Field
		PersonName        respjson.Field
		ProductiveMinutes respjson.Field
		ThruDate          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkloadSummaryListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkloadSummaryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadSummaryListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得対象範囲（YYYY-MM）
	YearMonth string `query:"year_month,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 取得対象従業員の検索スコープです。 <ul> <li> all を指定した場合は全従業員が対象
	// です。 絞り込みを行わない場合にご使用ください。 person_ids, team_ids での絞り込
	// みはできません。 </li> <li> team を指定した場合はチーム単位で絞り込みが可能です
	// 。 team_ids での絞り込みができます。 person_ids での絞り込みはできません。 </li>
	// <li> employee を指定した場合は person_ids による絞り込みができます。 team_ids で
	// の絞り込みは行なえません。 </li> <li> scope を指定しない場合はログインユーザの情
	// 報のみの取得です。 </li> </ul>
	//
	// Any of "all", "team", "employee".
	EmployeesScope WorkloadSummaryListParamsEmployeesScope `query:"employees_scope,omitzero" json:"-"`
	// 取得対象従業員のユーザ ID
	PersonIDs []int64 `query:"person_ids,omitzero" json:"-"`
	// 取得対象のチーム ID
	TeamIDs []int64 `query:"team_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkloadSummaryListParams]'s query parameters as
// `url.Values`.
func (r WorkloadSummaryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// 取得対象従業員の検索スコープです。 <ul> <li> all を指定した場合は全従業員が対象
// です。 絞り込みを行わない場合にご使用ください。 person_ids, team_ids での絞り込
// みはできません。 </li> <li> team を指定した場合はチーム単位で絞り込みが可能です
// 。 team_ids での絞り込みができます。 person_ids での絞り込みはできません。 </li>
// <li> employee を指定した場合は person_ids による絞り込みができます。 team_ids で
// の絞り込みは行なえません。 </li> <li> scope を指定しない場合はログインユーザの情
// 報のみの取得です。 </li> </ul>
type WorkloadSummaryListParamsEmployeesScope string

const (
	WorkloadSummaryListParamsEmployeesScopeAll      WorkloadSummaryListParamsEmployeesScope = "all"
	WorkloadSummaryListParamsEmployeesScopeTeam     WorkloadSummaryListParamsEmployeesScope = "team"
	WorkloadSummaryListParamsEmployeesScopeEmployee WorkloadSummaryListParamsEmployeesScope = "employee"
)
