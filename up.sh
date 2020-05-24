#!/usr/bin/env bash
# usage: up.sh api v0.2.4
#ssh dev@167.71.214.233 "cd /home/dev/sheets && ./up.sh $1 $2"

ssh ec2-user@ec2-52-76-145-82.ap-southeast-1.compute.amazonaws.com "cd /home/ec2-user/sheet && ./up.sh $1 $2"
