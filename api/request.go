package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"troncy.fr/go-webserver-template/conf"
)

var basePath = "https://BASE.PATH"

func requestGET(url string) string {

	resp, err := http.Get(basePath + url)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	return sb
}

func requestPOST(path string, requestBody string) string {

	config := conf.LoadConfig()

	method := "POST"

	client := &http.Client{}

	var jsonStr = []byte(requestBody)

	req, err := http.NewRequest(method, basePath+path, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("PARAMHEADER1", config.Param1)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println("Status : " + resp.Status)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	return sb
}
