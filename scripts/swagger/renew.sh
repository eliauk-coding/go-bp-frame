#!/bin/bash
export CURR=$(cd `dirname $0`; pwd)
export ROOT=$CURR/../..

cd $ROOT
cat scripts/swagger/head.tpl > docs/api.docs.go
cat docs/swagger.json >> docs/api.docs.go
cat scripts/swagger/tail.tpl >> docs/api.docs.go
cd $CURR
