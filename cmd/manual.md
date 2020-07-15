### use http_cmd cli manual

```
./agent-otf httpclient invitation --from-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --to-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --invitation-data '{     "@type": "appuser-002",     "@id": "8",     "lable": "001",     "did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",     "router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"] }'
```


```
./agent-otf httpclient connect --from-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --to-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --connect-data '{
    "@id": "009",
    "@type": "spec/connections/1.0/request",
    "label": "bob",
    "connection": {
        "my_did": "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr",
        "my_router":["did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#1"],
        "their_did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",
        "their_router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"]
    },
    "invitation_id": "8"
}'
```

```
./agent-otf httpclient sendmsg --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --send-msg '{
    "content":"test agent",
    "connection": {
        "my_did": "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr",
        "my_router":["did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#1"],
        "their_did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",
        "their_router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"]
    }
}'
```

```
./agent-otf httpclient reqcredential --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --req-credential '{
    "@type":"spec/issue-credential/1.0/request-credential",
    "@id":"11",
    "comment":"request 002",
    "connection": {
        "my_did": "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr",
        "my_router":["did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#1"],
        "their_did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",
        "their_router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"]
    },
    "formats":[
        {
            "attach_id":"1",
            "format":"string"
        }
    ],
    "requests_attach":[
        {
            "@id":"1",
            "data":{
                "json":{"name":"age","value":"greater than 18"}
            }
        }   
    ]
}'
```

```
./agent-otf httpclient reqpresentation --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --req-presentation '{
    "@type":"spec/issue-credential/1.0/propose-credential",
    "@id":"13",
    "comment":"proposal1",
    "connection": {
         "my_did": "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr",
        "my_router":["did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#1"],
        "their_did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",
        "their_router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"]
    }
}'
```

```
./agent-otf httpclient querycredential --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --credential-id 11
Password:
==============credential==============
{"code":0,"msg":"","data":{"message_type":0,"content":{"formats":[{"attach_id":"1","format":"base64"}],"credentials~attach":[{"@id":"1","lastmod_time":"2020-07-15T15:20:18.630702+08:00","data":{"base64":"eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VFFGbWZyYlFib0RVU2VWOTg5WnA4NjdyNkRhd2IxTVBTRiNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRRRm1mcmJRYm9EVVNlVjk4OVpwODY3cjZEYXdiMU1QU0YiLCJleHAiOjE1OTQ4ODQwMTcsIm5iZiI6MTU5NDc5NzYxOCwiaWF0IjoxNTk0Nzk3NjE4LCJqdGkiOiJ1cm46dXVpZDo5MzA3ZDdiNS1iMDcxLTRiNTktOWFlMy1iNTMzMWNiZDI2YWMiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA3OjIwOjE4WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.QfhkbGsYeYW+irt7PKZzU31DW1N4KWloZUYxc5ja/DWTYM/+nGyGgxPYPnTz2WMHsRyYSKprQqdqcxPSh7tFiw=="}}]}}}
==============credential==============
```