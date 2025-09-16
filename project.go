// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/zchee/freee-go/internal/apijson"
	"github.com/zchee/freee-go/internal/apiquery"
	"github.com/zchee/freee-go/internal/requestconfig"
	"github.com/zchee/freee-go/option"
	"github.com/zchee/freee-go/packages/pagination"
	"github.com/zchee/freee-go/packages/param"
	"github.com/zchee/freee-go/packages/respjson"
)

// ProjectService contains methods and other services that help with interacting
// with the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// プロジェクトを登録することができます。
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ID に該当するプロジェクトの詳細情報を返します。
func (r *ProjectService) Get(ctx context.Context, id int64, query ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// この事業所のプロジェクトの一覧情報を返します。 運用ステータス、マネージャー、発
// 注先、発注元で絞り込みできます。
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.ProjectsOffset[ProjectListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "projects"
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

// この事業所のプロジェクトの一覧情報を返します。 運用ステータス、マネージャー、発
// 注先、発注元で絞り込みできます。
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.ProjectsOffsetAutoPager[ProjectListResponse] {
	return pagination.NewProjectsOffsetAutoPager(r.List(ctx, query, opts...))
}

// ページネーションのメタ情報
type Meta struct {
	// リクエストのオフセット件数
	CurrentOffset int64 `json:"current_offset"`
	// 次ページのオフセット件数
	NextOffset int64 `json:"next_offset,nullable"`
	// 前ページのオフセット件数
	PrevOffset int64 `json:"prev_offset,nullable"`
	// 全レコード件数
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentOffset respjson.Field
		NextOffset    respjson.Field
		PrevOffset    respjson.Field
		TotalCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meta) RawJSON() string { return r.JSON.raw }
func (r *Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// プロジェクト
type ProjectDetail struct {
	// プロジェクト ID
	ID int64 `json:"id"`
	// 招待リンク
	AssignmentURLEnabled bool `json:"assignment_url_enabled"`
	// 収支管理詳細
	Balance ProjectDetailBalance `json:"balance"`
	// プロジェクトコード
	Code string `json:"code"`
	// カラー
	Color string `json:"color"`
	// 発注先
	Contractors []ProjectDetailContractor `json:"contractors"`
	// プロジェクト概要
	Description string `json:"description,nullable"`
	// 期間 from
	FromDate string `json:"from_date"`
	// プロジェクトマネージャー
	Manager ProjectDetailManager `json:"manager"`
	// プロジェクトメンバー
	Members []ProjectDetailMember `json:"members"`
	// プロジェクト名
	Name string `json:"name"`
	// 運用ステータス
	OperationalStatus string `json:"operational_status"`
	// 発注元
	Orderers []ProjectDetailOrderer `json:"orderers"`
	// 従業員への公開設定
	PublishToEmployee bool `json:"publish_to_employee"`
	// 受注ステータス
	SalesOrderStatus string `json:"sales_order_status"`
	// 期間 to
	ThruDate string `json:"thru_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AssignmentURLEnabled respjson.Field
		Balance              respjson.Field
		Code                 respjson.Field
		Color                respjson.Field
		Contractors          respjson.Field
		Description          respjson.Field
		FromDate             respjson.Field
		Manager              respjson.Field
		Members              respjson.Field
		Name                 respjson.Field
		OperationalStatus    respjson.Field
		Orderers             respjson.Field
		PublishToEmployee    respjson.Field
		SalesOrderStatus     respjson.Field
		ThruDate             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetail) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 収支管理詳細
type ProjectDetailBalance struct {
	// 人件費
	LaborCosts ProjectDetailBalanceLaborCosts `json:"labor_costs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LaborCosts  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalance) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 人件費
type ProjectDetailBalanceLaborCosts struct {
	// 各メンバーの内訳
	Details []ProjectDetailBalanceLaborCostsDetail `json:"details"`
	// 各月
	Monthly []ProjectDetailBalanceLaborCostsMonthly `json:"monthly"`
	// 全体
	Total ProjectDetailBalanceLaborCostsTotal `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Details     respjson.Field
		Monthly     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCosts) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCosts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailBalanceLaborCostsDetail struct {
	// 各月
	Monthly []ProjectDetailBalanceLaborCostsDetailMonthly `json:"monthly"`
	// ユーザー ID
	PersonID int64 `json:"person_id"`
	// ユーザー名
	PersonName string                                    `json:"person_name"`
	Total      ProjectDetailBalanceLaborCostsDetailTotal `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Monthly     respjson.Field
		PersonID    respjson.Field
		PersonName  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetail) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailBalanceLaborCostsDetailMonthly struct {
	// 実績
	ActualResult ProjectDetailBalanceLaborCostsDetailMonthlyActualResult `json:"actual_result"`
	// 予算
	Budget ProjectDetailBalanceLaborCostsDetailMonthlyBudget `json:"budget"`
	// 年月（YYYY-MM）
	YearMonth string `json:"year_month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualResult respjson.Field
		Budget       respjson.Field
		YearMonth    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailMonthly) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailMonthly) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 実績
type ProjectDetailBalanceLaborCostsDetailMonthlyActualResult struct {
	// 金額
	Cost int64 `json:"cost"`
	// 時間
	Minutes int64 `json:"minutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Minutes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailMonthlyActualResult) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailMonthlyActualResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 予算
type ProjectDetailBalanceLaborCostsDetailMonthlyBudget struct {
	// 金額
	Cost int64 `json:"cost"`
	// 時間
	Minutes int64 `json:"minutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Minutes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailMonthlyBudget) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailMonthlyBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailBalanceLaborCostsDetailTotal struct {
	// 実績
	ActualResult ProjectDetailBalanceLaborCostsDetailTotalActualResult `json:"actual_result"`
	// 予算
	Budget ProjectDetailBalanceLaborCostsDetailTotalBudget `json:"budget"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualResult respjson.Field
		Budget       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailTotal) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 実績
type ProjectDetailBalanceLaborCostsDetailTotalActualResult struct {
	// 金額
	Cost int64 `json:"cost"`
	// 時間
	Minutes int64 `json:"minutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Minutes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailTotalActualResult) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailTotalActualResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 予算
type ProjectDetailBalanceLaborCostsDetailTotalBudget struct {
	// 金額
	Cost int64 `json:"cost"`
	// 時間
	Minutes int64 `json:"minutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Minutes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsDetailTotalBudget) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsDetailTotalBudget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailBalanceLaborCostsMonthly struct {
	// 実績
	ActualResult int64 `json:"actual_result"`
	// 予算
	Budget int64 `json:"budget"`
	// 年月（YYYY-MM）
	YearMonth string `json:"year_month"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualResult respjson.Field
		Budget       respjson.Field
		YearMonth    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsMonthly) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsMonthly) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 全体
type ProjectDetailBalanceLaborCostsTotal struct {
	// 実績
	ActualResult int64 `json:"actual_result"`
	// 予算
	Budget int64 `json:"budget"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualResult respjson.Field
		Budget       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailBalanceLaborCostsTotal) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailBalanceLaborCostsTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailContractor struct {
	// 取引先コード
	Code string `json:"code,nullable"`
	// 取引先 ID
	PartnerID int64 `json:"partner_id"`
	// 取引先名
	PartnerName string `json:"partner_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		PartnerID   respjson.Field
		PartnerName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailContractor) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailContractor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// プロジェクトマネージャー
type ProjectDetailManager struct {
	// ユーザ ID
	PersonID int64 `json:"person_id"`
	// 氏名
	PersonName string `json:"person_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PersonID    respjson.Field
		PersonName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailManager) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailMember struct {
	// ユーザ ID
	PersonID int64 `json:"person_id"`
	// 氏名
	PersonName string `json:"person_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PersonID    respjson.Field
		PersonName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailMember) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectDetailOrderer struct {
	// 取引先コード
	Code string `json:"code,nullable"`
	// 取引先 ID
	PartnerID int64 `json:"partner_id"`
	// 取引先名
	PartnerName string `json:"partner_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		PartnerID   respjson.Field
		PartnerName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectDetailOrderer) RawJSON() string { return r.JSON.raw }
func (r *ProjectDetailOrderer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewResponse struct {
	// プロジェクト
	Project ProjectDetail `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetResponse struct {
	// プロジェクト
	Project ProjectDetail `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// プロジェクト
type ProjectListResponse struct {
	// プロジェクト ID
	ID int64 `json:"id"`
	// 招待リンク
	AssignmentURLEnabled bool `json:"assignment_url_enabled"`
	// プロジェクトコード
	Code string `json:"code"`
	// カラー
	Color string `json:"color"`
	// 発注先
	Contractors []ProjectListResponseContractor `json:"contractors"`
	// プロジェクト概要
	Description string `json:"description,nullable"`
	// 期間 from
	FromDate string `json:"from_date"`
	// プロジェクトマネージャー
	Manager ProjectListResponseManager `json:"manager"`
	// プロジェクトメンバー
	Members []ProjectListResponseMember `json:"members"`
	// プロジェクト名
	Name string `json:"name"`
	// 運用ステータス
	OperationalStatus string `json:"operational_status"`
	// 発注元
	Orderers []ProjectListResponseOrderer `json:"orderers"`
	// プロジェクトタグ
	ProjectTags []ProjectListResponseProjectTag `json:"project_tags"`
	// 従業員への公開設定
	PublishToEmployee bool `json:"publish_to_employee"`
	// 受注ステータス
	SalesOrderStatus string `json:"sales_order_status"`
	// 期間 to
	ThruDate string `json:"thru_date"`
	// プロジェクトで使える工数タグのグループ
	WorkloadTagGroups []ProjectListResponseWorkloadTagGroup `json:"workload_tag_groups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AssignmentURLEnabled respjson.Field
		Code                 respjson.Field
		Color                respjson.Field
		Contractors          respjson.Field
		Description          respjson.Field
		FromDate             respjson.Field
		Manager              respjson.Field
		Members              respjson.Field
		Name                 respjson.Field
		OperationalStatus    respjson.Field
		Orderers             respjson.Field
		ProjectTags          respjson.Field
		PublishToEmployee    respjson.Field
		SalesOrderStatus     respjson.Field
		ThruDate             respjson.Field
		WorkloadTagGroups    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseContractor struct {
	// 取引先コード
	PartnerCode string `json:"partner_code,nullable"`
	// 取引先 ID
	PartnerID int64 `json:"partner_id"`
	// 取引先名
	PartnerName string `json:"partner_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PartnerCode respjson.Field
		PartnerID   respjson.Field
		PartnerName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseContractor) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseContractor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// プロジェクトマネージャー
type ProjectListResponseManager struct {
	// ユーザ ID
	PersonID int64 `json:"person_id"`
	// 氏名
	PersonName string `json:"person_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PersonID    respjson.Field
		PersonName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseManager) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseMember struct {
	// ユーザ ID
	PersonID int64 `json:"person_id"`
	// 氏名
	PersonName string `json:"person_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PersonID    respjson.Field
		PersonName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseMember) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseOrderer struct {
	// 取引先コード
	PartnerCode string `json:"partner_code,nullable"`
	// 取引先 ID
	PartnerID int64 `json:"partner_id"`
	// 取引先名
	PartnerName string `json:"partner_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PartnerCode respjson.Field
		PartnerID   respjson.Field
		PartnerName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseOrderer) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseOrderer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseProjectTag struct {
	// タググループ名
	TagGroupName string `json:"tag_group_name"`
	// タグ名
	TagName string `json:"tag_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TagGroupName respjson.Field
		TagName      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseProjectTag) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseProjectTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseWorkloadTagGroup struct {
	// 工数登録時に必須かどうかのフラグ
	Required bool `json:"required"`
	// タググループ ID
	TagGroupID int64 `json:"tag_group_id"`
	// タググループ名
	TagGroupName string `json:"tag_group_name"`
	// タグ
	Tags []ProjectListResponseWorkloadTagGroupTag `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Required     respjson.Field
		TagGroupID   respjson.Field
		TagGroupName respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseWorkloadTagGroup) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseWorkloadTagGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseWorkloadTagGroupTag struct {
	// タグ ID
	ID int64 `json:"id"`
	// タグ名
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseWorkloadTagGroupTag) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseWorkloadTagGroupTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// プロジェクトコード
	Code string `json:"code,required"`
	// 事業所 ID
	CompanyID int64 `json:"company_id,required"`
	// プロジェクト開始日
	FromDate string `json:"from_date,required"`
	// プロジェクト名
	Name string `json:"name,required"`
	// プロジェクトマネージャーのコスト(円)
	PmBudgetsCost int64 `json:"pm_budgets_cost,required"`
	// プロジェクト終了日
	ThruDate string `json:"thru_date,required"`
	// プロジェクトマネージャーの従業員 ID このパラメータはシステム管理者かプロジェクト
	// マネージャーでログインしているときのみ指定可能。（デフォルト：指定しない場合はロ
	// グインユーザ）
	ManagerPersonID param.Opt[int64] `json:"manager_person_id,omitzero"`
	// プロジェクトの招待リンク機能設定プロジェクトの招待リンクを発行できるようにするか
	// どうかを設定します。
	AssignmentURLEnabled param.Opt[bool] `json:"assignment_url_enabled,omitzero"`
	// プロジェクトの色を指定可能（デフォルト：orange） { orange: 1, blue_green: 2,
	// green: 3, blue: 4, purple: 5, red: 6, yellow: 7 }
	ColorID param.Opt[int64] `json:"color_id,omitzero"`
	// プロジェクト概要
	Description param.Opt[string] `json:"description,omitzero"`
	// 従業員への公開設定公開するとプロジェクト一覧に表示され、従業員がアサインリクエス
	// トを送れるようになります。（詳細画面は閲覧不可）
	PublishToEmployee param.Opt[bool] `json:"publish_to_employee,omitzero"`
	// 受注ステータス ID
	SalesOrderStatusID param.Opt[int64] `json:"sales_order_status_id,omitzero"`
	// アサインするユーザの配列
	Members []ProjectNewParamsMember `json:"members,omitzero"`
	// 発注先として指定する取引先 ID の配列
	ContractorIDs []int64 `json:"contractor_ids,omitzero"`
	// 発注元として指定する取引先 ID の配列
	OrdererIDs []int64 `json:"orderer_ids,omitzero"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BudgetsCost, PersonID, UnitCostID are required.
type ProjectNewParamsMember struct {
	// 予算計算用の単価(円)
	BudgetsCost int64 `json:"budgets_cost,required"`
	// 従業員 ID
	PersonID int64 `json:"person_id,required"`
	// このプロジェクトでの実績単価 ID
	UnitCostID int64 `json:"unit_cost_id,required"`
	// 標準従業員単価の利用（デフォルト：false）
	UseStandardUnitCost param.Opt[bool] `json:"use_standard_unit_cost,omitzero"`
	paramObj
}

func (r ProjectNewParamsMember) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParamsMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParamsMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectGetParams]'s query parameters as `url.Values`.
func (r ProjectGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 発注先の取引先 ID
	ContractorIDs []int64 `query:"contractor_ids,omitzero" json:"-"`
	// マネージャのユーザ ID
	ManagerIDs []int64 `query:"manager_ids,omitzero" json:"-"`
	// 運用ステータス
	//
	// Any of "planning", "awaiting_approval", "in_progress", "rejected", "done".
	OperationalStatus ProjectListParamsOperationalStatus `query:"operational_status,omitzero" json:"-"`
	// 発注元の取引先 ID
	OrdererIDs []int64 `query:"orderer_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// 運用ステータス
type ProjectListParamsOperationalStatus string

const (
	ProjectListParamsOperationalStatusPlanning         ProjectListParamsOperationalStatus = "planning"
	ProjectListParamsOperationalStatusAwaitingApproval ProjectListParamsOperationalStatus = "awaiting_approval"
	ProjectListParamsOperationalStatusInProgress       ProjectListParamsOperationalStatus = "in_progress"
	ProjectListParamsOperationalStatusRejected         ProjectListParamsOperationalStatus = "rejected"
	ProjectListParamsOperationalStatusDone             ProjectListParamsOperationalStatus = "done"
)
