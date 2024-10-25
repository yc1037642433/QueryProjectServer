#!/bin/bash
nohup go run user_crypto.go  -f etc/user_crypto-api.yaml > server.log 2>&1 &
# nohup ./user_crypto -f etc/user_crypto-api.yaml > server.log 2>&1 &
