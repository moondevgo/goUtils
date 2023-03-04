cd C:\MoonDev\withLang\inGo\goUtils\guDatabase
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guDatabase
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy