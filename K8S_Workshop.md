This part of the workshop , assumes you already have a working K8S Cluster.
We will be using a the same demo App that we used for docker workshop.
You can find reference kube manifest files in the following repo.

1. Create a NameSpace
```
apiVersion: v1
kind: Namespace
metadata:
  name: fortune-app-ns
```
Save this in a file `fortune-app-ns.yml`

### Now apply the manifest to K8S Cluster
```
kubectl apply -f fortune-app-ns.yml
```

2. Deploy the app under the above created namespace.
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: fortune-app-nodejs
  namespace: fortune-app-ns
  labels:
    app: fortune-app
    lang: nodejs
spec:
  replicas: 2
  selector:
    matchLabels:
      app: fortune-app
  template:
    metadata:
      labels:
        app: fortune-app
        lang: nodejs
    spec:
      containers:
        - image: balanus/fortune-app:latest
          name: fortune-app
          env:
          - name: PORT
            value: "3000"
          ports:
            - containerPort: 3000
              name: fortune-app
          imagePullPolicy: Always

```

Save this in a file `fortune-app-deployment.yml`

### Now apply the manifest to K8S Cluster
```
kubectl apply -f fortune-app-deployment.yml
```

### Verify the deployment:
```
kubectl get deployment --namespace = fortune-app-ns
```

### Also you can look at the no of pods,created for this deployment
```
kubectl get pods --namespace = fortune-app-ns
```

Now to reach the APP running inside the cluster, you need to create a SVC.

3. Expose the APP via K8S Service.
```
apiVersion: v1
kind: Service
metadata:
  name: fortune-app-svc-nodejs
  namespace: fortune-app-ns
  labels:
    app: fortune-app-svc
    lang: nodejs

spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 3000
      protocol: TCP
  selector:
    app: fortune-app
    lang: nodejs
```

Service type must be a LoadBalancer in order to get a public dns to reach your app.
Also remember to destory this LoadBalancer after used.

### Verify the service:
```
kubectl get svc --namespace = fortune-app-ns
```

4. Two approaches to Scale:

- Scale the deployment using `kubectl scale`
```
kubectl scale deployment/fortune-app-nodejs --replica=4 --namespace=fortune-app-ns
```

- Scale the deployment by updating the deployment manifest
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: fortune-app-nodejs
  namespace: fortune-app-ns
  labels:
    app: fortune-app
    lang: nodejs
spec:
  replicas: 4
  selector:
    matchLabels:
      app: fortune-app
  template:
    metadata:
      labels:
        app: fortune-app
        lang: nodejs
    spec:
      containers:
        - image: balanus/fortune-app:latest
          name: fortune-app
          env:
          - name: PORT
            value: "3000"
          ports:
            - containerPort: 3000
              name: fortune-app
          imagePullPolicy: Always

```

Update the content of `fortune-app-deployment.yml`. And Apply it using kubectl
```
kubectl apply -f fortune-app-deployment.yml
```