# lab-01: internal services

## Context

The engineering team is developing an annoying service called **speaker-api** that cannot stop talking. This service exposes a REST api that must be accessible from other services within the cluster (internally). There is only one endpoint to retrieve the total amount of words it said.

```
GET /words
```

The docker image is `platform101/speaker-api` and want to run 2 replicas of this container in production for high availability reasons. Create a manifest file and apply it with `kubectl`. In the containers section, remember to expose the port 8080 as this is where the service is listening to :)
```
containers:
  - name: speaker-api
    ports:
      - containerPort: 8080
```

Check if everything is ok:
```
kubectl -n platform101 describe deployment/speaker-api
```

Get running replicas:
```
kubectl -n platform101 get pods -l app=speaker-api

```

Check the logs:
```
kubectl -n platform101 logs -f -l app=speaker-api
```

Now, we have to expose the pods inside the cluster at port 80. We will use a ClusterIP service for that

```
apiVersion: v1
kind: Service
metadata:
  name: speaker-api
spec:
  type: ClusterIP # by default
  selector:
    app: speaker-api
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

Check the kubernetes service is working fine:
```
# Forward your local port 8080 to destination port 80
kubectl -n platform101 port-forward service/speaker-api 8080:80

# Check it!
curl localhost:80/words
```

## References
- Kubernetes controllers: https://kubernetes.io/docs/concepts/architecture/controller/
- Pods: https://kubernetes.io/docs/concepts/workloads/pods
- Deployments: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- Services: https://kubernetes.io/docs/concepts/services-networking/service/
