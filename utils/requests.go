package utils

import (
	"io/ioutil"
	"log"
	"namecheap-cli/consts"
	"net/http"
	"net/url"
)

// PostRequest
// @Description: Post to an url and return Response
// @Params: url string, payload map[string]string
// @Return: http.Response

func PostRequest(url string, payload map[string]string) []byte {
	params := ParamsBuilder(payload)
	resp, err := http.PostForm(url, params)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

// ParamsBuilder
// @Description: Build Common Params
// @Params: payload map[string]string
// @Return: url.Values

func ParamsBuilder(payload map[string]string) url.Values {
	// get current ip
	c := make(chan string)
	go GetCurrentIP(c)
	ip := <-c

	params := url.Values{}
	params.Add("ApiUser", consts.ApiUser)
	params.Add("ApiKey", consts.ApiKey)
	params.Add("UserName", consts.UserName)
	params.Add("ClientIp", ip)

	for key, value := range payload {
		// add the payload into the common params
		params.Add(key, value)
	}

	return params
}
