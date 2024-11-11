#!/bin/bash

VERSION_INFO="./version.info"

# 读取当前版本号并加1
if [ -f "$VERSION_INFO" ]; then
    VERSION=$(cat $VERSION_INFO)
    VERSION=$((VERSION + 1))
else
    VERSION=1
fi

TARGET_VERSION="0.0.${VERSION}"

# 将新版本号写回文件
echo $VERSION > $VERSION_INFO

# 定义 tar 文件名
TAR_NAME="aliyun_firm_push_v${TARGET_VERSION}.tar.gz"

# 清理旧的 tar.gz 文件和二进制文件
rm ./*.tar.gz 
rm ./bin/*

# 编译 Go 程序
go build -o ./aliyun_firm_push ./

# 打包新版本的 tar.gz 文件
tar -zcvf $TAR_NAME ./aliyun_firm_push ./config

# 将二进制文件移动到 bin 目录
mv ./aliyun_firm_push ./bin

# 输出生成的文件名
echo "Created package: $TAR_NAME"
