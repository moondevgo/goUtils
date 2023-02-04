package network

import (
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "net/url"
	// "strings"
	// "time"
	// "bytes"
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"context"

	"github.com/machinebox/graphql"
)

// func HttpStr(method string, uri string, header map[string]string, data string) ([]byte, error) {
// 	var req *http.Request
// 	method = strings.ToUpper(method)
// 	client := http.Client{Timeout: 10 * time.Second}

// 	switch method {
// 	case "POST":
// 		jsonStr := []byte(data)
// 		req, _ = http.NewRequest(method, uri, bytes.NewBuffer(jsonStr))
// 		fmt.Printf("req: %v", req)
// 		// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 		req.Header.Set("Content-Type", "application/json")
// 	}

// 	for k, v := range header {
// 		req.Header.Add(k, v)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("[HttpPost] HTTP POST Error", err, resp.StatusCode)
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	return ioutil.ReadAll(resp.Body)
// }

// func Http(method string, uri string, header map[string]string, data url.Values) ([]byte, error) {
func Http(method string, uri string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var req *http.Request
	method = strings.ToUpper(method)
	client := http.Client{Timeout: 10 * time.Second}

	switch method {
	case "GET":
		if len(data) > 0 {
			params := url.Values{}
			for k, v := range data {
				params.Add(k, v.(string))
			}
			uri += "?" + params.Encode()
		}
		// fmt.Printf("uri: %v\n", uri)
		req, _ = http.NewRequest(method, uri, nil)
	case "POST":
		jsonStr, _ := json.Marshal(data) // jsonStr = []byte(`{"nickname": "mon", "email": "mon@gmail.com"}`)
		req, _ = http.NewRequest(method, uri, bytes.NewBuffer(jsonStr))
		fmt.Printf("req: %v", req)
		// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("[HttpPost] HTTP POST Error", err, resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpBearer(method string, uri string, header map[string]string, data map[string]interface{}, callback func(map[string]interface{}) string) ([]byte, error) {
	header["Authorization"] = "Bearer " + callback(data) // NOTE: JwtFunc 를 콜백 함수, 매개변수로 설정
	uri = strings.Split(uri, "?")[0]
	return Http(method, uri, header, data)
}

// func HttpBearer(method string, uri string, header map[string]string, data map[string]interface{}) ([]byte, error) {
// 	header["Authorization"] = "Bearer " + UpbitJwt(data)
// 	uri = strings.Split(uri, "?")[0]
// 	return Http(method, uri, header, data)
// }

// https://www.thepolyglotdeveloper.com/2020/02/interacting-with-a-graphql-api-with-golang/
func HttpGraphql(uri, query string) {
	// graphqlClient := graphql.NewClient("https://<GRAPHQL_API_HERE>")
	// graphqlRequest := graphql.NewRequest(`
	//     {
	//         people {
	//             firstname,
	//             lastname,
	//             website
	//         }
	//     }
	// `)
	graphqlClient := graphql.NewClient(uri)
	graphqlRequest := graphql.NewRequest(query)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)
}

// func HttpDelete(url string, header map[string]string, data url.Values) ([]byte, error) {
// 	client := http.Client{Timeout: 10 * time.Second}

// 	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(data.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	for k, v := range header {
// 		req.Header.Add(k, v)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("[HttpGet] HTTP DELETE Error", err, resp.StatusCode)
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	return ioutil.ReadAll(resp.Body)
// }
