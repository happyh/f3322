package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var domain, password string
	flag.StringVar(&domain, "d", "", "domain")
	flag.StringVar(&password, "p", "", "password")
	flag.Parse()

	if domain == "" || password == "" {
		flag.Usage()
		return
	}

	client := &http.Client{}

	url := "http://members.3322.net/dyndns/update?system=dyndns&hostname=" + domain

	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	auth := "root:" + password

	reqest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	resp, err := client.Do(reqest)
	if err == nil && resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println("OK:", string(bodyBytes))
		} else {
			fmt.Println("err:", err)
		}
	} else if err != nil {
		fmt.Println("err:", err)
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println("resp:", resp.StatusCode, "body:", string(bodyBytes))
		} else {
			fmt.Println("err:", err)
		}
	}

}
