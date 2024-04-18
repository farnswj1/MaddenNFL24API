import requests
import json

with open('nfl.json') as f:
    a = json.load(f)

for i in a:
    r = requests.post(
        'http://localhost:8080/players',
        data=i,
        headers={'Content-Type': 'application/json'}
    )
    print(r.content)
    r.raise_for_status()
