FROM alpine

# 移动到工作目录：/app
WORKDIR /app

# 下载release已经编译好的二进制文件
RUN wget https://example.com/release/app-linux-amd64 -O app

# 修改二进制文件权限
RUN chmod +x app

# 声明服务端口
EXPOSE 80

# 启动容器时运行的命令
CMD ["./app"]