#!/bin/bash

# run the main bee script 
$HOME/.local/bin/scripthub/aliasengine

lastCommandFile=$HOME/.local/bin/scripthub/lastcommand
# if a last command file exists
if [ -f "$lastCommandFile" ]
then
  # read whatever command is there
  lastCommand=$(<$lastCommandFile)
  # remove the file
  rm $lastCommandFile

  # wait for any other params or changes
  read -p "$lastCommand" fullCommand

  # execute everything
  eval "$lastCommand $fullCommand" 
fi

