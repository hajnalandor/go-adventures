package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var data Data

// Data result struct
type Data struct {
	WebsiteURL         string
	SessionID          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             // Seconds
}

// Dimension of window
type Dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

// RequestBody struct
type RequestBody struct {
	EventType  string    `json:"eventType"`
	WebsiteURL string    `json:"websiteUrl"`
	SessionID  string    `json:"sessionId"`
	Pasted     bool      `json:"pasted"`
	FormID     string    `json:"formId"`
	ResizeFrom Dimension `json:"resizeFrom"`
	ResizeTo   Dimension `json:"resizeTo"`
	Time       int       `json:"time"`
}

func main() {
	data.CopyAndPaste = make(map[string]bool)
	fmt.Println("starting...")
	fmt.Println("server listening port is 8080...")
	fmt.Println("waiting resize,copypaste or submit action...")
	http.HandleFunc("/send", HandleRequest)
	http.ListenAndServe(":8080", nil)

}

// HandleRequest ...
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
	var requestBody RequestBody
	err = json.Unmarshal([]byte(body), &requestBody)
	if err != nil {
		fmt.Println("err", err)
	}
	switch requestBody.EventType {
	case "timeTaken":
		submit(requestBody)
	case "copyAndPaste":
		copyAndPaste(requestBody)
	case "resize":
		resize(requestBody)
	}

	fmt.Println(data)
}

func submit(requestBody RequestBody) {
	if requestBody.Time != 0 {
		data.FormCompletionTime = requestBody.Time
	}
	if requestBody.WebsiteURL != "" {
		data.WebsiteURL = requestBody.WebsiteURL
	}
	if requestBody.SessionID != "" {
		data.SessionID = requestBody.SessionID
	}
}

func copyAndPaste(requestBody RequestBody) {
	if requestBody.WebsiteURL != "" {
		data.WebsiteURL = requestBody.WebsiteURL
	}
	if requestBody.SessionID != "" {
		data.SessionID = requestBody.SessionID
	}

	data.CopyAndPaste[requestBody.FormID] = requestBody.Pasted

}

func resize(requestBody RequestBody) {
	if requestBody.WebsiteURL != "" {
		data.WebsiteURL = requestBody.WebsiteURL
	}
	if requestBody.SessionID != "" {
		data.SessionID = requestBody.SessionID
	}
	data.ResizeFrom = requestBody.ResizeFrom
	data.ResizeTo = requestBody.ResizeTo

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
