# Install CLI
```
wget https://github.com/kubeless/kubeless/releases/download/v1.0.0-alpha.8/kubeless_linux-amd64.zip
unzip kubeless_linux-amd64.zip
sudo mv bundles/kubeless_linux-amd64/kubeless /usr/local/bin
rm -r bundles/
```

# Deploy kubeless
kubectl create ns kubeless
kubectl create -f https://github.com/kubeless/kubeless/releases/download/v1.0.0-alpha.8/kubeless-v1.0.0-alpha.8.yaml 

# Example function

## python
```
kubeless function deploy hello --runtime python2.7 \
                               --from-file python-example/example.py \
                               --handler test.hello
```
## NodeJS
```
kubeless function deploy myfunction --runtime nodejs6 \
                                --dependencies node-example/package.json \
                                --handler test.myfunction \
                                --from-file node-example/example.js
```

# Commands

## List Function
```
kubeless function ls
```
## Call Function
```
kubeless function call myfunction --data 'This is some data'
```

## Expose function
```
kubectl create -f nginx-ingress-controller-with-elb.yml
kubeless trigger http create myfunction --function-name myfunction --hostname myfunction.kubernetes.newtech.academy
```
