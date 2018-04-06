@echo off
cd /D %~dp0
call build.bat
pause
"../bin/mockappback.exe"
pause