package main

import (
	"os"
	"fmt"
	"time"
	"errors"
	"context"
	"syscall"	
	"os/signal"
	"net/http"
	"golang.org/x/sync/errgroup"
)

var done = make(chan int)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	
	//http server
	svr := http.NewServer()
	group.Go(func() error {
		fmt.Println("http server start！")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done！")
			svr.Shutdown(context.TODO())
		}()
		return svr.Start()
	})

	//Signal
	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			fmt.Println("signal ctx done！")
			return ctx.Err()
		case <-sig:
			fmt.Println("receive signal！")
			return nil	
		}
	})

	err := group.Wait()
	fmt.Println(err)
}