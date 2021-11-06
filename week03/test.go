package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"golang.org/x/sync/errgroup"
)

// 启动 HTTP server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

// 句柄操作
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	// 创建上下文，并定义取消操作
	ctx, cancel := context.WithCancel(context.Background())

	// 生成一个新的Group和派生上下文
	group, errCtx := errgroup.WithContext(ctx)

	// 启动http服务
	srv := &http.Server{Addr: ":8081"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-chanel:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}

