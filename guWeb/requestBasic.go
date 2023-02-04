package guWeb

import (
	"encoding/json"
	"io"
	"net/http"
	// basic "github.com/moondevgo/goUtils/guBasic"
)

// GET 호출
func RequestGet(url string) string {
	return string(RequestGetByte(url))
}

func RequestGetByte(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return data
}

func RequestGetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
