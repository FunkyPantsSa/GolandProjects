package main

import (
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem") // 创建图书数据存储模块实例
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s) // 创建http服务实例

	errChan, err := srv.ListenAndServe() // 运行http服务
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select { // 监视来自errChan以及c的事件
	case err = <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Srv.Shutdown(ctx) // 优雅关闭http服务实例
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}
	log.Println("bookstore program exit ok")
}
