---
title: Installation with NGINX App Protect DoS
description: "This document explains the steps to take to use NGINX App Protect DoS with NGINX Ingress Controller."
weight: 100
doctypes: [""]
toc: true
docs: "DOCS-583"
---

{{< custom-styles >}}

## Before you start

{{< note >}} To use NGINX App Protect DoS with NGINX Ingress Controller, you must have NGINX Plus. {{< /note >}}

### Get the NGINX Plus Controller Image

{{<note>}}Always use the most up-to-date stable release listed on the [releases page]({{< relref "releases.md" >}}).{{</note>}}

Choose one of the following methods to get the NGINX Plus Ingress Controller image:

- Download the image using your NGINX Ingress Controller subscription certificate and key. See the [Getting the F5 Registry NGINX Ingress Controller Image]({{< relref "installation/nic-images/pulling-ingress-controller-image.md" >}}) guide.
- Use your NGINX Ingress Controller subscription JWT token to get the image: Instructions are in [Getting the NGINX Ingress Controller Image with JWT]({{< relref "installation/nic-images/using-the-jwt-token-docker-secret.md" >}}).
- Build your own image: To build your own image, follow the [Building NGINX Ingress Controller]({{< relref "installation/building-nginx-ingress-controller.md" >}}) guide.

### Clone the repository

Clone the NGINX Ingress Controller repository and go to the _deployments_ folder. Replace `<version_number>` with the specific release you want to use.

```shell
git clone https://github.com/nginxinc/kubernetes-ingress.git --branch <version_number>
cd kubernetes-ingress/deployments
```

For example, if you want to use version 3.2.1, the command would be `git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1`. 

This guide assumes you are using the latest release.

---


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

---

## Set up role-based access control (RBAC) {#set-up-rbac}

{{< include "rbac/set-up-rbac.md" >}}

---

## Create common resources {#create-common-resources}

{{< include "installation/create-common-resources.md" >}}

---

## Deploy NGINX Ingress Controller {#deploy-ingress-controller}

You have two options for deploying NGINX Ingress Controller:

- **Deployment**. Choose this method for the flexibility to dynamically change the number of NGINX Ingress Controller replicas.
- **DaemonSet**. Choose this method if you want NGINX Ingress Controller to run on all nodes or a subset of nodes.

Before you start, update the [command-line arguments]({{< relref "configuration/global-configuration/command-line-arguments.md" >}}) for the NGINX Ingress Controller container in the relevant manifest file to meet your specific requirements.

### Using a Deployment

{{< include "installation/manifests/deployment.md" >}}

### Using a DaemonSet

{{< include "installation/manifests/daemonset.md" >}}

---

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

---

## Enable NGINX App Protect DoS module

To enable the NGINX App Protect DoS Module:

1. Add the `enable-app-protect-dos` [command-line argument]({{< relref "configuration/global-configuration/command-line-arguments.md#cmdoption-enable-app-protect-dos" >}}) to your Deployment or DaemonSet file.

---

## Confirm NGINX Ingress Controller is running

{{< include "installation/manifests/verify-pods-are-running.md" >}}


For more information, see the [Configuration guide]({{< relref "installation/integrations/app-protect-dos/configuration.md" >}}),the [NGINX Ingress Controller with App Protect DoS example for VirtualServer](https://github.com/nginxinc/kubernetes-ingress/tree/v3.2.1/examples/custom-resources/app-protect-dos) and the [NGINX Ingress Controller with App Protect DoS example for Ingress](https://github.com/nginxinc/kubernetes-ingress/tree/v3.2.1/examples/ingress-resources/app-protect-dos).
