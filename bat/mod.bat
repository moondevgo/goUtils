cd C:\MoonDev\withLang\inGo\goUtils
REM go clean -modcache
go mod init github.com/moondevgo/goUtils
go mod edit -replace github.com/moondevgo/goUtils/guBasic=./guBasic
go mod edit -replace github.com/moondevgo/goUtils/guWin32=./guWin32
go mod edit -replace github.com/moondevgo/goUtils/guCloud=./guCloud
go mod tidy