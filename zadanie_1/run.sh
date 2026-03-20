#!/bin/bash

set -e

echo "Building Docker image..."
docker build -t pascal-sorting-app .
echo "Docker image built."

echo "Launching application..."
docker run --rm pascal-sorting-app

