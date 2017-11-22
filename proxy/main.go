package main

import (
	"context"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			log.Println(ctx.Session)
			ctx.UserData = "hello"
			c := r.Context()
			c = context.WithValue(c, "bar", "foo")
			r = r.WithContext(c)
			return r, nil
		})
	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			r := ctx.Req
			c := r.Context()

			log.Println(ctx.Session)         // session cannot be used
			log.Println(ctx.UserData)        // userdata can be used
			log.Println(r.Header.Get("FOO")) // original request can be used
			log.Println(c.Value("bar"))      // requests context can not be used
			return resp
		})
	http.ListenAndServe(":8080", proxy)
}
