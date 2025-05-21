#!/bin/bash
set -e  # 一旦出错立即退出

echo -e "\033[1;34m[1/5] 构建后端 Go 项目...\033[0m"
if [ -d "../server" ]; then
  cd ../server
  go env -w GOPROXY=https://goproxy.cn,direct
  go mod tidy
  CGO_ENABLED=0 go build -o mimi-server
  cd - >/dev/null
else
  echo -e "\033[0;31m[错误] ../server 目录不存在，构建失败！\033[0m"
  exit 1
fi

echo -e "\033[1;34m[2/5] 构建并启动 MySQL 服务...\033[0m"
docker-compose -f docker-compose-mysql.yaml up -d --build

echo -e "\033[1;34m[3/5] 构建并启动 mimi 服务...\033[0m"
docker-compose -f docker-compose-mimi.yaml up -d --build

echo -e "\033[1;34m[4/5] 容器状态检查...\033[0m"
docker-compose -f docker-compose-mimi.yaml ps

echo -e "\033[1;34m[5/5] 重启异常退出的容器（如有）...\033[0m"
EXITED_CONTAINERS=$(docker-compose -f docker-compose-mimi.yaml ps -q | xargs docker inspect --format '{{.Name}} {{.State.Status}}' | grep "exited" | awk '{print $1}' | sed 's/\///')
if [ -n "$EXITED_CONTAINERS" ]; then
  for name in $EXITED_CONTAINERS; do
    echo -e "\033[0;33m重启 $name ...\033[0m"
    docker-compose -f docker-compose-mimi.yaml restart "$name"
  done
else
  echo -e "\033[0;32m所有容器运行正常，无需重启。\033[0m"
fi

echo -e "\033[1;32m✅ 部署完成！\033[0m"
