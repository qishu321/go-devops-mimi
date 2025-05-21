FROM alpine

COPY mimi-server /app/server
COPY config /app/config

WORKDIR /app
ENV ENV=test_docker
EXPOSE 8888
RUN echo "https://mirrors.aliyun.com/alpine/v3.18/main" > /etc/apk/repositories
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD cd /app  &&./mimi-server
