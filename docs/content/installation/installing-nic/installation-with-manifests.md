---
title: Installing with Manifests
description: "This document describes how to install the NGINX Ingress Controller in your Kubernetes cluster using Kubernetes manifests."
weight: 100
doctypes: [""]
aliases:
    - /installation/
toc: true
docs: "DOCS-603"
---

{{<custom-styles>}}

## Before you start

{{<note>}}Always use the most up-to-date stable release listed on the [releases page]({{< relref "releases.md" >}}).{{</note>}}

1. Get the NGINX Ingress Controller image:

    - For NGINX: Get the image `nginx/nginx-ingress` from [DockerHub](https://hub.docker.com/r/nginx/nginx-ingress).
    - For NGINX Plus: Follow the steps in the [Getting the F5 Registry NGINX Ingress Controller Image]({{< relref "installation/nic-images/pulling-ingress-controller-image.md" >}}) guide.
    - To pull from the F5 Container registry in your Kubernetes cluster: Follow the steps to [Getting the NGINX Ingress Controller Image with JWT]({{< relref "installation/nic-images/using-the-jwt-token-docker-secret.md" >}}).
    - To build your own image: Follow the steps in [Building NGINX Ingress Controller]({{< relref "installation/building-nginx-ingress-controller.md" >}}).

2. Clone the NGINX Ingress Controller repository and go the _deployments_ folder:

    ```shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch <version_number>. Replace `<version_number>` with the specific release you want to use.
    cd kubernetes-ingress/deployments
    ```

    Note: For example, if you want to use version 3.2.1, the command would be `git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1`. 

    This guide is based on the latest release.

---

## Set up role-based access control (RBAC) {#configure-rbac}

{{< include "rbac/set-up-rbac.md" >}}

---

## Create common resources {#create-common-resources}

{{< include "installation/create-common-resources.md" >}}

---

## Create custom resources {#create-custom-resources}

{{<note>}}
To ensure that NGINX Ingress Controller pods reach the `Ready` state, you need to create custom resource definitions for VirtualServer, VirtualServerRoute, TransportServer, and Policy. If you prefer to skip this, set the [`-enable-custom-resources`]({{< relref "configuration/global-configuration/command-line-arguments.md#cmdoption-global-configuration.md" >}}) command-line argument to `false`.
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

3. To use the App Protect WAF module, create custom resource definitions for `APPolicy`, `APLogConf` and `APUserSig`:

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

## Deploy NGINX Ingress Controller {#deploy-ingress-controller}

You can deploy NGINX Ingress Controller in two ways:

- **Deployment**. Choose this method if you want the flexibility to change the number of NGINX Ingress Controller replicas dynamically.
- **DaemonSet**. Choose this option if you want NGINX Ingress Controller to run on all nodes or a specific set of nodes.

{{<note>}}Before setting up a Deployment or DaemonSet resource, update the [command-line arguments]({{< relref "configuration/global-configuration/command-line-arguments.md" >}}) for the NGINX Ingress Controller container in the relevant manifest file to meet your specific needs.{{</note>}}

### Run NGINX Ingress Controller

#### Using a deployment

{{< include "installation/manifests/deployment.md" >}}

#### Using a DaemonSet

{{< include "installation/manifests/daemonset.md" >}}

### Confirm NGINX Ingress Controller is running

{{< include "installation/manifests/verify-pods-are-running.md" >}}

---

## How to access NGINX Ingress Controller

### For DaemonSet installations

Connect to ports 80 and 443 using the IP address of any node in the cluster where NGINX Ingress Controller is running.

### For Deployment installations

For deployments, you have two options for accessing NGINX Ingress Controller pods.

#### Option 1: Create a NodePort service

To create a service with the type *NodePort*, run:

```shell
kubectl create -f service/nodeport.yaml
```

Kubernetes will automatically allocate two ports on every node in the cluster. To access NGINX Ingress Controller, use any node's IP address combined with these allocated ports.

{{<note>}}For more details on _NodePort_, refer to the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport). {{</note>}}

#### Option 2: Create a LoadBalancer service

1. To create a service using a manifest for your cloud provider:

    - For GCP or Azure, run:

        ```shell
        kubectl apply -f service/loadbalancer.yaml
        ```

    - For AWS, run:

        ```shell
        kubectl apply -f service/loadbalancer-aws-elb.yaml
        ```

        Kubernetes will set up a Classic Load Balancer (ELB) in TCP mode. This load balancer will have the PROXY protocol enabled to pass along the client's IP address and port. To make NGINX work with this, you need to adjust NGINX's configuration:

        - Add the following keys to the `nginx-config.yaml` ConfigMap file, which you created in the [Create common resources](#create-common-resources) section.

            ```yaml
            proxy-protocol: "True"
            real-ip-header: "proxy_protocol"
            set-real-ip-from: "0.0.0.0/0"
            ```

        - Update the ConfigMap:

            ```shell
            kubectl apply -f common/nginx-config.yaml
            ```

        {{<note>}}AWS users have more customization options for their load balancers. These include choosing the load balancer type and configuring SSL termination. Refer to the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer) to learn more. {{</note>}}

2. To access NGINX Ingress Controller, you'll need the public IP of your load balancer. Here's how to find it:

    - For GCP or Azure, run:

        ```shell
        kubectl get svc nginx-ingress --namespace=nginx-ingress
        ```

    - For AWS ELB, `kubectl` won't show the public IP because ELB uses dynamic IP addresses. Generally, you can rely on the ELB DNS name. But for testing, you can get the DNS name like this:

        ```shell
        kubectl describe svc nginx-ingress --namespace=nginx-ingress
        ```

        You can resolve the DNS name into an IP address using `nslookup`:

        ```shell
        nslookup <dns-name>
        ```

    You can also find more details about the public IP in the status section of an ingress resource. For more details, refer to the [Reporting Resources Status doc]({{< relref "configuration/global-configuration/reporting-resources-status.md" >}}).

{{<note>}}For more information on the LoadBalancer service refer to the [Kubernetes documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer). {{</note>}}

## Uninstall NGINX Ingress Controller

{{<important>}}Be cautious when performing these steps, as they will remove NGINX Ingress Controller and all related resources, potentially affecting your running services.{{</important>}}

1. o uninstall NGINX Ingress Controller and remove all auxiliary resources, delete the nginx-ingress namespace:

    ```shell
    kubectl delete namespace nginx-ingress
    ```

1. Next, remove the cluster role and cluster role binding:

    ```shell
    kubectl delete clusterrole nginx-ingress
    kubectl delete clusterrolebinding nginx-ingress
    ```

1. Lastly, delete the Custom Resource Definitions:

    {{<note>}} Performing this step will also remove all associated Custom Resources. {{</note>}}

    ```shell
    kubectl delete -f common/crds/
    ```

---

## Optional Deploy NGINX Ingress Controller with NGINX App Protect DoS

To deploy NGINX Ingress Controller with NGINX App Protect DoS:

1. Follow the guidelines in [Installation with NGINX App Protect DoS]({{< relref "installation/integrations/app-protect-dos/installation.md#build-docker-image" >}}) to build your custom image and upload it to your private Docker registry.

2. Start the Arbitrator as a Kubernetes deployment and service:

   ```shell
   kubectl apply -f deployment/appprotect-dos-arb.yaml
   kubectl apply -f service/appprotect-dos-arb-svc.yaml
   ```
