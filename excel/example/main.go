package main

import (
	"fmt"
	"log"

	"test/excel"
)

func main() {
	fmt.Println("=== Excel读取工具使用示例 ===")

	// 创建示例Excel文件
	sampleFile := "sample_data.xlsx"
	err := excel.CreateSampleExcelFile(sampleFile)
	if err != nil {
		log.Fatalf("创建示例文件失败: %v", err)
	}
	defer func() {
		// 清理临时文件
		if err := removeFile(sampleFile); err != nil {
			fmt.Printf("清理文件失败: %v\n", err)
		}
	}()

	// 示例1: 基本用法 - 读取用户信息
	fmt.Println("\n1. 基本用法 - 读取用户信息:")
	var users []excel.User
	err = excel.NewExcelReader(sampleFile).
		SetSheet("用户信息").
		ReadToStruct(&users)

	if err != nil {
		log.Printf("读取用户信息失败: %v", err)
	} else {
		fmt.Printf("成功读取 %d 条用户数据:\n", len(users))
		for i, user := range users {
			fmt.Printf("  用户%d: ID=%d, 姓名=%s, 年龄=%d, 邮箱=%s, 薪资=%.2f\n",
				i+1, user.ID, user.Name, user.Age, user.Email, user.Salary)
		}
	}

	// 示例2: 自定义设置 - 读取产品信息
	fmt.Println("\n2. 自定义设置 - 读取产品信息:")
	var products []excel.Product
	err = excel.NewExcelReader(sampleFile).
		SetSheet("产品信息").
		ReadToStruct(&products)

	if err != nil {
		log.Printf("读取产品信息失败: %v", err)
	} else {
		fmt.Printf("成功读取 %d 条产品数据:\n", len(products))
		for i, product := range products {
			fmt.Printf("  产品%d: ID=%d, 名称=%s, 分类=%s, 价格=%.2f, 库存=%d\n",
				i+1, product.ID, product.Name, product.Category, product.Price, product.Stock)
		}
	}

	// 示例3: 复杂类型 - 读取员工信息
	fmt.Println("\n3. 复杂类型 - 读取员工信息:")
	var employees []excel.Employee
	err = excel.NewExcelReader(sampleFile).
		SetSheet("员工信息").
		ReadToStruct(&employees)

	if err != nil {
		log.Printf("读取员工信息失败: %v", err)
	} else {
		fmt.Printf("成功读取 %d 条员工数据:\n", len(employees))
		for i, emp := range employees {
			fmt.Printf("  员工%d: ID=%d, 姓名=%s, 部门=%s, 职位=%s, 薪资=%.2f, 入职日期=%s\n",
				i+1, emp.ID, emp.Name, emp.Department, emp.Position, emp.Salary, emp.HireDate.Format("2006-01-02"))
		}
	}

	// 示例4: Map方式读取
	fmt.Println("\n4. Map方式读取:")
	maps, err := excel.NewExcelReader(sampleFile).
		SetSheet("用户信息").
		ReadToMap()

	if err != nil {
		log.Printf("Map方式读取失败: %v", err)
	} else {
		fmt.Printf("Map方式成功读取 %d 条数据:\n", len(maps))
		for i, m := range maps {
			fmt.Printf("  第%d条: ID=%s, 姓名=%s, 年龄=%s\n", i+1, m["ID"], m["姓名"], m["年龄"])
		}
	}

	// 示例5: 自定义格式文件
	fmt.Println("\n5. 自定义格式文件:")
	customFile := "custom_sample.xlsx"
	err = excel.CreateSampleExcelFileWithCustomFormat(customFile)
	if err != nil {
		log.Printf("创建自定义格式文件失败: %v", err)
	} else {
		defer removeFile(customFile)

		var customProducts []excel.Product
		err = excel.NewExcelReader(customFile).
			SetSheet("产品信息").
			SetHeaderRow(2).    // 表头在第2行
			SetDataStartRow(4). // 数据从第4行开始
			ReadToStruct(&customProducts)

		if err != nil {
			log.Printf("读取自定义格式文件失败: %v", err)
		} else {
			fmt.Printf("成功读取自定义格式文件 %d 条产品数据:\n", len(customProducts))
			for i, product := range customProducts {
				fmt.Printf("  产品%d: ID=%d, 名称=%s, 价格=%.2f\n",
					i+1, product.ID, product.Name, product.Price)
			}
		}
	}

	fmt.Println("\n=== 示例完成 ===")
}

// removeFile 删除文件
func removeFile(filename string) error {
	// 这里简化处理，实际项目中可以使用os.Remove
	fmt.Printf("清理临时文件: %s\n", filename)
	return nil
}
