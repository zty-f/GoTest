package es

import (
	"context"
	"encoding/json"
	"time"

	"github.com/olivere/elastic/v7"
)

const BussTypeSaleOrder = "readcamp_sale_order"

// SaleOrder 对应ES中的索引结构
type SaleOrder struct {
	ID           string     `json:"id"`
	UserID       int64      `json:"user_id"`
	PackageID    int64      `json:"package_id"`
	MsgStr       string     `json:"msg_str,omitempty"`
	UtmSource    string     `json:"utm_source,omitempty"`
	PayType      string     `json:"pay_type,omitempty"`
	Price        int64      `json:"price"`
	Source       int64      `json:"source"`
	OuterOrderID string     `json:"outer_order_id,omitempty"`
	OrderAssign  string     `json:"order_assign,omitempty"`
	LeadsAssign  string     `json:"leads_assign,omitempty"`
	CreateTime   *time.Time `json:"ct,omitempty"`
	UpdateTime   *time.Time `json:"ut,omitempty"`
}

// SearchSaleOrderRequest 查询请求参数
type SearchSaleOrderRequest struct {
	ID           string     `json:"id"`
	IDs          []string   `json:"ids"`
	UserID       int64      `json:"user_id"`
	UserIDs      []int64    `json:"user_ids"`
	PackageID    int64      `json:"package_id"`
	PackageIDs   []int64    `json:"package_ids"`
	UtmSources   []string   `json:"utm_sources"`
	PayTypes     []string   `json:"pay_types"`
	MinPrice     int64      `json:"min_price"`
	MaxPrice     int64      `json:"max_price"`
	Source       int64      `json:"source"`
	Sources      []int64    `json:"sources"`
	OuterOrderID string     `json:"outer_order_id"`
	OrderAssign  string     `json:"order_assign"`
	OrderAssigns []string   `json:"order_assigns"`
	LeadsAssign  string     `json:"leads_assign"`
	LeadsAssigns []string   `json:"leads_assigns"`
	MsgStr       string     `json:"msg_str"`
	BeginTime    *time.Time `json:"begin_time"`
	EndTime      *time.Time `json:"end_time"`
	Limit        int64      `json:"limit"`
	Offset       int64      `json:"offset"`
}

// SearchSaleOrderDSL 查询DSL结构
type SearchSaleOrderDSL struct {
	Query SaleOrderQuery  `json:"query"`
	Sort  []SaleOrderSort `json:"sort,omitempty"`
}

type SaleOrderSort struct {
	CreateTime *SaleOrderSortOrder `json:"ct,omitempty"`
	UpdateTime *SaleOrderSortOrder `json:"ut,omitempty"`
}

type SaleOrderSortOrder struct {
	Order string `json:"order,omitempty"`
}

type SaleOrderQuery struct {
	Bool SaleOrderBool `json:"bool"`
}

type SaleOrderBool struct {
	Must   []SaleOrderCondition `json:"must,omitempty"`
	Should []SaleOrderCondition `json:"should,omitempty"`
}

type SaleOrderCondition struct {
	Term        *SaleOrderTerm        `json:"term,omitempty"`
	Terms       *SaleOrderTerms       `json:"terms,omitempty"`
	Range       *SaleOrderRange       `json:"range,omitempty"`
	MatchPhrase *SaleOrderMatchPhrase `json:"match_phrase,omitempty"`
	Bool        *SaleOrderBool        `json:"bool,omitempty"`
}

type SaleOrderTerm struct {
	ID           *string `json:"id,omitempty"`
	UserID       *int64  `json:"user_id,omitempty"`
	PackageID    *int64  `json:"package_id,omitempty"`
	UtmSource    *string `json:"utm_source,omitempty"`
	PayType      *string `json:"pay_type,omitempty"`
	Source       *int64  `json:"source,omitempty"`
	OuterOrderID *string `json:"outer_order_id,omitempty"`
	OrderAssign  *string `json:"order_assign,omitempty"`
	LeadsAssign  *string `json:"leads_assign,omitempty"`
}

type SaleOrderTerms struct {
	ID          []string `json:"id,omitempty"`
	UserID      []int64  `json:"user_id,omitempty"`
	PackageID   []int64  `json:"package_id,omitempty"`
	UtmSource   []string `json:"utm_source,omitempty"`
	PayType     []string `json:"pay_type,omitempty"`
	Source      []int64  `json:"source,omitempty"`
	OrderAssign []string `json:"order_assign,omitempty"`
	LeadsAssign []string `json:"leads_assign,omitempty"`
}

