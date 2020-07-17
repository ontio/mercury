# mercury

# Ontology DID Communication

[TOC]

## 简介

[Ontology DID Communication Framework](../doc/Detail Design.md)



## 使用

### build project

1. clone 项目
2. 在项目根目录下执行 ```./make```


### Start project

启动命令

```./agent-otf```

参数:

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

参数说明:

**loglevel**:日志级别

**rest-ip**:启动rest服务的ip地址,默认为127.0.0.1

**http-port**:rest服务的端口, 默认为8080

**chain-addr**:链接ontology节点的rpc地址,默认为polaris测试网地址

**enable-https**:是否开启https

**enable-package**:是否开启消息加密



### did 和 client

请参照: [http_cmd](https://git.ont.io/ontid/otf/src/master/cmd/manual.md)



