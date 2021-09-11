#!/bin/bash
export CURR=$(cd `dirname $0`; pwd)
export ROOT=$CURR/../..

cd $ROOT
swag init --generalInfo main.go --output docs
rm -f ./docs/swagger.yaml
mv ./docs/docs.go ./docs/api.docs.go
cd $CURR
