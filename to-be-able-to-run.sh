#!/bin/sh

oc new-project test
sleep 0.2
oc create quota cpu-quota --namespace=test --hard=cpu=1
sleep 0.2
oc create secret generic my-secret --from-literal=username=myusername --from-literal=password=mypassword --namespace=test