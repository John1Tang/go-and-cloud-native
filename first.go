package main

import (
	"flag"
	
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	
	"log"
)

func main() {
	
	
	http.HandleFunc("/example", func (w http.ResponseWriter, r *httpRequest){
		w.Header().Set("VERSION", os.Getenv("VERSION"))
	
	})
	
	
	http.HandleFunc("/healthz", func (w httpResponseWriter, r *http.Request){
		w.WriteHeader(200)
	})
	
	result := http.ListenAndServe(":8080", nil)
	
	// client call
	resp1, er1 := http.Get("http://localhost:8080" + "/example")
	if err != nil {
		fmt.Println("query cluster failed", err1.Error())
		return
	}
	defer resp1.Body.Close()
	
	body1, err1 := ioutil.readAll(resp1.Body)
	fmt.Println(resp1.statusCode, string(body1))
	
	resp2, err2 := http.Get("http://localhost:8080" + "/healthz")
	if err != nil {
		fmt.Println("query cluster failed", err2.Error())
		return
	}
	defer resp2.Body.Close()
	body2, err2 := ioutil.readAll(resp2.Body)
	fmt.Println(resp2.statusCode, string(body2))

}
