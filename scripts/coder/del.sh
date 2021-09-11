#!/bin/bash
export CURR=$(cd `dirname $0`; pwd)
export ROOT=$CURR/../..
function usage
{
    echo
    echo $1
    echo "使用方法 -- scripts/coder/del APP_NAME MODULE_NAME"
    echo
    exit
}

if [ "$1" == "" ]; then
    usage "参数错误 -- APP_NAME 不能为空!!"
fi

if [ "$2" == "" ]; then
    usage "参数错误 -- MODULE_NAME 不能为空!!"
fi

echo "将要删除如下目录和文件："
echo "  apps/$1/actions/$2" 
echo "  apps/$1/models/$2" 
echo "  apps/$1/services/$2" 
echo "  apps/$1/router/$2.go"
echo "确认删除？[Y/N] "
read OPT

if [[ "$OPT" == "Y" || "$OPT" == "y" ]];then
    cd $ROOT
    rm -f -r "apps/$1/actions/$2" 
    rm -f -r "apps/$1/models/$2" 
    rm -f -r "apps/$1/services/$2" 
    rm -f "apps/$1/router/$2.go"
    cd $CURR
fi
