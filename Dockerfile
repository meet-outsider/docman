FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 设置环境变量
ENV PORT=8888

# 暴露端口
EXPOSE $PORT

# 复制可执行文件到容器
COPY main /app

# 复制配置文件到容器
COPY configs /app/configs

# 运行应用
CMD ["./main","--prod"]