# Forward and Encrypt message

## 1. 背景

通常 agent 直接无法直接通信,需要通过中继和转发才能将消息传递到目标 agent, 在传递过程中,原数据不能泄露给中继agent.

![](./images/forward1.png)



## 2. 设计方案

agent直接的消息格式定义为 Envelope:

```json
{
    "message":{
        "data":"encrypt the message data",
        "msgtype": type of message(int),
        "sign":"signature of data"
    },
    "connection":{
        "data":"encrypt connection data",
        "sign":"signature of data"
    },
    "fromdid":"sender did",
    "todid":"receiver did"
}
```

connection的格式为:

```json
{
    "my_did": "did:ont:alice",
    "my_router":["did:ont:cloudA#serviceid","did:ont:cloundB#serviceid"],
    "their_did": "did:ont:Bob",
    "their_router":["did:ont:cloudD#serviceid","did:ont:cloudC#serviceid"]
}
```

说明:

1. my_router 和 their_router必须非空, my_router的第一个元素为代理节点,其他为转发节点, their_router 的第一个元素为接收节点,其他为转发节点
2. ​





