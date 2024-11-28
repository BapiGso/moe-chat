# 使用ARG指令接收架构参数
ARG ARCH=amd64

# 根据架构选择对应的Alpine镜像
FROM alpine:${ARCH}

# 移动到工作目录：/app
WORKDIR /app

# 安装必要的依赖
RUN apk add --no-cache bash curl

# 下载并运行安装脚本
RUN bash <(curl -s https://raw.githubusercontent.com/BapiGso/moe-chat/master/shell/install_moe-chat.sh)

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["./app"]