# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UserGetMeResponse">UserGetMeResponse</a>

Methods:

- <code title="get /users/me">client.Users.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UserService.GetMe">GetMe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UserGetMeResponse">UserGetMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#Meta">Meta</a>
- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectDetail">ProjectDetail</a>
- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{id}">client.Projects.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectGetParams">ProjectGetParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#ProjectsOffset">ProjectsOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#ProjectListResponse">ProjectListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# UnitCosts

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UnitCostListResponse">UnitCostListResponse</a>

Methods:

- <code title="get /unit_costs">client.UnitCosts.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UnitCostService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UnitCostListParams">UnitCostListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#UnitCostsOffset">UnitCostsOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#UnitCostListResponse">UnitCostListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# People

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PersonListResponse">PersonListResponse</a>

Methods:

- <code title="get /people">client.People.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PersonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PersonListParams">PersonListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#PeopleOffset">PeopleOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PersonListResponse">PersonListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Partners

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PartnerListResponse">PartnerListResponse</a>

Methods:

- <code title="get /partners">client.Partners.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PartnerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PartnerListParams">PartnerListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#PartnersOffset">PartnersOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#PartnerListResponse">PartnerListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workloads

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#Workload">Workload</a>
- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadNewResponse">WorkloadNewResponse</a>

Methods:

- <code title="post /workloads">client.Workloads.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadNewParams">WorkloadNewParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadNewResponse">WorkloadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workloads">client.Workloads.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadListParams">WorkloadListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#WorkloadsOffset">WorkloadsOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#Workload">Workload</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WorkloadSummaries

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadSummaryListResponse">WorkloadSummaryListResponse</a>

Methods:

- <code title="get /workload_summaries">client.WorkloadSummaries.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadSummaryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadSummaryListParams">WorkloadSummaryListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#WorkloadSummariesOffset">WorkloadSummariesOffset</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#WorkloadSummaryListResponse">WorkloadSummaryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#TeamListResponse">TeamListResponse</a>

Methods:

- <code title="get /teams">client.Teams.<a href="https://pkg.go.dev/github.com/zchee/freee-go#TeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#TeamListParams">TeamListParams</a>) (<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go/packages/pagination#TeamsOffsetPage">TeamsOffsetPage</a>[<a href="https://pkg.go.dev/github.com/zchee/freee-go">freee</a>.<a href="https://pkg.go.dev/github.com/zchee/freee-go#TeamListResponse">TeamListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
