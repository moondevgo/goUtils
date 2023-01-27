cd C:\MoonDev\withLang\inGo\goUtils\guCloud
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guCloud
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy