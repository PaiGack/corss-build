#!/bin/bash

# > sudo apt-get update
# 32 位
# > sudo apt-get install gcc-mingw-w64-i686
i686-w64-mingw32-gcc main.c -o main_32.exe
# > file main_32.exe 
# main_32.exe: PE32 executable (console) Intel 80386, for MS Windows


# 64 位
# > sudo apt-get install gcc-mingw-w64-x86-64
x86_64-w64-mingw32-gcc main.c -o main_64.exe
# > file main_64.exe
# main_64.exe: PE32+ executable (console) x86-64, for MS Windows