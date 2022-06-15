FROM golang:1.18.2
ENV TZ=Asia/Shanghai
RUN go env -w GOPROXY="https://proxy.golang.com.cn,direct"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN echo 'PATH=$PATH:/go/bin' >> /etc/bash.bashrc
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y git vim unzip zlib1g-dev cmake protobuf-compiler
WORKDIR /data