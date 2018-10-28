@for /f "usebackq delims=" %%a in (`timestamp.exe -prefix="0."`) do @set VERSION=%%a
@echo version = %VERSION%
@set LDFLAGS=-w -s -X main.Version=%VERSION%
@echo ldflags = %LDFLAGS%
rem go clean -r -x -cache
cd pkg
go build ./...
go install ./...
cd ..
go env
go build -v -ldflags "%LDFLAGS%" -o dist\lpdp-server\win-lpdp-server.exe cmd\lpdp-server\main.go
set GOOS=linux
go build -v -ldflags "%LDFLAGS%" -o dist\lpdp-server\lnx-lpdp-server cmd\lpdp-server\main.go
set GOOS=darwin
go build -v -ldflags "%LDFLAGS%" -o dist\lpdp-server\mac-lpdp-server cmd\lpdp-server\main.go
set GOOS=windows 