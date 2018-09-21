# Clusterctl CLI

### Introduction
This is a CLI that is currently built in the vendor specific repo.  As mentioned in the introduction, the cluster api was designed as Kubernetes API controllers and can be accessed using kubectl.  It follows the standard operator pattern of Kubernetes.  However, kubectl does not solve the problem of how to get the bootstrap cluster up and running.  This is the problem that clusterctl attempts to solve.  Outside of provisioning the bootstrap cluster, nearly all other operations that clusterctl execute are performed by shelling out to kubectl.


### The bootstrap problem

Since the cluster API is designed as a Kubernetes API server, it presents a chicken and egg problem.  How does the initial cluster get created in order to deploy the cluster APIs that will be used to create and manage new clusters?  As stated in the previous section, clusterctl CLI solves this by creating a bootstrap cluster using minikube; however, there are actually 3 possible workflows.  There are actually a few other possible workflows discussed later.

1. Bootstrap with a minikube cluster and pivot to target cluster
2. Bootstrap using an existing cluster and pivot to target cluster
3. Bootstrap and DON'T pivot to target cluster

##### Bootstrap with a minikube cluster and pivot to target cluster

This is the original workflow built into clusterctl.

1. The CLI creates a **bootstrap** minikube cluster, that is initialized using kubeadm.
2. The CLI deploys the cluster API components onto this **bootstrap** cluster.
3. The CLI creates a **target** cluster using the cluster API
4. The CLI pivots to the **target** cluster and pulls down the kubeconfig for it.
5. The CLI deletes the bootstrap cluster.

These steps are the basic workflow for most clusterctl CLI commands.  There are some differences, depending on the command, but the workflow are similar.  Below are diagrams illustrating the create and delete workflows.

*Diagram 1: Cluster create with bootstrap minikube and pivot:*

![Cluster Create with CLI](images/cluster_create_with_cli.png)

*Diagram 2: Cluster delete with bootstrap minikube and pivot:*

![Cluster Delete with CLI](images/cluster_delete_with_cli.png)

While this workflow works for a developer persona, it assumes the administrator's role is simply to provision the vSphere resource pool.  In addition, the workflow is very slow due to bootstrap cluster creation time.  There can be other approaches to creating this bootstrap cluster to speed up the workflow, such as deploying a [Kind cluster](https://github.com/kubernetes-sigs/kind) as the bootstrap cluster; however, it is not certain how much more development will go into the minikube bootstrap approach in the future.

##### Bootstrap using an existing cluster and pivot to target cluster



##### Bootstrap and DO NOT pivot to target cluster


### Potential Future Directions

