# Hello Codefresh

## 1.) Create an app

Let's create a Hello World! app in [Go](https://golang.org/) to keep things simple :D

```go
// hello_codefresh.go
package main

import (
  "fmt"
  "net/http"
  "runtime"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(rw, `
      <h1>Hello World!</h1>
      <p>Greetings from %s with an %s CPU.
      <p>Check out the <a href="https://github.com/DrewBro-Designs/hello_codefresh">source code</a>.
      `, runtime.GOOS, runtime.GOARCH)
}

func main(){
  http.HandleFunc("/", helloHandler)
  http.ListenAndServe(":3000", nil)
}
```

## 2.) Create a [Docker](https://www.docker.com/) image  

Create the Dockerfile  

```Dockerfile
FROM golang

ADD . /go/src/github.com/DrewBro-Designs/hello_codefresh
RUN go install github.com/DrewBro-Designs/hello_codefresh
ENTRYPOINT /go/bin/hello_codefresh

EXPOSE 3000
```

Build the image

```sh
docker build -t hello_codefresh .
```

Run a container with the image we just built
```
docker run -p 3000:3000 hello_codefresh
```
## 3.) Add Source Control

If you haven't already, initialize a git repo, and push your source to somewhere like GitHub

```
git init
git add -A
git commit -m 'I <3 my mom!'
git remote add origin git@github.com:<github-username>/<reponame>.git
git push origin master
```

## 4.) Get started with [Codefresh](https://docs.codefresh.io/docs/create-an-account)

You can sign up for an account with your Github, Bitbucket, or Gitlab account (you will need to sign in with the account from the repo host you want to use).  
  
Next up:  
  * [Create a Basic Pipeline](https://docs.codefresh.io/docs/getting-started-create-a-basic-pipeline)  
  * [Push to Docker registry](https://docs.codefresh.io/docs/push-image-to-a-docker-registry)  

## 5.) Add [Integration Tests](https://docs.codefresh.io/docs/integration-tests)

We want to run some integration tests, so first let's [Create a Composition](https://docs.codefresh.io/docs/create-composition).

docker-compose.yml
```
version: '3'
services:
  hello_codefresh:
    build: .
    ports:
    - "3000:3000"
```
Codefresh compositions don't support port mapping. Also, it will try to guess from your service name what the registry image is called.  
  
In this case, we should modify the codefresh composition compose file to
```
version: '3'
services:
  hello_codefresh:
    ports:
      - '3000'
    image: 'drewbrodesigns/hellocodefresh:master'
```

## 6.) Get Ready to Deploy

Let's get ready to [deploy to Kubernetes](https://docs.codefresh.io/docs/get-ready-to-deploy)  
  
Sign into [console.cloud.google.com](https://console.cloud.google.com/cloud-resource-manager) and create a new project.  
  
Enable [Container Engine API](https://console.cloud.google.com/apis/api/container/overview?project=codefresh-hellogo)  
  
Enable [Compute Engine API](https://console.cloud.google.com/apis/api/compute_component/overview)  
  
To create a Kubernetes cluster, visit this [short guide](https://cloud.google.com/container-engine/docs/quickstart).  

```shell
# Set default compute zone
gcloud config set compute/zone us-west1-a

# create cluster
gcloud container clusters create hello-codefresh-cluster # this takes a long time (5 minutes+ ?)

# Run container
kubectl run hello-go --image=drewbrodesigns/hellocodefresh:master --port=3000


# Expose the container
kubectl expose deployment hello-codefresh --type="LoadBalancer" 

# Get public IP
kubectl get service hello-codefresh # takes about 1 minute to change from pending to having an IP

# the external IP will resolve about 5 minutes after creating it
```

**Note**: Clean up the cluster (you will be billed)

```
kubectl delete service hello-codefresh
gcloud container clusters delete hello-codefresh-cluster
```
  
Next, [deploy to GKE](https://docs.codefresh.io/docs/codefresh-kubernetes-integration-beta)






