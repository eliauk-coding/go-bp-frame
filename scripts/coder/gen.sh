#!/bin/bash
export CURR=$(cd `dirname $0`; pwd)
export ROOT=$CURR/../..

function usage
{
    echo
    echo $1
    echo "使用方法 -- scripts/coder/gen APP_NAME MODULE_NAME"
    echo
    exit
}

function create
{
    if [ ! -d "$1" ]; then
    mkdir -p "$1"
    fi

    if [ ! -f "$1/$2.go" ]; then
    echo "package $3">"$1/$2.go"
    fi
}

if [ "$1" == "" ]; then
    usage "参数错误 -- APP_NAME 不能为空!!"
fi

if [ "$2" == "" ]; then
    usage "参数错误 -- MODULE_NAME 不能为空!!"
fi

cd $ROOT
create "apps/$1/actions/$2" "$2" "$2"
create "apps/$1/models/$2" "$2" "$2"
create "apps/$1/services/$2" "$2" "$2"
create "apps/$1/router" "$2" "router"
cd $CURR