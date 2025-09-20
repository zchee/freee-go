// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package freee

import (
	"context"
	"net/http"
	"slices"

	"github.com/zchee/freee-go/internal/apijson"
	"github.com/zchee/freee-go/internal/requestconfig"
	"github.com/zchee/freee-go/option"
	"github.com/zchee/freee-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the zchee API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// このリクエストの認可セッションにおけるログインユーザーの情報を返します。 freee
// 工数管理では一人のログインユーザーを複数の事業所に関連付けられるため、このユーザ
// ーと関連のあるすべての事業所の情報をリストで返します。 他の API のパラメータとし
// て company_id が求められる場合は、この API で取得した company_id を使用します。
func (r *UserService) GetMe(ctx context.Context, opts ...option.RequestOption) (res *UserGetMeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ログインユーザー情報
type UserGetMeResponse struct {
	// ユーザー ID
	ID int64 `json:"id"`
	// ユーザーが属する事業所の一覧
	Companies []UserGetMeResponseCompany `json:"companies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Companies   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetMeResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetMeResponseCompany struct {
	// 事業所 ID
	ID int64 `json:"id"`
	// 事業所に所属する従業員の表示名
	DisplayName string `json:"display_name"`
	// 事業所番号(半角英数字 10 桁)
	ExternalCid string `json:"external_cid"`
	// 事業所名
	Name string `json:"name"`
	// ログインユーザー情報
	PersonMe UserGetMeResponseCompanyPersonMe `json:"person_me"`
	// 事業所におけるロール
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		ExternalCid respjson.Field
		Name        respjson.Field
		PersonMe    respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetMeResponseCompany) RawJSON() string { return r.JSON.raw }
func (r *UserGetMeResponseCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ログインユーザー情報
type UserGetMeResponseCompanyPersonMe struct {
	// 従業員 ID
	ID int64 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 氏名
	Name string `json:"name"`
	// 事業所におけるロール
	Role string `json:"role"`
	// ステータス
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetMeResponseCompanyPersonMe) RawJSON() string { return r.JSON.raw }
func (r *UserGetMeResponseCompanyPersonMe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
