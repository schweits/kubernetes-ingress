---
title: Getting the F5 Registry NGINX Ingress Controller Image
description: "This guide walks you through the process of pulling an F5 NGINX Plus Ingress Controller image from the F5 Docker registry and uploading it to your private registry."
weight: 10
doctypes: ["install"]
toc: true
docs: "DOCS-605"
---

To complete the steps in this guide, you'll need your NGINX Ingress Controller subscription certificate and key. Keep in mind that an NGINX Plus certificate and key won't work for this. If you're looking for alternative methods, here are some:

- [Install using a JWT token in a Docker Config Secret]({{< relref "using-the-jwt-token-docker-secret" >}}).
- [Build the Ingress Controller image]({{< relref "building-ingress-controller-image" >}}) using the source code from the GitHub repository and your NGINX Plus subscription certificate and key.
- For NGINX Ingress Controller based on NGINX OSS, you can pull the [nginx/nginx-ingress image](https://hub.docker.com/r/nginx/nginx-ingress/) from DockerHub.

## Prerequisites

Before you start, you'll need these installed on your machine:

- [Docker](https://www.docker.com/products/docker) v18.09 or higher.
- An NGINX Ingress Controller subscription. Download both the certificate (*nginx-repo.crt*) and key (*nginx-repo.key*) from [MyF5](https://my.f5.com).

## Pull an image using Docker and push it to a private registry

1. tart by configuring Docker to communicate with the F5 Container registry at `private-registry.nginx.com`. If you're using Linux, create a directory under `/etc/docker/certs.d` and name it `private-registry.nginx.com`. Then copy your certificate and key into this folder, renaming the certificate to have a _.cert_ extension. Follow these commands:

    ```shell
    mkdir -p /etc/docker/certs.d/private-registry.nginx.com
    cp <path-to-your-nginx-repo.crt> /etc/docker/certs.d/private-registry.nginx.com/client.cert
    cp <path-to-your-nginx-repo.key> /etc/docker/certs.d/private-registry.nginx.com/client.key
    ```

    <br>

    {{<tip>}}The example above is for Linux. For Mac or Windows, consult the [Docker for Mac](https://docs.docker.com/docker-for-mac/#add-client-certificates) or [Docker for Windows](https://docs.docker.com/docker-for-windows/#how-do-i-add-client-certificates) documentation. For more details on Docker Engine security, you can refer to the [Docker Engine Security documentation](https://docs.docker.com/engine/security/).
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

   <br>

   or for NGINX App Protect WAF enabled image:

   ```shell
   docker tag private-registry.nginx.com/nginx-ic-nap/nginx-plus-ingress:3.2.1 <my-docker-registry>/nginx-ic-nap/nginx-plus-ingress:3.2.1
   docker push <my-docker-registry>/nginx-ic-nap/nginx-plus-ingress:3.2.1
   ```

   <br>

   or for NGINX App Protect DoS enabled image:

   ```shell
   docker tag private-registry.nginx.com/nginx-ic-dos/nginx-plus-ingress:3.2.1 <my-docker-registry>/nginx-ic-dos/nginx-plus-ingress:3.2.1
   docker push <my-docker-registry>/nginx-ic-dos/nginx-plus-ingress:3.2.1
   ```

<br>

## Troubleshooting

If you run into issues while following this guide, here are some solutions for the most common problems.

### Can't pull the image from F5 Container Registry

- **Possible Reason**: Incorrect certificate or key.
- **Solution**: Double-check that you've downloaded the correct NGINX Ingress Controller certificate and key from [MyF5](https://my.f5.com). An NGINX Plus certificate and key won't work here.

### Docker throws "unauthorized" error while pulling image

- **Possible Reason**: Authentication failure.
- **Solution**: Make sure you've placed the certificate and key in the correct directory and named them correctly. Also, confirm you're targeting `private-registry.nginx.com`.

### Can't tag or upload the image to your private registry

- **Possible Reason**: Not logged in or incorrect path.
- **Solution**: Run `docker login <my-docker-registry>` to make sure you're logged in. Check your registry's path and confirm that you've replaced `<my-docker-registry>` in the example commands with the correct path.

### Docker throws "TLS handshake timeout" or "certificate signed by unknown authority"

- **Possible Reason**: Certificate issue.
- **Solution**: Verify the certificate and key are correctly placed in the Docker certificate directory. Also, make sure you renamed the certificate to have a _.cert_ extension. Refer to the [Docker Engine Security documentation](https://docs.docker.com/engine/security/) for additional guidance.

### `curl` command returns errors while listing available image tags

- **Possible Reason**: Wrong key or certificate path in the `curl` command.
- **Solution**: Double-check the paths you're using in the `curl` command. Make sure you're using the correct key and certificate to authenticate against the Docker registry API.

### Image not found in your private registry after upload

- **Possible Reason**: Failed upload or wrong tag.
- **Solution**: Make sure you've tagged the image correctly before pushing it to your private registry. Double-check the output of the `docker push` command to see if the upload succeeded.
