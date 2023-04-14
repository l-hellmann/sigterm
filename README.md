# Sigterm

```shell
$ go get github.com/l-hellmann/sigterm
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/l-hellmann/sigterm"
)

func main() {
	ctx := sigterm.Context()
	s := &http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello world")
		}),
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	<-ctx.Done()
	_ = s.Shutdown(context.Background())
}
```