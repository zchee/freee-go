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

// TeamService contains methods and other services that help with interacting with
// the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r TeamService) {
	r = TeamService{}
	r.Options = opts
	return
}

// 登録されているチームの一覧を返します。
func (r *TeamService) List(ctx context.Context, query TeamListParams, opts ...option.RequestOption) (res *pagination.TeamsOffsetPage[TeamListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "teams"
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

// 登録されているチームの一覧を返します。
func (r *TeamService) ListAutoPaging(ctx context.Context, query TeamListParams, opts ...option.RequestOption) *pagination.TeamsOffsetPageAutoPager[TeamListResponse] {
	return pagination.NewTeamsOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// チーム情報
type TeamListResponse struct {
	// チーム ID
	ID int64 `json:"id,required"`
	// メンバー数
	MemberCount int64 `json:"member_count,required"`
	// チーム名
	Name string `json:"name,required"`
	// チームに登録されているメンバーの配列
	Members []TeamListResponseMember `json:"members"`
	// memo
	Memo string `json:"memo,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MemberCount respjson.Field
		Name        respjson.Field
		Members     respjson.Field
		Memo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamListResponse) RawJSON() string { return r.JSON.raw }
func (r *TeamListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamListResponseMember struct {
	// リーダフラグ、リーダならば true
	IsLeader bool `json:"is_leader,required"`
	// チームに所属している従業員 ID
	PersonID int64 `json:"person_id,required"`
	// チームに所属している従業員名
	PersonName string `json:"person_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsLeader    respjson.Field
		PersonID    respjson.Field
		PersonName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamListResponseMember) RawJSON() string { return r.JSON.raw }
func (r *TeamListResponseMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamListParams struct {
	// 事業所 ID
	CompanyID int64 `query:"company_id,required" json:"-"`
	// 取得レコードの件数（デフォルト：50, 最小：1, 最大 100）
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 取得レコードのオフセット（デフォルト：0）
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamListParams]'s query parameters as `url.Values`.
func (r TeamListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
