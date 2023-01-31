cd C:\MoonDev\withLang\inGo\goUtils\guTa
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guTa
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy