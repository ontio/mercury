# OTF RestAPI Document

[TOC]

## 1. Overview

This document the introduction for cloud agent RESTful API.

**Note**: The startup parameter ```--enable-pack```, all messages will be encrypt and packed to envelope .



## 2. Rest API List

| Name                      | Method | URL                              | Description                 |
| ------------------------- | ------ | -------------------------------- | --------------------------- |
| invitation                | POST   | /api/v1/ invitation              | create a invitation         |
| connection request        | POST   | /api/v1/connectionrequest        | request connection          |
| send message              | POST   | /api/v1/sendbasicmsg           | send a basic message      |
| send proposal credential  | POST   | /api/v1/sendproposalcredential   | send a proposal credential  |
| send request credential   | POST   | /api/v1/sendrequestcredential    | send a request credential   |
| send request presentation | POST   | /api/v1/ sendrequestpresentation | send a request presentation |
| query credential          | POST   | /api/v1/querycredential          | query a credential          |
| query presentation        | POST   | /api/v1/querypresentation        | query a presentation        |
| query basic message       | POST   | /api/v1/querybasicmsg          | query basic message       |
| send disconnect           | POST   | /api/v1/senddisconnect           | send disconnect request     |

### 2.1 Invitation

POST

```
/api/v1/invitation
```

Request body example:

```json
{
	"@type":"spec/connection/1.0/invitation",
	"@id":"A000000020",
	"label":"alice",
	"did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
	"router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
}
```

Response:

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "message_type": 0,
        "content": {
            "@type": "spec/connection/1.0/invitation",
            "@id": "A000000020",
            "label": "alice",
            "did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
            "router": [
                "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
            ]
        }
    }
}
```



### 2.2 connection request

POST

```
/api/v1/connectionrequest
```

Request body example:

```json
{
    "@id": "000019",
    "@type": "spec/connections/1.0/request",
    "label": "bob",
    "connection": {
        "my_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
    },
    "invitation_id": "A000000019"
}

