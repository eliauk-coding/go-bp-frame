@echo off
SET CURR=%CD%
SET ROOT=%~dp0..\..

cd /D %ROOT%
swag init --generalInfo main.go --output docs
del /F ".\docs\swagger.yaml"
mv .\docs\docs.go .\docs\api.docs.go
cd %CURR%
