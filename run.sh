#!/bin/bash

cd /app/golang-cms
echo "Current version is $(git tag --sort=committerdate | tail -1)"
bee run
