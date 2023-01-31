cd C:\MoonDev\withLang\inGo\goUtils\guPlot
REM go clean -modcache
go mod init github.com/moondevgo/goUtils/guPlot
go mod edit -replace github.com/moondevgo/goUtils/guBasic=../guBasic
go mod tidy