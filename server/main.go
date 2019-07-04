package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var datas map[string]*Data

// Data result struct
type Data struct {
	WebsiteURL         string          `json:",omitempty"`
	SessionID          string          `json:",omitempty"`
	ResizeFrom         Dimension       `json:",omitempty"`
	ResizeTo           Dimension       `json:",omitempty"`
	CopyAndPaste       map[string]bool `json:",omitempty"`
	FormCompletionTime int             `json:",omitempty"`
}

// Dimension of window
type Dimension struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
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
	datas = make(map[string]*Data)
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
		fmt.Printf("There was an error during reading the body: %s", err)
	}
	var requestBody RequestBody
	err = json.Unmarshal([]byte(body), &requestBody)
	if err != nil {
		fmt.Printf("Error during unmarshalling the body: %s", err)
	}

	_, exist := datas[requestBody.SessionID]
	if !exist {
		datas[requestBody.SessionID] = &Data{}
		datas[requestBody.SessionID].CopyAndPaste = make(map[string]bool)
	}

	switch requestBody.EventType {
	case "timeTaken":
		datas[requestBody.SessionID].submit(requestBody)
	case "copyAndPaste":
		datas[requestBody.SessionID].copyAndPaste(requestBody)
	case "resize":
		datas[requestBody.SessionID].resize(requestBody)
	}

	//fmt.Println(*datas[requestBody.SessionID])
	fmt.Printf("%+v\n", *datas[requestBody.SessionID])
}

func (data *Data) submit(requestBody RequestBody) {
	if requestBody.Time != 0 {
		data.FormCompletionTime = requestBody.Time
	}
	if requestBody.WebsiteURL != "" {
		data.WebsiteURL = requestBody.WebsiteURL
	}
	if requestBody.SessionID != "" {
		data.SessionID = requestBody.SessionID
	}
	hashURL(data.WebsiteURL)
}

func (data *Data) copyAndPaste(requestBody RequestBody) {
	if requestBody.WebsiteURL != "" {
		data.WebsiteURL = requestBody.WebsiteURL
	}
	if requestBody.SessionID != "" {
		data.SessionID = requestBody.SessionID
	}
	data.CopyAndPaste[requestBody.FormID] = requestBody.Pasted

}

func (data *Data) resize(requestBody RequestBody) {
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

func hashURL(url string) {
	hash := 0
	for i := 0; i < len(url); i++ {
		hash += int(url[i])
	}
	fmt.Println("URL hash:", hash*len(url)*31)
}
