package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

//MyResponse is a struct representing a simple code anwer json object
type MyResponse struct {
	Code    string "json:code"
	Message string "json:answer"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request")
		resp := MyResponse{
			Code:    "200",
			Message: "You hit the endpoint: " + html.EscapeString(r.URL.Path),
		}
		mapResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(mapResp)+"\n")
	})

	log.Fatal(http.ListenAndServe(":80", nil))

}
