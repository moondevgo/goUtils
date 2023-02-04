# go mod

```bash
$ cd ~/dev/inGo/go-mods/crawl

# mod init
~/dev/inGo/go-mods/crawl$ go mod init github.com/moondevgo/go-mods/crawl

# mod replace
~/dev/inGo/go-mods/crawl$ go mod edit -replace github.com/moondevgo/go-mods/basic=../basic

# mod tidy
~/dev/inGo/go-mods/crawl$ go mod tidy
```



[go 및 colly를 사용하여 웹 스크레이퍼를 빌드하는 방법 - golang + gocolly 자습서](https://www.youtube.com/watch?v=bfVxq-oQA3c)

[gocolly login](http://go-colly.org/docs/examples/login/)