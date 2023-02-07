package main

import (
	"github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/dal"
	relation "github.com/aldlss/MiniTikTok-Relation/app/kitex_gen/pb/relation/relationservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	log "github.com/sirupsen/logrus"
	"net"
)

func init() {
	dal.Init()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:19198")
	if err != nil {
		log.Error(err)
	}
	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "minitiktok-relation-grpc",
		}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}