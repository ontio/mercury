# Ontology DID Communication Framework(Nezha)


English | [中文](README_CN.md)

Ontology DID communication framework is a trusted, decentralized peer to peer communication framework, based on DID, we also define protocols for connection , basic message, verifiable credential and presentation proofs.



## Features

- Connection protocol based on DID
- Encrypted channel message exchange 
- Verifiable Credential protocol
- Presentation proof protocol
- Support message routing 
- Scalable to support multiple DID systems 

## Detail Design

[Detail design](https://git.ont.io/ontid/otf/src/master/doc/Detail%20Design.md)



## How to use

### Build binary file

1. git clone project
2. use command  ``` ./make```

### Run agent 

Use CLI

```
./agent-otf
```

CLI parameters:

```
GLOBAL OPTIONS:
   --loglevel <level>  Set the log level to <level> (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel (default: 1)
   --rest-ip value     Set http rest ip addr default:127.0.0.1 (default: "127.0.0.1")
   --http-port value   Set http rest port default:8080 (default: "8080")
   --chain-addr value  Set block chain rpc addr default:127.0.0.1:20334 (default: "http://polaris2.ont.io:20336")
   --https-port value  Set https rest port default:8443 (default: "8443")
   --enable-https      start https restful service
   --enable-package    start package msg
   --help, -h          show help

```

By default , agent will connect polaris (ontology testnet) for querying DID, you can change   ```chain-addr```  to connect mainnet node or you local sync node.



### Tools

We provide some CLI to help create did and other functions. 

Detail please refer to :[Tools cli](https://git.ont.io/ontid/otf/src/master/cmd/manual.md)



### Restful API

Agent also provides restful APIs for clients

Detail please refer to :[Restful API](https://git.ont.io/ontid/otf/src/master/doc/OTF%20RestAPI%20Document.md)