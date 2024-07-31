package main

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/akali/http/resp"
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/platform/idl/xerr"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"math/rand"
	"strings"
	excel "test/generate-excel"
	"test/gin-demo/model"
	"time"
)

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin!",
	})
}

type Test struct {
	Name   string `form:"name" json:"name"`
	Brand  string `form:"brand" json:"brand"`
	UserId int    `form:"user_id" json:"user_id"`
}

type Test1 struct {
	Name         string `form:"name" json:"name"`
	Brand        string `form:"brand" json:"brand"`
	UserIdString string `form:"user_id" json:"user_id"`
}

func TestShouldBind(c *gin.Context) {
	test := &Test{}
	err := c.ShouldBindWith(test, binding.Form) // 这个不会影响bind次数
	if err != nil {
		fmt.Println(err)
		return
	}
	test2 := &Test{}
	err2 := c.ShouldBind(test2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// ShouldBind绑定参数时，如果参数类型是form-data/x-www-form-urlencoded时,可以多次使用ShouldBind
	// 但是如果参数类型是json时，只能使用一次ShouldBind
	// test3 := &Test{}
	// err3 := c.ShouldBindJSON(test3)
	// if err3 != nil {
	//	fmt.Println(err3)
	//	return
	// }
	// test4 := &Test{}
	// err4 := c.ShouldBindJSON(test4)
	// if err4 != nil {
	//	fmt.Println(err4)
	//	return
	// }
	// ShouldBindJSON不能使用多次，尽管是针对不同地址空间的结构体,也不能和shouldBind共同使用多次
	// 1.单次解析，追求性能使用 ShouldBindJson，因为多次绑定解析会出现EOF
	// 2.多次解析，避免报EOF，使用ShouldBindBodyWith
	c.JSON(200, gin.H{
		"message": "hello TestShouldBind!",
	})
}

// 假设的任务函数，有一定几率失败
func task(ctx context.Context) error {
	// 模拟50%几率成功或失败
	if rand.Intn(2) == 0 {
		return fmt.Errorf("task failed")
	}
	fmt.Println(ctx.Value("x_trace_id"))
	fmt.Println("task succeeded")
	return nil
}

// 异步执行任务，带有重试逻辑
func asyncTaskWithRetry(ctx context.Context, maxRetries int, delay time.Duration, t func(ctx context.Context) error) {
	go func() {
		for retries := 0; retries < maxRetries; retries++ {
			err := t(ctx)
			if err != nil {
				fmt.Println(err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
		fmt.Println("子流程执行完成")
	}()
}

func TestRetry(c *gin.Context) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "x_trace_id", "123456")

	// 启动异步任务，最多重试3次，每次重试间隔1秒
	asyncTaskWithRetry(ctx, 5, 20*time.Second, task)

	// 主流程继续执行其他任务
	fmt.Println("主流程继续执行...")
	time.Sleep(2 * time.Second) // 假设主流程有其他耗时任务
	fmt.Println("主流程结束")
	c.JSON(200, gin.H{
		"code":    200,
		"message": "done",
	})
}

/*
 subject STRING COMMENT '学科',
 grade STRING COMMENT '年级',
 source STRING COMMENT '来源（语文是古诗名）',
 content STRING COMMENT '题目',
 candidate_list STRING COMMENT '4个选项信息',
 right_idx STRING COMMENT '正确答案位置')
*/

func ImportCheckinQuestion(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		resp.WriteErrJSON(ctx, xerr.ErrorInternalError(fmt.Sprintf("get form err: %s", err.Error())))
		return
	}
	// 读取文件
	fileBytes, err := file.Open()
	if err != nil {
		resp.WriteErrJSON(ctx, xerr.ErrorInternalError(fmt.Sprintf("read file err: %s", err.Error())))
		return
	}
	defer fileBytes.Close()
	res, err := importCheckinQuestion(ctx.Request.Context(), fileBytes, file.Filename)
	if err != nil {
		resp.WriteErrJSON(ctx, xerr.ErrorInternalError(""))
		return
	}
	resp.WriteSuccessJSON(ctx, res)
}

// ImportCheckinQuestion 导入签到题目
func importCheckinQuestion(ctx context.Context, file io.Reader, fileName string) (*model.ImportCheckinQuestionResp, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("文件【%s】打开失败，error: %v", fileName, err)
	}
	resp := &model.ImportCheckinQuestionResp{
		Filename: fileName,
	}
	datas := make([]*model.Question, 0)
	// 获取 Sheet 上所有单元格数据
	sheetName := f.GetSheetName(0)
	rows, _ := f.GetRows(sheetName)
	if len(rows) == 0 {
		log.Println(fmt.Sprintf("Sheet【%s】没有数据", sheetName))
		return nil, nil
	}
	total := len(rows) - 1
	log.Println(fmt.Sprintf("Sheet【%s】总共有【%d】行数据", sheetName, total))
	succNum := 0
	for j, row := range rows {
		j++
		// 第一行是标题，不处理
		if j == 1 {
			continue
		}
		data, err := ParseQuestionData(ctx, sheetName, row)
		if err != nil {
			log.Println(fmt.Sprintf("Sheet【%s】第【%d】行 %s", sheetName, j, err.Error()))
			continue
		}
		datas = append(datas, data)
		succNum++
	}
	// 生成表格
	ExcelProcess(datas)
	sheet := model.ImportCheckinQuestionSheet{
		SheetName: sheetName,
		ImportCommonItem: model.ImportCommonItem{
			TotalNum:   total,
			SuccessNum: succNum,
			FailNum:    total - succNum,
		},
	}
	resp.Sheets = append(resp.Sheets, sheet)
	return resp, nil
}

