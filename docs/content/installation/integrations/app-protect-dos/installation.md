---
title: Installation with NGINX App Protect DoS
description: "This document provides an overview of the steps required to use NGINX App Protect DoS with your NGINX Ingress Controller deployment."
weight: 100
doctypes: [""]
toc: true
docs: "DOCS-583"
---

{{< custom-styles >}}

{{< note >}} The F5 NGINX Kubernetes Ingress Controller integration with F5 NGINX App Protect DoS requires the use of F5 NGINX Plus. {{< /note >}}

This document provides an overview of the steps required to use NGINX App Protect DoS with your NGINX Ingress Controller deployment. You can visit the linked documents to find additional information and instructions.

## Prerequisites

1. Make sure you have access to the NGINX Ingress Controller image:
    - For NGINX Plus Ingress Controller, see [here]({{< relref "installation/nic-images/pulling-ingress-controller-image" >}}) for details on how to pull the image from the F5 Docker registry.
    - To pull from the F5 Container registry in your Kubernetes cluster, configure a docker registry secret using your JWT token from the MyF5 portal by following the instructions from [here]({{< relref "installation/nic-images/using-the-jwt-token-docker-secret" >}}).
    - It is also possible to build your own image and push it to your private Docker registry by following the instructions from [here]({{< relref "installation/building-nginx-ingress-controller.md" >}})).
2. Clone the NGINX Ingress Controller repo:

    ``` shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1
    cd kubernetes-ingress/deployments
    ```

## Install the App Protect DoS Arbitrator

### Helm Chart

The App Protect DoS Arbitrator can be installed using the [NGINX App Protect DoS Helm Chart](https://github.com/nginxinc/nap-dos-arbitrator-helm-chart).
If you have the NGINX Helm Repository already added, you can install the App Protect DoS Arbitrator by running the following command:

```shell
helm install my-release-dos nginx-stable/nginx-appprotect-dos-arbitrator
```

### YAML Manifests

Alternatively, you can install the App Protect DoS Arbitrator using the YAML manifests provided in the NGINX Ingress Controller repo.

1. Create the namespace and service account:

    ```shell
      kubectl apply -f common/ns-and-sa.yaml
    ```

2. Deploy the NGINX App Protect Arbitrator as a Deployment and service:

    ```shell
    kubectl apply -f deployment/appprotect-dos-arb.yaml
    kubectl apply -f service/appprotect-dos-arb-svc.yaml
    ```

## Build the Docker Image {#build-docker-image}

Take the steps below to create the Docker image that you'll use to deploy NGINX Ingress Controller with App Protect DoS in Kubernetes.

- [Build the NGINX Ingress Controller image]({{< relref "installation/building-nginx-ingress-controller.md" >}}).

  When running the `make` command to build the image, be sure to use the `debian-image-dos-plus` target. For example:

    ```shell
    make debian-image-dos-plus PREFIX=<your Docker registry domain>/nginx-plus-ingress
    ```

    Alternatively, if you want to run on an [OpenShift](https://www.openshift.com/) cluster, use the `ubi-image-dos-plus` target.

    If you want to include the App Protect WAF module in the image, you can use the `debian-image-nap-dos-plus` target or the `ubi-image-nap-dos-plus` target for OpenShift.

- [Push the image to your local Docker registry]({{< relref "installation/building-nginx-ingress-controller.md#build-image-push-to-private-repo " >}}).

## Install NGINX Ingress Controller {#install-nic}

Take the steps below to set up and deploy the NGINX Ingress Controller and App Protect DoS module in your Kubernetes cluster.

### Set up role-based access control (RBAC) {#set-up-rbac}

{{< include "rbac/set-up-rbac.md" >}}

### Create Common Resources

{{< include "installation/create-common-resources.md" >}}

### Enable NGINX App Protect DoS module

To enable the NGINX App Protect DoS Module:

1. Add the `enable-app-protect-dos` [command-line argument]({{< relref "configuration/global-configuration/command-line-arguments.md#cmdoption-enable-app-protect-dos" >}}) to your Deployment or DaemonSet file.

### Deploy NGINX Ingress Controller

You can deploy NGINX Ingress Controller in two ways:

- **Deployment**. Choose this method if you want the flexibility to change the number of NGINX Ingress Controller replicas dynamically.
- **DaemonSet**. Choose this option if you want NGINX Ingress Controller to run on all nodes or a specific set of nodes.

{{<note>}}Before setting up a Deployment or DaemonSet resource, update the [command-line arguments]({{< relref "configuration/global-configuration/command-line-arguments.md" >}}) for the NGINX Ingress Controller container in the relevant manifest file to meet your specific needs.{{</note>}}

#### Deploy as a Deployment

{{< include "installation/manifests/deployment.md" >}}

#### Deploy as a DaemonSet

{{< include "installation/manifests/daemonset.md" >}}

#### Confirm NGINX Ingress Controller is running

{{< include "installation/manifests/verify-pods-are-running.md" >}}


For more information, see the [Configuration guide]({{< relref "installation/integrations/app-protect-dos/configuration.md" >}}),the [NGINX Ingress Controller with App Protect DoS example for VirtualServer](https://github.com/nginxinc/kubernetes-ingress/tree/v3.2.1/examples/custom-resources/app-protect-dos) and the [NGINX Ingress Controller with App Protect DoS example for Ingress](https://github.com/nginxinc/kubernetes-ingress/tree/v3.2.1/examples/ingress-resources/app-protect-dos).
