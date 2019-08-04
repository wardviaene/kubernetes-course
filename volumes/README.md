# Create volumes

## Create Volume in AWS

```
aws ec2 create-volume --size 10 --region eu-west-1 --availability-zone eu-west-1a --volume-type gp2 --tag-specifications 'ResourceType=volume, Tags=[{Key= KubernetesCluster, Value=kubernetes.domain.tld}]'
```

aws ec2 create-volume --size 10 --region us-east-1 --availability-zone us-east-1a --volume-type gp2

aws ec2  delete-volume --volume-id vol-id
kubectl delete -f volumes/helloworld-with-colume.yml