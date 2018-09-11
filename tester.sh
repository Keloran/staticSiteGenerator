#!/usr/bin/env bash

usage()
{
cat << EOF
    This script can be used to build and to run the app

     usage: $0 options

    OPTIONS:
        -h show this message
        -b build app
        -r build and run app
        -c clean the app
EOF
}

LOC=$(pwd)

buildIt()
{
    buildGoApp
    buildNodeApp
    cleanGo
}

cleanIt()
{
    cd $LOC
    rm -rf staticSite.app
    rm -rf go/staticSite.app
    rm -rf node/dist
    rm -rf node/.cache
}

runIt()
{
    cd $LOC
    open staticSite.app
}

buildGoApp()
{
    cd $LOC/go
    go build -o staticSite.app/Contents/MacOS/staticSite
    cp -r staticSite.app ../
}

buildNodeApp()
{
    cd $LOC/node
    yarn build
    cp -r dist/ ../staticSite.app/Contents/MacOS/dist
}

cleanGo()
{
     cd $LOC/go
     rm -rf staticSite.app
}

while getopts ":?cbrh" OPTION; do
    case $OPTION in
        c)
            cleanIt
            ;;
        b)
            buildIt
            ;;
        r)
            buildIt
            runIt
            ;;
        h)
            usage
            exit
            ;;
        ?)
            usage
            exit
            ;;
    esac
done