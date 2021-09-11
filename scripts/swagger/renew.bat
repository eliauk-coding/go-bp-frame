@echo off
@echo off
SET CURR=%CD%
SET ROOT=%~dp0..\..

cd /D %ROOT%
type scripts\swagger\head.tpl > docs/api.docs.go
type docs\swagger.json >> docs/api.docs.go
type scripts\swagger\tail.tpl >> docs/api.docs.go
cd /D %CURR%
