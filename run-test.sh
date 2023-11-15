#!/bin/sh

echo "Logging in..."
oc login --server=https://localhost:7777 --token=sha256~iqhrpiwhtpiwehtipwehtwe --insecure-skip-tls-verify=true

echo "Creating projects..."
oc new-project test1
oc new-project test2
oc new-project test3
oc new-project test4

echo "Listing projects..."
oc projects

echo "Setting default project..."
oc project test2