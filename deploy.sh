#!/bin/bash

# 设置脚本出错时退出
set -e

# 配置参数
APP_NAME="picbed-logic"  # 替换为你的应用程序名称
BUILD_DIR="build"        # 编译产物的目录
OUTPUT_BINARY="$BUILD_DIR/$APP_NAME" # 输出的二进制文件路径

# 创建编译目录
mkdir -p "$BUILD_DIR"

# 清理之前的构建
if [ -f "$OUTPUT_BINARY" ]; then
    echo "清理旧的二进制文件..."
    rm "$OUTPUT_BINARY"
fi

# 编译 Go 程序
echo "开始编译 Go 程序..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_BINARY" ./

if [ $? -eq 0 ]; then
    echo "编译完成，二进制文件位于 $OUTPUT_BINARY"
else
    echo "编译失败，请检查代码。"
    exit 1
fi

# 压缩二进制文件和配置文件为 tar.gz 格式
PACKAGE_NAME="$APP_NAME-$(date +%Y%m%d).tar.gz"
cd "$BUILD_DIR"

# 复制 config.yaml 到 BUILD_DIR
cp ../config.yaml .

echo "打包文件为 $PACKAGE_NAME..."
tar -czf "$PACKAGE_NAME" "$APP_NAME" "config.yaml"

if [ $? -eq 0 ]; then
    echo "打包完成，文件位于 $BUILD_DIR/$PACKAGE_NAME"
else
    echo "打包失败。"
    exit 1
fi

# 完成
echo "Go 程序打包完成！"
exit 0
