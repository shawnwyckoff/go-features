package main

/*
https://www.jianshu.com/p/0dc7596ba90a

什么是context
从go1.7开始，golang.org/x/net/context包正式作为context包进入了标准库。那么，这个包到底是做什么的呢？根据官方的文档说明：

Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.

也就是说，通过context，我们可以方便地对同一个请求所产生地goroutine进行约束管理，可以设定超时、deadline，甚至是取消这个请求相关的所有goroutine。形象地说，假如一个请求过来，需要A去做事情，而A让B去做一些事情，B让C去做一些事情，A、B、C是三个有关联的goroutine，那么问题来了：假如在A、B、C还在处理事情的时候请求被取消了，那么该如何优雅地同时关闭goroutine A、B、C呢？这个时候就轮到context包上场了。
*/

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe(":8989", nil)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	log.Println(A(ctx))
	select {}
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	ctx, _ = context.WithCancel(ctx)
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	}
	return ""
}
