package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	http.HandleFunc("/", hi5Handler)

	http.ListenAndServe(":80", nil)
}

type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func hi5Handler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	decoded, _ := url.QueryUnescape(string(b))
	for _, s := range strings.Split(decoded, "&") {
		fmt.Println(s)
	}

	slackResponse := SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Sorry, slash commando, that didn't work. Please try again.",
	}

	js, _ := json.Marshal(slackResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// /hi5 @fernando.carvalho teste