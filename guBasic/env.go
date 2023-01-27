package guBasic

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// * CONST: 디폴트값
const (
	DEFAULT_ROOT_FOLDER = `C:\MoonDev\_config`
	DEFAULT_ROOT_NAME   = "DEV_CFG_ROOT"
)

// TODO: path/filepath 이용
func SetFilePath(names ...string) string {
	path := strings.ReplaceAll(strings.Join(names, "/"), `\`, "/")
	return strings.ReplaceAll(path, "//", "/")
}

// * 환경변수를 먼저 찾고, 없으면 기본값을 사용
func GetRootFolder(names ...string) string {
	root_folder := DEFAULT_ROOT_FOLDER
	root_name := DEFAULT_ROOT_NAME

	if len(names) > 0 {
		root_name = strings.ToUpper(strings.Join(names, "_"))
	}

	if os.Getenv(root_name) != "" {
		root_folder = os.Getenv(root_name)
	}

	return root_folder
}

// * 파일을 읽어서 []byte 형태로 리턴
func BytesFromFile(path string) []byte {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	return buf
}

// * unmarshal
func MapFromFile(buf []byte, configs map[string]interface{}, path string) map[string]interface{} {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(buf, &configs); err != nil {
			return nil
		}
	case ".json":
		if err := json.Unmarshal(buf, &configs); err != nil {
			return nil
		}
	}

	return configs
}

// * org Map에서 key에 해당하는 Sub Map을 리턴
func GetSubMap(org map[string]interface{}, keys ...string) map[string]interface{} {
	for _, key := range keys {
		org = org[key].(map[string]interface{})
	}

	return org
}

// * path에 해당하는 파일을 읽어서 map[string]interface{} 형태로 리턴한다.
func GetConfigMap(path string, keys ...string) (configs map[string]interface{}) {
	return GetSubMap(MapFromFile(BytesFromFile(path), configs, path), keys...)
}

// ** 시스템 환경 변수 추가/삭제
// * 시스템 환경 설정 추가
// ex) key: "DEV_CFG_ROOT", value: `C:\MoonDev\_config`
// TODO: AddSystemEnv 구현
func AddUserEnv(key, value string) {
	cmd := exec.Command("setx", key, value)
	err := cmd.Run()
	log.Println(err)
}
