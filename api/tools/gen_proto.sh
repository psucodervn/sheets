#!/usr/bin/env bash

protoc -I . proto/*.proto --go_out=plugins=grpc:.
