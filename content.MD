# **Content**

# Introduction
- Kubernetes
- Monolithic
- Microservice
- Setup
    - Minikube (single node without ha)
    - Kops (Aws) (Production)
        - Kubernetes Operation
    - Kubeadm (Generic) (Production)
- Docker
    - Build Docker Image
    - Push container
- Run app in Kubernetes
    - Create Pod (pod-app.yaml)
    - Push pod to cluster

# Architecture
- Kubernetes Architecture
    - Master
    - Worker
    - Pod
    - Container
- Master
    - ApiServer
    - Scheduler
    - Controller
        - Node Controller
        - Replication Controller
        - Endpoint Controller
        - Service Controller
    - Etcd
- Worker
    - Kubelet
    - Kube-Proxy
    - Pod
    - Container
- Container
    - Virtual Machine
    - Container
    - Docker
- Pod
    - Pod Life cycle
        - Pod state ($ kubectl get pod < name >)
            - Running
            - Pending
            - Failed
            - Succeeded
            - Unknown
        - Pod Condition ($ kubectl describe pod < name > -n kube-system)
            - PodScheduled
            - Ready
            - Initialized
            - UnSchedulable
            - ContainerReady
        - Definition/Spec
            - initContainer
            - mainContainer
            - postStart hook
            - readinessProbe (Health Check)
            - livenessProbe (Health Check)
            - postStop hook
    - Networking
        - Inter Pod Network
        - Intra Pod Network

# Basic
- Pod
    - Introduction
    - Deployment
    - Pod Life cycle
    - Config
        - Pod Manifest

- Scaling/Replica
    - Scaling
        - Stateless (No sync between app)
            - Horizontal scaling
            - web application
        - Statefull (Sync between app)
            - Vertical scaling
            - Volume (volume need to save state)
            - Db application
    - Replication Controller
        - Introduction
        - Load Balancer
        - Spec
        - Kubectl cmd
    - ReplicaSet
        - Introduction
        - Spec
    - Replication Controller vs ReplicaSet
        - Selector
        - Set Based Selection
        - Equality Based Selection

- Deployment
    - Introduction
    - Type
        - Recreation
        - Rolling Update
        - Canary
        - Blue/Green
    - Spec
    - Kubectl Command

- StatefulSet (Depend on Volume)
    - Introduction
    - Naming Convention
    - Example

- Labels
    - Basic
    - Selector
        - Equality based selector
        - Set Based selector
    - Label
        - Definition in Pod (metadata)
    - NodeSelector
        - Node Label
    - Kubectl Command
        - Get pod
        - Set label to node


# Service

- Service
    - Introduction
    - Type
        - Service Discovery / NodePort (Expose port outside)
            - DNS (nslookup)
            - NodePort
            - Port
            - TargetPort
            - Spec
            - Kubectl Command
        - Load Balancer (Multiple node)
            - Introduction
            - Spec
            - Kubectl Command
            - Ingress
                - Introduction
                - Ingress controller
                    - nginx ingress controller
                - Ingress rules
                - Spec/Example
        - External DNS (WIP)
        - ClusterIp (Vip)
            - Introduction
            - Spec
            - Kubectl Command

# Storage

- Secrets
    - Introduction
    - Definition/Spec
    - Type
        - volume
            - Access from Env variable
        - volumeMounts
            - Access from File in Container
    - Available
        - Internal (Kubernetes etcd)
        - External (External vault storage)
    - Kubectl Command
    - Example
        - Wordpress

- ConfigMap
    - Introduction
    - Create
        - Dictionary
        - Files
        - Literals
    - Usage in Pod
        - Volume/VolumeMount
    - Spec
    - Kubectl Command

- Volume
    - Type
        - Ephernal (Volume with Pod lifetime)
        - Durable (Volume beyond Pod lifetime)
    - Kubernetes Supported Volume
        - Cloud Based Volume
            - VolumeID
            - Used in pod spec
            - Example
                - Azure
                - AWS ECB
                - Google Disk
        - Network Based Volume
            - Example
                - NFS, ISCSI, CHEPFS
        - Local Node Based Volume
            - Example
                - Local disk
    - Volume Example
        - EmptyDir, hostPath, Secrets, AWS, NFS, ...

- AutoProvision Persistence Volume
    - AutoProvision Lifecycle of Volume
        - Provision (Create Volume)
        - Binding (Claim Volume)
        - Using (Use Volume in Pod)
        - Reclaiming (Delete or Reclaiming Volume)
    - Volume Type
        - Static (Persistence Volume)
            - Spec
                - Persistence Volume Manifest
                - Persistence Volume Claim
                - Refer Volume in Pod
            - Kubectl command
        - Dynamic (Persistence Volume Claim)
            - Spec
                - Persistence Volume Manifest
                - Persistence Volume Claim
                - Refer Volume in Pod
            - Kubectl command

- Pod Presets
    - Introduction
    - Injection Resource
        - Environmental Variable
        - Volumes
        - configMaps
        - Secrets
    - Usage
    - Example

# Monitoring

