# lab-01: first deployment

Let's play with bare pods and pods controlled by a deployment

This a **bare pod** definition:

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

To deploy it:
```
kubectl -n platform101 apply -f resources/speaker-pod.yaml
```

Check if everything is ok:

```
kubectl -n platform101 describe pod/speaker
```

Check the logs:
```
kubectl -n platform101 logs -f pod/speaker
```

### Doing it the right way
What's the problem with this configuration? Easy! It does not offer any garantees if the pod crashes or the node where the pod is running goes down. The pod will
just dissapear and nobody will notice it except for the final users.

That's where upper level controllers come into play. To keep it simple, let's say we *never* run bare pods on production. Instead, we use higher level objects that will ensure our pods are running as we desire. In our case, we will use a **deployment**.

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
