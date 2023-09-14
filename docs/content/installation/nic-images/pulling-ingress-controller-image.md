---
title: Getting the F5 Registry NGINX Ingress Controller Image
description: "This guide walks you through the process of pulling an F5 NGINX Plus Ingress Controller image from the F5 Docker registry and uploading it to your private registry."
weight: 10
doctypes: [""]
toc: true
docs: "DOCS-605"
---

To complete the steps in this guide, you'll need your NGINX Ingress Controller subscription certificate and key. Keep in mind that an NGINX Plus certificate and key won't work for this. If you're looking for alternative methods, here are some:

- [Install using a JWT token in a Docker Config Secret]({{< relref "using-the-jwt-token-docker-secret" >}}).
- [Build the Ingress Controller image]({{< relref "building-ingress-controller-image" >}}) using the source code from the GitHub repository and your NGINX Plus subscription certificate and key.
- For NGINX Ingress Controller based on NGINX OSS, you can pull the [nginx/nginx-ingress image](https://hub.docker.com/r/nginx/nginx-ingress/) from DockerHub.

## Prerequisites

Before you start, you'll need these installed on your machine:

- [Docker](https://www.docker.com/products/docker) v18.09 or higher
- An NGINX Ingress Controller subscription. Download both the certificate (*nginx-repo.crt*) and key (*nginx-repo.key*) from [MyF5](https://my.f5.com).

## Pull an image using Docker and push it to a private registry

1. Start by setting up Docker to communicate with the F5 Container registry, `private-registry.nginx.com`. In a Linux environment, create a directory under `/etc/docker/certs.d` named `private-registry.nginx`.com. Add a certificate (*client.cert*) and a key (*client.key*) to it. Here's how:

    ```shell
    mkdir -p /etc/docker/certs.d/private-registry.nginx.com
    cp nginx-repo.crt /etc/docker/certs.d/private-registry.nginx.com/client.cert
    cp nginx-repo.key /etc/docker/certs.d/private-registry.nginx.com/client.key
    ```

    <br>

    {{<tip>}}The example above is for Linux. For Mac or Windows, consult the [Docker for Mac](https://docs.docker.com/docker-for-mac/#add-client-certificates) or [Docker for Windows](https://docs.docker.com/docker-for-windows/#how-do-i-add-client-certificates) documentation.
    {{</tip>}}


2. Now, pull the image you need from `private-registry.nginx.com`. You can find the right image for your needs in the [Tech Specs guide]({{< relref "technical-specifications#images-with-nginx-plus" >}}). For example:

   - For NGINX Plus Ingress Controller, run the following command:

       ```shell
       docker pull private-registry.nginx.com/nginx-ic/nginx-plus-ingress:3.2.1
       ```

   - For NGINX Plus Ingress Controller with App Protect WAF, run:

       ```shell
       docker pull private-registry.nginx.com/nginx-ic-nap/nginx-plus-ingress:3.2.1
       ```

   - For NGINX Plus Ingress Controller with App Protect DoS:

       ```shell
       docker pull private-registry.nginx.com/nginx-ic-dos/nginx-plus-ingress:3.2.1
       ```

   - You can also use the Docker registry API to list the available image tags. For example:

        ```json
        $ curl https://private-registry.nginx.com/v2/nginx-ic/nginx-plus-ingress/tags/list --key <path-to-client.key> --cert <path-to-client.cert> | jq
        {
          "name": "nginx-ic/nginx-plus-ingress",
          "tags": [
            "3.2.1-alpine",
            "3.2.1-ubi",
            "3.2.1"
          ]
        }

       $ curl https://private-registry.nginx.com/v2/nginx-ic-nap/nginx-plus-ingress/tags/list --key <path-to-client.key> --cert <path-to-client.cert> | jq
        {
          "name": "nginx-ic-nap/nginx-plus-ingress",
          "tags": [
            "3.2.1-ubi",
            "3.2.1"
          ]
        }

       $ curl https://private-registry.nginx.com/v2/nginx-ic-dos/nginx-plus-ingress/tags/list --key <path-to-client.key> --cert <path-to-client.cert> | jq
        {
          "name": "nginx-ic-dos/nginx-plus-ingress",
          "tags": [
            "3.2.1-ubi",
            "3.2.1"
          ]
        }
       ```

3. Once you've pulled the image, tag and upload it to your private registry:

   - First, log in to your registry with `docker login <my-docker-registry>`.
   - Then, replace `<my-docker-registry>` in the examples below with your registry's path.

   ```shell
   docker tag private-registry.nginx.com/nginx-ic/nginx-plus-ingress:3.2.1 <my-docker-registry>/nginx-ic/nginx-plus-ingress:3.2.1
   docker push <my-docker-registry>/nginx-ic/nginx-plus-ingress:3.2.1
   ```

   or for NGINX App Protect WAF enabled image

   ```shell
   docker tag private-registry.nginx.com/nginx-ic-nap/nginx-plus-ingress:3.2.1 <my-docker-registry>/nginx-ic-nap/nginx-plus-ingress:3.2.1
   docker push <my-docker-registry>/nginx-ic-nap/nginx-plus-ingress:3.2.1
   ```

   or for NGINX App Protect DoS enabled image

   ```shell
   docker tag private-registry.nginx.com/nginx-ic-dos/nginx-plus-ingress:3.2.1 <my-docker-registry>/nginx-ic-dos/nginx-plus-ingress:3.2.1
   docker push <my-docker-registry>/nginx-ic-dos/nginx-plus-ingress:3.2.1
   ```
