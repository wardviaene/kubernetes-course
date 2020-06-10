# Files
There is no quickstart-for-gke.sh anymore. This has been replaced by quickstart.sh.

# setup storage
```
kubectl create -f storage.yml
```

# setup Operator
```
./quickstart.sh
./set-path.sh 
```

After these commands you'll need to logout and login again.

# port forwarding

```
kubectl get pods -n pgo
kubectl port-forward -n pgo postgres-operator-xxx-yyy 8443:8443
```

# Test command

```
pgo version
```

# Create cluster

```
pgo create cluster mycluster
```

# show secrets
```
pgo show cluster mycluster
```

# connect to psql
```
pgo show user mycluster
kubectl run -it --rm --image=postgres:10.4 psql-client -- psql -h mycluster.pgo -U testuser -W postgres
```


# Create read replic
```
pgo scale mycluster
```

# manually failover
```
pgo failover mycluster --query
pgo failover mycluster --target=mycluster-xxx
kubectl get pgtasks mycluster-failover -o yaml
```
