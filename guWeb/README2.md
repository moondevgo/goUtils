# go mod

```bash
$ cd ~/dev/inGo/go-mods/scrap

# mod init
~/dev/inGo/go-mods/scrap$ go mod init github.com/moondevgo/go-mods/scrap

# mod replace
~/dev/inGo/go-mods/scrap$ go mod edit -replace github.com/moondevgo/go-mods/basic=../basic
# go mod edit -replace github.com/moondevgo/go-mods/database=../database


~/dev/inGo/go-mods/scrap$ go mod tidy
```