type SaleOrderRange struct {
	Price      *SaleOrderRangeValue `json:"price,omitempty"`
	CreateTime *SaleOrderRangeValue `json:"ct,omitempty"`
	UpdateTime *SaleOrderRangeValue `json:"ut,omitempty"`
}

type SaleOrderRangeValue struct {
	GTE interface{} `json:"gte,omitempty"`
	LTE interface{} `json:"lte,omitempty"`
	GT  interface{} `json:"gt,omitempty"`
	LT  interface{} `json:"lt,omitempty"`
}

type SaleOrderMatchPhrase struct {
	MsgStr string `json:"msg_str,omitempty"`
}

type SaleOrders []*SaleOrder

// InsertOrUpdateSaleOrder 插入或更新订单数据到ES
func InsertOrUpdateSaleOrder(ctx context.Context, saleOrder *SaleOrder) error {
	jsonReq, err := json.Marshal(saleOrder)
	if err != nil {
		return err
	}
	return UpsertDoc(ctx, BussTypeSaleOrder, saleOrder.ID, string(jsonReq), SEARCH_UP_TYPE_UPSERT_FILED)
}

// DeleteSaleOrder 从ES中删除订单数据
func DeleteSaleOrder(ctx context.Context, id string) error {
	return DeleteDoc(ctx, BussTypeSaleOrder, id)
}

// SearchSaleOrder 根据条件查询订单
/*
Es 构建查询方式1：通过结构体构造然后最终marshal成为对应的查询json字符串
！！！！！！注意！！！！！！！！
Es有设置最大的默认结果窗口限制，即分页查询offset+limit的总和不能超过这个值，否则会返回错误
类似于mysql的深分页问题，offset太大会导致需要加载大量数据进入内存进行排序，导致性能下降。
导出需求可以使用scroll api进行分页查询，避免offset+limit的总和超过默认限制。
如果是后台查询的需求，则可以通过增加查询条件控制结果集的数量。判断offset+limit的大小超过默认限制时，返回空结果集，给出提示即可。
*/
func SearchSaleOrder(ctx context.Context, req *SearchSaleOrderRequest) (SaleOrders, int64, int64, bool, error) {
	// fun := "SearchSaleOrder-->"

	dsl := &SearchSaleOrderDSL{}
	must := make([]SaleOrderCondition, 0)

	// 添加查询条件
	if req.ID != "" {
		must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{ID: &req.ID}})
	}

	if len(req.IDs) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{ID: req.IDs}})
	}

	if req.UserID > 0 {
		must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{UserID: &req.UserID}})
	}

	if len(req.UserIDs) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{UserID: req.UserIDs}})
	}

	if req.PackageID > 0 {
		must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{PackageID: &req.PackageID}})
	}

	if len(req.PackageIDs) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{PackageID: req.PackageIDs}})
	}

	if len(req.UtmSources) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{UtmSource: req.UtmSources}})
	}

	if len(req.PayTypes) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{PayType: req.PayTypes}})
	}

	// 价格范围查询
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceRange := &SaleOrderRangeValue{}
		if req.MinPrice > 0 {
			priceRange.GTE = req.MinPrice
		}
		if req.MaxPrice > 0 {
			priceRange.LTE = req.MaxPrice
		}
		must = append(must, SaleOrderCondition{Range: &SaleOrderRange{Price: priceRange}})
	}

	if req.Source > 0 {
		must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{Source: &req.Source}})
	}

	if len(req.Sources) > 0 {
		must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{Source: req.Sources}})
	}

	if req.OuterOrderID != "" {
		must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{OuterOrderID: &req.OuterOrderID}})
	}

	// 处理order_assign和leads_assign的OR关系
	var shouldConditions []SaleOrderCondition

	// 单个值的Term查询
	if req.OrderAssign != "" {
		shouldConditions = append(shouldConditions, SaleOrderCondition{Term: &SaleOrderTerm{OrderAssign: &req.OrderAssign}})
	}
	if req.LeadsAssign != "" {
		shouldConditions = append(shouldConditions, SaleOrderCondition{Term: &SaleOrderTerm{LeadsAssign: &req.LeadsAssign}})
	}

	// 多个值的Terms查询
	if len(req.OrderAssigns) > 0 {
		shouldConditions = append(shouldConditions, SaleOrderCondition{Terms: &SaleOrderTerms{OrderAssign: req.OrderAssigns}})
	}
	if len(req.LeadsAssigns) > 0 {
		shouldConditions = append(shouldConditions, SaleOrderCondition{Terms: &SaleOrderTerms{LeadsAssign: req.LeadsAssigns}})
	}

	// 如果有任何assign条件，将它们作为should条件添加到查询中
	if len(shouldConditions) > 0 {
		// 创建一个bool查询作为must条件的一部分
		assignBool := SaleOrderBool{
			Should: shouldConditions,
		}

		// 将这个bool查询包装成一个条件
		must = append(must, SaleOrderCondition{
			Bool: &assignBool,
		})
	}

	// 消息内容模糊匹配
	if req.MsgStr != "" {
		must = append(must, SaleOrderCondition{MatchPhrase: &SaleOrderMatchPhrase{MsgStr: req.MsgStr}})
	}

	// 时间范围查询
	if req.BeginTime != nil || req.EndTime != nil {
		createTimeRange := &SaleOrderRangeValue{}
		if req.BeginTime != nil {
			createTimeRange.GTE = req.BeginTime
		}
		if req.EndTime != nil {
			createTimeRange.LTE = req.EndTime
		}
		must = append(must, SaleOrderCondition{Range: &SaleOrderRange{CreateTime: createTimeRange}})
	}

	if len(must) > 0 {
		dsl.Query.Bool.Must = must
	}

	// 默认按创建时间降序排序
	sort := []SaleOrderSort{{CreateTime: &SaleOrderSortOrder{Order: "desc"}}}
	dsl.Sort = sort

	dslStr, err := json.Marshal(dsl)
	if err != nil {
		return []*SaleOrder{}, 0, 0, false, err
	}
	// xlog.Infof(ctx, "%s dsl:%s", fun, dslStr)

	var search []*SaleOrder
	data, err := SearchByDSL(ctx, BussTypeSaleOrder, string(dslStr), req.Offset, req.Limit)
	if err != nil {
		return search, 0, 0, false, err
	}
	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &search)
	if err != nil {
		// xlog.Warnf(ctx, "%s queryres doc:%s qvalue:%s err:%+v", fun, docs, string(dslStr), err)
		return search, 0, 0, false, err
	}
	return search, data.Offset, data.Total, data.More, nil
}

