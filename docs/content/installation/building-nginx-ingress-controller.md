---
title: Building NGINX Ingress Controller
description: "Learn how to build an NGINX Ingress Controller image from source code. In this document, you'll find step-by-step instructions for building the image for NGINX or NGINX Plus and uploading it to a private registry. You'll also find information on the Makefile targets and variables."
weight: 200
doctypes: ["installation"]
toc: true
---

<br>

{{<see-also>}}If you'd rather use a pre-built image, you have several options:

- For NGINX Plus: See [Using NGINX Ingress Controller Plus JWT token in a Docker Config Secret]({{< relref "installation/nic-images/using-the-jwt-token-docker-secret.md" >}}) and [Getting the F5 Registry NGINX Ingress Controller Image]({{< relref "installation/nic-images/pulling-ingress-controller-image" >}}).
- For NGINX OSS: Get images from [DockerHub](https://hub.docker.com/r/nginx/nginx-ingress/) or [GitHub Container](https://github.com/nginxinc/kubernetes-ingress/pkgs/container/kubernetes-ingress).
{{</see-also>}}

## Prerequisites

To get started, you'll need the following software installed on your machine:

- [Docker v19.03 or higher](https://docs.docker.com/engine/release-notes/19.03/)
- [GNU Make](https://www.gnu.org/software/make/)
- [git](https://git-scm.com/)
- [OpenSSL](https://www.openssl.org/), optionally, if you would like to generate a self-signed certificate and a key for the default server.
- For NGINX Plus users, download the certificate (_nginx-repo.crt_) and key (_nginx-repo.key_) from [MyF5](https://my.f5.com).

Although NGINX Ingress Controller is written in Golang, you don't need to have Golang installed. You can either download the precompiled binary file or build NGINX Ingress Controller in a Docker container.

## Build the image and push it to a private registry

Here's how to create the NGINX Ingress Controller binary, build the image, and upload that image to your private repository.

{{<note>}}If you have a local Golang environment and would like to build the binary yourself, remove `TARGET=download` from the `make` commands. <br> If you don't have Golang but still want to build the binary, then use `TARGET=container`.{{</note>}}

### Initial steps

We'll guide you through building NGINX Ingress Controller v3.2.1. To build a different version, simply replace `v3.2.1` with your chosen version.

1. Begin by signing into your private registry using the `docker login` command. Replace `<my-docker-registry>` with the actual path to your private registry. If you're using Google Container Registry, you'll also need to run `gcloud auth login` and `gcloud auth configure-docker`.

    ```shell
    docker login <my-docker-registry>
    ```

2. Next, clone the NGINX Ingress Controller's GitHub repository, specifying the version you want. For instance, here's how to clone version `v3.2.1`:

    ```shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1
    cd kubernetes-ingress
    ```

### For NGINX

1. Build the image. Replace `<my-docker-registry>` with your private registry's path.

    ```shell
    make debian-image PREFIX=<my-docker-registry>/nginx-ingress TARGET=download
    ```

    or for an Alpine-based image:

    ```shell
    make alpine-image PREFIX=<my-docker-registry>/nginx-ingress TARGET=download
    ```

    <br>

    **Result**: The image `<my-docker-registry>/nginx-ingress:3.2.1` is built. The tag `3.2.1` comes from the `VERSION` variable defined in the [_Makefile_](#makefile-details).

2. Upload the image you've built. If you're using a custom tag, add `TAG=your-tag` to the end of the command.

    ```shell
    make push PREFIX=<my-docker-registry>/nginx-ingress
    ```

### For NGINX Plus

1. Place your NGINX Plus license files (`nginx-repo.crt` and `nginx-repo.key`) in the project's root folder.
2. Build the image. Replace `<my-docker-registry>` with your private registry's path.

    ```shell
    make debian-image-plus PREFIX=<my-docker-registry>/nginx-plus-ingress TARGET=download
    ```

    <br>

    **Result**: The image `<my-docker-registry>/nginx-ingress:3.2.1` is built. The tag `3.2.1` comes from the `VERSION` variable defined in the [_Makefile_](#makefile-details).

3. Upload the image you've built. If you're using a custom tag, add `TAG=your-tag` to the end of the command.

    ```shell
    make push PREFIX=<my-docker-registry>/nginx-plus-ingress
    ```

### Next Steps

For details on the available _Makefile_ targets and variables, proceed to the next section.

## Makefile details {#makefile-details}

### Makefile targets

{{<tip>}}Run the `make` command without a target or use `make help` to see a list of available _Makefile_ targets. {{</tip>}}

Here are some key targets:

{{<bootstrap-table "table table-striped table-bordered">}}
| <div style="width:200px">Target | Description                                                                                                                                                                                                  |
|---------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **build**                       | Creates the NGINX Ingress Controller binary with your local Go environment.                                                                                                                                  |
| **alpine-image**                | Builds an Alpine-based image with NGINX.                                                                                                                                                                     |
| **alpine-image-plus**           | Builds an Alpine-based image with NGINX Plus.                                                                                                                                                                |
| **alpine-image-plus-fips**      | Builds an Alpine-based image with NGINX Plus and FIPS.                                                                                                                                                       |
| **debian-image**                | Builds a Debian-based image with NGINX.                                                                                                                                                                      |
| **debian-image-plus**           | Builds a Debian-based image with NGINX Plus.                                                                                                                                                                 |
| **debian-image-nap-plus**       | Builds a Debian-based image with NGINX Plus and the [NGINX App Protect WAF](/nginx-app-protect/) module.                                                                                                     |
| **debian-image-dos-plus**       | Builds a Debian-based image with NGINX Plus and the [NGINX App Protect DoS](/nginx-app-protect-dos/) module.                                                                                                 |
| **debian-image-nap-dos-plus**   | Builds a Debian-based image with NGINX Plus, [NGINX App Protect WAF](/nginx-app-protect/) and [NGINX App Protect DoS](/nginx-app-protect-dos/) modules.                                                      |
| **ubi-image**                   | Builds a UBI-based image with NGINX for [OpenShift](https://www.openshift.com/) clusters.                                                                                                                    |
| **ubi-image-plus**              | Builds a UBI-based image with NGINX Plus for [OpenShift](https://www.openshift.com/) clusters.                                                                                                               |
| **ubi-image-nap-plus**          | Builds a UBI-based image with NGINX Plus and the [NGINX App Protect WAF](/nginx-app-protect/) module for [OpenShift](https://www.openshift.com/) clusters.                                                   |
| **ubi-image-dos-plus**          | Builds a UBI-based image with NGINX Plus and the [NGINX App Protect DoS](/nginx-app-protect-dos/) module for [OpenShift](https://www.openshift.com/) clusters.                                               |
| **ubi-image-nap-dos-plus**      | Builds a UBI-based image with NGINX Plus, [NGINX App Protect WAF](/nginx-app-protect/) and the [NGINX App Protect DoS](/nginx-app-protect-dos/) module for [OpenShift](https://www.openshift.com/) clusters. |
{{</bootstrap-table>}}

<br>

{{<important>}}
Store your RHEL organization and activation keys in a file named _rhel_license_ in the project root. For example:

```text
RHEL_ORGANIZATION=1111111
RHEL_ACTIVATION_KEY=your-key
```

{{</important>}}

<br>

Other useful targets:

{{<bootstrap-table "table table-striped table-bordered">}}
| <div style="width:200px">Target</div> | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **push**                              | Pushes the built image to the Docker registry. Uses `PREFIX` and `TAG` as settings.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| **all**                               | Runs `test`, `lint`, `verify-codegen`, `update-crds`, and `debian-image`. Stops and reports an error if any of the targets fail.                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| **test**                              | Runs unit tests.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| **certificate-and-key**               | The default HTTP/HTTPS server needs a certificate and a key. You have a few options: <ul><li>Reference them using a TLS Secret through a command-line argument when starting NGINX Ingress Controller.</li><li>Add them directly to the image as a PEM-formatted file at the path `/etc/nginx/secrets/default`.</li><li>Generate a self-signed certificate and key with this target.</li></ul> Note, if you choose to add the certificate and key directly, you must include an `ADD` instruction in your Dockerfile to copy the certificate and key to the image. |
{{</bootstrap-table>}}

<br>

### Makefile variables

The _Makefile_ includes the following main variables. You can customize these variables by changing them directly in the _Makefile_ or overriding them when running the `make` command.

{{<bootstrap-table "table table-striped table-bordered">}}
| <div style="width:200px">Variable</div> | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|-----------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **ARCH**                                | Sets the image and binary architecture. Defaults to `amd64`. Other common choices include `arm64`, `arm`, `ppc64le`, and `s390x`.                                                                                                                                                                                                                                                                                                                                                                              |
| **PREFIX**                              | Names the image. Defaults to `nginx/nginx-ingress`.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| **TAG**                                 | Assigns a tag to the image, usually set to the NGINX Ingress Controller's version.                                                                                                                                                                                                                                                                                                                                                                                                                             |
| **DOCKER_BUILD_OPTIONS**                | Provides extra [options](https://docs.docker.com/engine/reference/commandline/build/#options) for the `docker build` command, such as `--pull`.                                                                                                                                                                                                                                                                                                                                                                |
| **TARGET**                              | Sets the build environment. By default, NGINX Ingress Controller is compiled locally using a local Golang environment. If you choose this option, make sure the NGINX Ingress Controller's repo is in your `$GOPATH`. <br> To compile NGINX Ingress Controller using a Docker [Golang](https://hub.docker.com/_/golang/) container, set `TARGET=container`. If you've checked out a specific tag or are on the latest commit on the `main` branch, you can set `TARGET=download` to skip compiling the binary. |
{{</bootstrap-table>}}

<br>
