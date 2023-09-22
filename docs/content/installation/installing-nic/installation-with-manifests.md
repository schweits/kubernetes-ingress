---
title: Installation with Manifests
description: "This document describes how to install the NGINX Ingress Controller in your Kubernetes cluster using Kubernetes manifests."
weight: 100
doctypes: [""]
aliases:
    - /installation/
toc: true
docs: "DOCS-603"
---

{{<custom-styles>}}

## Prerequisites

{{<note>}} Always use the latest stable release as shown on the [the releases page](https://github.com/nginxinc/kubernetes-ingress/releases). {{</note>}}

1. Get the NGINX Ingress Controller image:

    - For NGINX: Get the image `nginx/nginx-ingress` from [DockerHub](https://hub.docker.com/r/nginx/nginx-ingress).
    - For NGINX Plus: Follow the steps in the [Getting the F5 Registry NGINX Ingress Controller Image]({{< relref "installation/nic-images/pulling-ingress-controller-image.md" >}}) guide.
    - To pull from the F5 Container registry in your Kubernetes cluster: Follow the steps to [Configure a Docker registry secret with your JWT token]({{< relref "installation/using-the-jwt-token-docker-secret.md" >}}).
    - Or build your own image: Follow the steps in [Building NGINX Ingress Controller]({{< relref "installation/building-nginx-ingress-controller.md" >}}).

2. Clone the NGINX Ingress Controller repository and go the _deployments_ folder:

    ```shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1
    cd kubernetes-ingress/deployments
    ```

    {{<note>}}The command above clones the latest release, which is what this guide is based on.{{</note>}}

---

## Set Up role-based access control (RBAC)

{{<call-out "important" "Admin access required" >}}You must be a cluster admin to perform the steps in this section. Refer to the documentation for your Kubernetes platform to set up admin access. For Google Kubernetes Engine (GKE), see their [Role-Based Access Control guide](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control).{{</call-out>}}

1. Create a namespace and a service account:

    ```shell
    kubectl apply -f common/ns-and-sa.yaml
    ```

2. Create a cluster role and binding for the service account:

    ```shell
    kubectl apply -f rbac/rbac.yaml
    ```

3. (App Protect only) Create the *App Protect* role and binding:

    ```shell
    kubectl apply -f rbac/ap-rbac.yaml
    ```

4. (App Protect DoS only) Create the *App Protect DoS* role and binding:

    ```shell
    kubectl apply -f rbac/apdos-rbac.yaml
    ```

---

## Create common resources

In this section, we create resources common for most of NGINX Ingress Controller installations:

1. (Optional) Create a secret that contains both a TLS certificate and a key for the default NGINX server. Make sure you're in the `kubernetes-ingress/deployment` directory and run:

    ```shell
    kubectl apply -f ../examples/shared-examples/default-server-secret/default-server-secret.yaml
    ```

    {{<call-out "important" "Optional step" >}}
You only need to install the `default-server-secret.yaml` if you're using the  [default server TLS secret]({{< relref "configuration/global-configuration/command-line-arguments#cmdoption-default-server-tls-secret.md" >}}) command-line argument. Otherwise, you can skip this step. We recommend you use your own certificate.
{{</call-out>}}

    {{<note>}}By default, the server returns a _404 Not Found_ page for all requests where no ingress rules are set up. For testing, we've included a self-signed certificate and key, but you should use your own. {{</note>}}

2. Create a ConfigMap to customize your NGINX settings:

    ```shell
    kubectl apply -f common/nginx-config.yaml
    ```

3. Create an IngressClass resource:

    ```shell
    kubectl apply -f common/ingress-class.yaml
    ```

    To make this NGINX Ingress Controller instance the default, remove the comment from the annotation ingressclass.kubernetes.io/is-default-class. Doing this will automatically assign this IngressClass to any new Ingresses that don't specify an ingressClassName.

    {{<note>}}NGINX Ingress Controller won't start if you don't create an IngressClass resource. {{</note>}}

---

## Create custom resources

{{<note>}}
To make sure NGINX Ingress Controller pods become `Ready`, you'll need to create custom resource definitions for VirtualServer, VirtualServerRoute, TransportServer, and Policy. If you want to skip this, set the [`-enable-custom-resources`]({{< relref "configuration/global-configuration/command-line-arguments.md#cmdoption-global-configuration.md" >}}) command-line argument to `false`.
{{</note>}}

1. Create custom resource definitions for [VirtualServer and VirtualServerRoute]({{< relref "configuration/virtualserver-and-virtualserverroute-resources.md" >}}), [TransportServer]({{< relref "configuration/transportserver-resource.md" >}}), and [Policy]({{< relref "configuration/policy-resource.md" >}}):

    ```shell
    kubectl apply -f common/crds/k8s.nginx.org_virtualservers.yaml
    kubectl apply -f common/crds/k8s.nginx.org_virtualserverroutes.yaml
    kubectl apply -f common/crds/k8s.nginx.org_transportservers.yaml
    kubectl apply -f common/crds/k8s.nginx.org_policies.yaml
    ```

2. To use TCP and UDP load balancing, create a custom resource definition for [GlobalConfiguration]({{< relref "configuration/global-configuration/globalconfiguration-resource.md" >}}):

    ```shell
    kubectl apply -f common/crds/k8s.nginx.org_globalconfigurations.yaml
    ```

3. To use the App Protect WAF module, you need to create custom resource definitions for `APPolicy`, `APLogConf` and `APUserSig`:

    ```shell
    kubectl apply -f common/crds/appprotect.f5.com_aplogconfs.yaml
    kubectl apply -f common/crds/appprotect.f5.com_appolicies.yaml
    kubectl apply -f common/crds/appprotect.f5.com_apusersigs.yaml
    ```

4. To use the App Protect DoS module, create custom resource definitions for `APDosPolicy`, `APDosLogConf` and `DosProtectedResource`:

   ```shell
   kubectl apply -f common/crds/appprotectdos.f5.com_apdoslogconfs.yaml
   kubectl apply -f common/crds/appprotectdos.f5.com_apdospolicy.yaml
   kubectl apply -f common/crds/appprotectdos.f5.com_dosprotectedresources.yaml
   ```

---

## Deploying NGINX Ingress Controller

You have two ways to deploy NGINX Ingress Controller:

- **Deployment**. Choose this option if you need to dynamically change the number of NGINX Ingress Controller replicas.
- **DaemonSet**. Use this option if you want NGINX Ingress Controller to run on every node or a subset of nodes.

If you're planning to use the NGINX App Protect DoS module, you must deploy the Arbitrator. See the steps [below](#deploy-arbitrator).

{{<note>}}Before you set up a Deployment or DaemonSet resource, make sure to update the [command-line arguments]({{< relref "configuration/global-configuration/command-line-arguments.md" >}}) for the NGINX Ingress Controller container in the corresponding manifest file to meet your specific needs. {{</note>}}

### Deploying Arbitrator for NGINX App Protect DoS {#deploy-arbitrator}

Follow these steps to deploy NGINX Ingress Controller with the NGINX App Protect DoS module:

1. Follow the instructions in [Installation with NGINX App Protect DoS]({{< relref "installation/integrations/app-protect-dos/installation.md#build-docker-image" >}}) to build your custom image and upload it to your private Docker registry.

2. Run the Arbitrator as Kubernetes deployment and service:

   ```shell
   kubectl apply -f deployment/appprotect-dos-arb.yaml
   kubectl apply -f service/appprotect-dos-arb-svc.yaml
   ```

### Running NGINX Ingress Controller

#### With a deployment

By default, Kubernetes will set up a single NGINX Ingress Controller pod when you use a Deployment.

- For NGINX, run:

    ```shell
    kubectl apply -f deployment/nginx-ingress.yaml
    ```

- For NGINX Plus, run:

    ```shell
    kubectl apply -f deployment/nginx-plus-ingress.yaml
    ```

    {{<note>}} Update the `nginx-plus-ingress.yaml` with the chosen image from the F5 Container registry; or the container image that you have built. {{</note>}}

#### With a DaemonSet

When you run the Ingress Controller by using a DaemonSet, Kubernetes creates an Ingress Controller pod on every node in the cluster.

{{<note>}}Read the Kubernetes [DaemonSet docs](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) to learn how to run NGINX Ingress Controller on a subset of nodes instead of on every node of the cluster.{{</note>}}

For NGINX, run:

```shell
kubectl apply -f daemon-set/nginx-ingress.yaml
```

For NGINX Plus, run:

```shell
kubectl apply -f daemon-set/nginx-plus-ingress.yaml
```

{{<note>}}Update `nginx-plus-ingress.yaml` with the chosen image from the F5 Container registry; or the container image that you have built.{{</note>}}

### Check that NGINX Ingress Controller is running

Run the following command to make sure that the NGINX Ingress Controller pods are running:

```shell
kubectl get pods --namespace=nginx-ingress
```

---

## Getting access to NGINX Ingress Controller

**If you created a daemonset**, ports 80 and 443 of NGINX Ingress Controller container are mapped to the same ports of the node where the container is running. To access NGINX Ingress Controller, use those ports and an IP address of any node of the cluster where the Ingress Controller is running.

**If you created a deployment**, there are two options for accessing NGINX Ingress Controller pods:

### Create a service for the NGINX Ingress Controller pods

#### Using a NodePort service

Create a service with the type *NodePort*:

```shell
kubectl create -f service/nodeport.yaml
```

Kubernetes will randomly allocate two ports on every node of the cluster. To access the Ingress Controller, use an IP address of any node of the cluster along with the two allocated ports.

{{<note>}} Read more about the type NodePort in the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport). {{</note>}}

#### Using a LoadBalancer service

1. Create a service using a manifest for your cloud provider:
    - For GCP or Azure, run:

        ```shell
        kubectl apply -f service/loadbalancer.yaml
        ```

    - For AWS, run:

        ```shell
        kubectl apply -f service/loadbalancer-aws-elb.yaml
        ```

        Kubernetes will allocate a Classic Load Balancer (ELB) in TCP mode with the PROXY protocol enabled to pass the client's information (the IP address and the port). NGINX must be configured to use the PROXY protocol:
        - Add the following keys to the config map file `nginx-config.yaml` from the Step 2:

            ```yaml
            proxy-protocol: "True"
            real-ip-header: "proxy_protocol"
            set-real-ip-from: "0.0.0.0/0"
            ```

        - Update the config map:

            ```shell
            kubectl apply -f common/nginx-config.yaml
            ```

        {{<note>}} For AWS, additional options regarding an allocated load balancer are available, such as its type and SSL termination. Read the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer) to learn more. {{</note>}}

    Kubernetes will allocate and configure a cloud load balancer for load balancing the Ingress Controller pods.
