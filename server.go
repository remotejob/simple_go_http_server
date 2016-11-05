package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {

	exist := check_logfile_existense("log.csv")

	if exist {

		log.Println("log file exist")

	}

}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		h.ServeHTTP(rw, req)
	})
}

func loghandler(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(rw, "Hello request")
}

func main() {

	h := middleware(http.HandlerFunc(loghandler))
	log.Fatal(http.ListenAndServe(":8080", h))

}
