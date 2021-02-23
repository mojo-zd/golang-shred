x509证书主要包含 .crt(证书文件), .key(私钥文件), .csr(可以使用openssl基于csr.conf生成), 私钥文件一般不会暴露给用户。
所有的服务器端证书或客户端证书都是基于根证书生成的
### 证书使用流程
#### 创建根证书
1. 新建根证书私钥
2. 创建crs文件
3. 基于步骤1，2创建更证书
#### 基于根证书签发客户端证书
1. 新建私钥文件
2. 创建crs文件
3. 基于根证书创建客户端证书, 从安全角度来讲也可以使用根证书签发中间证书。使用中间证书来创建服务器端、客户端证书

#### 常见的证书类型
1. x509二进制格式证书, 常见后缀.cer、.crt
2. x509文本格式证书, 常见后缀.pem
3. 有些证书只包含服务器公钥  eg: .cer、.crt、.pem
4. 有些证书既包含公钥又包含私钥 eg: .pfx、.p12

#### 自建证书流程
1. 生成自签根证书
```
openssl genrsa -out ca.key 2048 // 生成CA私钥
openssl req -new -key ca.key -out ca.csr // 生成CA证书请求文件
openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt // 生成CA证书
```
2. 生成用户证书
```
openssl genrsa -des3 -out server.key 1024 
openssl req -new -key server.key -out server.csr
openssl ca -in server.csr -out server.crt -cert ca.crt -keyfile ca.key
```

## 安全协议
### TLS
传输层安全协议, 前身为SSL

### SSL
安全套接字