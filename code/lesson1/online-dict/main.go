package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Sourse    string `json:"sourse"`
	// UserID    string `json:"user_id"`
}

func query(word string) {

	client := &http.Client{}
	var data2 = strings.NewReader(`{"trans_type":"en2zh","source":"hello"}`)
	// {"trans_type":"en2zh","sourse":"good"}
	// var word string = "good"
	request := DictRequest{TransType: "en2zh", Sourse: word}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	var data = bytes.NewReader(buf)
	fmt.Println(data)
	fmt.Println(data2)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data2)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// if resp.Status != "200" {
	// 	log.Fatal("bad Status code", resp.StatusCode, "body", string(bodyText))
	// }
	// fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse

	err = json.Unmarshal(bodyText, &dictResponse)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v\n", dictResponse)
	fmt.Println("good", "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)

	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage:simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}

	word := os.Args[1]
	query(word)
}
