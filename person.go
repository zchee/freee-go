// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freeepm

import (
	"context"
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

// PersonService contains methods and other services that help with interacting
// with the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	Options []option.RequestOption
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.Options = opts
	return
}

// このリクエストで指定した ID の事業所の従業員一覧を返します。 権限・ステータス・
// 従業員 ID で取得する情報を絞り込むことができます。
func (r *PersonService) List(ctx context.Context, query PersonListParams, opts ...option.RequestOption) (res *pagination.PeopleOffset[PersonListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "people"
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

// このリクエストで指定した ID の事業所の従業員一覧を返します。 権限・ステータス・
// 従業員 ID で取得する情報を絞り込むことができます。
func (r *PersonService) ListAutoPaging(ctx context.Context, query PersonListParams, opts ...option.RequestOption) *pagination.PeopleOffsetAutoPager[PersonListResponse] {
	return pagination.NewPeopleOffsetAutoPager(r.List(ctx, query, opts...))
}

// 従業員情報
type PersonListResponse struct {
	// 従業員 ID
	ID int64 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 氏名
	Name string `json:"name"`
	// 人事労務側従業員 ID
	PayrollEmployeeID int64 `json:"payroll_employee_id"`
	// 事業所におけるロール
	Role string `json:"role"`
	// 事業所におけるロールの表示名
	RoleDisplayName string `json:"role_display_name"`
	// ステータス
	Status string `json:"status"`
	// 標準単価
	UnitCost PersonListResponseUnitCost `json:"unit_cost"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Email             respjson.Field
		Name              respjson.Field
		PayrollEmployeeID respjson.Field
		Role              respjson.Field
		RoleDisplayName   respjson.Field
		Status            respjson.Field
		UnitCost          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 標準単価
type PersonListResponseUnitCost struct {
	// 標準単価 ID
	ID int64 `json:"id"`
	// 名前
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
func (r PersonListResponseUnitCost) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseUnitCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 役割
	Role param.Opt[string] `query:"role,omitzero" json:"-"`
	// 従業員 ID
	PersonIDs []int64 `query:"person_ids,omitzero" json:"-"`
	// ステータス（招待中・利用中・無効）
	//
	// Any of "sent", "accepted", "inactive".
	Status PersonListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ステータス（招待中・利用中・無効）
type PersonListParamsStatus string

const (
	PersonListParamsStatusSent     PersonListParamsStatus = "sent"
	PersonListParamsStatusAccepted PersonListParamsStatus = "accepted"
	PersonListParamsStatusInactive PersonListParamsStatus = "inactive"
)
