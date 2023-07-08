package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lamhieo02/socialnetapp/cache"
	"github.com/lamhieo02/socialnetapp/config"
	authenandposthandler "github.com/lamhieo02/socialnetapp/internal/authen_and_post/handler"
	authenandpostservice "github.com/lamhieo02/socialnetapp/internal/authen_and_post/service"
	authenandpoststorage "github.com/lamhieo02/socialnetapp/internal/authen_and_post/storage"
	mysqlpkg "github.com/lamhieo02/socialnetapp/pkg/mysql"
	authen_and_post "github.com/lamhieo02/socialnetapp/pkg/proto/authen_and_post/pkg/proto"
	redispkg "github.com/lamhieo02/socialnetapp/pkg/redis"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Get config error", err)
		return
	}

	db, err := mysqlpkg.NewMYSQLConnection(&conf.Mysql)
	if err != nil {
		log.Fatalln("Connect to mysql error", err)
		return
	}

	redis, err := redispkg.NewRedisConnection(&conf.Redis)
	if err != nil {
		log.Fatalln("Connect to redis error", err)
		return
	}
	cache := cache.NewCache(redis)

	authenAndPostStorage := authenandpoststorage.NewAuthenticateAndPostStorage(db)
	authenAndPostService := authenandpostservice.NewAuthenticateAndPostService(authenAndPostStorage, cache)
	authenAndPostHandler := authenandposthandler.NewAuthenticateAndPostHandler(authenAndPostService)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.AuthenAndPostService.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	authen_and_post.RegisterAuthenticateAndPostServer(grpcServer, authenAndPostHandler)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}

}
