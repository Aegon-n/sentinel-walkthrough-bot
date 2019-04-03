#!/bin/sh

cd /root;
nohup mongod >> /dev/null & sleep 1;
nohup dante >> /dev/null &
nohup dumb-init sockd >> /dev/null &
gaiacli advanced rest-server --chain-id Sentinel-testnet-1.1 --node http://tm-lcd.sentinelgroup.io:26657 --home /root/.sentinel & sleep 1;
python app.py
