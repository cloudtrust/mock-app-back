@echo off
cd /D %~dp0
call build.bat
pause
call run.bat
pause