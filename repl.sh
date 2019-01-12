#!/bin/bash
# 
src="..\/..\/res\/"
dst="\/static\/plug\/fly\/"

# find views/ -name \*.html | xargs sed -i "s/radiumme.com\/dragon/radiumme.com\/story/g"
find . -name main* | xargs sed -i '' "s/${src}/${dst}/g"