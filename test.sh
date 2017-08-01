#!/usr/bin/env bash

echo curl http://localhost:8080
curl http://localhost:8080

curl http://localhost:8080/ping
echo
curl http://localhost:8080/api/programs/hello
echo