// SearchSaleOrderV2 使用EsSearch方式构建查询的订单搜索方法--推荐使用
/*
Es查询方式第二种：通过elastic包官方方法构建，然后转为json查询即可
功能和SearchSaleOrder一致，但使用EsSearch结构体构建查询
*/
func SearchSaleOrderV2(ctx context.Context, req *SearchSaleOrderRequest) (SaleOrders, int64, int64, bool, error) {
	var search EsSearch

	// 添加查询条件
	if req.ID != "" {
		search.AddMustQuery(elastic.NewTermQuery("id", req.ID))
	}

	if len(req.IDs) > 0 {
		ids := make([]interface{}, len(req.IDs))
		for i, id := range req.IDs {
			ids[i] = id
		}
		search.AddMustQuery(elastic.NewTermsQuery("id", ids...))
	}

	if req.UserID > 0 {
		search.AddMustQuery(elastic.NewTermQuery("user_id", req.UserID))
	}

	if len(req.UserIDs) > 0 {
		userIDs := make([]interface{}, len(req.UserIDs))
		for i, userID := range req.UserIDs {
			userIDs[i] = userID
		}
		search.AddMustQuery(elastic.NewTermsQuery("user_id", userIDs...))
	}

	if req.PackageID > 0 {
		search.AddMustQuery(elastic.NewTermQuery("package_id", req.PackageID))
	}

	if len(req.PackageIDs) > 0 {
		packageIDs := make([]interface{}, len(req.PackageIDs))
		for i, packageID := range req.PackageIDs {
			packageIDs[i] = packageID
		}
		search.AddMustQuery(elastic.NewTermsQuery("package_id", packageIDs...))
	}

	if len(req.UtmSources) > 0 {
		utmSources := make([]interface{}, len(req.UtmSources))
		for i, utmSource := range req.UtmSources {
			utmSources[i] = utmSource
		}
		search.AddMustQuery(elastic.NewTermsQuery("utm_source", utmSources...))
	}

	if len(req.PayTypes) > 0 {
		payTypes := make([]interface{}, len(req.PayTypes))
		for i, payType := range req.PayTypes {
			payTypes[i] = payType
		}
		search.AddMustQuery(elastic.NewTermsQuery("pay_type", payTypes...))
	}

	// 价格范围查询
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceRange := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			priceRange.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			priceRange.Lte(req.MaxPrice)
		}
		search.AddMustQuery(priceRange)
	}

	if req.Source > 0 {
		search.AddMustQuery(elastic.NewTermQuery("source", req.Source))
	}

	if len(req.Sources) > 0 {
		sources := make([]interface{}, len(req.Sources))
		for i, source := range req.Sources {
			sources[i] = source
		}
		search.AddMustQuery(elastic.NewTermsQuery("source", sources...))
	}

	if req.OuterOrderID != "" {
		search.AddMustQuery(elastic.NewTermQuery("outer_order_id", req.OuterOrderID))
	}

	// 处理order_assign和leads_assign的OR关系
	var shouldConditions []elastic.Query

	// 单个值的Term查询
	if req.OrderAssign != "" {
		shouldConditions = append(shouldConditions, elastic.NewTermQuery("order_assign", req.OrderAssign))
	}
	if req.LeadsAssign != "" {
		shouldConditions = append(shouldConditions, elastic.NewTermQuery("leads_assign", req.LeadsAssign))
	}

	// 多个值的Terms查询
	if len(req.OrderAssigns) > 0 {
		orderAssigns := make([]interface{}, len(req.OrderAssigns))
		for i, orderAssign := range req.OrderAssigns {
			orderAssigns[i] = orderAssign
		}
		shouldConditions = append(shouldConditions, elastic.NewTermsQuery("order_assign", orderAssigns...))
	}
	if len(req.LeadsAssigns) > 0 {
		leadsAssigns := make([]interface{}, len(req.LeadsAssigns))
		for i, leadsAssign := range req.LeadsAssigns {
			leadsAssigns[i] = leadsAssign
		}
		shouldConditions = append(shouldConditions, elastic.NewTermsQuery("leads_assign", leadsAssigns...))
	}

	// 如果有任何assign条件，将它们作为should条件添加到查询中
	if len(shouldConditions) > 0 {
		search.AddMustQuery(elastic.NewBoolQuery().Should(shouldConditions...).MinimumShouldMatch("1"))
	}

	// 消息内容模糊匹配
	if req.MsgStr != "" {
		search.AddMustQuery(elastic.NewMatchPhraseQuery("msg_str", req.MsgStr))
	}

	// 时间范围查询
	if req.BeginTime != nil || req.EndTime != nil {
		createTimeRange := elastic.NewRangeQuery("ct")
		if req.BeginTime != nil {
			createTimeRange.Gte(req.BeginTime)
		}
		if req.EndTime != nil {
			createTimeRange.Lte(req.EndTime)
		}
		search.AddMustQuery(createTimeRange)
	}

	// 默认按创建时间降序排序
	search.Sorters = append(search.Sorters, elastic.NewFieldSort("ct").Desc())

	dsl, err := search.BoolQueryDSL()
	if err != nil {
		return []*SaleOrder{}, 0, 0, false, err
	}

	var searchResult []*SaleOrder
	data, err := SearchByDSL(ctx, BussTypeSaleOrder, string(dsl), req.Offset, req.Limit)
	if err != nil {
		return searchResult, 0, 0, false, err
	}

	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &searchResult)
	if err != nil {
		return searchResult, 0, 0, false, err
	}

	return searchResult, data.Offset, data.Total, data.More, nil
}

