#!/usr/bin/python

from dataclasses import dataclass
import sys
from typing import Optional
from unittest import result
import requests
import json
import os

PATH_PREFIX = '/usr/local/bin/scripthub'
DOMAIN = 'https://scriptexchange.herokuapp.com' # 'http://127.0.0.1:8000'
LIST_COMMAND = 'ls'
UPDATE_COMMAND = 'update'


def print_help_message(message: str):
  help = f"""
  {message}
  
  ScriptHub usage:
    scripthub update <API_KEY>  # download all of your scripts
    scripthub ls                # list all of your Collections
    scripthub ls Collection     # list all of the aliases in a Collection
  """
  sys.exit(help)


def parse_response_content_as_json(text: str) -> Optional[any]:
  try:
    result = json.loads(text)
    return result['content'] if 'content' in result else None
  except:
    return None
  

def get_collections_for_user(api_key: str):
  url = f'{DOMAIN}/alias/collections'
  headers = {'Authorization': f'Api-Key {api_key}'}
  response = requests.request('GET', url, headers=headers)
  status_code = response.status_code
  
  if status_code == 200:
    result = parse_response_content_as_json(response.text)
    return result if result else []
  elif status_code == 403:
    print_help_message('Please provide a valid API Key')
  elif status_code == 404:
    print_help_message('User does not exist')
  else:
    print_help_message('Unknown error happened whilst trying to get Scripts')


def get_scripts_for_collection(api_key: str, collection_id: str):
  url = f'{DOMAIN}/alias/collections/{collection_id}'
  headers = {'Authorization': f'Api-Key {api_key}'}
  response = requests.request('GET', url, headers=headers)
  status_code = response.status_code
  
  if status_code == 200:
    result = parse_response_content_as_json(response.text)
    return result if result else []
  elif status_code == 403:
    print_help_message('Please provide a valid API Key')
  elif status_code == 404:
    print_help_message('User does not exist')
  else:
    print_help_message('Unknown error happened whilst trying to get Scripts')


def get_content_for_script(api_key: str, script_id: str, ) -> Optional[str]:
  url = f'{DOMAIN}/alias/scripts/{script_id}/content'
  headers = {'Authorization': f'Api-Key {api_key}'}
  response = requests.request('GET', url, headers=headers)
  status_code = response.status_code
  
  if status_code == 200:
    return parse_response_content_as_json(response.text)
  elif status_code == 403:
    print_help_message('Please provide a valid API Key')
  elif status_code == 404:
    return None  # soft exit
  else:
    return None  # soft exit
  

def clear_existing_folder():
  clear_command = f'rm -rf {PATH_PREFIX}'
  os.system(clear_command)
  print(f'Cleared cache at {PATH_PREFIX}')


def create_scripthub_folder():
  create_command = f'mkdir {PATH_PREFIX}/'
  os.system(create_command)
  print(f'Created cache at {PATH_PREFIX}')


def update(api_key: str):
  # for collection, script and content tree
  collection_list = get_collections_for_user(api_key)
  for collection in collection_list:
    script_list = get_scripts_for_collection(api_key, collection['id'])
    for script in script_list:
      content = get_content_for_script(api_key, script['id'])
      script['content'] = content
    
    collection['script_list'] = script_list
  
  # recreate all data
  clear_existing_folder()
  create_scripthub_folder()
  
  # write file data
  alias_list = []
  collection_name_list = []
  
  for collection in collection_list:
    name: str = collection['name']
    folder = name.replace(' ', '_')
    
    collection_name_list.append(name)
    
    create_folder_command = f'mkdir {PATH_PREFIX}/{folder}'
    os.system(create_folder_command)
    
    for script in collection['script_list']:
      alias: str = script['alias']
      content: str = script['content']
      path = f'{PATH_PREFIX}/{folder}/{alias}'
      
      f = open(path, 'w')
      f.write(content)
      f.close()
      
      format_script_command = f"sed -i '' 's/\r//g' {path}"
      os.system(format_script_command)
      
      permission_command = f'chmod 777 {path}'
      os.system(permission_command)
      
      print(f'Wrote file {path}')
      
      alias_command = f'alias {alias}={path}'
      alias_list.append(alias_command)
      
  # write the alias meta data
  aliases_file_path = f'{PATH_PREFIX}/.aliases'
  aliases_content = '\n'.join(alias_list)
  f = open(aliases_file_path, 'w')
  f.write(aliases_content)
  f.close()
  
  # write the collection meta data
  collection_file_path = f'{PATH_PREFIX}/.collections'
  collection_content = '\n'.join(collection_name_list)
  f = open(collection_file_path, 'w')
  f.write(collection_content)
  f.close()
  

def list_collections():
  list_command = f'cat {PATH_PREFIX}/.collections'
  os.system(list_command)
  os.system('echo \n')
  

def list_collection_aliases(collection_name_list):
  full_collection_name = ' '.join(collection_name_list)
  folder_name = full_collection_name.replace(' ', '_')
  list_command = f'ls {PATH_PREFIX}/{folder_name}/'
  os.system(list_command)


# get start data
arguments = sys.argv
arg_len = len(sys.argv) - 1 # excluding 'self' as the first argument


# main business logic
if arg_len < 0:
  print_help_message("Unexpected number of arguments.")
elif arg_len >= 1:
  second_argument = arguments[1]
  if second_argument == UPDATE_COMMAND:
    if arg_len >= 2:
      third_argument = arguments[2]
      update(third_argument)
    else:
      print_help_message("Missing API_KEY argument.")
  elif second_argument == LIST_COMMAND:
    if arg_len == 1:
      list_collections()
    elif arg_len > 1:
      name_list = arguments[2:]
      list_collection_aliases(name_list)
  else:
    print_help_message("Invalid arguments.")