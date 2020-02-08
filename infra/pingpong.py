import requests

r = requests.post("http://localhost:9080/api/v1/ping/", data={"foo": "bar"})
print(r.text)

r = requests.post("http://localhost:9080/api/v1/nodes", data={"token": "0"})
print(r.text)

r = requests.post("http://localhost:9080/api/v1/systems", data={"token": "1"})
print(r.text)

