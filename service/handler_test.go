package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGql(t *testing.T) {
	url := "https://api.thegraph.com/subgraphs/name/volodymyrzolotukhin/sbtidentityverifier-polygon"
	method := "POST"

	payload := strings.NewReader(`{"query":"query {users(where:{senderAddr: \"0x45929D79A6DDdaA3C8154D4F245d17d1D80DbBcc\"}){id}}","variables":{}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("body: %s\n", string(body))
}
