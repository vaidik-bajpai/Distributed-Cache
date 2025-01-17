package main

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/vaidik-bajpai/d-cache/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCacheServer(t *testing.T) {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	client := pb.NewCacheClient(conn)

	client.Put(context.TODO(), &pb.PutReq{
		Key:   []byte("key1"),
		Value: []byte("val1"),
	})

	res, _ := client.Get(context.TODO(), &pb.GetReq{Key: []byte("key1")})
	log.Println(string(res.Value))

	assert.Equal(t, []byte("val1"), res.Value, "creating a key value pair and checking its correctness")

	client.Put(context.TODO(), &pb.PutReq{
		Key:   []byte("key1"),
		Value: []byte("val10000000"),
	})

	res, _ = client.Get(context.TODO(), &pb.GetReq{Key: []byte("key1")})
	log.Println(string(res.Value))

	assert.Equal(t, []byte("val10000000"), res.Value, "does put correctly overwrites an already present key")
}
