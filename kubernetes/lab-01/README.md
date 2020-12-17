# LAB-01: first deployment

Let's play with bare pods and pods controlled by a deployment

This a bare pod definition:

```
apiVersion: v1
kind: Pod
metadata:
  name: speaker
  namespace: platform101
  labels:
    project: platform101
spec:
  containers:
    - name: speaker
      image: platform101/speaker
```

This is how we would properly deploy it on production:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: speaker
  namespace: platform101
  labels:
    project: platform101
    lab: lab01
    app: speaker
spec:
  replicas: 2
  selector:
    matchLabels:
      app: speaker
  template:
    metadata:
      labels:
        app: speaker
    spec:
      containers:
      - name: speaker
        image: platform101/speaker
```
