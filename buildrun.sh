#!/bin/zsh
# buildrun.sh
pkill lang1 && echo "Sent kill"
rm -f ./lang1 && echo "Removed old binary"
echo "Going to build and run..."
printf '—%.0s' {1..$COLUMNS} ##zsh
go build && cat input.txt | ./lang1
printf '—%.0s' {1..$COLUMNS} ##zsh