```

Response:

```json
{
    "code": 0,
    "msg": ""
}
```



### 2.3 send message

POST

```
/api/v1/sendbasicmsg
```

Request body example:

```json
{
	"content":"hello world",
    "connection": {
        "my_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
    }
}
```

Response :

```json
{
    "code": 0,
    "msg": ""
}
```

### 2.4 send proposal credential

POST

```
/api/v1/sendproposalcredential
```

Request body example:

```json
{
	"@type":"spec/issue-credential/1.0/propose-credential",
	"@id":"P000002",
	"comment":"proposal1",
    "connection": {
       "my_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
    }
}
```

Response:

```json
{
    "code": 0,
    "msg": ""
}
```

### 2.5 send request credential

POST

```
/api/v1/sendrequestcredential
```

Request body example:

```json
{
	"@type":"spec/issue-credential/1.0/request-credential",
	"@id":"RC00000031",
	"comment":"request 020",
    "connection": {
        "my_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
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
}
```

Response:

```json
{
    "code": 0,
    "msg": ""
}
```

### 2.6 send request presentation

POST

```
/api/v1/sendrequestpresentation
```

Request body example

```json
{
	"@type":"spec/present-proof/1.0/request-presentation",
	"@id":"RP00000019",
	"comment":"test0001",
    "connection": {
        "my_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
    },
      "formats":[
    	{
    		"attach_id":"1",
    		"format":"base64"
    	}
    ],
    "request_presentation_attach":[
    	{
	    	"@id":"1",
	    	"data":{
	    		"base64":"UkMwMDAwMDAzMQ=="
	    	}
    	}
    ]
}
```

Response

```json
{
    "code": 0,
    "msg": ""
}
```

### 2.7 query credential

POST

```
/api/v1/querycredential
```

Request body example:

```json
{
	"did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
	"id":"RC00000030"
}
```

Response

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "message_type": 0,
        "content": {
            "formats": [
                {
                    "attach_id": "1",
                    "format": "base64"
                }
            ],
            "credentials~attach": [
                {
                    "@id": "1",
                    "lastmod_time": "2020-07-15T16:54:58.511900118+08:00",
                    "data": {
                        "base64": "eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEdBOFlXcHF3eGU5TERRQ2RUR0M3d214VG11bUVROUdqeCNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRHQThZV3Bxd3hlOUxEUUNkVEdDN3dteFRtdW1FUTlHangiLCJleHAiOjE1OTQ4ODk2OTcsIm5iZiI6MTU5NDgwMzI5OCwiaWF0IjoxNTk0ODAzMjk4LCJqdGkiOiJ1cm46dXVpZDoxYWY3ZGJmZC1mODRjLTQ3NzctOTgzZC1iNTIzZGZlYTA0NmUiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA4OjU0OjU4WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.qrgt72U95rby3L+Ox+ZV0rr7doJ8T/yv1OYZcvx7BH63oJY2npxWl9X9sGOTfOHskAUyauwAvzGOIi53oKHlhA=="
                    }
                }
            ]
        }
    }
}
```

### 2.8 query basic message

POST

```
/api/v1/querypresentation
```

Request body example:

```json
{
	"did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
	"id":"RP00000019"
}
```

Response

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "message_type": 0,
        "content": {
            "formats": [
                {
                    "attach_id": "1",
                    "format": "base64"
                }
            ],
            "presentations~attach": [
                {
                    "@id": "1",
                    "lastmod_time": "2020-07-15T17:11:23.379131151+08:00",
                    "data": {
                        "base64": "eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VFFBaWFlZmtkeXBTQmlDU1Y5aDlNZkJKMllweTlmYTdMWSNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRRQWlhZWZrZHlwU0JpQ1NWOWg5TWZCSjJZcHk5ZmE3TFkiLCJhdWQiOiIiLCJqdGkiOiJ1cm46dXVpZDo2ZDkwYmQzYi00NWEyLTRkMWUtYTI0Ni0yYWJmODFhYzllZWIiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sInZlcmlmaWFibGVDcmVkZW50aWFsIjpbImV5SmhiR2NpT2lKRlV6STFOaUlzSW10cFpDSTZJbVJwWkRwdmJuUTZWRWRCT0ZsWGNIRjNlR1U1VEVSUlEyUlVSME0zZDIxNFZHMTFiVVZST1VkcWVDTnJaWGx6TFRFaUxDSjBlWEFpT2lKS1YxUWlmUT09LmV5SnBjM01pT2lKa2FXUTZiMjUwT2xSSFFUaFpWM0J4ZDNobE9VeEVVVU5rVkVkRE4zZHRlRlJ0ZFcxRlVUbEhhbmdpTENKbGVIQWlPakUxT1RRNE9UQTJNaklzSW01aVppSTZNVFU1TkRnd05ESXlNaXdpYVdGMElqb3hOVGswT0RBME1qSXlMQ0pxZEdraU9pSjFjbTQ2ZFhWcFpEcGtOalkwTURFMU1pMDVORFEyTFRRd05URXRZV00xTUMxbE9UYzVNRFU1WWpSbFl6SWlMQ0oyWXlJNmV5SkFZMjl1ZEdWNGRDSTZXeUpvZEhSd2N6b3ZMM2QzZHk1M015NXZjbWN2TWpBeE9DOWpjbVZrWlc1MGFXRnNjeTkyTVNJc0ltaDBkSEJ6T2k4dmIyNTBhV1F1YjI1MExtbHZMMk55WldSbGJuUnBZV3h6TDNZeElpd2lZMjl1ZEdWNGRERWlMQ0pqYjI1MFpYaDBNaUpkTENKMGVYQmxJanBiSWxabGNtbG1hV0ZpYkdWRGNtVmtaVzUwYVdGc0lpd2liM1JtSWwwc0ltTnlaV1JsYm5ScFlXeFRkV0pxWldOMElqcGJleUp1WVcxbElqb2lZV2RsSWl3aWRtRnNkV1VpT2lKbmNtVmhkR1Z5SUhSb1lXNGdNVGdpZlYwc0ltTnlaV1JsYm5ScFlXeFRkR0YwZFhNaU9uc2lhV1FpT2lJd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdJaXdpZEhsd1pTSTZJa0YwZEdWemRFTnZiblJ5WVdOMEluMHNJbkJ5YjI5bUlqcDdJbU55WldGMFpXUWlPaUl5TURJd0xUQTNMVEUxVkRBNU9qRXdPakl5V2lJc0luQnliMjltVUhWeWNHOXpaU0k2SW1GemMyVnlkR2x2YmsxbGRHaHZaQ0o5ZlgwPS4rbU03by9hQW1yYnVacWpwUkJ3TzJBSkJ5aTZmS0FhOXpBYUNpN2FhTDhFUFF0dVRpdldOc1k1cEpEak5QVVRMcjNsb0xKby9Ld1VjY3JvclllKzN0dz09Il0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA5OjExOjIzWiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.SdNgkzazOTzdkEMisSZ6Ew0ObiuzFRPzCuvN+yq5jf2FAlv6YiCIDKKvzhJR5m9lPc2f48yg74IXBgWUqr5VmQ=="
                    }
                }
            ]
        }
    }
}
```

### 2.9 query basic message

POST

```
 /api/v1/querypresentation
```

Request body example:

```json
{
	"did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
	"latest":false,
	"remove_after_read":false
}
```

**latest**: true : return the latest message, false:return all messages.

**remove_after_read**:true :remove the message in storage.

Response

```
/api/v1/querybasicmsg
```

Request body example:

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "message_type": 0,
        "content": [
            {
                "@type": "spec/didcomm/1.0/basicmessage",
                "@id": "6ff22592-3476-42a8-8e50-2d76cf771cb7",
                "send_time": "0001-01-01T00:00:00Z",
                "content": "124 adfasdfasefa",
                "~I10n": {
                    "locale": "en"
                },
                "connection": {
                    "my_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
                    "my_router": [
                        "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
                    ],
                    "their_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
                    "their_router": [
                        "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
                    ]
                }
            }
        ]
    }
}
```

### 2.10  send disconnect

Delete connection

POST

```
/api/v1/senddisconnect
```

Request body example:

```json
{
	"@type":"disconnect",
	"@id":"someid",
	"connection": {
        "my_did": "did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
        "my_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"],
        "their_did": "did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
        "their_router":["did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"]
    }
}
```

Response

```json
{
    "code": 0,
    "msg": ""
}
```

