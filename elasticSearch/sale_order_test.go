package es

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/samber/lo"

	"github.com/stretchr/testify/assert"
)

func TestSaleOrderMarshalUnmarshal(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")
	now := time.Now()
	order := &SaleOrder{
		ID:           "1001",
		UserID:       2001,
		PackageID:    301,
		MsgStr:       "测试订单消息",
		UtmSource:    "wechat",
		PayType:      "alipay",
		Price:        9900,
		Source:       1,
		OuterOrderID: "OUT20250915001",
		OrderAssign:  "user1",
		LeadsAssign:  "leads1",
		CreateTime:   &now,
		UpdateTime:   &now,
	}

	// 测试序列化
	jsonData, err := json.Marshal(order)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// 测试反序列化
	var newOrder SaleOrder
	err = json.Unmarshal(jsonData, &newOrder)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, newOrder.ID)
	assert.Equal(t, order.UserID, newOrder.UserID)
	assert.Equal(t, order.PackageID, newOrder.PackageID)
	assert.Equal(t, order.MsgStr, newOrder.MsgStr)
	assert.Equal(t, order.UtmSource, newOrder.UtmSource)
	assert.Equal(t, order.PayType, newOrder.PayType)
	assert.Equal(t, order.Price, newOrder.Price)
	assert.Equal(t, order.Source, newOrder.Source)
	assert.Equal(t, order.OuterOrderID, newOrder.OuterOrderID)
	assert.Equal(t, order.OrderAssign, newOrder.OrderAssign)
	assert.Equal(t, order.LeadsAssign, newOrder.LeadsAssign)
}

func TestBuildSearchDSL(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	// 构建查询请求
	beginTime := time.Date(2025, 9, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2025, 9, 15, 23, 59, 59, 0, time.Local)

	req := &SearchSaleOrderRequest{
		ID:          "order123",
		UserID:      1001,
		PackageIDs:  []int64{101, 102, 103},
		UtmSources:  []string{"wechat", "website"},
		PayTypes:    []string{"alipay", "wechat"},
		MinPrice:    5000,
		MaxPrice:    10000,
		OrderAssign: "user1",
		LeadsAssign: "leads1",
		BeginTime:   &beginTime,
		EndTime:     &endTime,
		Limit:       10,
		Offset:      0,
	}

	// 构建DSL
	dsl := &SearchSaleOrderDSL{}
	must := make([]SaleOrderCondition, 0)

	id := req.ID
	must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{ID: &id}})
	must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{UserID: &req.UserID}})
	must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{PackageID: req.PackageIDs}})
	must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{UtmSource: req.UtmSources}})
	must = append(must, SaleOrderCondition{Terms: &SaleOrderTerms{PayType: req.PayTypes}})

	// 价格范围
	priceRange := &SaleOrderRangeValue{
		GTE: req.MinPrice,
		LTE: req.MaxPrice,
	}
	must = append(must, SaleOrderCondition{Range: &SaleOrderRange{Price: priceRange}})

	// 订单分配和线索分配
	orderAssign := req.OrderAssign
	leadsAssign := req.LeadsAssign
	must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{OrderAssign: &orderAssign}})
	must = append(must, SaleOrderCondition{Term: &SaleOrderTerm{LeadsAssign: &leadsAssign}})

	// 时间范围
	createTimeRange := &SaleOrderRangeValue{
		GTE: req.BeginTime.Format("2006-01-02 15:04:05"),
		LTE: req.EndTime.Format("2006-01-02 15:04:05"),
	}
	must = append(must, SaleOrderCondition{Range: &SaleOrderRange{CreateTime: createTimeRange}})

	dsl.Query.Bool.Must = must

	// 排序
	sort := []SaleOrderSort{{CreateTime: &SaleOrderSortOrder{Order: "desc"}}}
	dsl.Sort = sort

	// 序列化DSL
	dslStr, err := json.Marshal(dsl)
	assert.NoError(t, err)
	assert.NotEmpty(t, dslStr)

	// 这里只是验证DSL生成是否正确，不实际调用ES
	t.Logf("Generated DSL: %s", string(dslStr))
}

func TestSearchSaleOrder_Integration(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 创建测试订单
	now := time.Now()
	order := &SaleOrder{
		ID:           "test9999",
		UserID:       8888,
		PackageID:    7777,
		MsgStr:       "测试订单",
		UtmSource:    "test",
		PayType:      "test",
		Price:        1000,
		Source:       1,
		OuterOrderID: "TEST_OUTER_ID",
		OrderAssign:  "test_user",
		LeadsAssign:  "test_leads",
		CreateTime:   &now,
		UpdateTime:   &now,
	}

	// 插入测试数据
	err := InsertOrUpdateSaleOrder(ctx, order)
	assert.NoError(t, err)

	// 查询测试
	req := &SearchSaleOrderRequest{
		ID: order.ID,
	}

	orders, offset, total, more, err := SearchSaleOrder(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.False(t, more)
	assert.Equal(t, int64(0), offset)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, order.ID, orders[0].ID)

	// 测试完成后删除测试数据
	// err = DeleteSaleOrder(ctx, order.ID)
	assert.NoError(t, err)
}

