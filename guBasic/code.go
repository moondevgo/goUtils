package guBasic

import (
	"os"
	"regexp"
	"strings"
)

// * 함수 Body부분 특수문자 변환(백업)
func EncodeFnBody(str string) string {
	str = strings.Replace(str, "\r", "", -1) // ? 필요없을 수도 있으나
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, "\t", `\\t`)
	str = strings.ReplaceAll(str, "\n", `\\n`)
	return str
}

// * 함수 Body부분 특수문자 변환(복원)
func DecodeFnBody(str string) string {
	str = strings.ReplaceAll(str, `\\t`, "\t")
	str = strings.ReplaceAll(str, `\\n`, "\n")
	return str
}

// * 문자열 블록 -> 파일 내에 있는 함수 추출
func FindFns(block string) []map[string]string {
	block = strings.Replace(block, "\r", "", -1)
	rx := `func (.+)\((.*)\) (.*) *?\{\s*\n([\S\n\s]+?)\n\}`
	matches := regexp.MustCompile(rx).FindAllStringSubmatch(block, -1)
	fns := []map[string]string{}
	for _, match := range matches {
		fn := map[string]string{}
		fn["name"] = match[1]
		fn["params"] = match[2]
		fn["rturns"] = match[3]
		fn["body"] = EncodeFnBody(match[4])
		fns = append(fns, fn)
	}
	return fns
}

// * 파일 경로 -> 파일 내에 있는 함수 추출
func ExtractFns(path string) []map[string]string {
	data, _ := os.ReadFile(path)
	return FindFns(string(data))
}

// * tpl 이름 -> tpl로 생성되는 파일 경로(ext: 파일 타입)
func FindPathsInTpl(tplRoot, srcRoot, name, ext string) (tplPath, srcPath string) {
	tplPath = FindFilePathByName(tplRoot, name, "tpl")
	srcPath = strings.Replace(tplPath, tplRoot, srcRoot, 1)
	srcPath = strings.Replace(srcPath, ".tpl", "."+ext, 1)
	return
}

// * template -> file
func GenFileOnTpl(pathTpl, pathDst string, areas, contents []string) {
	_data, _ := os.ReadFile(pathTpl)
	data := string(_data)

	for i, area := range areas {
		data = strings.Replace(data, area, contents[i], 1)
	}

	// ? "\n"가 3개 이상 -> 2개로 변경
	data = regexp.MustCompile(`\n{3,}`).ReplaceAllString(data, "\n\n")

	os.WriteFile(pathDst, []byte(data), 0644)
}

// * maps -> map
func MapFromMaps(maps []map[string]string, fields []string) map[string]string {
	rMap := map[string]string{}
	for _, map_ := range maps {
		rMap[map_[fields[0]]] = map_[fields[1]]
	}
	return rMap
}

// * maps []map[string]string, fields []string -> 템플릿 data 변경
// fields: []string{"key", "val"} sheet field 이름 / 첫번째 문자열: key가 되는 필드이름, 두번째 문자열: value가 되는 필드이름
func FillAreasByMaps(data string, maps []map[string]string, fields []string) string {
	for _, map_ := range maps {
		key := map_[fields[0]]
		area := "{{" + key + "}}"
		data = strings.Replace(data, area, map_[fields[1]], -1)
	}
	return data
}

// * template -> file
func GenFileOnTplByMaps(pathTpl, pathDst string, maps []map[string]string, fields []string) {
	_data, _ := os.ReadFile(pathTpl)
	data := string(_data)

	os.WriteFile(pathDst, []byte(FillAreasByMaps(data, maps, fields)), 0644)
}
