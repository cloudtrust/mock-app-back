@echo off
cd /D %~dp0/../bin/
echo Running mock-app-back...
mockappback.exe
cd /D %~dp0