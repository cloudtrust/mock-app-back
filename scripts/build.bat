@echo off
cd /D %~dp0/../cmd
echo Building mock-app-back...
go build -o ../bin/mockappback.exe
cd /D %~dp0