# Stage1:  Pull & Run images as containers

## Run directly

- hello-world
    `docker run hello-world`

    ```
    bala:~/environment/fortune_app (master) $ docker run hello-world
    Unable to find image 'hello-world:latest' locally
    latest: Pulling from library/hello-world
    9db2ca6ccae0: Pull complete 
    Digest: sha256:4b8ff392a12ed9ea17784bd3c9a8b1fa3299cac44aca35a85c90c5e3c7afacdc
    Status: Downloaded newer image for hello-world:latest

    Hello from Docker!
    This message shows that your installation appears to be working correctly.

    To generate this message, Docker took the following steps:
    1. The Docker client contacted the Docker daemon.
    2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
        (amd64)
    3. The Docker daemon created a new container from that image which runs the
        executable that produces the output you are currently reading.
    4. The Docker daemon streamed that output to the Docker client, which sent it
        to your terminal.

    To try something more ambitious, you can run an Ubuntu container with:
    $ docker run -it ubuntu bash

    Share images, automate workflows, and more with a free Docker ID:
    https://hub.docker.com/

    For more examples and ideas, visit:
    https://docs.docker.com/engine/userguide/

    ```

## Pull & Run 
- docker/whalesay
    `docker pull docker/whalesay`

    `docker run docker/whalesay cowsay hello`
    
    ```
    bala:~/environment/fortune_app (master) $ docker run docker/whalesay cowsay hello
        _______ 
        < hello >
        ------- 
            \
            \
            \     
                            ##        .            
                    ## ## ##       ==            
                ## ## ## ##      ===            
            /""""""""""""""""___/ ===        
        ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~   
            \______ o          __/            
                \    \        __/             
                \____\______/   
    ```


# Stage2:  Build your own image

Download the App from the following repo.
We will be using this demo fortune_app for this docker workshop. Its a simple Node.js app, which tell random sfortune. The app will be started on `PORT 3000` by default. But you can still override by passing PORT env variable.
```
git clone https://github.com/balaprasanna/fortune_app.git
```

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

### EXPECTED OUTPUT

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

### EXPECTED OUTPUT
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
### Expected Output

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

## Pushing images to Docker Hub

Step 1: You need to create a docker hub accont for this. Please register if you dont have an account.
- [https://hub.docker.com/](https://hub.docker.com/)
  
Step 2: Do a docker login in command line.
```
docker login
```
When promted please enter username & password.

### Expected Output:
```
bala:~/environment/fortune_app (master) $ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: <your username>
Password: 
Login Succeeded
```

Step 3: Tag your image with your username infront
```
docker tag myapp:latest balanus/fortune-app:latest
```
Step 4: Verify the tagged image by running `docker images`
```
bala:~/environment/fortune_app (master) $ docker images
REPOSITORY            TAG                 IMAGE ID            CREATED             SIZE
myapp                 latest              1eeea450f741        29 minutes ago      81.7MB
balanus/fortune-app   latest              1eeea450f741        29 minutes ago      81.7MB
node                  alpine              4e50ad7c0e0b        10 days ago         70.6MB
lambci/lambda         nodejs4.3           6c30c5c1b1e0        8 weeks ago         969MB
lambci/lambda         python2.7           377732dd7a1f        8 weeks ago         974MB
lambci/lambda         python3.6           acf16b1d5297        8 weeks ago         1.1GB
lambci/lambda         nodejs6.10          da301bf4fe34        8 weeks ago         1.02GB
```

Step 5: Finally, push the image using `docker push`

```
docker push balanus/fortune-app:latest
```

### Expected Output
```
bala:~/environment/fortune_app (master) $ docker push balanus/fortune-app:latest
The push refers to repository [docker.io/balanus/fortune-app]
f41d5bed9291: Pushed 
d9df74bc7450: Pushed 
c8f8f4a9df78: Pushed 
48cb9c80d7bb: Pushed 
287ef32bfa90: Mounted from library/node 
ce291010afac: Mounted from library/node 
73046094a9b8: Mounted from library/node 
latest: digest: sha256:c0112775a24091482665ec43fd6fda0b854dd0417ade411aba458a3b6c01955c size: 1784
```


### Note:
```
By default, all docker images that your push will be a private image.
Please login to your docker hub account and change it to a public repo.
From Private repo to Public repo. So that , we can use this image to run in K8S later.
```