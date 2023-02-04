# go mod

```bash
$ cd ~/dev/inGo/go-mods/network

# mod init
~/dev/inGo/go-mods/network$ go mod init github.com/moondevgo/go-mods/network

# mod replace
~/dev/inGo/go-mods/network$ go mod edit -replace github.com/moondevgo/go-mods/basic=../basic


~/dev/inGo/go-mods/network$ go mod tidy
```