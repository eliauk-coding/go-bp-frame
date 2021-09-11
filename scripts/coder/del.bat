@echo off
setlocal enabledelayedexpansion
set CURR=%CD%
set ROOT=%~dp0..\..

if "%1"=="" (
  echo.
  echo 参数错误 -- APP_NAME 不能为空!!
  goto :USAGE
)

if "%2"=="" (
  echo.
  echo 参数错误 -- MODULE_NAME 不能为空!!
  goto :USAGE
)
cd /D %ROOT%

SET APP=%1
set MOD=%2
SET GO_FILE=%MOD%.go
CALL :FILE_NAME GO_FILE

set APP_DIR=apps\%APP%
set ACT_DIR=%APP_DIR%\actions\%MOD%
set ACT_FILE=%ACT_DIR%\%GO_FILE%
set MOD_DIR=%APP_DIR%\models\%MOD%
set MOD_FILE=%MOD_DIR%\%GO_FILE%
set SVC_DIR=%APP_DIR%\services\%MOD%
set SVC_FILE=%SVC_DIR%\%GO_FILE%
set ROUTER_DIR=%APP_DIR%\router
set ROUTER_INDX=%ROUTER_DIR%\router.go
set ROUTER_FILE=%ROUTER_DIR%\%GO_FILE%

set MOD_NAME=%MOD%
call :CAMEL_CASE_CAP MOD_NAME
call :UNREG_ROUTER %ROUTER_INDX% init%MOD_NAME%Route^(^)

echo 将要删除如下目录和文件：
echo   %ACT_DIR%
echo   %MOD_DIR%
echo   %SVC_DIR%
echo   %ROUTER_FILE%
set /p OPT=确认删除？[Y/N] 

if "%OPT%"=="Y" goto :CLEANUP
if "%OPT%"=="y" goto :CLEANUP
goto :EOF

:CLEANUP
if exist %ACT_FILE% del /F %ACT_FILE% 2>nul
if exist %ACT_DIR% rmdir %ACT_DIR% 2>nul
if exist %MOD_FILE% del /F %MOD_FILE% 2>nul
if exist %MOD_DIR% rmdir %MOD_DIR% 2>nul
if exist %SVC_FILE% del /F %SVC_FILE% 2>nul
if exist %SVC_DIR% rmdir %SVC_DIR% 2>nul 
if exist %ROUTER_FILE% del /F %ROUTER_FILE% 2>nul

call :UNREG_ROUTER %ROUTER_INDX% init%MOD_NAME%Route^(^)
goto :EOF

:UNREG_ROUTER
set FILE=%1
set CONT=%2
type %FILE%|find "%CONT%">nul&&(type %FILE%|find /v "%CONT%")>%FILE%.tmp
move /Y %FILE%.tmp %FILE% >nul 2>nul
set FILE=
set CONT=
goto :EOF

:CAMEL_CASE_CAP
set "CH=!%1:~,1!"
for %%b in (A B C D E F G H I J K L M N O P Q R S T U V W X Y Z) do (if /I "!CH!"=="%%b" set "CH=%%b")
set "TMP=!CH!!%1:~1!"
for %%i IN ("_a=A" "_b=B" "_c=C" "_d=D" "_e=E" "_f=F" "_g=G" "_h=H" "_i=I" "_j=J" "_k=K" "_l=L" "_m=M" "_n=N" "_o=O" "_p=P" "_q=Q" "_r=R" "_s=S" "_t=T" "_u=U" "_v=V" "_w=W" "_x=X" "_y=Y" "_z=Z") DO CALL set "TMP=%%TMP:%%~i%%"
set "%1=%TMP%"
set CH=
set TMP=
goto :EOFggGddg

:FILE_NAME
SET "%1=!%1:_=.!"
goto :EOF

:USAGE
echo 使用方法 -- scripts\coder\del APP_NAME MODULE_NAME
echo.

:EOF
cd %CURR%
SET ROOT=
SET CURR=
SET APP=
SET MOD=
SET OPT=
SET GO_FILE=
SET APP_DIR=
SET ACT_DIR=
SET ACT_FILE=
SET MOD_DIR=
SET MOD_FILE=
SET SVC_DIR=
SET SVC_FILE=
SET ROUTER_DIR=
SET ROUTER_FILE=
SET ROUTER_INDX=
endlocal