# 绘本数据导出工具

这个工具可以从绘本API获取数据，并将数据导出到Excel文件中，同时下载相关的图片和音频文件。

## 功能特性

- 🔄 循环调用API获取多个绘本数据
- 📊 将数据导出到Excel文件，每个绘本对应一个工作表
- 🖼️ **图片直接插入Excel**：下载的图片直接嵌入到Excel单元格中
- 🎵 **音频文件超链接**：音频文件创建超链接，点击即可播放
- 📁 自动下载图片和音频文件到本地
- ⚙️ 支持配置文件自定义参数
- 🛡️ 错误处理和重试机制

## 文件结构

```
export/
├── main.go              # 主程序入口
├── models.go            # 数据结构定义
├── api_client.go        # API客户端
├── file_downloader.go   # 文件下载器
├── excel_exporter.go    # Excel导出器
├── config.json          # 配置文件示例
└── README.md           # 说明文档
```

## 使用方法

### 1. 配置参数

编辑 `config.json` 文件：

```json
{
  "book_ids": [465442971693328, 123456789012345],
  "output_dir": "./output",
  "excel_file": "./output/绘本数据导出.xlsx",
  "api_url": "https://sea.pri.ibanyu.com/picturebookopapi/ugc/picturebook/boutique/book/page/get",
  "cookie": "your_cookie_here"
}
```

参数说明：
- `book_ids`: 要导出的绘本ID列表
- `output_dir`: 输出目录
- `excel_file`: Excel文件路径
- `api_url`: API接口地址
- `cookie`: 请求Cookie（从浏览器开发者工具中复制）

### 2. 运行程序

```bash
cd export
go run .
```

### 3. 输出结果

程序运行后会在指定目录生成：

```
output/
├── 绘本数据导出.xlsx                    # Excel文件
├── The_Lunch_Box/                      # 绘本文件夹
│   ├── 图片/
│   │   ├── 封面.jpg                    # 绘本封面
│   │   ├── 1.jpg                       # 第1页图片
│   │   ├── 2.jpg                       # 第2页图片
│   │   └── ...
│   └── 音频/
│       ├── 1.mp3                       # 第1页音频
│       ├── 2.mp3                       # 第2页音频
│       └── ...
└── ...
```

## Excel文件结构

每个绘本对应Excel中的一个工作表，包含以下内容：

| 列名 | 说明 | 数据来源 | 显示方式 |
|------|------|----------|----------|
| 页码 | 页面序号 | pageinfos[].index | 数字 |
| 绘本内容 | 背景图片 | pageinfos[].bg_picture | **直接插入图片** |
| 英文字幕原文 | 英文文本 | pageinfos[].listentext | 文本 |
| 中文翻译原文 | 中文翻译 | pageinfos[].translation | 文本 |
| 英文字母音频 | 音频文件 | pageinfos[].listenaudiourl | **超链接** |

### 特殊功能说明

- **图片插入**：下载的图片直接嵌入到Excel单元格中，转发后其他人可以直接看到图片
- **音频超链接**：音频文件创建超链接，点击单元格可以播放对应的音频文件
- **自动调整**：Excel会自动调整行高以适应插入的图片

## 依赖包

```bash
go mod tidy
```

主要依赖：
- `github.com/xuri/excelize/v2` - Excel文件操作
- `net/http` - HTTP请求
- `encoding/json` - JSON处理

## 注意事项

1. **Cookie有效性**: 请确保Cookie有效，如果失效需要重新获取
2. **网络连接**: 确保网络连接正常，能够访问API和下载文件
3. **存储空间**: 确保有足够的磁盘空间存储下载的文件
4. **文件名**: 程序会自动清理文件名中的特殊字符
5. **请求频率**: 程序会在请求间添加延迟，避免过于频繁的请求

## 错误处理

程序包含完善的错误处理机制：
- API请求失败会记录错误并继续处理下一个绘本
- 文件下载失败会记录错误但不影响整体流程
- Excel导出失败会记录错误信息

## 自定义配置

可以通过修改 `main.go` 中的配置结构来自定义更多参数：

```go
type Config struct {
    BookIDs    []int64 `json:"book_ids"`
    OutputDir  string  `json:"output_dir"`
    ExcelFile  string  `json:"excel_file"`
    APIURL     string  `json:"api_url"`
    Cookie     string  `json:"cookie"`
    // 可以添加更多配置项
    Timeout    int     `json:"timeout"`    // 请求超时时间
    RetryCount int     `json:"retry_count"` // 重试次数
}
```
