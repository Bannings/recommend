package internal_api

import (
	"context"
	"google.golang.org/grpc"
	"recommend/golbal"
	"recommend/internal_api/proto"
	"recommend/log"
	"recommend/model"
	"sync"
	"time"
)

var (
	middlewareOnce sync.Once
	middlewareRPC  *grpc.ClientConn
	coverOnce      sync.Once
	coverRPC       *grpc.ClientConn
)

func NewGrpcConnect(address string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Errorf("connect fail: %v", err)
		return nil, err
	}
	return conn, nil
}

func GetMiddlewareRPC() *grpc.ClientConn {
	middlewareOnce.Do(func() {
		conf := golbal.GetConfig()
		conn, err := NewGrpcConnect(conf.MiddlewareRPC)
		if err != nil {
			log.Fatal("did not connect: %v", err)
		}
		middlewareRPC = conn
	})

	return middlewareRPC
}

func GetCoverRPC() *grpc.ClientConn {
	coverOnce.Do(func() {
		conf := golbal.GetConfig()
		conn, err := NewGrpcConnect(conf.CoverRPC)
		if err != nil {
			log.Fatal("did not connect: %v", err)
		}
		coverRPC = conn
	})

	return coverRPC
}

func GetCoverRecommend(info model.UserInfo) *proto.CoverDivision {
	conn := GetCoverRPC()
	client := proto.NewDivisionRPCClient(conn)
	var gender proto.DivisionGender
	if info.Gender == 0 {
		gender = proto.DivisionGender_MALE
	} else {
		gender = proto.DivisionGender_FEMALE
	}
	user := proto.User{Uid: info.Uid, Udid: info.UdId, Channel: info.Channel, Version: info.AppVersion, Gender: gender}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	res, err := client.GetCoverRecommend(ctx, &user)
	if err != nil {
		log.Error(err)
		return nil
	}

	if res != nil {
		return res.Data
	} else {
		return nil
	}
}

func GetCoverRecommends(info model.UserInfo) []*proto.MultiCoverDivision {
	conn := GetCoverRPC()
	client := proto.NewDivisionRPCClient(conn)
	var gender proto.DivisionGender
	if info.Gender == 0 {
		gender = proto.DivisionGender_MALE
	} else {
		gender = proto.DivisionGender_FEMALE
	}
	user := proto.User{Uid: info.Uid, Udid: info.UdId, Channel: info.Channel, Version: info.AppVersion, Gender: gender}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	res, err := client.GetMultiCoverRecommend(ctx, &user)
	if err != nil {
		log.Error(err)
		return nil
	}

	if res != nil {
		return res.Data
	} else {
		return nil
	}
}
