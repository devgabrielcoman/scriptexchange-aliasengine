#!/bin/bash

# get python file name
FILE_NAME=$1
NAME="${FILE_NAME%.*}"

# transform from .py file to .c file
cython $NAME.py --embed

# set PYTHONLIBVER version 
PYTHONLIBVER=python$(python3 -c 'import sys; print(".".join(map(str, sys.version_info[:2])))')$(python3-config --abiflags)

# compile .c file to binary
gcc -Os $(python3-config --includes) $NAME.c -o build/$NAME $(python3-config --ldflags) -l$PYTHONLIBVER

# cleanup
rm $NAME.c 