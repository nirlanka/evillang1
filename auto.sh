#!/bin/zsh

find . -type f \( -name "*.go" -o -name "*.tmpl" -o -name "*.txt" \) | \
	entr -r ./buildrun.sh
