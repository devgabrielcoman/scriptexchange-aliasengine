#!/bin/bash

# define variables used by the script
PATH_PREFIX=/usr/local/bin/scripthub
DOMAIN="https://scriptexchange.herokuapp.com" # "http://127.0.0.1:8000"
API_KEY=$1

# delete all existing aliases
rm -rf $PATH_PREFIX/*
rm $PATH_PREFIX/.aliases
mkdir $PATH_PREFIX

# get the current list of scripts for this particular user
SCRIPT_RESULT=$(curl --silent --location --request GET "$DOMAIN/alias/scripts" --header "Authorization: Api-Key $API_KEY")
SCRIPT_LIST=$( jq -r '.content' <<< "${SCRIPT_RESULT}" ) 

# for each script object in the list
for row in $(echo "${SCRIPT_LIST}" | jq -r '.[] | @base64'); do
  _jq() {
    echo ${row} | base64 --decode | jq -r ${1}
  }

  # get ID and ALIAS
  ID=$(_jq '.id')
  ALIAS=$(_jq '.alias')
  # download the content of the script
  CONTENT=$(curl --silent --location --request GET "$DOMAIN/alias/scripts/$ID/content"  --header "Authorization: Api-Key $API_KEY" | jq -r '.content')
  echo "Downloaded Script" $ID "with Alias" $ALIAS

  # form the final path where we'll save the script and write it
  FULL_PATH=$PATH_PREFIX/$ALIAS.sh
  echo $CONTENT > $FULL_PATH
  # additionally replace unwanted characters
  sed -i '' 's/\r /\n/g' $FULL_PATH
  # give corresponding permissions
  chmod 777 $FULL_PATH
  # and fill out the main alias file
  echo "alias $ALIAS=$FULL_PATH" >> $PATH_PREFIX/.aliases
done