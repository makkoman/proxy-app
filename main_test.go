package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"proxy-app/api/handlers"
	"proxy-app/api/server"
	"proxy-app/api/utils"
	"sync"
	"testing"
)

type Response struct {
	Status   int    `json:"status,omitempty"`
	Response string `json:"result,omitempty"`
	ResponseText []ResponseText `json:"res,omitempty"`
}
type ResponseText struct {
	Domain string
}

func TestAlgorithm(t *testing.T) {
	cases := []struct {
		Domain string
		Output string
	}{
		{
			Domain: "alpha", Output: "[\"alpha\"]",
		},
		{
			Domain: "omega", Output: "[\"alpha\",\"omega\"]",
		},
		{
			Domain: "beta", Output: "[\"beta\",\"alpha\",\"omega\"]",
		},
		{
			Domain: "", Output: "error",
		},
	}

	for _, singleCase := range cases {
		client := http.Client{}
		valuesToCompare := &Response{}

		req, _ := http.NewRequest("GET", "http://localhost:8001/ping", nil)
		req.Header.Add("domain", singleCase.Domain)
		response, err := client.Do(req)
		reader, _ := ioutil.ReadAll(response.Body)
		fmt.Println("------------")
		err = json.Unmarshal(reader, valuesToCompare)
		fmt.Println("Response from test", valuesToCompare.Response)
		assert.Nil(t, err)
		assert.Equal(t, singleCase.Output, valuesToCompare.Response)
	}
}

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		utils.LoadEnv()
		app := server.SetUp()
		handlers.HandlerRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("Server Running")
}
