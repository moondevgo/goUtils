cd C:\MoonDev\withLang\inGo\goUtils\guWin32
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guWin32
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy