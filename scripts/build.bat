@echo off
cd /D %~dp0/../cmd
go build -o ../bin/mockappback.exe
cd /D %~dp0