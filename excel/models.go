package excel

import "time"

// User 用户信息结构体
type User struct {
	ID       int       `excel:"ID"`   // 用户ID
	Name     string    `excel:"姓名"`   // 用户姓名
	Age      int       `excel:"年龄"`   // 年龄
	Email    string    `excel:"邮箱"`   // 邮箱
	Phone    string    `excel:"电话"`   // 电话
	Salary   float64   `excel:"薪资"`   // 薪资
	IsActive bool      `excel:"是否激活"` // 是否激活
	Birthday time.Time `excel:"生日"`   // 生日
	CreateAt time.Time `excel:"创建时间"` // 创建时间
}

// Product 产品信息结构体
type Product struct {
	ID          int     `excel:"产品ID"`
	Name        string  `excel:"产品名称"`
	Category    string  `excel:"分类"`
	Price       float64 `excel:"价格"`
	Stock       int     `excel:"库存"`
	Description string  `excel:"描述"`
	IsAvailable bool    `excel:"是否可用"`
}

// Order 订单信息结构体
type Order struct {
	OrderID     string    `excel:"订单号"`
	UserID      int       `excel:"用户ID"`
	ProductID   int       `excel:"产品ID"`
	Quantity    int       `excel:"数量"`
	TotalAmount float64   `excel:"总金额"`
	Status      string    `excel:"状态"`
	OrderTime   time.Time `excel:"下单时间"`
	Remark      string    `excel:"备注"`
}

// Employee 员工信息结构体（用于测试复杂字段）
type Employee struct {
	ID           int       `excel:"员工ID"`
	Name         string    `excel:"姓名"`
	Department   string    `excel:"部门"`
	Position     string    `excel:"职位"`
	Salary       float64   `excel:"薪资"`
	HireDate     time.Time `excel:"入职日期"`
	IsManager    bool      `excel:"是否经理"`
	PhoneNumber  string    `excel:"电话号码"`
	EmailAddress string    `excel:"邮箱地址"`
}

// SimpleStruct 简单结构体（用于测试基本类型）
type SimpleStruct struct {
	ID    int    `excel:"ID"`
	Name  string `excel:"名称"`
	Value int    `excel:"值"`
}
