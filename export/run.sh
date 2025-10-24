#!/bin/bash

# 绘本数据导出工具使用示例

echo "=== 绘本数据导出工具 ==="
echo ""

# 检查是否安装了Go
if ! command -v go &> /dev/null; then
    echo "错误: 未安装Go语言环境"
    echo "请先安装Go: https://golang.org/dl/"
    exit 1
fi

# 进入项目目录
cd "$(dirname "$0")"

echo "1. 检查依赖..."
go mod tidy

echo "2. 编译程序..."
go build -o export_tool .

if [ $? -eq 0 ]; then
    echo "编译成功！"
    echo ""
    echo "3. 运行程序..."
    echo "程序将使用 config.json 中的配置"
    echo ""
    ./export_tool
else
    echo "编译失败！"
    exit 1
fi