func ExcelProcess(list []*model.Question) {
	begin := time.Now().Unix()
	err := excel.ExcelProcess(list).
		Headers("subject", "grade", "source", "content", "candidate_list", "right_idx").
		Columns("subject", "grade", "source", "content", "candidate_list", "right_idx").
		Sheet("Sheet1").
		// Style(func(currentSheet string, f *excelize.File) error {
		// 	styleId, err := f.NewStyle(&excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 20}})
		// 	if err != nil {
		// 		return err
		// 	}
		// 	return f.SetCellStyle(currentSheet, "A1", "H1", styleId)
		// }).
		SavePath("demo.xlsx").ToExcel().Error
	if err != nil {
		log.Println("err:", err)
	}
	end := time.Now().Unix()
	fmt.Println("表格生成耗费时长：", end-begin, "s")
}

func ParseQuestionData(ctx context.Context, sheetName string, row []string) (*model.Question, error) {
	if len(row) < 4 {
		return nil, fmt.Errorf("数据不完整")
	}
	content := &model.Content{}
	err := json.Unmarshal([]byte(row[3]), content)
	if err != nil {
		return nil, err
	}
	candidate := &model.Candidate{}
	err = json.Unmarshal([]byte(row[4]), candidate)
	if err != nil {
		return nil, err
	}
	rightIdx := string(row[5][2])
	data := &model.Question{
		Subject:       row[0],
		Grade:         row[1],
		Source:        row[2],
		Content:       content.Content,
		CandidateList: strings.Join(candidate.Content, ","),
		RightIdx:      rightIdx,
	}
	return data, nil
	// if sheetName == constant.CheckinQuestionIdiom {
	// 	if len(row) < 4 {
	// 		return nil, fmt.Errorf("数据不完整")
	// 	}
	// 	options, err := GetQuestionOption(ctx, row[2])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	data := &model.CheckinQuestion{
	// 		Type:    constant.CheckinQuestionTypeIdiom,
	// 		Content: row[1],
	// 		Options: options,
	// 		Answer:  row[3],
	// 		Status:  constant.CheckinQuestionStatusEnable,
	// 	}
	// 	return data, nil
	// }
	// if sheetName == constant.CheckinQuestionArithmetic {
	// 	if len(row) < 3 {
	// 		return nil, fmt.Errorf("数据不完整")
	// 	}
	// 	options, err := GetQuestionOption(ctx, row[1])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	data := &model.CheckinQuestion{
	// 		Type:    constant.CheckinQuestionTypeArithmetic,
	// 		Content: row[0],
	// 		Options: options,
	// 		Answer:  row[2],
	// 		Status:  constant.CheckinQuestionStatusEnable,
	// 	}
	// 	return data, nil
	// }
	// if sheetName == constant.CheckinQuestionEnglish {
	// 	if len(row) < 6 {
	// 		return nil, fmt.Errorf("数据不完整")
	// 	}
	// 	content := gjson.Get(row[0], "content").String()
	// 	if content == "" {
	// 		return nil, fmt.Errorf("英语题目content为空")
	// 	}
	// 	optionObjArr := make([]model.CheckinQuestionOption, 0)
	// 	optionObjArr = append(optionObjArr,
	// 		model.CheckinQuestionOption{
	// 			Label:   "A",
	// 			Content: row[1],
	// 		},
	// 		model.CheckinQuestionOption{
	// 			Label:   "B",
	// 			Content: row[2],
	// 		},
	// 		model.CheckinQuestionOption{
	// 			Label:   "C",
	// 			Content: row[3],
	// 		},
	// 		model.CheckinQuestionOption{
	// 			Label:   "D",
	// 			Content: row[4],
	// 		},
	// 	)
	// 	options, err := json.Marshal(optionObjArr)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if _, ok := constant.CheckinQuestionOptionMap[row[5]]; !ok {
	// 		return nil, fmt.Errorf("答案不正确")
	// 	}
	// 	data := &model.CheckinQuestion{
	// 		Type:    constant.CheckinQuestionTypeEnglish,
	// 		Content: content,
	// 		Options: options,
	// 		Answer:  constant.CheckinQuestionOptionMap[row[5]],
	// 		Status:  constant.CheckinQuestionStatusEnable,
	// 	}
	// 	return data, nil
	// }
	// if sheetName == constant.CheckinQuestionFun {
	// 	if len(row) < 3 {
	// 		return nil, fmt.Errorf("数据不完整")
	// 	}
	// 	options, err := GetQuestionOption(ctx, row[1])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	data := &model.CheckinQuestion{
	// 		Type:    constant.CheckinQuestionTypeFun,
	// 		Content: row[0],
	// 		Options: options,
	// 		Answer:  row[2],
	// 		Status:  constant.CheckinQuestionStatusEnable,
	// 	}
	// 	return data, nil
	// }
	return nil, fmt.Errorf("未知的sheet名")
}
func GetQuestionOption(ctx context.Context, option string) ([]byte, error) {
	if option == "" {
		return nil, fmt.Errorf("选项为空")
	}
	optionObjArr := make([]model.CheckinQuestionOption, 0)
	optionArr := strings.Split(option, "\n")
	for _, v := range optionArr {
		arr := strings.Split(v, "、")
		if len(arr) < 2 {
			continue
		}
		optionObjArr = append(optionObjArr, model.CheckinQuestionOption{
			Label:   arr[0],
			Content: arr[1],
		})
	}
	return json.Marshal(optionObjArr)
}

func AddCheckinQuestion(ctx context.Context, data *model.CheckinQuestion) error {
	err := DB.Table("checkin_question").Create(data).Error
	if err != nil {
		logger.Ex(ctx, "data: AddCheckinQuestion", "data[%v], err[%v]", data, err)
		return err
	}
	return nil
}
