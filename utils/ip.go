package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCurrentIP(c chan string) {
	resp, err := http.Get("https://ifconfig.io/ip")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	ip := strings.TrimSpace(string(body))

	c <- ip
}
