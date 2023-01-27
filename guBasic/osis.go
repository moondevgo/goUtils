package guBasic

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ReadFileBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("\nError: %v\n", err)
	}
	return data
}

func ReadFileStr(path string) string {
	return string(ReadFileBytes(path))
}

// * 라인 단위 읽기
func ReadLines(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		log.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		log.Println(line)
	}

	// log.Println(fileLines)
	return fileLines
}

func GetFolder(path string) string {
	sep := "/"
	if strings.Contains(path, `\`) {
		sep = `\`
	}
	pieces := strings.Split(path, sep)
	return strings.Join(pieces[:len(pieces)-1], sep)
}

func FindFilePathByName(root, name, ext string) string {
	rPath := ""
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Base(path) == name+"."+ext {
			rPath = path
			return nil
		}

		return nil
	})
	return rPath

}

func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// * pattern에 맞는 file 리스트 리턴
// pattern: ex) C:/Dev/withGit/_docs/chris/*.html C:\Dev\withGit\_docs\chris
func FileGlob(pattern string) (paths []string) {
	matches, _ := filepath.Glob(pattern)
	for _, path := range matches {
		paths = append(paths, path)
	}
	return paths
}

// * root에 해당하는 파일들 중에, reg(정규식)에 적합한 파일 리스트 리턴
// root: "C:/eBEST/xingAPI/Res/*"
// reg: `C:[^_]+\d+.res`
func FileGlob2(root, reg string) (paths []string) {
	matches, _ := filepath.Glob(root)
	for _, path := range matches {
		paths = append(paths, path)
	}
	return regexp.MustCompile(reg).FindAllString(strings.Join(paths, " "), -1)
}

// func ExistFile(path string) bool {
// 	_, err := os.Stat(path)
// 	return err == nil
// }

// func Env(key string) string {
// 	return os.Getenv(key)
// }

// func CurrDir() string {
// 	cwd, _ := os.Getwd()
// 	return cwd
// }

// func HomeDir() string {
// 	return Env("USERPROFILE")
// }

// func GOPATH() string {
// 	GOPATH := Env("GOPATH")

// 	if GOPATH == "" {
// 		GOPATH = HomeDir() + `/Go` // Go 1.8부터 생긴 디폴트 GOPATH
// 	}

// 	return GOPATH
// }

// func GOROOT() (GOROOT string) {
// 	if GOROOT = Env("GOROOT"); GOROOT == "" {
// 		if ExistFile(`C:\Go\bin\go.exe`) {
// 			GOROOT = `C:\Go`
// 		} else if ExistFile(`C:\Program Files\Go\bin\go.exe`) {
// 			GOROOT = `C:\Program Files\Go`
// 		} else {
// 			// ! 파일검색 추가
// 			GOROOT = `C:\Program Files\Go`
// 		}
// 	}

// 	return GOROOT
// }
