# lab-01: bare pods and deployments

## Context

The engineering team is developing an annoying service that cannot stop talking. To deploy it, we can use a **bare pod** definition:

```
apiVersion: v1
kind: Pod
metadata:
  name: speaker
  namespace: platform101
  labels:
    project: platform101
    lab: lab01
    app: speaker
spec:
  containers:
    - name: speaker
      image: platform101/speaker
```

**Manifest »** [speaker-pod.yaml](./resources/speaker-pod.yaml)

Deploy it:
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

Delete the pod and see how no replacement is created:
```
kubectl -n platform101 delete pod/speaker
```

### Doing it the right way
What's the problem with this configuration? Easy! It does not offer any garantees if the pod crashes or the node where the pod is running goes down. The pod will
just dissapear and nobody will notice it except the final users.

That's where upper level controllers come into play. To keep it simple, let's say we *never* run bare pods on production. Instead, we use higher level objects that will ensure our pods are running as we desire, even in the case of a failure . So, we will use a **deployment** to run our speaker container.

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
  replicas: 1
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
**Manifest »** [speaker-deployment.yaml](./resources/speaker-deployment.yaml)

Deploy it:
```
kubectl -n platform101 apply -f resources/speaker-deployment.yaml
```

Check if everything is ok:
```
kubectl -n platform101 describe deployment/speaker
```

Get running replicas:
```
kubectl -n platform101 get pods -l app=speaker
```

Check the logs:
```
kubectl -n platform101 logs -f -l app=speaker
```

Delete one replica:
```
kubectl -n platform101 delete pod <POD_NAM>
```

Check a replacement is started:
```
kubectl -n platform101 get pods -l app=speaker
```

Play a little bit with the desired replicas:
```
kubectl -n platform101 scale deployment/speaker --replicas=<NUMBER>
```

## References
- Kubernetes controllers: https://kubernetes.io/docs/concepts/architecture/controller/
- Pods: https://kubernetes.io/docs/concepts/workloads/pods
- Deployments: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
