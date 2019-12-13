# go-serve

https://hub.docker.com/repository/docker/youpy/go-serve

## Usage

```Dockerfile
FROM node:alpine as builder

COPY . ./
RUN yarn && yarn build

FROM youpy/go-serve
COPY --from=0 /build ./build/
```

```
$ docker build -t image-name .
$ docker run --rm -p 5000:5000 image-name
```
