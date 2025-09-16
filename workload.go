// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/freee-go/internal/apijson"
	"github.com/stainless-sdks/freee-go/internal/apiquery"
	"github.com/stainless-sdks/freee-go/internal/requestconfig"
	"github.com/stainless-sdks/freee-go/option"
	"github.com/stainless-sdks/freee-go/packages/param"
	"github.com/stainless-sdks/freee-go/packages/respjson"
)

// WorkloadService contains methods and other services that help with interacting
// with the freee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkloadService] method instead.
type WorkloadService struct {
	Options []option.RequestOption
}

// NewWorkloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkloadService(opts ...option.RequestOption) (r WorkloadService) {
	r = WorkloadService{}
	r.Options = opts
	return
}

// 工数を登録することが出来ます。
func (r *WorkloadService) New(ctx context.Context, body WorkloadNewParams, opts ...option.RequestOption) (res *WorkloadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "workloads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// 取得対象の従業員の工数実績の詳細を返します。 取得対象従業員と年月の取得範囲で絞
// り込みできます。
func (r *WorkloadService) List(ctx context.Context, query WorkloadListParams, opts ...option.RequestOption) (res *WorkloadListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "workloads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// 工数実績詳細
type Workload struct {
	// 工数登録日
	Date time.Time `json:"date,required" format:"date"`
	// 工数実績（分）
	Minutes int64 `json:"minutes,required"`
	// 対象従業員のユーザー ID
	PersonID int64 `json:"person_id,required"`
	// 対象従業員氏名
	PersonName string `json:"person_name,required"`
	// プロジェクトコード
	ProjectCode string `json:"project_code,required"`
	// 対象プロジェクト ID
	ProjectID int64 `json:"project_id,required"`
	// プロジェクト名
	ProjectName string `json:"project_name,required"`
	// 工数実績 ID
	ID int64 `json:"id"`
	// 業務内容
	Memo string `json:"memo"`
	// 工数タグ
	WorkloadTags []WorkloadWorkloadTag `json:"workload_tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date         respjson.Field
		Minutes      respjson.Field
		PersonID     respjson.Field
		PersonName   respjson.Field
		ProjectCode  respjson.Field
		ProjectID    respjson.Field
		ProjectName  respjson.Field
		ID           respjson.Field
		Memo         respjson.Field
		WorkloadTags respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workload) RawJSON() string { return r.JSON.raw }
func (r *Workload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadWorkloadTag struct {
	// タググループの ID
	TagGroupID int64 `json:"tag_group_id,required"`
	// タググループ名
	TagGroupName string `json:"tag_group_name,required"`
	// タグの ID
	TagID int64 `json:"tag_id,nullable"`
	// タグ名
	TagName string `json:"tag_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TagGroupID   respjson.Field
		TagGroupName respjson.Field
		TagID        respjson.Field
		TagName      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkloadWorkloadTag) RawJSON() string { return r.JSON.raw }
func (r *WorkloadWorkloadTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadNewResponse struct {
	// 工数実績詳細
	Workload Workload `json:"workload,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workload    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkloadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkloadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadListResponse struct {
	// ページネーションのメタ情報
	Meta      Meta       `json:"meta,required"`
	Workloads []Workload `json:"workloads,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		Workloads   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkloadListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadNewParams struct {
	// 事業所 ID
	CompanyID int64 `json:"company_id,required"`
	// 対象日
	Date time.Time `json:"date,required" format:"date"`
	// 記録時間（分）
	Minutes int64 `json:"minutes,required"`
	// 対象プロジェクト ID
	ProjectID int64 `json:"project_id,required"`
	// 業務内容
	Memo param.Opt[string] `json:"memo,omitzero"`
	// 対象従業員 ID このパラメータは管理者かチームリーダーでログインしているときのみ指
	// 定可能。指定しない場合はログインユーザに登録
	PersonID     param.Opt[int64]               `json:"person_id,omitzero"`
	WorkloadTags []WorkloadNewParamsWorkloadTag `json:"workload_tags,omitzero"`
	paramObj
}

func (r WorkloadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkloadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkloadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties TagGroupID, TagID are required.
type WorkloadNewParamsWorkloadTag struct {
	// 工数タググループ ID
	TagGroupID int64 `json:"tag_group_id,required"`
	// 工数タグ ID
	TagID int64 `json:"tag_id,required"`
	paramObj
}

func (r WorkloadNewParamsWorkloadTag) MarshalJSON() (data []byte, err error) {
	type shadow WorkloadNewParamsWorkloadTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkloadNewParamsWorkloadTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadListParams struct {
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
	EmployeesScope WorkloadListParamsEmployeesScope `query:"employees_scope,omitzero" json:"-"`
	// 取得対象従業員のユーザ ID
	PersonIDs []int64 `query:"person_ids,omitzero" json:"-"`
	// 取得対象のチーム ID
	TeamIDs []int64 `query:"team_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkloadListParams]'s query parameters as `url.Values`.
func (r WorkloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
type WorkloadListParamsEmployeesScope string

const (
	WorkloadListParamsEmployeesScopeAll      WorkloadListParamsEmployeesScope = "all"
	WorkloadListParamsEmployeesScopeTeam     WorkloadListParamsEmployeesScope = "team"
	WorkloadListParamsEmployeesScopeEmployee WorkloadListParamsEmployeesScope = "employee"
)