func TestSearchSaleOrderV2_Integration(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 创建测试订单
	now := time.Now()
	order := &SaleOrder{
		ID:           "test9999_v2",
		UserID:       8889,
		PackageID:    7778,
		MsgStr:       "测试订单V2",
		UtmSource:    "test_v2",
		PayType:      "test_v2",
		Price:        2000,
		Source:       1,
		OuterOrderID: "TEST_OUTER_ID_V2",
		OrderAssign:  "test_user_v2",
		LeadsAssign:  "test_leads_v2",
		CreateTime:   &now,
		UpdateTime:   &now,
	}

	// 插入测试数据
	err := InsertOrUpdateSaleOrder(ctx, order)
	assert.NoError(t, err)

	// 查询测试 - 使用SearchSaleOrderV2
	req := &SearchSaleOrderRequest{
		ID: order.ID,
	}

	orders, offset, total, more, err := SearchSaleOrderV2(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.False(t, more)
	assert.Equal(t, int64(0), offset)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, order.ID, orders[0].ID)

	// 测试完成后删除测试数据
	// err = DeleteSaleOrder(ctx, order.ID)
	assert.NoError(t, err)
}

func TestSearchSaleOrderV2_CompareWithOriginal(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 创建测试订单
	now := time.Now()
	order := &SaleOrder{
		ID:           "test_compare",
		UserID:       9999,
		PackageID:    8888,
		MsgStr:       "比较测试订单",
		UtmSource:    "compare_test",
		PayType:      "compare_test",
		Price:        3000,
		Source:       1,
		OuterOrderID: "COMPARE_TEST_ID",
		OrderAssign:  "compare_user",
		LeadsAssign:  "compare_leads",
		CreateTime:   &now,
		UpdateTime:   &now,
	}

	// 插入测试数据
	err := InsertOrUpdateSaleOrder(ctx, order)
	assert.NoError(t, err)

	// 使用相同的查询条件测试两个方法
	req := &SearchSaleOrderRequest{
		UserID: order.UserID,
		Limit:  10,
		Offset: 0,
	}

	// 测试原始方法
	orders1, offset1, total1, more1, err1 := SearchSaleOrder(ctx, req)
	assert.NoError(t, err1)

	// 测试V2方法
	orders2, offset2, total2, more2, err2 := SearchSaleOrderV2(ctx, req)
	assert.NoError(t, err2)

	// 比较结果
	assert.Equal(t, total1, total2, "Total count should be the same")
	assert.Equal(t, offset1, offset2, "Offset should be the same")
	assert.Equal(t, more1, more2, "More flag should be the same")
	assert.Equal(t, len(orders1), len(orders2), "Result count should be the same")

	// 测试完成后删除测试数据
	// err = DeleteSaleOrder(ctx, order.ID)
	assert.NoError(t, err)
}

func TestInsertOrUpdateReadcampLeadsOrder(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")
	ctx := context.Background()
	order := &SaleOrder{
		ID:     "test9999",
		UserID: 8888,
	}
	err := InsertOrUpdateSaleOrder(ctx, order)
	assert.NoError(t, err)
}

