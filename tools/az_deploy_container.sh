#!/bin/bash
set -xeu -o pipefail

export API_NAME="juliencorner"
export ACR_NAME="juliencorner"

docker build -t "${ACR_NAME}.azurecr.io/${API_NAME}":latest .

# TODO: Can I check that I am already logged in?
az acr login -n "$ACR_NAME"
docker push "${ACR_NAME}.azurecr.io/${API_NAME}":latest
