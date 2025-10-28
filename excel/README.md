# Excel读取工具

这是一个用于读取Excel文件并将数据解析到Go结构体中的工具库。

## 功能特性

### 读取功能
- ✅ 支持将Excel数据解析到结构体切片
- ✅ 支持从文件路径、文件流（io.ReadSeeker）和URL读取Excel
- ✅ 支持Map方式读取Excel数据
- ✅ 支持自定义表头行和数据开始行
- ✅ 支持多种数据类型：字符串、整数、浮点数、布尔值、时间等
- ✅ 支持自定义时间格式解析
- ✅ 支持指针类型字段
- ✅ 完整的错误处理

### 导出功能
- ✅ 支持将结构体切片导出为Excel文件
- ✅ 支持自定义列顺序和表头
- ✅ 支持使用excel标签作为表头
- ✅ 支持自定义时间格式

### 测试
- ✅ 丰富的测试用例覆盖所有功能

## 安装依赖

```bash
go get github.com/xuri/excelize/v2
```

## 基本用法

### 1. 定义结构体

```go
type User struct {
    ID       int       `excel:"ID"`       // 用户ID
    Name     string    `excel:"姓名"`      // 用户姓名
    Age      int       `excel:"年龄"`      // 年龄
    Email    string    `excel:"邮箱"`     // 邮箱
    Salary   float64   `excel:"薪资"`     // 薪资
    IsActive bool      `excel:"是否激活"`  // 是否激活
    Birthday time.Time `excel:"生日"`     // 生日
}
```

### 2. 读取Excel到结构体

#### 方式1: 从文件路径读取
```go
var users []User
err := excel.NewExcelReader("data.xlsx").
    SetSheet("用户信息").
    ReadToStruct(&users)

if err != nil {
    log.Fatal(err)
}
```

#### 方式2: 从文件流读取（前端上传）
```go
import "bytes"

// 假设fileBytes是前端上传的Excel文件字节流
fileBytes, _ := os.ReadFile("data.xlsx")  // 实际中来自HTTP请求体
reader := bytes.NewReader(fileBytes)

var users []User
err := excel.NewExcelReaderFromReader(reader).
    SetSheet("用户信息").
    ReadToStruct(&users)

if err != nil {
    log.Fatal(err)
}
```

#### 方式3: 从URL读取
```go
var users []User
err := excel.NewExcelReaderFromURL("http://example.com/data.xlsx").
    SetSheet("用户信息").
    ReadToStruct(&users)

if err != nil {
    log.Fatal(err)
}
```

// 使用数据
for _, user := range users {
    fmt.Printf("用户: %s, 年龄: %d, 薪资: %.2f\n", user.Name, user.Age, user.Salary)
}

### 3. Map方式读取

```go
maps, err := excel.NewExcelReader("data.xlsx").
    SetSheet("用户信息").
    ReadToMap()

if err != nil {
    log.Fatal(err)
}

for _, m := range maps {
    fmt.Printf("ID: %s, 姓名: %s\n", m["ID"], m["姓名"])
}
```

## 高级用法

### 自定义表头和数据行

```go
var products []Product
err := excel.NewExcelReader("data.xlsx").
    SetSheet("产品信息").
    SetHeaderRow(2).     // 表头在第2行
    SetDataStartRow(4).  // 数据从第4行开始
    ReadToStruct(&products)
```

### 支持的数据类型

- `string` - 字符串
- `int`, `int8`, `int16`, `int32`, `int64` - 整数类型
- `uint`, `uint8`, `uint16`, `uint32`, `uint64` - 无符号整数
- `float32`, `float64` - 浮点数
- `bool` - 布尔值
- `time.Time` - 时间类型
- `time.Duration` - 时间间隔
- 指针类型（如 `*string`, `*int` 等）

### 时间格式支持

工具会自动尝试解析以下时间格式：
- `2006-01-02 15:04:05`
- `2006-01-02`
- `2006/01/02 15:04:05`
- `2006/01/02`
- `01/02/2006 15:04:05`
- `01/02/2006`
- `2006-01-02T15:04:05Z`
- `2006-01-02T15:04:05.000Z`
- RFC3339格式
- Unix时间戳（秒和毫秒）

## 结构体标签

使用 `excel` 标签指定Excel列名：

```go
type User struct {
    ID   int    `excel:"ID"`     // 对应Excel中的"ID"列
    Name string `excel:"姓名"`    // 对应Excel中的"姓名"列
    Age  int    `excel:"年龄"`    // 对应Excel中的"年龄"列
}
```

如果没有指定 `excel` 标签，将使用字段名作为列名。

## 错误处理

工具提供详细的错误信息：

```go
err := excel.NewExcelReader("data.xlsx").
    SetSheet("用户信息").
    ReadToStruct(&users)

if err != nil {
    // 处理错误
    log.Printf("读取失败: %v", err)
}
```

常见错误：
- 文件不存在
- Sheet不存在
- 表头行超出范围
- 数据类型转换失败
- 字段设置失败

## 测试

运行测试：

```bash
cd excel
go test -v
```

运行性能测试：

```bash
go test -bench=.
```

## 示例

查看 `example/main.go` 文件了解完整的使用示例。

## 导出功能

### 将结构体导出为Excel

```go
var users []User = []User{
    {ID: 1, Name: "张三", Age: 25, Salary: 5000.5},
    {ID: 2, Name: "李四", Age: 30, Salary: 6000.0},
}

// 基本用法：使用excel标签作为表头
err := excel.NewStructExcelWriter(users).
    UseTagHeaders().
    TimeFormat("2006-01-02").
    SavePath("export.xlsx").
    Sheet("用户信息").
    ToExcel().
    Error

// 自定义列和表头
err := excel.NewStructExcelWriter(products).
    Columns("ID", "Name", "Price").  // 指定导出的字段
    Headers("产品ID", "产品名称", "价格").  // 自定义表头
    SavePath("products.xlsx").
    ToExcel().
    Error

// 支持指针切片
var usersPtr []*User = []*User{
    &User{ID: 1, Name: "张三", Age: 25},
}
err := excel.NewStructExcelWriter(usersPtr).
    UseTagHeaders().
    SavePath("export.xlsx").
    ToExcel().
    Error
```

## 注意事项

1. 确保Excel文件格式正确，表头和数据行清晰
2. 结构体字段类型要与Excel数据类型匹配
3. 时间格式要符合支持的标准格式
4. 大文件读取时注意内存使用
5. 建议在生产环境中添加适当的错误处理和日志记录
6. URL读取需要网络访问权限，建议添加超时设置

## 许可证

MIT License
