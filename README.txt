=====================
WEB SERVER IN GOLANG
=====================

=====================
Golang references
=====================
Read file to string: https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable
Read file and print the file: https://stackoverflow.com/questions/36111777/how-to-read-a-text-file

=====================
GOLANG COMPILER IMAGE
=====================

- Existing image on laptop:
$ docker images | grep golang
golang                                                          1.13          84125009cb55   14 months ago   803MB

$ docker run -ti golang:1.13 /bin/sh
# go version
go version go1.13.12 linux/amd64

- Pull latest version
https://hub.docker.com/_/golang?tab=tags&page=1&ordering=last_updated

$ docker pull golang:alpine3.14
$ docker images |grep golang
golang                                                          alpine3.14    4d3587ec7acf   8 days ago      315MB

$ docker run -ti golang:alpine3.14 /bin/sh
/go # go version
go version go1.17 linux/amd64


=========================================
SET UP GO MODULE
=========================================
- Initialize
$ go mod init goweb1
$ ls 
go.mod

- Create main.go
$ ls
go.mod
main.go

- Run the program once to download dependencies and create go.sum
$ go run .

$ ls 
go.mod
go.sum
main.go

=========================================
CREATE IMAGE WITH YOUR WEB SERVER PROGRAM
=========================================
https://docs.docker.com/language/golang/build-images/
- https://github.com/olliefr/docker-gs-ping (code)

- Create a Dockefile

- Build image
$ docker image build . -t image02-golang:1.0

$ docker images |grep golang
image02-golang                                                  1.0           2d5e04fe6d68   2 minutes ago   467MB

[+] Building 23.6s (12/12) FINISHED                                                                                                                                   
 => [internal] load build definition from Dockerfile                                                                                                             0.0s
 => => transferring dockerfile: 321B                                                                                                                             0.0s
 => [internal] load .dockerignore                                                                                                                                0.0s
 => => transferring context: 2B                                                                                                                                  0.0s
 => [internal] load metadata for docker.io/library/golang:alpine3.14                                                                                             0.0s
 => [1/7] FROM docker.io/library/golang:alpine3.14                                                                                                               0.0s
 => [internal] load build context                                                                                                                                0.0s
 => => transferring context: 6.63kB                                                                                                                              0.0s
 => [2/7] WORKDIR /app                                                                                                                                           0.0s
 => [3/7] COPY go.mod ./                                                                                                                                         0.0s
 => [4/7] COPY go.sum ./                                                                                                                                         0.0s
 => [5/7] RUN go mod download                                                                                                                                   19.1s
 => [6/7] COPY *.go ./                                                                                                                                           0.0s
 => [7/7] RUN go build -o /myapp                                                                                                                                 3.3s
 => exporting to image                                                                                                                                           1.0s
 => => exporting layers                                                                                                                                          1.0s
 => => writing image sha256:2d5e04fe6d686701288fc5eb04fddd1fec5e0cbc92a59c2019f6b165e895de50                                                                     0.0s
 => => naming to docker.io/library/image02-golang:1.0                                                                                                            0.0s

=================================
RUN THE IMAGE AS A CONTAINER --> AFTER 'messages' golang code was introduced
=================================
- Create message file directory on host
$ mkdir /tmp/hostmessages --> to create message1.txt file per main.go code

$ echo "This is from message file again and again" > /tmp/hostmessages/message1.txt

- Run container with volume mounted 
$ docker run --name image02-golang1 -v /tmp/hostmessages:/tmp/messages -d --rm -p 8081:8081/tcp image02-golang:1.0

- Verify
curl http://localhost:8081/message
{"message":"This is from message file again and again\n"}

=================================
RUN THE IMAGE AS A CONTAINER --> BEFORE 'messages' golang code was introduced
=================================
https://docs.docker.com/engine/reference/commandline/run/
Example: docker run -p 127.0.0.1:80:8080/tcp ubuntu bash --> This binds port 8080 of the container to TCP port 80 on 127.0.0.1 of the host machine. 

- Run the image (quick test)
$ docker run -d --rm image02-golang:1.0

$ docker ps -a
CONTAINER ID   IMAGE                 COMMAND                  CREATED          STATUS                        PORTS      NAMES
86d48f4f75c3   image02-golang:1.0    "/myapp"                 15 seconds ago   Up 15 seconds                 8081/tcp   recursing_snyder

$ docker stop 86d48f4f75c3
$ docker rm 86d48f4f75c3

- Run the image
$ docker run --name image02-golang1 -d --rm -p 8081:8081/tcp image02-golang:1.0 
b9446820177c8f6d229d30d2b07b39a51a0b9ee16f37f11f0b730da012773794

$ docker ps -a
CONTAINER ID   IMAGE                COMMAND    CREATED         STATUS         PORTS                                       NAMES
b9446820177c   image02-golang:1.0   "/myapp"   4 seconds ago   Up 2 seconds   0.0.0.0:8081->8081/tcp, :::8081->8081/tcp   image02-golang1

-- Verify
$ curl localhost:8081
{"message":"Welcome!"}

==================================
GET INSIDE THE CONTAINER AND LOOK AROUND
==================================

$ docker exec -ti image02-golang1 /bin/sh
/app # ls -l
total 16
-rw-r--r--    1 root     root            96 Aug 25 09:32 go.mod
-rw-r--r--    1 root     root          5389 Aug 25 09:32 go.sum
-rw-r--r--    1 root     root          1036 Aug 25 09:00 main.go

/app # cd /
/ # ls -l
total 8996
drwxr-xr-x    1 root     root          4096 Aug 25 09:39 app
drwxr-xr-x    1 root     root          4096 Aug 17 01:38 bin
drwxr-xr-x    5 root     root           340 Aug 25 10:01 dev
drwxr-xr-x    1 root     root          4096 Aug 25 10:01 etc
drwxrwxrwx    1 root     root          4096 Aug 25 09:38 go
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 home
drwxr-xr-x    1 root     root          4096 Aug 17 01:38 lib
drwxr-xr-x    5 root     root          4096 Aug  5 12:25 media
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 mnt
-rwxr-xr-x    1 root     root       9143655 Aug 25 09:39 myapp
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 opt
dr-xr-xr-x  199 root     root             0 Aug 25 10:01 proc
drwx------    1 root     root          4096 Aug 25 10:01 root
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 run
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 sbin
drwxr-xr-x    2 root     root          4096 Aug  5 12:25 srv
dr-xr-xr-x   13 root     root             0 Aug 25 10:01 sys
drwxrwxrwt    1 root     root          4096 Aug 25 09:39 tmp
drwxr-xr-x    1 root     root          4096 Aug 17 01:38 usr
drwxr-xr-x    1 root     root          4096 Aug  5 12:25 var

/ # cd /tmp
/tmp # ls -l
total 0

/tmp # uname -a
Linux 7d7872d0b166 5.10.25-linuxkit #1 SMP Tue Mar 23 09:27:39 UTC 2021 x86_64 Linux

/ # exit


