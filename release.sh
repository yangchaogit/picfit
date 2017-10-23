#!/bin/bash

bee pack -be GOOS=linux -exr="server" -exr="application" -exr="storage" -exr="release.sh" -exr="run.dev.sh" -exr="bin" -exr="config" -exr="engine" -exr="image" -exr="kvstore" -exr="middleware" -exr="payload" -exr="signature" -exr="views" -exr="Dockerfile.build" -exr="Makefile" -exr="README.rst" -exr="main.go" -exr="Gopkg.lock" -exr="Gopkg.toml" -exr="vendor" -exr="CHANGES.rst" -exr="Dockerfile" -exr="LICENSE" -exr="Vagrantfile" -exr="constants" -exr="errs" -exr="hash" -exr="logger" -exr="modd.conf" -exr="provisioning" -exr="util"

