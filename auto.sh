#!/bin/zsh
# auto.sh
find . -type f \( -name "*.go" -o -name "*.tmpl" \) | \
	entr -r ./buildrun.sh

## Source: https://www.reddit.com/r/golang/comments/32mdo2/reliable_autocompileandrun_your_project_on_save/

## brew install entr