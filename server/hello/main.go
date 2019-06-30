package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Website            string    `json:"websiteurl"`
	ResizeFrom         dimension `json:"resizeFrom"`
	ResizeTo           dimension `json:"resizeTo"`
	CopyAndPaste       bool      `json:"copyAndPaste"`
	FormCompletionTime float64   `json:"formCompletionTime"`
}

type dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func main() {
	fmt.Println("starting...")
	fmt.Println("server listening port is 8080...")
	fmt.Println("waiting resize,copypaste or submit action...")
	http.HandleFunc("/resize", HandleRequest)
	http.HandleFunc("/paste", HandleRequest)
	http.HandleFunc("/submit", HandleRequest)

	http.ListenAndServe(":8080", nil)

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	enableCors(&w)
	if r.Method != "POST" {
		fmt.Println("not a POST method!!!")
		return
	}
	if err != nil {
		fmt.Println("err", err)
	}
	var data data
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(data)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
