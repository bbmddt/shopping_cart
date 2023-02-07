package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 優雅的停止server，當收到一個 os.Interrupt 或者 syscall.SIGTERM 信號.
func ShutdownGin(instance *http.Server, timeout time.Duration) {

	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("關閉 Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("Server 關閉:", err)
	}
	// 超時5秒 ctx.Done().
	select {
	case <-ctx.Done():
		log.Println("超時5秒.")
	}
	log.Println("Server 退出")
}
