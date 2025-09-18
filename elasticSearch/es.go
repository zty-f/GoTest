package es

import (
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

type ESSearchAgg struct {
	Name string
	Agg  elastic.Aggregation
}

type EsSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	Aggregations []ESSearchAgg
	Size         *int
	// 不使用这里的分页机制，rpc.SearchByDSL 接口支持了 limit/offset 参数
	// From         int //分页
	// Size         int
}

type EsBucket struct {
	Key      int64 `json:"key"`
	DocCount int64 `json:"doc_count"`
}

func (e *EsSearch) BoolQueryDSL() ([]byte, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(e.MustQuery...)
	boolQuery.MustNot(e.MustNotQuery...)
	boolQuery.Should(e.ShouldQuery...)
	boolQuery.Filter(e.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(e.MustQuery) == 0 && len(e.MustNotQuery) == 0 && len(e.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}

	source := elastic.NewSearchSource().Query(boolQuery)
	if len(e.Sorters) > 0 {
		source = source.SortBy(e.Sorters...)
	}
	if e.Size != nil {
		source.Size(*e.Size)
	}

	for aggName := range e.Aggregations {
		agg := e.Aggregations[aggName]
		source = source.Aggregation(agg.Name, agg.Agg)
	}

	return json.Marshal(source)
}

func (e *EsSearch) AddMustQuery(query elastic.Query) {
	e.MustQuery = append(e.MustQuery, query)
}

func (e *EsSearch) AddMustNotQuery(query elastic.Query) {
	e.MustNotQuery = append(e.MustNotQuery, query)
}

func (e *EsSearch) AddFilter(query elastic.Query) {
	e.Filters = append(e.Filters, query)
}

func (e *EsSearch) AddShouldQuery(query elastic.Query) {
	e.ShouldQuery = append(e.ShouldQuery, query)
}
