rem @echo off
set gopath=C:\Users\TSorriaux\go\
cls
go run -o pts.exe -ldflags "-H windowsgui" db.go html.go main.go
pts.exe