# To run the app locally.

```
npm install
npm start
```

# For Docker

## Build your own image
```
docker build -t fortune-app:latest .
```

```
docker run -d -it --name app -p 3000:3000 fortune-app:latest
```

## Test your APP.
```
curl localhost:3000
```


## Push your image to docker HUB
Signup for dockerhub if you dont have an account.

Step :1
```
docker login
```

Step :2
```
docker tag fortune-app  balanus/fortune-app
```

Step :3
```
docker push  balanus/fortune-app
```