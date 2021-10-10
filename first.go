package main

import (
	_ "flag"
	"io"

	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "time"

	_ "log"
)

func main() {

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("VERSION", os.Getenv("VERSION"))

	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

	// client call
	resp1, err1 := http.Get("http://localhost:8080" + "/example")
	if err1 != nil {
		fmt.Println("query cluster failed", err1.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp1.Body)

	body1, err1 := ioutil.ReadAll(resp1.Body)
	fmt.Println(resp1.StatusCode, string(body1))

	resp2, err2 := http.Get("http://localhost:8080" + "/healthz")
	if err2 != nil {
		fmt.Println("query cluster failed", err2.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp2.Body)

	body2, err2 := ioutil.ReadAll(resp2.Body)
	fmt.Println(resp2.StatusCode, string(body2))

}
