#!/bin/bash

# define variables used by the script
PATH_PREFIX=$HOME/.local/bin/scripthub
MAIN_SCRIPT=scripthub
MAIN_SCRIPT_URL=https://raw.githubusercontent.com/devgabrielcoman/scriptexchange-aliasengine/main/scripthub

# delete everything in install path
rm -rf $PATH_PREFIX/
echo "Cleaned out source folder" $PATH_PREFIX
mkdir -p $PATH_PREFIX/
echo "Recreated source folder" $PATH_PREFIX

# copy remote script to local
main_script_path=$PATH_PREFIX/$MAIN_SCRIPT

# download
curl $MAIN_SCRIPT_URL --output $main_script_path
chmod 777 $main_script_path
echo "Downloaded ScriptHub app"

# setup default command
echo "alias scripthub=$main_script_path" >> $PATH_PREFIX/.aliases
echo "Setup ScriptHub"