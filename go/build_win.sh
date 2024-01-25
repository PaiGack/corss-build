#!/bin/bash

# sudo apt-get update
# sudo apt-get install gcc-mingw-w64-x86-64
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o main.exe
# > file main.exe
# main.exe: PE32+ executable (console) x86-64, for MS Windows
