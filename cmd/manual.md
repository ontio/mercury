# agent-otf cli manual


## 目录

[TOC]

##  1. did cmd

### 1.1 生成并绑定一个新的DID

```

./agent-otf did newdid

Password:
Did:  did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ,  Hash:4c4d3b8bafd22c16cc2b39e8383c0d060430e608d1726713b8a9ade72db95159
did:  did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ

```

### 1.2 为DID增加一个Service endPoint


```

./agent-otf did addsvr --did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --service-id "1" --type "1" --endpoint "http://127.0.0.1:8080"  --index 1

Password:
txHash:8d5ddcd9bc7050748336de8c81da689f9798c66b687c3eefa835337276a9272f

```

### 1.3 更新指定DID的 Service endPoint

```
./agent-otf did updatesvr --did did:ont:TWXpoTiCedMBUHiyrPtpwi171yp4gDSPyC --service-id "1" --type "1" --endpoint "http://127.0.0.1:8089" --index 1

Password:
txHash:44f5dc80012f516835b014afd4dbbe9cb255e1c6093e4a034d72133497ba7dc8

```

### 1.4 查询DID 对应的DID doc

```

./agent-otf did diddoc --did did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ


doc: &{[https://www.w3.org/ns/did/v1 https://ontid.ont.io/did/v1] did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ [map[controller:did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ id:did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ#keys-1 publicKeyHex:03afb755bf9c9a3a7577b6d210f07aeac2730ff9800b7af443917be80ef1ddd52f type:EcdsaSecp256r1VerificationKey2019]] [did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ#keys-1] <nil> <nil> [{did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ#1 1 http://127.0.0.1:8080}] <nil> 1.59480394e+09 1.594804114e+09 }
```

### 1.5 查询DID对应的 Service end points

```

./agent-otf did endpoint --did did:ont:TT2sekt32e4pDNrjmjFsJYcXJhGaiiurfQ

endPoints:[http://127.0.0.1:8080]

```

## 2、http_client cli cmd
client cli是模拟一个用户的终端agent行为的cli,提供如下功能:

### 2.1 创建一个 Invitation
```


./agent-otf httpclient invitation --from-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --to-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --invitation-data '{     "@type": "appuser-002",     "@id": "8",     "lable": "001",     "did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",     "router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"] }'


```

### 2.2 创建链接

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

### 2.3 发送一个通用消息
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

### 2.4 请求凭证
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

### 2.5 请求presentation
```
./agent-otf httpclient reqpresentation --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --req-presentation '{
    "@type":"spec/issue-credential/1.0/propose-credential",
    "@id":"15",
    "comment":"proposal1",
    "connection": {
         "my_did": "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr",
        "my_router":["did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#1"],
        "their_did": "did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF",
        "their_router":["did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF#1"]
    }
}'
```

### 2.6 查询一个凭证
```
./agent-otf httpclient querycredential --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --credential-id 11
Password:
==============credential==============
{"code":0,"msg":"","data":{"message_type":0,"content":{"formats":[{"attach_id":"1","format":"base64"}],"credentials~attach":[{"@id":"1","lastmod_time":"2020-07-15T15:20:18.630702+08:00","data":{"base64":"eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VFFGbWZyYlFib0RVU2VWOTg5WnA4NjdyNkRhd2IxTVBTRiNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRRRm1mcmJRYm9EVVNlVjk4OVpwODY3cjZEYXdiMU1QU0YiLCJleHAiOjE1OTQ4ODQwMTcsIm5iZiI6MTU5NDc5NzYxOCwiaWF0IjoxNTk0Nzk3NjE4LCJqdGkiOiJ1cm46dXVpZDo5MzA3ZDdiNS1iMDcxLTRiNTktOWFlMy1iNTMzMWNiZDI2YWMiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA3OjIwOjE4WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.QfhkbGsYeYW+irt7PKZzU31DW1N4KWloZUYxc5ja/DWTYM/+nGyGgxPYPnTz2WMHsRyYSKprQqdqcxPSh7tFiw=="}}]}}}
==============credential==============
```

### 2.7 查询一个Presentation

```
./agent-otf httpclient querypresentation --from-did did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr --to-did did:ont:TQFmfrbQboDUSeV989Zp867r6Dawb1MPSF --presentation-id 15


==============presentation==============
{"code":0,"msg":"","data":{"message_type":0,"content":{"formats":[{"attach_id":"1","format":"base64"}],"presentations~attach":[{"@id":"1","lastmod_time":"2020-07-17T15:21:23.109729+08:00","data":{"base64":"eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEw5ZDlKZGRleVVaem56OWVpVE53TEVXUUFpcFVMcjRtciNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRMOWQ5SmRkZXlVWnpuejllaVROd0xFV1FBaXBVTHI0bXIiLCJhdWQiOiIiLCJqdGkiOiJ1cm46dXVpZDoyYjlmYWYwYS1kY2Q3LTQ2YTEtOGMwMC1mNGRkY2FmYTEyNjEiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE3VDA3OjIxOjIyWiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.7Smhyzps3mt/LOQxCgvAbK8JgawMAwXoY7t4Un+6x8r4hdaVOPfngoebmUrhRiVWlIDqpVm7MNkq1f7t+V3BCA=="}}]}}}
==============presentation==============
```