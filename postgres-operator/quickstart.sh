#!/bin/bash 
echo "This script replaces quickstart-for-gke.sh"
echo ""
echo "This script will:"
echo "- Create the pgo namespace"
echo "- Apply postgres-operator.yml"
echo "- install the client"
echo ""
kubectl create namespace pgo
kubectl apply -f postgres-operator.yml
echo "wait until deploy job is complete..."
kubectl wait --for=condition=complete --timeout=180s -n pgo job/pgo-deploy
./client-setup.sh
