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
	"sync"
)

type DictResponseYouDao struct {
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	ErrorCode   int    `json:"errorCode"`
	Type        string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
}

type DictResponseCaiYun struct {
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

func youdao(word string) {
	client := &http.Client{}
	var letter string = "i=" + word + "&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=16519273859772&sign=b95105716cedfaaf0064d5cff6acc59c&lts=1651927385977&bv=247811f9b7fd387f154bf67d8ebd44f3&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME"
	var data = strings.NewReader(letter)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", `JSESSIONID=abcnM6o9rE-uR9BDl673x; _ntes_nnid=52020dba167a0694382d031afeb3e658,1644836233031; OUTFOX_SEARCH_USER_ID_NCOO=1692709003.6024466; OUTFOX_SEARCH_USER_ID="-1259618769@10.110.96.153"; fanyi-ad-id=305838; fanyi-ad-closed=1; ___rl__test__cookies=1651927385972`)
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
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
	// fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponseYouDao

	err = json.Unmarshal(bodyText, &dictResponse)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v\n", dictResponse)
	fmt.Println(word)

	for _, item := range dictResponse.SmartResult.Entries {
		fmt.Println(item)
	}

}

func queryCaiYun(word string) {

	client := &http.Client{}
	// var data2 = strings.NewReader(`{"trans_type":"en2zh","source":"hello"}`)
	// {"trans_type":"en2zh","sourse":"good"}
	// var word string = "good"
	request := DictRequest{TransType: "en2zh", Sourse: word}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	var data = bytes.NewReader(buf)
	// fmt.Println(data)
	// fmt.Println(data2)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
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
	var dictResponse DictResponseCaiYun

	err = json.Unmarshal(bodyText, &dictResponse)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v\n", dictResponse)
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)

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
	var a sync.WaitGroup
	a.Add(2)
	go func() {
		fmt.Println("CAIYUN")
		queryCaiYun(word)
		a.Done()
	}()
	go func() {
		fmt.Println("YOUDAO")
		youdao(word)
		a.Done()
	}()
	a.Wait()
}
