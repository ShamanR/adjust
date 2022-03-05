package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
)

const defaultWorkersCnt = 10

func main() {
	workersCnt := flag.Int("parallel", defaultWorkersCnt, "parallel workers count")
	flag.Parse()
	// limit workers by number of CPUs
	if *workersCnt > runtime.NumCPU()-1 {
		*workersCnt = runtime.NumCPU() - 1
	}
	// init
	args := flag.Args()
	urls := make(chan string, len(args))
	result := make(chan [2]string, len(args))
	// start async workers
	for i := 0; i < *workersCnt; i++ {
		go worker(urls, result, httpReq)
	}
	// fill in chan
	for _, addr := range args {
		urls <- addr
	}
	close(urls)
	// reading results
	for i := 0; i < len(args); i++ {
		p := <-result
		fmt.Printf("%s %s\n", p[0], p[1])
	}
	close(result)
}

// fan in fan out worker
func worker(in chan string, out chan [2]string, job func(addr string) (result string)) {
	for addr := range in {
		out <- [2]string{addr, job(addr)}
	}
}

// make httpReq to the url, and returns md5 from the result or plain error
func httpReq(address string) (result string) {
	uri, err := url.Parse(address)
	if err != nil {
		return err.Error()
	}
	if uri.Scheme == "" {
		uri.Scheme = "http" // default scheme, can be moved to constant
	}
	resp, err := http.Get(uri.String()) // simple http request, but the best way is to use custom clients, with timeout settings
	if err != nil {
		return err.Error()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return md5Hash(body)
}

// calculated md5 string from input bytes slice
func md5Hash(in []byte) string {
	hash := md5.Sum(in)
	return hex.EncodeToString(hash[:])
}
