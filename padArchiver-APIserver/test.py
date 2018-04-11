'''
    this script uses requests and provoj
    provoj is a simple library to test endpoints of an API RESTful
    provoj can be downloaded using: pip install provoj
    provoj repository: https://github.com/arnaucode/provoj

    To run this test, just need to run:
    python test.py
'''
import provoj
import requests


test = provoj.NewTest("testing padArchiver API Server")

url = "http://127.0.0.1:3080"



jsonData = {"link": "http://board.net/p/pad1", "dir": "Group1", "title": "Pad1"}
r = requests.post(url + "/repos/repo01/pad", json=jsonData)
test.rStatus("POST add new pad", r)
print(r.json())

jsonData = {"link": "http://board.net/p/pad2", "dir": "Group2", "title": "Pad2"}
r = requests.post(url + "/repos/repo01/pad", json=jsonData)
test.rStatus("POST add new pad", r)
print(r.json())

jsonData = {"link": "http://board.net/p/pad3", "dir": "Group2", "title": "Pad3"}
r = requests.post(url + "/repos/repo01/pad", json=jsonData)
test.rStatus("POST add new pad", r)
print(r.json())


r = requests.get(url + "/repos")
test.rStatus("GET repos list", r)
print(r.json())
reposList = r.json()
testRepo = reposList[0]

r = requests.get(url + "/repos/" + testRepo)
test.rStatus("GET repo " + testRepo + " list", r)
print(r.json())



test.printScores()
