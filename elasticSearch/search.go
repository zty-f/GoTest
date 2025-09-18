package es

import (
	"context"
)

type ESResult struct {
	Hits *ESResultHits `json:"hits"`
}

type ESResultHits struct {
	Total int64               `json:"total"`
	Hits  []*ESResultHitsHits `json:"hits"`
}

type ESResultHitsHits struct {
	Id string `json:"_id"`
}

func (e *ESResult) Ids() []string {
	if e == nil || e.Hits == nil || e.Hits.Hits == nil {
		return []string{}
	}
	var res []string
	for _, hit := range e.Hits.Hits {
		res = append(res, hit.Id)
	}
	return res
}

func (e *ESResult) Total() int64 {
	if e == nil || e.Hits == nil {
		return 0
	}
	return e.Hits.Total
}

// 如果走数据库平台申请 ES，busstype 与索引是等价的关系。
func BussQueryByGrpc(ctx context.Context, busstype string, query string) (string, error) {
	// fun := "BussQueryByGrpc"
	// resp := SearchAdapter.BussQueryByGrpc(ctx, &searchservice.BussQueryReq{
	// 	Busstype: busstype,
	// 	Body:     query,
	// })
	// if resp.Errinfo != nil {
	// 	xlog.Warnf(ctx, "%s SearchForLogAdapter.BussQueryByGrpc err: %v, busstype: %s, query: %s", fun, resp.Errinfo.Msg, busstype, query)
	// 	return "", NewGrpcErrInfo(resp.Errinfo)
	// }
	return "", nil
}

// http://confluence.pri.ibanyu.com/pages/viewpage.action?pageId=35162469
// upType 10：仅更新，若ID不存在则不执行 20：若ID存在则更新相关字段，不存在则插入 不传默认覆盖更新

type SearchUpType int32

const (
	SEARCH_UP_TYPE_OVERLAY_UPDATE_FILED = 0  // 默认覆盖更新
	SEARCH_UP_TYPE_UPDATE_FILED         = 10 // 仅更新，若ID不存在则不执行
	SEARCH_UP_TYPE_UPSERT_FILED         = 20 // 若ID存在则更新相关字段，不存在则插入
)

func UpsertDoc(ctx context.Context, bussType string, id string, doc string, upType SearchUpType) error {
	// fun := "UpsertDoc -->"
	var err error
	// res := SearchAdapter.UpsertDoc(ctx, &UpsertDocReq{
	// 	Busstype: bussType,
	// 	Id:       id,
	// 	Doc:      doc,
	// 	Uptype:   int32(upType),
	// })
	// if res.Errinfo != nil {
	// 	xlog.Warnf(ctx, "%s UpsertDoc busstype:%s id:%s uptype:%d doc:%s err:%s", fun, bussType, id, upType, doc, res.Errinfo)
	// 	return NewThriftRpcErrInfo(res.Errinfo)
	// }
	// xlog.Infof(ctx, "%s UpsertDoc busstype:%s id:%s uptype:%d doc:%s success", fun, bussType, id, upType, doc)
	return err
}

func SearchQueryDslBasic(ctx context.Context, bussType string, dsl string) (string, error) {
	// fun := "SearchQueryDslBasic -->"
	// xlog.Infof(ctx, "%s QueryDSL busstype:%s dsl:%s", fun, bussType, dsl)
	// res := SearchAdapter.BussQueryByGrpc(ctx, &searchservice.BussQueryReq{
	// 	Busstype: bussType,
	// 	Body:     dsl,
	// })
	// if res.Errinfo != nil {
	// 	xlog.Warnf(ctx, "%s BussQueryByGrpc busstype:%s dsl:%s, err:%s", fun, bussType, dsl, res.Errinfo)
	// 	return "", NewGrpcErrInfo(res.Errinfo)
	// }

	return "", nil
}

type CompoundQueryData struct {
	Docs   string `thrift:"docs,1" json:"docs"`
	Offset int64  `thrift:"offset,2" json:"offset"`
	More   bool   `thrift:"more,3" json:"more"`
	Total  int64  `thrift:"total,4" json:"total"`
	Ids    string `thrift:"ids,5" json:"ids"`
}

