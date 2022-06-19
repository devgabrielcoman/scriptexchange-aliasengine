#!/bin/bash

# define variables used by the script
PATH_PREFIX=$HOME/.local/bin/scripthub

SCRIPT_NAME=aliasengine
WRAPPER_NAME=bee
MAIN_SCRIPT_URL=https://github.com/devgabrielcoman/scriptexchange-aliasengine/raw/main/aliasengine/Build/Products/Debug/aliasengine
WRAPPER_SCRIPT_URL=https://raw.githubusercontent.com/devgabrielcoman/scriptexchange-aliasengine/main/beewrapper.sh

# create folder
mkdir -p $PATH_PREFIX/
echo "Create source folder" $PATH_PREFIX

# paths to put things in
main_script_path=$PATH_PREFIX/$SCRIPT_NAME
wrapper_script_path=$PATH_PREFIX/$WRAPPER_NAME

# download
echo "Downloading script"
curl -fsSL $MAIN_SCRIPT_URL --output $main_script_path
curl -fsSL $WRAPPER_SCRIPT_URL --output $wrapper_script_path
# make it executable
chmod +x $main_script_path
chmod +x $wrapper_script_path
echo "Downloaded script"