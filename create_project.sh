#!/bin/bash

# 项目目录名称
PROJECT_NAME="./"

# 创建项目目录
mkdir -p $PROJECT_NAME

# 创建基本目录结构
mkdir -p $PROJECT_NAME/cmd
mkdir -p $PROJECT_NAME/pkg
mkdir -p $PROJECT_NAME/internal
mkdir -p $PROJECT_NAME/api
mkdir -p $PROJECT_NAME/handlers
mkdir -p $PROJECT_NAME/models
mkdir -p $PROJECT_NAME/utils
mkdir -p $PROJECT_NAME/config
mkdir -p $PROJECT_NAME/docs
mkdir -p $PROJECT_NAME/scripts



# 打印目录结构
echo "创建的项目目录结构如下:"
tree $PROJECT_NAME

echo "Gin项目目录结构创建完毕!"