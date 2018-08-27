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

## Run a container from your image
```
docker run -d -p 3000:3000 --name app0 myapp:latest
```

Explainations:
```
--name          Name for your container
-p 3000:3000    Port mapping from host 3000 to container 3000
-d              Deamonized (Run the container in Background)
myapp:latest    Name of the image (myapp) & its tag (latest)
```

## List all the running containers
```
docker ps
```
## Expected Output

```
bala:~/environment/fortune_app (master) $ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
92b2ebaa21b5        myapp:latest        "npm start"         2 seconds ago       Up 1 second         0.0.0.0:3000->3000/tcp   app0
```

## Reach the app running inside the container
`curl` `localhost:3000`, Thats where the application is running.
```
bala:~/environment/fortune_app (master) $ curl localhost:3000
In marriage, as in war, it is permitted to take every advantage of the enemy.
```

```
bala:~/environment/fortune_app (master) $ curl localhost:3000
Almost anything derogatory you could say about today's software design
would be accurate.
                -- K.E. Iverson
bala:~/environment/fortune_app (master) $ 
```

## Optionally: Use Public ip
Hint: You can use the EC2 public ip of the EC2 machine which supports the Cloud9 IDE to reach your app.
Since you already exposed 3000 port, make sure inbound rules are set to reach your EC2 machine on PORT 3000


## More commands: Start,Stop

### To stop

Use the container id to stop the container
```
docker stop 92b2ebaa21b5
```

Use the container name to stop the container
```
docker stop app0
```

### To start
Use can either use container name or id to start / stop contianers.
```
docker start app0
```

## Last run multiple copies on different ports:

```
docker run -d -p 3001:3000 --name app1 myapp:latest
docker run -d -p 3002:3000 --name app2 myapp:latest
docker run -d -p 3003:3000 --name app3 myapp:latest
```