#!/bin/sh

git rm -r --cached .
git add .

git commit -am "Remove ignored files"