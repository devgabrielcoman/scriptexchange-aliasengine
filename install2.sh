#!/bin/bash

# define variables used by the script
PATH_PREFIX=$HOME/.local/bin/scripthub
SCRIPT_NAME=bee
MAIN_SCRIPT_URL=https://github.com/devgabrielcoman/scriptexchange-aliasengine/raw/main/aliasengine/Build/Products/Debug/aliasengine

# create folder
mkdir -p $PATH_PREFIX/
echo "Create source folder" $PATH_PREFIX

# copy remote script to local
main_script_path=$PATH_PREFIX/$SCRIPT_NAME

# download
echo "Downloading script"
curl -fsSL $MAIN_SCRIPT_URL --output $main_script_path
# make it executable
chmod +x $main_script_path
echo "Downloaded script"