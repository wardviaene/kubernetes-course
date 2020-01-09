# Setup EKS
```
eksctl create cluster --name=cluster-2 --nodes=2 --region=eu-west-1 --managed
```

# Setup IAM Roles for Service Accounts

Enable IAM Roles for Service Accounts on the EKS cluster

```
eksctl utils associate-iam-oidc-provider --cluster=cluster-2
eksctl utils associate-iam-oidc-provider --cluster=cluster-2 --approve
```

Create new IAM Role using eksctl
```
eksctl create iamserviceaccount --cluster=cluster-2 --name=myserviceaccount --namespace=default --attach-policy-arn=arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
eksctl create iamserviceaccount --cluster=cluster-2 --name=myserviceaccount --namespace=default --attach-policy-arn=arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess --approve
```
