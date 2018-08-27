# Build your own image


**Link for our base image**
- [https://hub.docker.com/_/node/](https://hub.docker.com/_/node/)

## Create a file `Dockerfile`
```
FROM node:alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
```

## Build an image
```
docker build -t myapp:latest .
```

## EXPECTED OUTPUT

```
bala:~/environment/fortune_app (master) $ docker build -t myapp:latest .
Sending build context to Docker daemon  56.32kB
Step 1/7 : FROM node:alpine
alpine: Pulling from library/node
8e3ba11ec2a2: Pull complete 
0cbaf23af500: Pull complete 
c53714e0a6af: Pull complete 
Digest: sha256:e9a83aa3e7ea576b93a21195f2e318ef8075ceccbebaa8f6c333294e9aa51dbd
Status: Downloaded newer image for node:alpine
---> 4e50ad7c0e0b
Step 2/7 : WORKDIR /app
Removing intermediate container 18e6f811388f
---> b75d246ba795
Step 3/7 : COPY package.json .
---> 8aaa70ec520e
Step 4/7 : RUN npm install
---> Running in 04fd854efa40

> fortune-teller@0.1.2 postinstall /app/node_modules/fortune-teller
> node lib/convert

npm notice created a lockfile as package-lock.json. You should commit this file.
npm WARN scs_docker_workshop@1.0.0 No description
npm WARN scs_docker_workshop@1.0.0 No repository field.

added 51 packages from 48 contributors and audited 120 packages in 2.371s
found 0 vulnerabilities

Removing intermediate container 04fd854efa40
---> 76bd0fd7119d
Step 5/7 : COPY . .
---> 3969cb19cf2b
Step 6/7 : EXPOSE 3000
---> Running in 3193faafab0e
Removing intermediate container 3193faafab0e
---> f4335a3a0cf0
Step 7/7 : CMD ["npm", "start"]
---> Running in ac62595f0093
Removing intermediate container ac62595f0093
---> 1eeea450f741
Successfully built 1eeea450f741
Successfully tagged myapp:latest
```


## List all your images
```
docker images
```

## EXPECTED OUTPUT
```
bala:~/environment/fortune_app (master) $ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
myapp               latest              1eeea450f741        3 minutes ago       81.7MB
node                alpine              4e50ad7c0e0b        10 days ago         70.6MB
lambci/lambda       nodejs4.3           6c30c5c1b1e0        8 weeks ago         969MB
lambci/lambda       python2.7           377732dd7a1f        8 weeks ago         974MB
lambci/lambda       python3.6           acf16b1d5297        8 weeks ago         1.1GB
lambci/lambda       nodejs6.10          da301bf4fe34        8 weeks ago         1.02GB
```