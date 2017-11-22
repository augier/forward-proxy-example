package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	proxyURL, _ := url.Parse("http://localhost:8080")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
	}
	request, _ := http.NewRequest("GET", "http://localhost:8081", nil)
	request.Header.Set("FOO", "BAR")
	resp, _ := client.Do(request)
	log.Println(resp.StatusCode)
}
