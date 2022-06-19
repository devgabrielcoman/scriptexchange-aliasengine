#!/bin/bash

# define variables used by the script
PATH_PREFIX=$HOME/.local/bin/scripthub

# create local build folder
mkdir -p build 

# build main app
cd bee 
go build -o ../build/bbee 
cd ../

# copy programs
cp build/bbee $PATH_PREFIX/bbee
cp beewrapper.sh $PATH_PREFIX/bee
chmod +x $PATH_PREFIX/bbee
chmod +x $PATH_PREFIX/bee

# clan data
rm -rf build