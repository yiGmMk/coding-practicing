package base

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Run("connect", func(t *testing.T) {
		resp, err := http.Get("https://pkg.go.dev/std")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()
	})
	var sites = []string{
		"https://golang.google.cn",
		"https://www.oracle.com/java",
		"https://python.org",
	}
	t.Run("close connect 1", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(len(sites))

		for _, site := range sites {
			site := site
			go func() {
				defer wg.Done()
				req, err := http.NewRequest("GET", site, nil)
				if err != nil {
					fmt.Println(err)
					return
				}
				req.Close = true

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer func() { _ = resp.Body.Close() }()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("get response from %s, resp length = %d\n", site, len(body))
			}()
		}
		wg.Wait()
	})

	t.Run("close connect 2", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(len(sites))

		tr := &http.Transport{
			DisableKeepAlives: true,
		}
		cli := &http.Client{
			Transport: tr,
		}

		for _, site := range sites {
			site := site
			go func() {
				defer wg.Done()

				resp, err := cli.Get(site)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer func() { _ = resp.Body.Close() }()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("get response from %s, resp length = %d\n", site, len(body))
			}()
		}
		wg.Wait()
	})
}
