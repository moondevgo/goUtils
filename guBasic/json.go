package guBasic

import (
	"encoding/json"
	"log"
	"os"
)

// * struct -> json []byte
func JsonBytesFromAny(src interface{}, indent ...string) []byte {
	indent_ := "  "
	if len(indent) > 0 {
		indent_ = indent[0]
	}
	data, _ := json.MarshalIndent(src, "", indent_)
	return data
}

// * any -> json string
func JsonStrFromAny(src interface{}, indent ...string) string {
	return string(JsonBytesFromAny(src, indent...))
}

// * any -> json file
func JsonFileFromAny(path string, src interface{}, indent ...string) {
	// log.Printf("\n%v\n", JsonStrFromAny(src, indent...))
	os.WriteFile(path, JsonBytesFromAny(src, indent...), os.FileMode(0644))
}

// * json []byte -> any
// st: pointer 변수, &변수이름 ex) &st
func AnyFromJsonBytes[T any](data []byte, st *T) {
	err := json.Unmarshal(data, st)
	if err != nil {
		log.Println(err)
	}
}

// * json file -> any
func AnyFromJsonFile[T any](path string, st *T) {
	data, _ := os.ReadFile(path)
	AnyFromJsonBytes(data, st)
}
