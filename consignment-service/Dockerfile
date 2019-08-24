# builder是我们对Golang镜像的命名
FROM golang:latest
# c创建根目录
RUN mkdir /app
# 设置工作目录
WORKDIR /app
# 将代码复制到工作目录中
COPY . /app/consignment-service
# 设置环境变量
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.io && export GOROOT=/usr/local/go && export PATH=$PATH:/usr/local/go/bin
RUN cd /app && go mod init github.com/yuwe1/learnmico && go mod tidy
# 编译
RUN cd /app/consignment-service && go build


# 在这里，第二次使用FROM，他告诉Docker，开始构建第二个流程
FROM Debian 

RUN mdkir /serapp && cd /serapp
WORKDIR /cliapp
COPY --from=golang:latest /app/consignment-service/consignment-service .

CMD [ "./consignment-service" ]


