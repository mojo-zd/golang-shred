#### 证书使用方式
1. 生成root ca
2. 安全角度讲生成intermediate ca
3. 签名服务器端使用的key、cert

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
