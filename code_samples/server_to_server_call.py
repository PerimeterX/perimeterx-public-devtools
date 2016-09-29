import httplib
import json
import time
import traceback
import sys

# Request info
user_ip = '1.2.3.4'
user_headers = {
    'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.41 Safari/537.36',
    'host': 'www.example.com'
}

uri = '/home'
full_url = 'http://www.example.com/home?a=b'
http_method = 'POST'
http_version = '1.1'
module_version = 'no_module'

# PX Auth Token
auth_token = 'AUTH_TOKEN'


def send_risk_request():
    body = prepare_risk_body()
    return send(body)


def send(body):
    http_client = httplib.HTTPSConnection('sapi.perimeterx.net:443', timeout=1)
    headers = {
        'Authorization': 'Bearer ' + auth_token,
        'Content-Type': 'application/json'
    }
    try:
        http_client.request('POST', '/api/v1/risk', body=json.dumps(body), headers=headers)
        r = http_client.getresponse()

        if r.status != 200:
            print 'Error performing risk call'
        response_body = r.read()

        return json.loads(response_body)
    except:
        print traceback.format_exception(*sys.exc_info())


def prepare_risk_body():
    body = {
        'request': {
            'ip': user_ip,
            'headers': format_headers(user_headers),
            'uri': uri,
            'url': full_url
        },
        'additional': {
            's2s_call_reason': 'no_cookie',
            'http_method': http_method,
            'http_version': http_version,
            'module_version': module_version,
        }
    }

    return body


def format_headers(headers):
    ret_val = []
    for key in headers.keys():
        ret_val.append({'name': key, 'value': headers[key]})
    return ret_val


# sending risk call
try:
    response = send_risk_request()
    if response:
        print 'risk score is ' + str(response['scores']['non_human'])
        print 'risk call uuid is ' + response['uuid']

except:
    print traceback.format_exception(*sys.exc_info())
