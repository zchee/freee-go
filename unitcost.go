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

// UnitCostService contains methods and other services that help with interacting
// with the freee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUnitCostService] method instead.
type UnitCostService struct {
	Options []option.RequestOption
}

// NewUnitCostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUnitCostService(opts ...option.RequestOption) (r UnitCostService) {
	r = UnitCostService{}
	r.Options = opts
	return
}

// 従業員の単価マスタを返します。
func (r *UnitCostService) List(ctx context.Context, query UnitCostListParams, opts ...option.RequestOption) (res *UnitCostListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "unit_costs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UnitCostListResponse struct {
	// ページネーションのメタ情報
	Meta      Meta                           `json:"meta,required"`
	UnitCosts []UnitCostListResponseUnitCost `json:"unit_costs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		UnitCosts   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitCostListResponse) RawJSON() string { return r.JSON.raw }
func (r *UnitCostListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 従業員単価マスタ
type UnitCostListResponseUnitCost struct {
	// 単価マスタ ID
	ID int64 `json:"id"`
	// 取得時点での適用金額
	CurrentCost int64 `json:"current_cost"`
	// 単価マスタ名
	Name string `json:"name"`
	// 期間ごとの適用金額の配列
	Rules []UnitCostListResponseUnitCostRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CurrentCost respjson.Field
		Name        respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitCostListResponseUnitCost) RawJSON() string { return r.JSON.raw }
func (r *UnitCostListResponseUnitCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitCostListResponseUnitCostRule struct {
	// 適用期間の金額
	Cost int64 `json:"cost"`
	// 適用開始日
	FromDate time.Time `json:"from_date" format:"date"`
	// 適用日終了日
	ThruDate time.Time `json:"thru_date" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		FromDate    respjson.Field
		ThruDate    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitCostListResponseUnitCostRule) RawJSON() string { return r.JSON.raw }
func (r *UnitCostListResponseUnitCostRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitCostListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UnitCostListParams]'s query parameters as `url.Values`.
func (r UnitCostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
