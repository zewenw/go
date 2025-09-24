package examples

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {

	t.Run("Error Group", func(t *testing.T) {
		g, ctx := errgroup.WithContext(context.Background())
		g.Go(func() error {
			log.Println("First Go")
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(2 * time.Second):
				log.Println("After wait Go")
			}
			return nil
		})

		g.Go(func() error {
			log.Println("Second Go")
			return errors.New("an error occurred")
		})
		if err := g.Wait(); err != nil {
			fmt.Println("Error:", err)
		}

	})

	t.Run("http client", func(t *testing.T) {
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get("http://localhost:8090/hello")
		if err != nil {
			log.Println("error out:", err)
			panic(err)
		}
		defer resp.Body.Close()

		log.Println("Response Status:", resp.Status)
		scanner := bufio.NewScanner(resp.Body)
		for range 5 {
			scanner.Scan()
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	})
}
