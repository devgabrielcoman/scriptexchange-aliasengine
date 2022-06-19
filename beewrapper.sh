#!/bin/bash

# run the main bee script - passing any arguments
$HOME/.local/bin/scripthub/bbee "$@"

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

