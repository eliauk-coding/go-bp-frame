@echo off
setlocal enabledelayedexpansion

SET CURR=%CD%
SET ROOT=%~dp0..\..

if "%1"=="" (
  echo.
  echo �������� -- APP_NAME ����Ϊ��!!
  goto :USAGE
)

if "%2%"=="" (
  echo.
  echo �������� -- MODULE_NAME ����Ϊ��!!
  goto :USAGE
)


SET APP=%1
SET MOD=%2
SET /p GO_MOD=<%ROOT%\go.mod
SET GO_MOD=%GO_MOD:~7%
SET GO_FILE=%MOD%.go
CALL :FILE_NAME GO_FILE
SET PKG_NAME=%MOD%
CALL :LOWER_CASE PKG_NAME

CD /D %ROOT%
SET APP_DIR=apps\%APP%
SET ACT_DIR=%APP_DIR%\actions\%MOD%
SET ACT_FILE=%ACT_DIR%\%GO_FILE%
SET MOD_DIR=%APP_DIR%\models\%MOD%
SET MOD_FILE=%MOD_DIR%\%GO_FILE%
SET SVC_DIR=%APP_DIR%\services\%MOD%
SET SVC_FILE=%SVC_DIR%\%GO_FILE%
SET ROUTER_DIR=%APP_DIR%\router
SET ROUTER_FILE=%ROUTER_DIR%\%GO_FILE%
SET ROUTER_INDX=%ROUTER_DIR%\router.go
SET SVC_NAME=%MOD%Service
CALL :CAMEL_CASE SVC_NAME

SET ROUTER_NAME=%MOD%Router
CALL :CAMEL_CASE ROUTER_NAME

SET MOD_NAME=%MOD%
CALL :CAMEL_CASE_CAP MOD_NAME

if not exist %ACT_DIR% mkdir %ACT_DIR%
if exist %ACT_FILE% (
  echo �Ѵ��ڣ�%ACT_FILE%
  goto :CREATE_MODEL
)
echo package %PKG_NAME%>%ACT_FILE%
echo.>>%ACT_FILE%
echo import (>>%ACT_FILE%
echo     "github.com/gin-gonic/gin">>%ACT_FILE%
echo. >>%ACT_FILE%
echo     %SVC_NAME% "%GO_MOD%/apps/%APP%/services/%MOD%">>%ACT_FILE%
echo     "%GO_MOD%/resp">>%ACT_FILE%
echo     "%GO_MOD%/utils/logger">>%ACT_FILE%
echo )>>%ACT_FILE%
echo.>>%ACT_FILE%
echo type Controller struct{}>>%ACT_FILE%
echo.>>%ACT_FILE%
echo func NewController() *Controller {>>%ACT_FILE%
echo     %SVC_NAME%.AutoMigrate()>>%ACT_FILE%
echo     return ^&Controller{}>>%ACT_FILE%
echo }>>%ACT_FILE%
echo. >>%ACT_FILE%
echo func (ctrl *Controller) List(ctx *gin.Context) {>>%ACT_FILE%
echo     logger.Debug("%PKG_NAME% controller list")>>%ACT_FILE%
echo     resp.Success(ctx, nil)>>%ACT_FILE%
echo }>>%ACT_FILE%

:CREATE_MODEL
if not exist %MOD_DIR% mkdir %MOD_DIR%
if exist %MOD_FILE% (
  echo �Ѵ��ڣ�%MOD_FILE%
  goto :CREATE_SERVICE
)
echo package %PKG_NAME%>%MOD_FILE%
echo.>>%MOD_FILE%
echo type %MOD_NAME% struct {>>%MOD_FILE%
echo 	ID       uint  `gorm:"primaryKey" json:"id"`>>%MOD_FILE%
echo 	UpdateAt int64 `gorm:"autoUpdateTime:milli" json:"utime"`>>%MOD_FILE%
echo 	CreateAt int64 `gorm:"autoCreateTime:milli" json:"ctime"`>>%MOD_FILE%
echo }>>%MOD_FILE%

:CREATE_SERVICE
if not exist %SVC_DIR% mkdir %SVC_DIR%
if exist %SVC_FILE% (
  echo �Ѵ��ڣ�%SVC_FILE%
  goto :CREATE_ROUTER
)
echo package %MOD%>%SVC_FILE%
echo.>>%SVC_FILE%
echo import (>>%SVC_FILE%
echo 	"%GO_MOD%/apps/%APP%/models/%MOD%">>%SVC_FILE%
echo 	"%GO_MOD%/server">>%SVC_FILE%
echo 	"%GO_MOD%/utils/logger">>%SVC_FILE%
echo )>>%SVC_FILE%
echo. >>%SVC_FILE%
echo func AutoMigrate() {>>%SVC_FILE%
echo 	mod := ^&%PKG_NAME%.%MOD_NAME%{}>>%SVC_FILE%
echo 	if err := server.DB().AutoMigrate(mod); err ^^!= nil {>>%SVC_FILE%
echo 		logger.Errorf("%PKG_NAME% service.AutoMigrate failed, %%v", err)>>%SVC_FILE%
echo 	}>>%SVC_FILE%
echo }>>%SVC_FILE%

