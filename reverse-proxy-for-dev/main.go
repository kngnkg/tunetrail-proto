package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"sync"
)

func runProxyServer(port int, forwardHost string) {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = forwardHost
	}

	rp := &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: func(resp *http.Response) error {
			// キャッシュ関連のヘッダーを追加/変更
			resp.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
			resp.Header.Set("Pragma", "no-cache")
			resp.Header.Set("Expires", "-1")
			return nil
		},
	}

	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: rp,
	}

	if err := server.ListenAndServeTLS("localhost.pem", "localhost-key.pem"); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		runProxyServer(443, "tunetrail-restapi:8080")
		wg.Done()
	}()

	go func() {
		runProxyServer(444, "tunetrail-webapp:3000")
		wg.Done()
	}()

	wg.Wait()
}
