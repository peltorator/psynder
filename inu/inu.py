#!/usr/bin/env python3

import argparse
from typing import List

import requests as req
import urllib3
import os
from http.client import responses as http_status_codes

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


def show(r: req.Response):
    print(f'{r.request.method} {r.request.url}')
    for header, value in r.request.headers.items():
        print(f'{header}: {value}')
    print()
    print(r.request.body)
    print()
    print(f'Status: {r.status_code} {http_status_codes[r.status_code]}')
    for header, value in r.headers.items():
        print(f'{header}: {value}')
    print()
    print(r.text)


def url_from_args(args: argparse.Namespace, path: str, port: str):
    return f'{args.protocol}://{args.host}:{port}{path}'


def dict_from_args(args: argparse.Namespace, props: List[str]):
    return {prop: getattr(args, prop) for prop in props if getattr(args, prop) is not None}


def make_req(args: argparse.Namespace, method: str, path: str, port: str, *,
             url_props: List[str] = None, json_props: List[str] = None, auth: bool = True, **kwargs):
    headers = {}
    if 'headers' in kwargs:
        headers = kwargs['headers']
    if auth:
        headers['Authorization'] = f'Bearer {read_token()}'
    return req.request(method, url_from_args(args, path, port), verify=False,
                       json=dict_from_args(args, json_props) if json_props is not None else None,
                       params=dict_from_args(args, url_props) if url_props is not None else None,
                       headers=headers)


def signup_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/signup', args.accounts_port,
                 json_props=['email', 'password', 'kind'], auth=False)
    show(r)


def login_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/login', args.accounts_port,
                 json_props=['email', 'password'], auth=False)
    show(r)
    if r.status_code != 200:
        return
    token = r.json()['token']
    with open(get_token_file_path(), 'w') as token_file:
        token_file.write(token)


def browse_psynas_action(args: argparse.Namespace):
    r = make_req(args, 'get', '/browse-psynas', args.swipes_port,
                 json_props=['breed', 'shelter_city', 'shelter_id'],
                 url_props=['limit', 'offset'])
    show(r)


def like_psynas_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/like-psyna', args.swipes_port,
                 json_props=['psynaId'])
    show(r)


def get_liked_psynas_action(args: argparse.Namespace):
    r = make_req(args, 'get', '/liked-psynas', args.swipes_port,
                 url_props=['limit', 'offset'])
    show(r)

def psyna_info_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/psyna-info', args.swipes_port,
                 json_props=['psynaId'])
    show(r)


def add_shelter_info_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/add-shelter-info', args.shelters_port,
                 json_props=['city', 'address', 'phone'])
    show(r)


def add_psyna_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/add-psyna', args.shelters_port,
                 json_props=['name', 'breed', 'description', 'photo_link'])
    show(r)


def delete_psyna_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/delete-psyna', args.shelters_port,
                 json_props=['id'])
    show(r)


def browse_my_psynas_action(args: argparse.Namespace):
    r = make_req(args, 'post', '/browse-my-psynas', args.shelters_port,
                 url_props=['limit', 'offset'])
    show(r)



def get_token_file_path():
    return os.path.join('/', 'tmp', 'psynder-token.txt')


def read_token():
    with open(get_token_file_path(), 'r') as token_file:
        token = token_file.read()
    return token


parser = argparse.ArgumentParser()
parser.add_argument('-P', '--protocol', '--proto', default='https', choices=['http', 'https'],
                    help='protocol to use for requests')
parser.add_argument('-shp', '--shelters_port', default=443,
                    help='port of the server')
parser.add_argument('-acp', '--accounts_port', default=444,
                    help='port of the server')
parser.add_argument('-swp', '--swipes_port', default=445,
                    help='port of the server')
parser.add_argument('-a', '--host', default='localhost',
                    help='hostname (or an IP address) of the server')


def add_pagination_params(parser):
    parser.add_argument('limit', type=int, default=10, nargs='?')
    parser.add_argument('offset', type=int, default=0, nargs='?')


subparsers = parser.add_subparsers(dest='action', required=True)

parser_signup = subparsers.add_parser('signup', help='make a new user account')
parser_signup.set_defaults(action=signup_action)
parser_signup.add_argument('email', help='account email')
parser_signup.add_argument('password', help='account password')
parser_signup.add_argument('kind', choices=['person', 'shelter'], help='account kind')

parser_login = subparsers.add_parser('login', help='log into an account an get an access token')
parser_login.set_defaults(action=login_action)
parser_login.add_argument('email', help='account email')
parser_login.add_argument('password', help='account password')

parser_browse_psynas = subparsers.add_parser('browse-psynas', help='load some info about psynas')
parser_browse_psynas.set_defaults(action=browse_psynas_action)
add_pagination_params(parser_browse_psynas)
parser_browse_psynas.add_argument('--breed', help='specific psyna breed')
parser_browse_psynas.add_argument('--shelter_city', help='specific shelter city')
parser_browse_psynas.add_argument('--shelter_id', help='specific shelter id')

parser_like_psyna = subparsers.add_parser('like-psyna', help='like a psyna')
parser_like_psyna.set_defaults(action=like_psynas_action)
parser_like_psyna.add_argument('psynaId', type=int, help='id of the psyna')

parser_get_liked_psynas = subparsers.add_parser('get-liked-psynas', help='get liked psynas')
parser_get_liked_psynas.set_defaults(action=get_liked_psynas_action)
add_pagination_params(parser_get_liked_psynas)

parser_psyna_info = subparsers.add_parser('psyna-info', help='get info about psyna')
parser_psyna_info.set_defaults(action=psyna_info_action)
parser_psyna_info.add_argument('psynaId', type=int, help='id of the psyna')

parser_add_shelter_info = subparsers.add_parser('add-shelter-info', help='load some info about shelter')
parser_add_shelter_info.set_defaults(action=add_shelter_info_action)
parser_add_shelter_info.add_argument('city', type=str, help='shelter city')
parser_add_shelter_info.add_argument('address', type=str, help='shelter address')
parser_add_shelter_info.add_argument('phone', type=str, help='shelter phone')

parser_add_psyna = subparsers.add_parser('add-psyna', help='load new psyna')
parser_add_psyna.set_defaults(action=add_psyna_action)
parser_add_psyna.add_argument('name', type=str, help='psyna name')
parser_add_psyna.add_argument('--breed', type=str, help='psyna breed')
parser_add_psyna.add_argument('description', type=str, help='psyna description')
parser_add_psyna.add_argument('photo_link', type=str, help='psyna photo link')

parser_delete_psyna = subparsers.add_parser('delete-psyna', help='delete psyna')
parser_delete_psyna.set_defaults(action=delete_psyna_action)
parser_delete_psyna.add_argument('id', type=int, help='id of psyna to delete')

parser_browse_my_psynas = subparsers.add_parser('browse-my-psynas', help='browse my psynas')
parser_browse_my_psynas.set_defaults(action=browse_my_psynas_action)
add_pagination_params(parser_browse_my_psynas)


if __name__ == '__main__':
    args = parser.parse_args()
    args.action(args)
