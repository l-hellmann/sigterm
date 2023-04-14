package sigterm

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func WrapContext(ctx context.Context) context.Context {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt,
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		<-interrupt
		cancel()
	}()
	return ctx
}


func Context() context.Context {
	return WrapContext(context.Background())
}
