# k8s-mock
`k8s-mock` is a tool designed to simulate the behavior of Kubernetes or OpenShift environments. This tool is particularly useful for testing purposes, allowing clients (kubectl/oc/etc.) to interact with a mock Kubernetes environment without the need for an actual cluster. 

### Overview

### Requirements
Before you run the cluster you will have to:
1. Generate an SSL certificate
```sh
cd ssl/
openssl genpkey -algorithm RSA -out private-key.pem
openssl req -new -key private-key.pem -out certificate.csr
openssl x509 -req -in certificate.csr -signkey private-key.pem -out certificate.crt
cd ../
```
2. Get Resource Definitions (needed to mock a given cluster)
```sh
mkdir -p files/apis/

# YOU MUST BE CONNECTED TO THE CLUSTER YOU WANT TO MOCK!

# this is to generate cache with cluster api definitions in json
kubectl get pods

# copy the api definitions
cd ./files/apis/
cp -R ~/.kube/cache/discovery/[cluster]/* .
cd ../
# get cluster version
kubectl version --output=json > version.json
# proxy the cluster to get http api access
kubectl proxy --port=8001
# get api versions
wget http://localhost:8001/api -O api.json
# get available apis
wget http://localhost:8001/apis -O apis.json
cd ../
# NOTE: this process will be streamlined in the future
```
3. Run
```sh
go run .
```