:CREATE_ROUTER
if not exist %ROUTER_DIR% (
  mkdir %ROUTER_DIR%
)
if not exist %ROUTER_INDX% (
  echo package router>%ROUTER_INDX%
  echo.>>%ROUTER_INDX%
  echo func Public^(^) {>>%ROUTER_INDX%
  echo }>>%ROUTER_INDX%
  echo.>>%ROUTER_INDX%
  echo func Protected^(^) {>>%ROUTER_INDX%
  echo }>>%ROUTER_INDX%
)
if exist %ROUTER_FILE% (
  echo �Ѵ��ڣ�%ROUTER_FILE%
  goto :EOF
)
echo package router>%ROUTER_FILE%
echo.>>%ROUTER_FILE%
echo import ^(>>%ROUTER_FILE%
echo     "%GO_MOD%/apps/%APP%/actions/%PKG_NAME%">>%ROUTER_FILE%
echo     "%GO_MOD%/server">>%ROUTER_FILE%
echo ^)>>%ROUTER_FILE%
echo.>>%ROUTER_FILE%
echo func init%MOD_NAME%Route() {>>%ROUTER_FILE%
echo     r := server.Server().Router()>>%ROUTER_FILE%
echo     ctrl := %PKG_NAME%.NewController()>>%ROUTER_FILE%
echo     r.GET("/", ctrl.List)>>%ROUTER_FILE%
echo }>>%ROUTER_FILE%
CALL :REG_ROUTER %ROUTER_INDX% init%MOD_NAME%Route^(^) "func Protected ^{" "^}"
goto :EOF

:REG_ROUTER
SET "FILE=%1"
SET "CONT=%2"
SET "AFTER=%3"
SET "BEFORE=%4"
SET INSERT_LINE=1

for /F "delims=:" %%a in ('findstr /n %AFTER% "%FILE%"') do (
    SET SKIP="skip=%%a"
    for /F "%SKIP% delims=:" %%b in ('findstr /n %BEFORE% "%FILE%"') do (SET /A INSERT_LINE=%%b)
    SET SKIP=
)
(for /f "tokens=1* delims=:" %%a in ('findstr /n ".*" "%FILE%"') do (
    SET /A n+=1
    if !n! equ !INSERT_LINE! echo     !CONT!
    echo,%%b
))>%FILE%.tmp
MOVE %FILE%.tmp %FILE%>nul 2>nul
SET FILE=
SET CONT=
SET AFTER=
SET BEFORE=
SET INSERT_LINE=

:LOWER_CASE
FOR %%i IN ("A=a" "B=b" "C=c" "D=d" "E=e" "F=f" "G=g" "H=h" "I=i" "J=j" "K=k" "L=l" "M=m" "N=n" "O=o" "P=p" "Q=q" "R=r" "S=s" "T=t" "U=u" "V=v" "W=w" "X=x" "Y=y" "Z=z") DO CALL SET "%1=%%%1:%%~i%%"
goto :EOF

:CAMEL_CASE
for %%i IN ("_a=A" "_b=B" "_c=C" "_d=D" "_e=E" "_f=F" "_g=G" "_h=H" "_i=I" "_j=J" "_k=K" "_l=L" "_m=M" "_n=N" "_o=O" "_p=P" "_q=Q" "_r=R" "_s=S" "_t=T" "_u=U" "_v=V" "_w=W" "_x=X" "_y=Y" "_z=Z") DO CALL SET "%1=%%%1:%%~i%%"
goto :EOF

:CAMEL_CASE_CAP
SET "CH=!%1:~,1!"
for %%b in (A B C D E F G H I J K L M N O P Q R S T U V W X Y Z) do (if /I "!CH!"=="%%b" SET "CH=%%b")
SET "TMP=!CH!!%1:~1!"
for %%i IN ("_a=A" "_b=B" "_c=C" "_d=D" "_e=E" "_f=F" "_g=G" "_h=H" "_i=I" "_j=J" "_k=K" "_l=L" "_m=M" "_n=N" "_o=O" "_p=P" "_q=Q" "_r=R" "_s=S" "_t=T" "_u=U" "_v=V" "_w=W" "_x=X" "_y=Y" "_z=Z") DO CALL SET "TMP=%%TMP:%%~i%%"
SET "%1=%TMP%"
SET CH=
SET TMP=
goto :EOF

:FILE_NAME
SET "%1=!%1:_=.!"
goto :EOF

:USAGE
echo ʹ�÷��� -- scripts\coder\gen APP_NAME MODULE_NAME
echo.
goto :EOF


:EOF
CD %CURR%

SET ROOT=
SET CURR=
SET APP=
SET MOD=
SET GO_MOD=
SET GO_FILE=
SET PKG_NAME=
SET APP_DIR=
SET ACT_DIR=
SET ACT_FILE=
SET MOD_DIR=
SET MOD_FILE=
SET SVC_DIR=
SET SVC_FILE=
SET ROUTER_DIR=
SET ROUTER_FILE=
SET ROUTER_NAME=
SET ROUTER_INDX=
SET SVC_NAME=
SET MOD_NAME=

endlocal
