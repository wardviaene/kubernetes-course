# External DNS

Project page: https://github.com/kubernetes-incubator/external-dns

## Create IAM Policy
```
./put-node-policy.sh
```

## start ingress
```
kubectl apply -f ../ingress/
```

## Create LoadBalancer for Ingress
```
kubectl apply -f service-l4.yaml
```

## Configure kops

Either turn off the Instance metadata service version 2, or enable IRSA (follow steps at https://github.com/kubernetes/kops/blob/master/docs/cluster_spec.md#service-account-issuer-discovery-and-aws-iam-roles-for-service-accounts-irsa).

To disable the Instance Metadata service 2 (easiest option for non-production clusters), run kops edit instancegroup nodes-eu-west-1a --state=... and modify the instanceMetadata (do this for every zone):
```
spec:
  instanceMetadata:
    httpTokens: optional
```

If you want to use IRSA, this config can be used during kops edit cluster:
```
spec:
  # enable IRSA
  serviceAccountIssuerDiscovery:
    discoveryStore: s3://publicly-readable-store
    enableAWSOIDCProvider: true
  # IAM policy for service account with external-dns
  iam:
    serviceAccountExternalPermissions: 
      - name: external-dns 
        namespace: default 
        aws: 
          inlinePolicy: |- 
            [  
              {  
                "Effect": "Allow", 
                "Action": ["route53:ListHostedZones", "route53:ListResourceRecordSets"], 
                "Resource": "*" 
              }, 
              {  
                "Effect": "Allow", 
                "Action": [ 
                  "route53:ChangeResourceRecordSets" 
                ], 
                "Resource": [ 
                  "arn:aws:route53:::hostedzone/*" 
                ] 
              } 
            ] 
```

## Create external DNS and ingress rules
```
kubectl apply -f external-dns.yaml
kubectl apply -f ingress.yaml
```
