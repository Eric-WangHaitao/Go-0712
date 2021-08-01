package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var done = make(chan int)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
			done <- 1
		})
		s := NewServer(":8080", mux)
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			s.Stop()
		}()
		return s.Start()
	})
	group.Go(func() error {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
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

//http服务
type httpServer struct {
	s       *http.Server
	handler http.Handler
	cxt     context.Context
}

func NewServer(address string, mux http.Handler) *httpServer {
	h := &httpServer{cxt: context.Background()}
	h.s = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	return h
}

func (h *httpServer) Start() error {
	fmt.Println("httpServer start")
	return h.s.ListenAndServe()
}

func (h *httpServer) Stop() error {
	_ = h.s.Shutdown(h.cxt)
	return fmt.Errorf("httpServer stop")
}