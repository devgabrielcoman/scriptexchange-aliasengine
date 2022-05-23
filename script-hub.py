#!/usr/bin/python

import sys
from typing import Optional
import requests
import json
import os

PATH_PREFIX = '/usr/local/bin/scripthub'
DOMAIN = 'https://scriptexchange.herokuapp.com' # 'http://127.0.0.1:8000'

def get_domain() -> str:
  return DOMAIN


def get_api_key() -> str:
  arguments = sys.argv
  if len(arguments) != 2:
    sys.exit('Please provide the API Key as an argument to the script')

  return str(sys.argv[1])


def parse_response_content_as_json(text: str) -> Optional[any]:
  try:
    result = json.loads(text)
    return result['content'] if 'content' in result else None
  except:
    return None


def get_scripts_for_user(api_key: str) -> list[dict]:
  url = f'{get_domain()}/alias/scripts'
  payload = {}
  headers = {'Authorization': f'Api-Key {api_key}'}
  response = requests.request('GET', url, headers=headers, data=payload)
  status_code = response.status_code
  
  if status_code == 200:
    result = parse_response_content_as_json(text=response.text)
    return result if result else []
  elif status_code == 403:
    sys.exit('Please provide a valid API Key')
  elif status_code == 404:
    sys.exit('User does not exist')
  else:
    sys.exit('Unknown error happened whilst trying to get Scripts')


def get_script_content(pk: str, api_key: str) -> Optional[str]:
  url = f'{get_domain()}/alias/scripts/{pk}/content'
  payload = {}
  headers = {'Authorization': f'Api-Key {api_key}'}
  response = requests.request('GET', url, headers=headers, data=payload)
  status_code = response.status_code
  
  if status_code == 200:
    return parse_response_content_as_json(text=response.text)
  elif status_code == 403:
    sys.exit('Please provide a valid API Key')
  elif status_code == 404:
    return None  # soft exit
  else:
    return None  # soft exit


def get_alias_file_full_path(path_prefix: str, alias: str) -> str:
  return f'{path_prefix}/{alias}.sh'
  

def get_aliases_file_full_path(path_prefix: str) -> str:
  return f'{path_prefix}/.aliases'


def clear_existing_alias_folder(path_prefix: str):
  clear_command = f'rm -rf {path_prefix}/*'
  os.system(clear_command) 
  print(f'Cleared current Alias cache at {path_prefix}')
  

def download_scripts(script_list: list[dict], api_key: str, path_prefix: str):
  for script in script_list:
    id = script['id']
    alias = script['alias']
    content: Optional[str] = get_script_content(pk=id, api_key=api_key)
    if content:
      path = get_alias_file_full_path(path_prefix=path_prefix, alias=alias)
      f = open(path, 'w')
      f.write(content)
      f.close()
      format_command = f"sed -i '' 's/\r//g' {path}"
      rights_command = f'chmod 777 {path}'
      os.system(format_command)
      os.system(rights_command)
      print(f'Wrote file {path}')
      

def form_alias_file(script_list: list[dict], path_prefix: str):
  alias_list: list[str] = []
  for script in script_list:
    alias = script['alias']
    path = get_alias_file_full_path(path_prefix=path_prefix, alias=alias)
    alias_command = f'alias {alias}={path}'
    alias_list.append(alias_command)
  
  alias_file_content = '\n'.join(alias_list)
  alias_file_path = get_aliases_file_full_path(path_prefix=path_prefix)
  f = open(alias_file_path, 'w')
  f.write(alias_file_content)
  f.close()
  print(f'Wrote main Alias file at {alias_file_path}')


api_key = get_api_key()
script_list = get_scripts_for_user(api_key=api_key)
print(f'Found {len(script_list)} Scripts')
clear_existing_alias_folder(path_prefix=PATH_PREFIX)
download_scripts(script_list=script_list, api_key=api_key, path_prefix=PATH_PREFIX)
form_alias_file(script_list=script_list, path_prefix=PATH_PREFIX)
    
    