FROM debian:latest

# 为我们的目录创建文件夹
RUN mkdir /app
# 设置工作目录
WORKDIR /app
ADD vessel-service /app/vessel-service

CMD ["./consignment-service"]
