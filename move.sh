#!/bin/bash

# usage: ./move.sh <dirname1/dirname2>

# 引数のチェック
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <dirPath>"
    exit 1
fi

dirPath=$1

mkdir -p "$dirPath"

mv main*.go $dirPath/ && mv input.txt $dirPath/

cd "$dirPath" && \
    go mod init github.com/kngnkg/algorithm-practice/"$dirPath" && \
    cd -

go work use $dirPath