func SearchByDSL(ctx context.Context, bussType string, dsl string, offset, limit int64) (*CompoundQueryData, error) {
	// fun := "SearchByDSL -->"
	// xlog.Infof(ctx, "%s CompoundQuery busstype:%s dsl:%s limit:%d offset:%d", fun, bussType, dsl, limit, offset)
	// var err error
	// data := &CompoundQueryData{}
	// res := SearchAdapter.CompoundQuery(ctx, &CompoundQueryReq{
	// 	Busstype: bussType,
	// 	Dsl:      dsl,
	// 	Offset:   offset,
	// 	Limit:    int32(limit),
	// })
	// if res.Errinfo != nil {
	// 	xlog.Warnf(ctx, "%s CompoundQuery busstype:%s dsl:%s limit:%d offset:%d err:%s", fun, bussType, dsl, limit, offset, res.Errinfo)
	// 	return data, NewThriftRpcErrInfo(res.Errinfo)
	// }
	// if res.Data == nil {
	// 	xlog.Warnf(ctx, "%s CompoundQuery busstype:%s dsl:%s limit:%d offset:%d data is nil", fun, bussType, dsl, limit, offset)
	// 	return data, NewThriftRpcErrInfo(res.Errinfo)
	// }
	// data = res.Data
	// xlog.Infof(ctx, "%s CompoundQuery busstype:%s dsl:%s limit:%d offset:%d success", fun, bussType, dsl, limit, offset)
	return nil, nil
}

func DeleteDoc(ctx context.Context, bussType string, id string) error {
	// fun := "DeleteDoc -->"
	var err error
	// res := SearchAdapter.DeleteDoc(ctx, &DeleteDocReq{
	// 	Busstype: bussType,
	// 	Id:       id,
	// })
	// if res.Errinfo != nil {
	// 	xlog.Warnf(ctx, "%s DeleteDoc busstype:%s id:%s err:%s", fun, bussType, id, res.Errinfo)
	// 	return NewThriftRpcErrInfo(res.Errinfo)
	// }
	// xlog.Infof(ctx, "%s DeleteDoc busstype:%s id:%s success", fun, bussType, id)
	return err
}

// UpdateByQueryResult ES update_by_query API的响应结构
type UpdateByQueryResult struct {
	Took             int64 `json:"took"`
	TimedOut         bool  `json:"timed_out"`
	Total            int64 `json:"total"`
	Updated          int64 `json:"updated"`
	Deleted          int64 `json:"deleted"`
	Batches          int64 `json:"batches"`
	VersionConflicts int64 `json:"version_conflicts"`
	Noops            int64 `json:"noops"`
	Retries          struct {
		Bulk   int64 `json:"bulk"`
		Search int64 `json:"search"`
	} `json:"retries"`
	ThrottledMillis      int64         `json:"throttled_millis"`
	RequestsPerSecond    float64       `json:"requests_per_second"`
	ThrottledUntilMillis int64         `json:"throttled_until_millis"`
	Failures             []interface{} `json:"failures"`
}

// UpdateByQuery 执行ES的update_by_query操作，通过查询条件批量更新文档
func UpdateByQuery(ctx context.Context, bussType string, query string) (*UpdateByQueryResult, error) {
	// fun := "UpdateByQuery -->"
	// apiPath := "/" + bussType + "/_update_by_query?conflicts=proceed"
	// // 调用ES API
	// resp := SearchAdapter.EsDirectByGrpc(ctx, &searchservice.EsDirectReq{
	// 	Busstype: bussType,
	// 	Uri:      apiPath,
	// 	Method:   http.MethodPost,
	// 	Body:     query,
	// })
	// if resp.Errinfo != nil {
	// 	xlog.Errorf(ctx, "%s EsDirectByGrpc failed, bussType: %s, query: %s, err: %v",
	// 		fun, bussType, query, resp.Errinfo)
	// 	return nil, errors.New(resp.Errinfo.Msg)
	// }
	//
	// if resp.Data == nil || resp.Data.Body == "" {
	// 	xlog.Errorf(ctx, "%s EsDirectByGrpc failed, bussType: %s, query: %s, data is nil",
	// 		fun, bussType, query)
	// 	return nil, errors.New("data is nil")
	// }
	//
	// // 解析响应
	// var result UpdateByQueryResult
	// if err := json.Unmarshal([]byte(resp.Data.Body), &result); err != nil {
	// 	xlog.Errorf(ctx, "%s unmarshal response failed, resp: %s, err: %v", fun, resp.Data.Body, err)
	// 	return nil, err
	// }
	//
	// xlog.Infof(ctx, "%s success, bussType: %s, updated: %d, total: %d",
	// 	fun, bussType, result.Updated, result.Total)

	return nil, nil
}
