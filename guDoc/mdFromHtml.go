package scrap

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/antchfx/htmlquery"
	// "github.com/antchfx/htmlquery"
)

func mdStrFromHtmlStr(str string) string {
	converter := md.NewConverter("", true, nil)
	mdStr, err := converter.ConvertString(str)
	if err != nil {
		log.Fatal(err)
	}

	return mdStr
}

func mdFromHtmlFile(srcPath string) {
	converter := md.NewConverter("", true, nil)
	data, _ := ioutil.ReadFile(srcPath)

	markdown, err := converter.ConvertString(string(data))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(markdown)
}

func mdFileFromLocalFiles(srcPaths []string) string {
	rst := ""
	replacements := [][2]string{
		{"![](/images/neovim.png)\n", ""},
		{"[neovim](/blog/category/neovim)", ""},
		{"\n\n```\n\n", "\n```\n\n"},
		{"\n\n\n", "\n\n"},
	}
	for _, srcPath := range srcPaths {
		doc, _ := htmlquery.LoadDoc(srcPath)
		div := htmlquery.FindOne(doc, "//main/div")

		// // NOTE: markdown
		mdStr := mdStrFromHtmlStr(htmlquery.OutputHTML(div, true))

		// NOTE: replace(DELETE)
		for _, repl := range replacements {
			mdStr = strings.Replace(mdStr, repl[0], repl[1], -1)
		}
		rst += mdStr + "\n----\n\n"
	}
	return rst
}