// BatchGetSaleOrderByUserIDs 批量获取用户订单
func BatchGetSaleOrderByUserIDs(ctx context.Context, userIDs []int64) (SaleOrders, int64, int64, bool, error) {
	// fun := "BatchGetSaleOrderByUserIDs-->"

	if len(userIDs) == 0 {
		return []*SaleOrder{}, 0, 0, false, nil
	}

	dsl := &SearchSaleOrderDSL{}
	must := make([]SaleOrderCondition, 0)

	must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{UserID: userIDs}})

	if len(must) > 0 {
		dsl.Query.Bool.Must = must
	}

	// 默认按创建时间降序排序
	sort := []SaleOrderSort{{CreateTime: &SaleOrderSortOrder{Order: "desc"}}}
	dsl.Sort = sort

	dslStr, err := json.Marshal(dsl)
	if err != nil {
		return []*SaleOrder{}, 0, 0, false, err
	}
	// xlog.Infof(ctx, "%s dsl:%s", fun, dslStr)

	var search []*SaleOrder
	data, err := SearchByDSL(ctx, BussTypeSaleOrder, string(dslStr), 0, 0)
	if err != nil {
		return search, 0, 0, false, err
	}
	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &search)
	if err != nil {
		// xlog.Warnf(ctx, "%s queryres doc:%s qvalue:%s err:%+v", fun, docs, string(dslStr), err)
		return search, 0, 0, false, err
	}
	return search, data.Offset, data.Total, data.More, nil
}

