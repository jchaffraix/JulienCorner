#!/bin/bash
set -xeu -o pipefail

export REGION="us-east1"
export PROJECT="juliencorner"
export SERVICE_NAME="$PROJECT"
export REPOSITORY="us-east1-docker.pkg.dev"
export ARTIFACTORY_REPO="docker"
export IMAGE_NAME="webapp"
export TAG=$(date +%Y%m%d-%H%M%S)
export LABEL="${REPOSITORY}/${PROJECT}/${ARTIFACTORY_REPO}/${IMAGE_NAME}:${TAG}"

DOCKER=`which podman || true`
if [ -z "$DOCKER" ]; then
  DOCKER=`which docker || true`
  if [ -z "$DOCKER" ]; then
    echo "No valid container engine. Install podman/docker..."
    exit 1
  fi
fi
echo "Using "$DOCKER" to manage containers"

# Check that we pass an authorization for a service account and log into it.
# Else for CLI invokation, we just use any existing credentials.
export SERVICE_ACCNT_FILE="${SERVICE_ACCNT_FILE:-}"
if [[ "$SERVICE_ACCNT_FILE" != "" ]]; then
  gcloud auth activate-service-account "${SERVICE_ACCNT}" --key-file="$SERVICE_ACCNT_FILE"
  gcloud config set project "$PROJECT"
  gcloud auth print-access-token --quiet \
    --impersonate-service-account "${SERVICE_ACCNT}" | "$DOCKER" login \
    -u oauth2accesstoken \
    --password-stdin "$REPOSITORY"
else
  ACTUAL_PROJECT=`gcloud info|grep Project|cut -d[ -f2`
  if [ "$PROJECT]" != "$ACTUAL_PROJECT" ]; then
    echo "Project mismatch, use gcloud init to select $PROJECT"
    exit 1
  fi

  gcloud auth print-access-token --quiet | "$DOCKER" login \
    -u oauth2accesstoken \
    --password-stdin "$REPOSITORY"
fi

$DOCKER build -t "$LABEL" .
$DOCKER push "$LABEL"

echo "Pushed image successfully"
echo ""

echo "Deploying the new container"
# TODO: I should probably push to a staging env first.
gcloud run deploy "$SERVICE_NAME" \
  --image="$LABEL" \
  --region="$REGION" \
  --revision-suffix="$TAG"

echo "The new container is deployed!"
echo "You can find more detail here: https://console.cloud.google.com/run/detail/${REGION}/${SERVICE_NAME}"
