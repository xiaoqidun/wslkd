set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -o wslkd.exe -trimpath -ldflags "-H windowsgui -s -w" wslkd.go