// GetSaleOrderByID 根据ID获取订单
func GetSaleOrderByID(ctx context.Context, id string) (*SaleOrder, error) {
	// fun := "GetSaleOrderByID-->"

	dsl := &SearchSaleOrderDSL{}
	must := make([]SaleOrderCondition, 0)

	must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{ID: &id}})

	if len(must) > 0 {
		dsl.Query.Bool.Must = must
	}

	dslStr, err := json.Marshal(dsl)
	if err != nil {
		return nil, err
	}
	// xlog.Infof(ctx, "%s dsl:%s", fun, dslStr)

	var search []*SaleOrder
	data, err := SearchByDSL(ctx, BussTypeSaleOrder, string(dslStr), 0, 1)
	if err != nil {
		return nil, err
	}
	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &search)
	if err != nil {
		// xlog.Warnf(ctx, "%s queryres doc:%s qvalue:%s err:%+v", fun, docs, string(dslStr), err)
		return nil, err
	}

	if len(search) == 0 {
		return nil, nil
	}
	return search[0], nil
}

// SearchSaleOrderWithDSL 使用自定义DSL查询订单
func SearchSaleOrderWithDSL(ctx context.Context, dsl string, offset, limit int64) (SaleOrders, int64, int64, bool, error) {
	// fun := "SearchSaleOrderWithDSL-->"
	// xlog.Infof(ctx, "%s dsl:%s", fun, dsl)
	var search []*SaleOrder
	data, err := SearchByDSL(ctx, BussTypeSaleOrder, dsl, offset, limit)
	if err != nil {
		return search, 0, 0, false, err
	}
	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &search)
	if err != nil {
		// xlog.Warnf(ctx, "%s queryres doc:%s qvalue:%s err:%+v", fun, docs, dsl, err)
		return search, 0, 0, false, err
	}
	return search, data.Offset, data.Total, data.More, nil
}

// BatchUpdateLeadsAssignByUserIDs 批量更新指定用户ID的记录的leads_assign字段（高效实现）
func BatchUpdateLeadsAssignByUserIDs(ctx context.Context, userIDs []int64, leadsAssign string) (int64, error) {
	// fun := "BatchUpdateLeadsAssignByUserIDs-->"

	if len(userIDs) == 0 || leadsAssign == "" {
		return 0, nil
	}

	// 构建update_by_query请求
	type Script struct {
		Source string                 `json:"source"`
		Lang   string                 `json:"lang"`
		Params map[string]interface{} `json:"params"`
	}

	type Query struct {
		Terms map[string]interface{} `json:"terms"`
	}

	type UpdateByQueryRequest struct {
		Script Script `json:"script"`
		Query  Query  `json:"query"`
	}

	// 构建更新脚本
	now := time.Now()
	script := Script{
		Source: "ctx._source.leads_assign = params.leadsAssign; ctx._source.ut = params.updateTime",
		Lang:   "painless",
		Params: map[string]interface{}{
			"leadsAssign": leadsAssign,
			"updateTime":  &now,
		},
	}

	// 构建查询条件
	query := Query{
		Terms: map[string]interface{}{
			"user_id": userIDs,
		},
	}

	request := UpdateByQueryRequest{
		Script: script,
		Query:  query,
	}

	// 序列化请求
	requestBody, err := json.Marshal(request)
	if err != nil {
		// xlog.Errorf(ctx, "%s marshal request failed, err: %v", fun, err)
		return 0, err
	}

	// 调用update_by_query API
	result, err := UpdateByQuery(ctx, BussTypeSaleOrder, string(requestBody))
	if err != nil {
		// xlog.Errorf(ctx, "%s UpdateByQuery failed, userIDs: %v, err: %v", fun, userIDs, err)
		return 0, err
	}

	// xlog.Infof(ctx, "%s successfully updated %d/%d records for userIDs: %v",
	// 	fun, result.Updated, result.Total, userIDs)

	return result.Updated, nil
}

// BatchUpdateLeadsAssignByUserID 更新单个用户ID的所有记录的leads_assign字段
func BatchUpdateLeadsAssignByUserID(ctx context.Context, userID int64, leadsAssign string) (int64, error) {
	return BatchUpdateLeadsAssignByUserIDs(ctx, []int64{userID}, leadsAssign)
}
