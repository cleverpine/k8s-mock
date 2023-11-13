#!/bin/sh

openssl genpkey -algorithm RSA -out private-key.pem
openssl req -new -key private-key.pem -out certificate.csr
openssl x509 -req -in certificate.csr -signkey private-key.pem -out certificate.crt

mv private-key* ssl/
mv certificate* ssl/