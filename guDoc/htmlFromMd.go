package scrap

import (
	"io/ioutil"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/moondevgo/go-mods/basic"
)

func htmlStrFromMdStr(mdStr string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	md := []byte(mdStr)
	html := markdown.ToHTML(md, parser, nil)
	return string(html)
}

// Markdown 문자열에 있는 Url 변경
func ConvUrlsInMdStr(strMarkdown, baseUrl, basePath, reg string) string {
	urls := FindSubmatches(strMarkdown, reg)
	for _, url := range urls {
		strMarkdown = strings.Replace(strMarkdown, url, GetDstPath(url, baseUrl, basePath), 1)
	}
	return strMarkdown
}

// Markdown 파일에 있는 Url 변경
func ConvUrlsInMdFile(srcPath, baseUrl, basePath, dstPath, reg string) {
	_data, _ := ioutil.ReadFile(srcPath)
	data := []byte(ConvUrlsInMdStr(string(_data), baseUrl, basePath, reg))
	basic.CreateFile(dstPath) // NOTE: directory가 없으면 생성, TODO: 좀더 효율적인 다른 방법 구현
	ioutil.WriteFile(dstPath, data, 0644)
}
