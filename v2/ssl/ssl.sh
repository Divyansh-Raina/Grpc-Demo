#!/bin/bash

SERVER_CN=localhost
MY_SUBJECT="/CN=${SERVER_CN}"


openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -sha256 -days 365 -key ca.key -out ca.crt -subj "/CN=ca"


openssl genrsa -passout pass:1111 -des3 -out server.key 4096

openssl req -passin pass:1111 -new -key server.key -out server.csr -subj ${MY_SUBJECT}

openssl x509 -req -extfile <(printf "subjectAltName=DNS:${SERVER_CN}") -passin pass:1111 -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt 

openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem