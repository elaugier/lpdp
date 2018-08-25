@for /f "usebackq delims=" %%a in (`timestamp.exe -prefix="0."`) do @set VERSION=%%a
@echo version = %VERSION%
@set LDFLAGS=-w -s -X main.Version=%VERSION%
@echo ldflags = %LDFLAGS%
go build ./...
go install ./...
go build -v -ldflags "%LDFLAGS%" -o dist\lpdp-server\lpdp-server.exe cmd\lpdp-server\lpdp-server.go
