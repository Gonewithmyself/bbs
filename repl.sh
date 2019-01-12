#!/bin/bash
# 
src="..\/..\/res\/"
dst="\/static\/plug\/fly\/"

# macos -i 后面需要指定备份文件
find views -name \*.html | xargs sed -i '' "s/${src}/${dst}/g"