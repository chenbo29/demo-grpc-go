FROM golang:1.18.2
RUN go env -w GOPROXY="https://proxy.golang.com.cn,direct"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN echo 'PATH=$PATH:/go/bin' >> /etc/bash.bashrc
WORKDIR /data