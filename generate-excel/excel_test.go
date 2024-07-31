package excel

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestExcelProcessStream(t *testing.T) {
	var list []OperationLog

	for i := 1; i <= 10; i++ {
		list = append(list, OperationLog{
			Id:       i,
			UserName: "xxx" + cast.ToString(i),
			IP:       "127.0.0.1" + cast.ToString(i),
			TypeStr:  "xxx" + cast.ToString(i),
			Module:   "xxx" + cast.ToString(i),
			Res:      "xxx" + cast.ToString(i),
			Time:     strconv.FormatInt(time.Now().Unix(), 10),
			Des:      "xxx" + cast.ToString(i),
		})
	}
	begin := time.Now().Unix()
	err := ExcelProcessStream(list).
		Headers("序号", "用户名", "登录IP", "用户类型", "操作模块", "操作结果", "操作时间", "描述").
		Columns("Id", "UserName", "IP", "TypeStr", "Module", "Res", "Time", "Des").
		SavePath("demo.xlsx").SetColWidth(2, 7, 25).SetColWidth(8, 8, 80).
		NewStyle(KeyStyle{
			Key:   "style1",
			Style: &excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 12}},
		}, KeyStyle{
			Key: "style2",
			Style: &excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 12}, Fill: excelize.Fill{
				Type:    "pattern",
				Pattern: 1,
				Color:   []string{"#E8E8E8"},
			}},
		}).Style(func(styleMap map[string]int, x int, y int, val interface{}) (int, error) {
		if y%2 == 0 {
			return styleMap["style2"], nil
		} else {
			return styleMap["style1"], nil
		}
	}).ToExcel().Error
	if err != nil {
		log.Println("err:", err)
	}
	end := time.Now().Unix()
	fmt.Println("表格生成耗费时长：", end-begin, "s")
}

func TestExcelProcess(t *testing.T) {
	var list []OperationLog
	for i := 1; i <= 500000; i++ {
		list = append(list, OperationLog{
			Id:       i,
			UserName: "xxx" + cast.ToString(i),
			IP:       "127.0.0.1" + cast.ToString(i),
			TypeStr:  "xxx" + cast.ToString(i),
			Module:   "xxx" + cast.ToString(i),
			Res:      "xxx" + cast.ToString(i),
			Time:     strconv.FormatInt(time.Now().Unix(), 10),
			Des:      "xxx" + cast.ToString(i),
		})
	}
	begin := time.Now().Unix()
	err := ExcelProcess(list).
		Headers("序号", "用户名", "登录IP", "用户类型", "操作模块", "操作结果", "操作时间", "描述").
		Columns("Id", "UserName", "IP", "TypeStr", "Module", "Res", "Time", "Des").
		Sheet("test").
		Style(func(currentSheet string, f *excelize.File) error {
			styleId, err := f.NewStyle(&excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 20}})
			if err != nil {
				return err
			}
			return f.SetCellStyle(currentSheet, "A1", "H1", styleId)
		}).
		SavePath("demo.xlsx").ToExcel().Error
	if err != nil {
		log.Println("err:", err)
	}
	end := time.Now().Unix()
	fmt.Println("表格生成耗费时长：", end-begin, "s")
}

func TestAppendExcel(t *testing.T) {
	f, _ := excelize.OpenFile("demo.xlsx")
	index, _ := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "B1", 100)
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)
	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestReadExcel(t *testing.T) {
	f, _ := excelize.OpenFile("demo.xlsx")
	rows, _ := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func HttpReadExcel() {
	f := func(read io.Reader) {
		file, err := excelize.OpenReader(read)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, item := range file.GetSheetList() {
			rows, _ := file.GetRows(item)
			// rows是一个二维数组
			for i := range rows {
				strs := rows[i]
				for j := range strs {
					str := rows[i][j]
					fmt.Println(str)
				}
			}
		}
	}
	http.HandleFunc("/excel", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		reader := bytes.NewReader(fileBytes)
		f(reader)
		w.Write([]byte("Hello Http Get!"))
	})
	http.ListenAndServe(":8080", nil)
}

func HttpDownloadExcel() {
	fun := func(fileName string) *bytes.Reader {
		f := excelize.NewFile()
		f.SetCellValue("Sheet1", "B1", 100)
		err := f.SaveAs(fileName)
		if err != nil {
			fmt.Println(err)
		}
		var buffer bytes.Buffer
		_ = f.Write(&buffer)
		return bytes.NewReader(buffer.Bytes())
	}
	http.HandleFunc("/downloadExcel", func(w http.ResponseWriter, r *http.Request) {
		reader := fun("demo.xlsx")
		// 重新设置文件名称
		w.Header().Set("Content-Disposition", "attachment; filename="+"Book2.xlsx")
		io.Copy(w, reader)
	})
	http.ListenAndServe(":8080", nil)
}
