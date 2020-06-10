#!/bin/bash
echo 'export PATH=$PATH:~/.pgo/pgo/' >> ~/.bashrc
echo 'export PGOUSER="${HOME?}/.pgo/pgo/pgouser"' >> ~/.bashrc
echo 'export PGO_CA_CERT="${HOME?}/.pgo/pgo/client.crt"' >> ~/.bashrc
echo 'export PGO_CLIENT_CERT="${HOME?}/.pgo/pgo/client.crt"' >> ~/.bashrc
echo 'export PGO_CLIENT_KEY="${HOME?}/.pgo/pgo/client.key"' >> ~/.bashrc
echo 'export PGO_APISERVER_URL='https://127.0.0.1:8443'' >> ~/.bashrc
echo 'export PGO_NAMESPACE=pgo' >> ~/.bashrc

