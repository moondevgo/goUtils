cd C:\MoonDev\withLang\inGo\goUtils\guDoc
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guDoc
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy