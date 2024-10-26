#!/bin/bash
nohup go run query_project.go  -f etc/query_project-api.yaml > server.log 2>&1 &
# nohup ./user_crypto -f etc/user_crypto-api.yaml > server.log 2>&1 &
