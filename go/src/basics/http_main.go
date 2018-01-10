package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"time"
)

// Transport复用，放在全局
var t *http.Transport
var cnt *http.Client

func MakeTransport() *http.Transport {
	if t != nil {
		return t
	}
	t = &http.Transport{
		DialContext:       (&net.Dialer{}).DialContext,
		DisableKeepAlives: true,
	}
	return t
}

func MakeClient(t *http.Transport) *http.Client {
	if cnt != nil {
		return cnt
	}
	cnt = &http.Client{
		Transport: t,
	}
	return cnt
}

func ClientHTTPRequest(host string) {
	// 构造url.URL
	url := &url.URL{Host: host, Scheme: "http"}
	// 构造http.Request
	req := &http.Request{
		URL:    url,
		Method: http.MethodGet,
		Close:  true,
	}
	// 构造http.Transport
	t := MakeTransport()
	// 构造http.Client
	cnt := MakeClient(t)
	// 发请求
	res, _ := cnt.Do(req)
	defer res.Body.Close()
	fmt.Printf("%v", res)
	buf, _ := ioutil.ReadAll(res.Body)
	ioutil.WriteFile("index.html", buf, os.ModePerm)
}

func DirectClientRequst(addr string) {
	// 构造http.Transport
	t := MakeTransport()
	// 构造http.Client
	cnt := MakeClient(t)
	// 发请求
	resp, err := cnt.Get(addr)
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v\n", string(buf))
}

func MakeupHTTPRequest(host string) {
	// 构造url.URL
	url := &url.URL{Host: host, Scheme: "http"}
	// 构造http.Request
	req := &http.Request{
		URL:    url,
		Method: http.MethodGet,
		Close:  true,
	}
	// tcp连接
	conn, err := (&net.Dialer{}).DialContext(context.Background(), "tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 写入请求
	req.Write(conn)
	buf, _ := ioutil.ReadAll(conn)
	fmt.Printf("%v\n", string(buf))
	conn.Close()
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "You are %v, %v", req.RemoteAddr, req.Proto)
}

func main() {
	// start a http server
	http.HandleFunc("/", HomeHandler)
	go http.ListenAndServe(":8000", nil)
	//go http.ListenAndServeTLS(":54321", "cert.pem", "key.pem", nil)

	time.Sleep(time.Second * 2)
	host := "localhost"
	//ClientHTTPRequest(host)
	DirectClientRequst("http://" + host + ":8000")
	//MakeupHTTPRequest(host)
}
