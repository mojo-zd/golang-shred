## csr-viewer 只读权限
### csr-viewer
1. 生成证书
```
openssl genrsa -out csr-viewer.key 2048
openssl req -new -key csr-viewer.key -subj "/CN=csr-viewer" -out csr-viewer.csr
```

2. 生成`request`内容
```
cat csr-viewer.csr | base64 | tr -d "\n"
```

3. 创建`csr`
`request`为上一步返回内容
```
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: csr-viewer
spec:
  request: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1dqQ0NBVUlDQVFBd0ZURVRNQkVHQTFVRUF3d0tZM055TFhacFpYZGxjakNDQVNJd0RRWUpLb1pJaHZjTgpBUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTVVBUWY4MmsrUCtGVityUnBEOVBDRFVYU1BxQXg4UGsrRktEbkhLCmowRmw1VjFveW1LaGhmeGtJV2xIaTJxSllhNlZ0cDNFOGpYK3JoenAwN0Uybyt3VndKSHl5MTVyRzN5aHJjSloKeXpIUjh6ODNSOHpmSVpUNGdVL3lkc3VKaFEzbFM1aE9kVFE1K2dVMEl6amtZVHpKanRZY2dXcGtNQmswWFpzSwo4M3NZK0EyNHArOW5NajRPcTcyc1dnemdybHJ0NnA0T2YxVTZkY29lZ2dMM29JU1MwS3d5V2pBemFncERvN2ZZCnlkdlJqRG1vQzlCRnlwN3ZUTER5ZlhsV1VPdHZ3aUUvS281UHcvNktlZmVOblloY0JrTDViS2diZUxlOWNNUVQKdTl5bXFSaXU0a3dORDlubjF4MkViR3orMzRaSTRjVmw0a0plUmhZMXpJNXNzcWtDQXdFQUFhQUFNQTBHQ1NxRwpTSWIzRFFFQkN3VUFBNElCQVFBdWVZeS9pMjB2VEhSSG1vRnZlZzdMdHlYWFRJNVBmWjZKYmRrNVNvNVpMOEdNCk9DbW1kdzlOZjJVd1NUT1FRYitlZkV6b3NMdjdISTFRNUliS0Z2a2dhRXJRY0ZBQkE3UnNFcmhiSEthcTR1aEEKWUhKWkN3T3JSTzJKcEt4ckhHTGVwWnpmR085UkwwbGpKTU53VUVHZ285SGk2a1d3MGwyVFc0RHFmVDZmSGRqaQp2ZmExTUNPbDhiV3JzMm0xUyt0S2JudHRWYTRPS0x6cmQ4Z3dkVjR5eEF3NmhkbWlndnhXdWZSSjFkWlBxNEN6CkxHK3RIQVZWWkMvNDZYUnUxbkZVYkNlc0lyUVB0TVFEVzJkbVNTQWV3Y3BHZVFhaUtsUWdWUWF1U2tUbUlYYWYKM09kdU1FSktxdGdHNU5BRjBEZFNiNVZSdXcwa1A3em1ySDU4TG4zWAotLS0tLUVORCBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0K
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
kubectl get csr
NAME     AGE   SIGNERNAME                            REQUESTOR                 CONDITION
csr-viewer   10s   kubernetes.io/kube-apiserver-client   100015915015-1618216441   Pending
```

4. `approve csr`
```
kubectl certificate approve csr-viewer
kubectl get csr
NAME         AGE     SIGNERNAME                            REQUESTOR                 CONDITION
csr-viewer   9m15s   kubernetes.io/kube-apiserver-client   100015915015-1618216441   Approved,Issued
```

5. 导出证书
```
kubectl get csr csr-viewer -o jsonpath='{.status.certificate}'| base64 -d > csr-viewer.crt
**查看证书有效期**
openssl x509 -text -in csr-viewer.crt
可以看到 Validity 为一年；
```

6. `clusterrole/clusterrolebinding`角色绑定

```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cr-csr-viewer
rules:
- apiGroups:
  - ""
  resources:
  - namespaces
  - configmaps
  - pods
  - persistentvolumeclaims
  - persistentvolumes
  - nodes
  - services
  - events
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - khaos.tencentcloud.com
  resources:
  - clusters
  - clusterpresets
  verbs:
  - get
  - list
  - watch
```

`clusterrolebinding`
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: crb-csr-viewer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cr-csr-viewer
subjects:
- kind: User
  name: csr-viewer
  apiGroup: rbac.authorization.k8s.io
```

7. 设置 `~/.kube/config`
```
kubectl config set-credentials csr-viewer --client-key=csr-viewer.key --client-certificate=csr-viewer.crt --embed-certs=true
kubectl config set-context csr-viewer --cluster=cls-6pznu3gw --user=csr-viewer
kubectl config use-context csr-viewer
```

8. 验证
```
kubectl get clsv2
Error from server (Forbidden): clusters.khaos.tencentcloud.com is forbidden: User "gateway" cannot list resource "clusters" in API group "khaos.tencentcloud.com" at the cluster scope
kubectl get po -n khaos-system
NAME                                       READY   STATUS    RESTARTS   AGE
chai-discovery-c747b5bcc-q9xtd             1/1     Running   0          26d
chai-oss-857dffbb87-bbw8g                  1/1     Running   0          5d6h
khaos-apiserver-95f86df77-f5jlz            1/1     Running   0          11d
khaos-apiserver-95f86df77-n2twl            1/1     Running   0          11d
```