- Health Checks
    - Introduction
    - Type
        - Inside Pod as command
        - http request to URL
    - livenessProbe
        - Introduction
        - Definition
    - readinessProbe
        - Introduction
        - Definition

- Resource Monitoring Usage
    - Introduction
    - Component
        - Heapster
        - InfluxDB
        - Grafana
    - Metrics Server
    - Kubectl Commands

# Scheduling

- Autoscaling
    - Introduction
    - Heapster
    - Calculation
    - Example

- Jobs
    - Introduction
    - Usage
        - Scheduler
        - Long Running Job
    - Spec

- DaemonSet
    - Introduction
    - Usage (Monitoring)
    - Spec

- Affinity/Anti-Affinity
    - Introduction
    - Type
        - Node Affinity/Anti-Affinity
            - Type
                - requiredDuringSchedulingIgnoredDuringExecution
                - preferredDuringSchedulingIgnoredDuringExecution
                    - Weight
            - Populated Labels
            - Example
        - InterPod Affinity/Anti-Affinity
            - Introduction
            - Type
                - requiredDuringSchedulingIgnoredDuringExecution
                - preferredDuringSchedulingIgnoredDuringExecution
                    - Weight
            - Populated Labels
            - Usage
                - Co-located
                - Availability Zone
            - Rule
                - topologyKey
                - Expression
            - Pod Anti-affinity
            - Operator
            - Example

- Taints and Toleration
    - Taints
    - Toleration
    - Type
        - NoSchedule
        - PreferNoSchedule
        - NoExecute
    - Key
        - node.kubernetes.io/not-ready
        - node.kubernetes.io/out-of-disk
        - node.kubernetes.io/unreachable
        - node.kubernetes.io/memory-pressure
        - node.kubernetes.io/disk-pressure
        - node.kubernetes.io/network-unavailable
        - node.kubernetes.io/unschedulables
    - Example


# Administration

- Resource Management
    - Introduction
    - Quotas
        - Resource Quotas
            - requests.cpu
            - requests.mem
            - requests.storage
        - Object Quotas
            - configmaps
            - secrets
            - pods
            - nodeports
            - loadbalancer
            - etc.
    - Namespace
        - kube-system
        - Kubectl commands
            - create, list, set default
        - Example/Spec
            - Resource Quota
            - Object Quota

- User Management
    - Introduction
    - Authentication and Authorization
        - Keys
    - User type
        - Normal user
        - Service user
    - Attribute
        - Username
        - UID
        - Groups
        - Metadata
    - Authentication
        - Client certificate
        - Barer Token
        - Authentication Proxy
        - HTTP Basic Authentication
        - OpenID
        - WebHooks
    - Authorization
            - AllowsAllow/AlwaysDeny
            - ABAC (Attribute Based Access Control)
            - RBAC (Role Based Access Control)
            - Webhooks (Authentication)
    - Kubectl Commands
    - RBAC (Role Based Access Control)
        - Introduction
        - Webhooks
        - User
        - Role
            - Role (single namespace)
            - ClusterRole (cluster-wide)
        - Role Binding
            - RoleBinding
            - ClusterRoleBinding
        - Demo

- Networking
    - Container to Container (Localhost)
    - Pod to Service (NodePort, DNS)
    - External to Service (LoadBalancer)
    - Pod to Pod Communication
        - Pod IP
        - Kubernetes Network
            - VPC
            - CNI (Container Network Interface)
                - Calico
                - Weave
                - Flannel

- Node Maintenance
    - Node Controller
    - Node Attribute
        - Metadata (Ip, Name, hostname)
        - Labels (cloud region, availability zone)
        - NodeCondition (Ready, OutOfDisk)
    - Delete Node


# High Availability

- Kubernetes HA
    - Introduction
    - Architecture
    - Requirement
        - Multiple Etcd node
        - Multiple Controller and Scheduler
        - LoadBalancer

- Etcd
    - Introduction
    - Backup
    - Restore

# Packaging & Deployment

- Helm
    - Introduction
    - Charts
        - Basic
        - Template
        - Example
        - Create Own charts
            - Command ($ helm create my_chart)
                - Chart.yaml
                - values.yaml
                - template
                    - Kubernetes resources
            - Install charts ($ helm install my_chart)
    - Kubectl Commands
        - init, reset, install, search, list, upgrade, rollback

- Continuous Development
    - Skafold
        - Introduction
        - Work Pipeline
            - Building
            - Pushing
            - Deploying
        - Design Flow
        - Demo
    - Flux
        - Introduction
        - GitOps
        - Architecture
        - Demo

- istio
    - Microservices
        - Example Design
    - Proxy (Sidecar)
        - Example Design
    - istio
        - Design
            - Envey Proxy
            - Pilicy check Telementry
            - Piolat (Config Data)
            - Citadel (TSL Certification)
        - Demo
    - istio security
        - TSL
            - Mutual TSL in istio
            - Demo
            - RBAC
                - Rules
                - Example
        - Origin Authentication (JWT)
            - Json web token
            - Token
                - Header
                - Payload
                - signed based hader+token
            - Demo