package examples

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	t.Run("Log demo", func(t *testing.T) {
		log.Println("standard logger")

		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
		log.Println("with micro")

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("with file/line")

		mylog := log.New(os.Stdout, "my:", log.LstdFlags)
		mylog.Println("from my log")

		mylog.SetPrefix("ohmy:")
		mylog.Println("from mylog")

		var buf bytes.Buffer
		buflog := log.New(&buf, "buf:", log.LstdFlags)

		buflog.Println("hello")

		fmt.Print("from buflog:", buf.String())

		ctx := context.Background()
		jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
		myslog := slog.New(jsonHandler)
		myslog = myslog.With(slog.Any("appid", 123))
		myslog.InfoContext(ctx, "appid")
		myslog.Info("hi there")
		myslog.Info("hello again", "key", "val", "age", 25)
	})
}