2. Use the public IP of the load balancer to access NGINX Ingress Controller. To get the public IP:
    - For GCP or Azure, run:

        ```shell
        kubectl get svc nginx-ingress --namespace=nginx-ingress
        ```

    - In case of AWS ELB, the public IP is not reported by `kubectl`, because the ELB IP addresses are not static. In general, you should rely on the ELB DNS name instead of the ELB IP addresses. However, for testing purposes, you can get the DNS name of the ELB using `kubectl describe` and then run `nslookup` to find the associated IP address:

        ```shell
        kubectl describe svc nginx-ingress --namespace=nginx-ingress
        ```

        You can resolve the DNS name into an IP address using `nslookup`:

        ```shell
        nslookup <dns-name>
        ```

    The public IP can be reported in the status of an ingress resource. See the [Reporting Resources Status doc]({{< relref "configuration/global-configuration/reporting-resources-status.md" >}}) for more details.

{{<note>}} Learn more about type LoadBalancer in the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer). {{</note>}}

## Uninstall NGINX Ingress Controller

1. Delete the `nginx-ingress` namespace to uninstall NGINX Ingress Controller along with all the auxiliary resources that were created:

    ```shell
    kubectl delete namespace nginx-ingress
    ```

1. Delete the ClusterRole and ClusterRoleBinding:

    ```shell
    kubectl delete clusterrole nginx-ingress
    kubectl delete clusterrolebinding nginx-ingress
    ```

1. Delete the Custom Resource Definitions:

    {{<note>}} This step will also remove all associated Custom Resources. {{</note>}}

    ```shell
    kubectl delete -f common/crds/
    ```
