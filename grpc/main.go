/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	finance "demo/protos/finance"
	helloworld "demo/protos/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
	finance.UnimplementedOrderServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName() + "[demo-grpc-go]"}, nil
}

func (s *server) Generate(ctx context.Context, in *finance.OrderRequest) (*finance.OrderResponse, error) {
	log.Printf("Received Order: type[%v] number[%v]", in.GetType(), in.GetNumber())
	return &finance.OrderResponse{
		Alipay: &finance.Alipay{Url: "aplipay url"},
		Wechat: &finance.Wechat{
			Config: &finance.WechatConfig{AppId: "app id", Noncestr: "nonce str", Signature: "signature", JsapiTicket: "jsapi ticket"},
			Pay:    &finance.WechatPay{AppId: "app id", Noncestr: "nonce str", Sign: "saeodlkejwo", SignType: "RSA", TimeStamp: 123312},
		},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &server{}

	helloworld.RegisterGreeterServer(s, server)
	finance.RegisterOrderServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
