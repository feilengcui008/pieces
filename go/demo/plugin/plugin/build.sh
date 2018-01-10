#!/bin/bash 
gcc -shared -fPIC cplugin.c -o cplugin.so
go build -buildmode=plugin goplugin.go
