cd C:\MoonDev\withLang\inGo\goUtils\guWeb
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guWeb
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy