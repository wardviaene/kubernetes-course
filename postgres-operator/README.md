# Files
quickstart-for-gke.sh. Tested version available in directory, or latest from: https://github.com/CrunchyData/postgres-operator/blob/master/examples/quickstart-for-gke.sh

# setup storage
```
kubectl create -f storage.yml
```

# setup Operator
```
./quickstart-for-gke.sh
./set-path.sh 
```

After these commands you'll need to logout and login again.

# port forwarding

```
kubectl port-forward postgres-operator-xxx-yyy 18443:8443
```

# Test command

```
pgo version
```

# Create cluster

```
pgo create cluster mycluster
pgo show cluster all
```

# show secrets
```
pgo show cluster mycluster --show-secrets=true
```

# connect to psql
```
kubectl run -it --rm --image=postgres:10.4 psql -- psql -h mycluster -U postgres -W
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
