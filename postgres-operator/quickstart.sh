#!/bin/bash 
echo "This script replaces quickstart-for-gke.sh"
echo ""
echo "This script will:"
echo "- Create the pgo namespace"
echo "- Apply postgres-operator.yml"
echo "- install the client"
echo ""
wget https://github.com/CrunchyData/postgres-operator-examples/archive/refs/heads/main.zip
unzip main.zip
kubectl apply -k postgres-operator-examples-main/kustomize/install
echo "wait until pod is ready"
sleep 15
kubectl -n postgres-operator wait pods   --selector=postgres-operator.crunchydata.com/control-plane=postgres-operator   --field-selector=status.phase=Running --for=condition=ready

echo "pgo 4 is not compatible anymore with newer kubernetes cluster. PGO 5 has been installed. Check the README.md in this directory for the commands to launch a postgres cluster"
