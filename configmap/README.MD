to create ConfigMap for nginx-app
kubectl create configmap nginx-config --from-file=reverseproxy.conf