func TestBatchUpdateLeadsAssignByUserIDs(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")
	ctx := context.Background()
	// 准备测试数据
	now := time.Now()
	testOrders := []*SaleOrder{
		{
			ID:          "test_batch_1",
			UserID:      9001,
			PackageID:   101,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
		{
			ID:          "test_batch_2",
			UserID:      9001,
			PackageID:   102,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
		{
			ID:          "test_batch_3",
			UserID:      9002,
			PackageID:   103,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
	}

	// 插入测试数据
	for _, order := range testOrders {
		err := InsertOrUpdateSaleOrder(ctx, order)
		assert.NoError(t, err)
	}

	// 执行批量更新
	userIDs := []int64{9001}
	newLeadsAssign := "new_assign"
	count, err := BatchUpdateLeadsAssignByUserIDs(ctx, userIDs, newLeadsAssign)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, int64(3), count) // 应该有3条记录被更新
	//
	// // 验证更新是否成功
	// for _, id := range []string{"test_batch_1", "test_batch_2", "test_batch_3"} {
	// 	order, err := GetSaleOrderByID(ctx, id)
	// 	assert.NoError(t, err)
	// 	assert.NotNil(t, order)
	// 	assert.Equal(t, newLeadsAssign, order.LeadsAssign)
	// }
	//
	// // 清理测试数据
	// for _, order := range testOrders {
	// 	err := DeleteSaleOrder(ctx, order.ID)
	// 	assert.NoError(t, err)
	// }
}

func TestBatchUpdateLeadsAssignByUserIDs_UpdateByQuery(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 准备测试数据
	now := time.Now()
	testOrders := []*SaleOrder{
		{
			ID:          "test_update_by_query_1",
			UserID:      9101,
			PackageID:   201,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
		{
			ID:          "test_update_by_query_2",
			UserID:      9101,
			PackageID:   202,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
		{
			ID:          "test_update_by_query_3",
			UserID:      9102,
			PackageID:   203,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
	}

	// 插入测试数据
	for _, order := range testOrders {
		err := InsertOrUpdateSaleOrder(ctx, order)
		assert.NoError(t, err)
	}

	// 执行批量更新
	userIDs := []int64{9101, 9102}
	newLeadsAssign := "new_update_by_query"
	count, err := BatchUpdateLeadsAssignByUserIDs(ctx, userIDs, newLeadsAssign)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, int64(3), count) // 应该有3条记录被更新

	// 验证更新是否成功 - 由于update_by_query是异步操作，可能需要等待一段时间
	time.Sleep(1 * time.Second) // 等待ES更新完成

	for _, id := range []string{"test_update_by_query_1", "test_update_by_query_2", "test_update_by_query_3"} {
		order, err := GetSaleOrderByID(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, newLeadsAssign, order.LeadsAssign)
	}

	// 清理测试数据
	for _, order := range testOrders {
		err := DeleteSaleOrder(ctx, order.ID)
		assert.NoError(t, err)
	}
}

func TestBatchUpdateLeadsAssignByUserID(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 准备测试数据
	now := time.Now()
	testOrders := []*SaleOrder{
		{
			ID:          "test_single_1",
			UserID:      9003,
			PackageID:   201,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
		{
			ID:          "test_single_2",
			UserID:      9003,
			PackageID:   202,
			LeadsAssign: "old_assign",
			CreateTime:  &now,
			UpdateTime:  &now,
		},
	}

	// 插入测试数据
	for _, order := range testOrders {
		err := InsertOrUpdateSaleOrder(ctx, order)
		assert.NoError(t, err)
	}

	// 执行单用户批量更新
	userID := int64(9003)
	newLeadsAssign := "new_single_assign"
	count, err := BatchUpdateLeadsAssignByUserID(ctx, userID, newLeadsAssign)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count) // 应该有2条记录被更新

	// 验证更新是否成功
	for _, id := range []string{"test_single_1", "test_single_2"} {
		order, err := GetSaleOrderByID(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, newLeadsAssign, order.LeadsAssign)
	}

	// 清理测试数据
	for _, order := range testOrders {
		err := DeleteSaleOrder(ctx, order.ID)
		assert.NoError(t, err)
	}
}

func TestDeleteSaleOrder(t *testing.T) {
	t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 搜索测试订单
	orders, offset, total, more, err := SearchSaleOrder(ctx, &SearchSaleOrderRequest{
		Limit:  20,
		Offset: 0,
	})
	fmt.Println(offset, total, more)

	// 执行删除
	for _, v := range orders {
		err = DeleteSaleOrder(ctx, v.ID)
		assert.NoError(t, err)
	}
}

func TestQueryLectures(t *testing.T) {
	// t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()

	// 搜索测试订单
	orders, offset, total, more, err := SearchSaleOrder(ctx, &SearchSaleOrderRequest{
		// MsgStr:       "体验课",
		LeadsAssigns: []string{"zhangtianyong26331"},
		Limit:        20,
		Offset:       0,
	})
	fmt.Println(offset, total, more, err)
	fmt.Println(len(orders))
	for _, v := range orders {
		fmt.Println(v.ID, v.UserID)
	}
}

func TestOrderAssignOrLeadsAssign(t *testing.T) {
	// t.Skip("集成测试，需要ES环境，默认跳过")

	ctx := context.Background()
	beginTime := lo.ToPtr(time.Now().AddDate(0, -2, 0))
	endTime := lo.ToPtr(time.Now())
	// 测试order_assign和leads_assign的OR关系
	req := &SearchSaleOrderRequest{
		OrderAssigns: []string{"zhusuqian26838"},
		LeadsAssigns: []string{"zhangtianyong26331"},
		BeginTime:    beginTime,
		EndTime:      endTime,
		Limit:        20,
		Offset:       0,
	}

	orders, offset, total, more, err := SearchSaleOrder(ctx, req)
	fmt.Println("查询结果:", offset, total, more, err)
	fmt.Println("记录数量:", len(orders))

	// 打印查询条件生成的DSL
	dsl := &SearchSaleOrderDSL{}
	must := make([]SaleOrderCondition, 0)

	// 处理order_assign和leads_assign的OR关系
	var shouldConditions []SaleOrderCondition

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

	// 添加消息内容模糊匹配
	if req.MsgStr != "" {
		must = append(must, SaleOrderCondition{MatchPhrase: &SaleOrderMatchPhrase{MsgStr: req.MsgStr}})
	}

	dsl.Query.Bool.Must = must

	// 打印DSL
	dslStr, _ := json.MarshalIndent(dsl, "", "  ")
	fmt.Println("生成的DSL:", string(dslStr))
}
