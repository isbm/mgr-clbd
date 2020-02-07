import requests

r = requests.post("http://localhost:9090/api/v1/ping/", data={"foo": "bar"})
print(r.text)

