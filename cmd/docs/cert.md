#### 证书使用方式
1. 生成root ca
2. 安全角度讲生成intermediate ca
3. 签名服务器端使用的key、cert

#### 常见的证书类型
1. x509二进制格式证书, 常见后缀.cer、.crt
2. x509文本格式证书, 常见后缀.pem
3. 有些证书只包含服务器公钥  eg: .cer、.crt、.pem
4. 有些证书既包含公钥又包含私钥 eg: .pfx、.p12