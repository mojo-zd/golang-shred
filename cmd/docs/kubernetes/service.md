## headless service

service被创建以后并不会创建vip, 而是通过dns直接访问ep。

### 示例

```
// edit svc并添加app=web的label
kubectl create svc clusterip sts-demo clusterip="None"

kubectl apply -f sts.yaml
// 测试
docker run -i --tty --image=busybox:1.28.4 dns-test --restart=Never --rm /bin/sh
nslookup web-0.sts-demo
```

```
// statefulset file  sts.yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app: web
  name: web
spec:
  serviceName: sts-demo
  replicas: 3
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
      - image: nginx:stable
        imagePullPolicy: IfNotPresent
        name: nginx
```