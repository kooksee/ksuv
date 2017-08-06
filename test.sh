#!/usr/bin/env bash

echo curl http://localhost:8080
curl http://localhost:8080

curl http://localhost:8080/ping

curl -X POST http://localhost:8080/api/programs \
-d {"name":"ss","cur_dir":"~","cmd":"ls -alh"}
# -d name="sss" \
#-d cur_dir="~" \
#-d cmd="ls -alh"




