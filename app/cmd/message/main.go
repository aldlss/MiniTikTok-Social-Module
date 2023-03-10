package main

import (
	"github.com/aldlss/MiniTikTok-Social-Module/app/cmd/message/dal"
	message "github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/message/messageservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	log "github.com/sirupsen/logrus"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	addr, err := net.ResolveTCPAddr("tcp", "[::]:15677")
	if err != nil {
		log.Fatal(err)
	}
	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-message-grpc",
		}))

	err = svr.Run()

	if err != nil {
		log.Error(err.Error())
	}
}
