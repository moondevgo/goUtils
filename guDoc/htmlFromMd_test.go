package scrap

import (
	"fmt"
	"github.com/moondevgo/go-mods/basic"
	"io/ioutil"
	// "strings"
	"testing"
)

func TestMdFromHtml(t *testing.T) {
	srcPaths := basic.FileGlob(`C:\Dev\_downloads\chris\html\*.html`)
	dstPath := `C:/Dev/_downloads/chris/md/neovim_chris.md`
	data := mdFileFromLocalFiles(srcPaths)
	fmt.Println(data)
	basic.CreateFile(dstPath)
	ioutil.WriteFile(dstPath, []byte(data), 0644)
}

func TestChangUrlsInMdFile(t *testing.T) {
	srcPath := `C:\Dev\withGit\_docs\git-tutorial.md`
	baseUrl := `https://backlog.com/git-tutorial/kr/`
	basePath := `./`
	dstPath := `C:\Dev\withGit\_docs\git-tutorial_1.md`
	reg := `!\[.+\]\((.+)\)` // NOTE: 이미지 정규표현식
	ConvUrlsInMdFile(srcPath, baseUrl, basePath, dstPath, reg)
}
