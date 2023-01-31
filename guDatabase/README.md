# go mod

```bash
$ cd ~/dev/inGo/go-mods/database

# mod init
~/dev/inGo/go-mods/database$ go mod init github.com/moondevgo/go-mods/database

# mod replace
~/dev/inGo/go-mods/database$ go mod edit -replace github.com/moondevgo/go-mods/basic=../basic

# mod tidy
~/dev/inGo/go-mods/database$ go mod tidy
```