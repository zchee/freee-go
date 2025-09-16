// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/zchee/freee-go"
	"github.com/zchee/freee-go/internal/apijson"
	"github.com/zchee/freee-go/internal/requestconfig"
	"github.com/zchee/freee-go/option"
	"github.com/zchee/freee-go/packages/param"
	"github.com/zchee/freee-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ProjectsOffsetMeta struct {
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
func (r ProjectsOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *ProjectsOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsOffset[T any] struct {
	Projects []T                `json:"projects"`
	Meta     ProjectsOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Projects    respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ProjectsOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *ProjectsOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ProjectsOffset[T]) GetNextPage() (res *ProjectsOffset[T], err error) {
	if len(r.Projects) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ProjectsOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ProjectsOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ProjectsOffsetAutoPager[T any] struct {
	page *ProjectsOffset[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewProjectsOffsetAutoPager[T any](page *ProjectsOffset[T], err error) *ProjectsOffsetAutoPager[T] {
	return &ProjectsOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ProjectsOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Projects) == 0 {
		return false
	}
	if r.idx >= len(r.page.Projects) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Projects) == 0 {
			return false
		}
	}
	r.cur = r.page.Projects[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ProjectsOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *ProjectsOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *ProjectsOffsetAutoPager[T]) Index() int {
	return r.run
}

type UnitCostsOffsetMeta struct {
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
func (r UnitCostsOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *UnitCostsOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitCostsOffset[T any] struct {
	UnitCosts []T                 `json:"unit_costs"`
	Meta      UnitCostsOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UnitCosts   respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r UnitCostsOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *UnitCostsOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *UnitCostsOffset[T]) GetNextPage() (res *UnitCostsOffset[T], err error) {
	if len(r.UnitCosts) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *UnitCostsOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &UnitCostsOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type UnitCostsOffsetAutoPager[T any] struct {
	page *UnitCostsOffset[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewUnitCostsOffsetAutoPager[T any](page *UnitCostsOffset[T], err error) *UnitCostsOffsetAutoPager[T] {
	return &UnitCostsOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *UnitCostsOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.UnitCosts) == 0 {
		return false
	}
	if r.idx >= len(r.page.UnitCosts) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.UnitCosts) == 0 {
			return false
		}
	}
	r.cur = r.page.UnitCosts[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *UnitCostsOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *UnitCostsOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *UnitCostsOffsetAutoPager[T]) Index() int {
	return r.run
}

type PeopleOffsetMeta struct {
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
func (r PeopleOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *PeopleOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PeopleOffset[T any] struct {
	People []T              `json:"people"`
	Meta   PeopleOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		People      respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PeopleOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *PeopleOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PeopleOffset[T]) GetNextPage() (res *PeopleOffset[T], err error) {
	if len(r.People) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PeopleOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PeopleOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PeopleOffsetAutoPager[T any] struct {
	page *PeopleOffset[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPeopleOffsetAutoPager[T any](page *PeopleOffset[T], err error) *PeopleOffsetAutoPager[T] {
	return &PeopleOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PeopleOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.People) == 0 {
		return false
	}
	if r.idx >= len(r.page.People) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.People) == 0 {
			return false
		}
	}
	r.cur = r.page.People[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PeopleOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *PeopleOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *PeopleOffsetAutoPager[T]) Index() int {
	return r.run
}

type PartnersOffsetMeta struct {
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
func (r PartnersOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *PartnersOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartnersOffset[T any] struct {
	Partners []T                `json:"partners"`
	Meta     PartnersOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Partners    respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PartnersOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *PartnersOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PartnersOffset[T]) GetNextPage() (res *PartnersOffset[T], err error) {
	if len(r.Partners) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PartnersOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PartnersOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PartnersOffsetAutoPager[T any] struct {
	page *PartnersOffset[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPartnersOffsetAutoPager[T any](page *PartnersOffset[T], err error) *PartnersOffsetAutoPager[T] {
	return &PartnersOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PartnersOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Partners) == 0 {
		return false
	}
	if r.idx >= len(r.page.Partners) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Partners) == 0 {
			return false
		}
	}
	r.cur = r.page.Partners[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PartnersOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *PartnersOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *PartnersOffsetAutoPager[T]) Index() int {
	return r.run
}

type WorkloadsOffsetMeta struct {
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
func (r WorkloadsOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *WorkloadsOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadsOffset[T any] struct {
	Workloads []freeepm.Workload  `json:"workloads"`
	Meta      WorkloadsOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workloads   respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WorkloadsOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *WorkloadsOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WorkloadsOffset[T]) GetNextPage() (res *WorkloadsOffset[T], err error) {
	if len(r.Workloads) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WorkloadsOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WorkloadsOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WorkloadsOffsetAutoPager[T any] struct {
	page *WorkloadsOffset[T]
	cur  freeepm.Workload
	idx  int
	run  int
	err  error
	paramObj
}

func NewWorkloadsOffsetAutoPager[T any](page *WorkloadsOffset[T], err error) *WorkloadsOffsetAutoPager[T] {
	return &WorkloadsOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WorkloadsOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Workloads) == 0 {
		return false
	}
	if r.idx >= len(r.page.Workloads) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Workloads) == 0 {
			return false
		}
	}
	r.cur = r.page.Workloads[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WorkloadsOffsetAutoPager[T]) Current() freeepm.Workload {
	return r.cur
}

func (r *WorkloadsOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *WorkloadsOffsetAutoPager[T]) Index() int {
	return r.run
}

type WorkloadSummariesOffsetMeta struct {
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
func (r WorkloadSummariesOffsetMeta) RawJSON() string { return r.JSON.raw }
func (r *WorkloadSummariesOffsetMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkloadSummariesOffset[T any] struct {
	WorkloadSummaries []T                         `json:"workload_summaries"`
	Meta              WorkloadSummariesOffsetMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WorkloadSummaries respjson.Field
		Meta              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WorkloadSummariesOffset[T]) RawJSON() string { return r.JSON.raw }
func (r *WorkloadSummariesOffset[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WorkloadSummariesOffset[T]) GetNextPage() (res *WorkloadSummariesOffset[T], err error) {
	if len(r.WorkloadSummaries) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WorkloadSummariesOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WorkloadSummariesOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WorkloadSummariesOffsetAutoPager[T any] struct {
	page *WorkloadSummariesOffset[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWorkloadSummariesOffsetAutoPager[T any](page *WorkloadSummariesOffset[T], err error) *WorkloadSummariesOffsetAutoPager[T] {
	return &WorkloadSummariesOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WorkloadSummariesOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.WorkloadSummaries) == 0 {
		return false
	}
	if r.idx >= len(r.page.WorkloadSummaries) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.WorkloadSummaries) == 0 {
			return false
		}
	}
	r.cur = r.page.WorkloadSummaries[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WorkloadSummariesOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *WorkloadSummariesOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *WorkloadSummariesOffsetAutoPager[T]) Index() int {
	return r.run
}

type TeamsOffsetPageMeta struct {
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
func (r TeamsOffsetPageMeta) RawJSON() string { return r.JSON.raw }
func (r *TeamsOffsetPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamsOffsetPage[T any] struct {
	Teams []T                 `json:"teams"`
	Meta  TeamsOffsetPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Teams       respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TeamsOffsetPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TeamsOffsetPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TeamsOffsetPage[T]) GetNextPage() (res *TeamsOffsetPage[T], err error) {
	if len(r.Teams) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Meta.NextOffset

	if next < r.Meta.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TeamsOffsetPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TeamsOffsetPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TeamsOffsetPageAutoPager[T any] struct {
	page *TeamsOffsetPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTeamsOffsetPageAutoPager[T any](page *TeamsOffsetPage[T], err error) *TeamsOffsetPageAutoPager[T] {
	return &TeamsOffsetPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TeamsOffsetPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Teams) == 0 {
		return false
	}
	if r.idx >= len(r.page.Teams) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Teams) == 0 {
			return false
		}
	}
	r.cur = r.page.Teams[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TeamsOffsetPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TeamsOffsetPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TeamsOffsetPageAutoPager[T]) Index() int {
	return r.run
}
