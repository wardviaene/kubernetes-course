# install jenkins
```
kubectl create -f serviceaccount.yaml
helm install jenkins --set rbac.create=true,master.runAsUser=1000,master.fsGroup=1000,agent.enabled=true bitnami/jenkins
helm create -f jenkins-role-binding.yaml
kubectl patch svc jenkins --type merge -p '{"spec":{"ports": [{"port": 50000,"name":"agent-listener", "protocol": "TCP", "targetPort": "agent-listener"}, {"port": 80, "name": "http", "targetPort": "http"}]}}'
```

# Plugins
Ensure you install the following plugins within jenkins:
* Pipelines
* Kubernetes
* Git

To configure the Kubernetes plugin correctly, navigate to Manage Jenkins -> Manage Nodes and Clouds -> Configure Clouds -> Add a new cloud. The name should be pre-filled with "kubernetes". Fill out http://jenkins as Jenkins URL in the "Kubernetes Cloud details" and click "Save" without filling out anything extra. The Kubernetes plugin will now automatically find the Kubernetes cluster jenkins is installed on.

If you get a timeout on port 50000 then go to "Manage Jenkins" -> "Configure Global Security" -> under "Agents" check whether TCP port for inbound agents is set to "Fixed" 50000.

Note: ensure that you are running jenkins privately (non-internet facing), as port 50000 will also be exposed in this setup.
