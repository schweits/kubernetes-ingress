---
title: Building NGINX Ingress Controller
description: "Learn how to build an NGINX Ingress Controller image from source codes and upload it to a private Docker registry. You'll also find information on the Makefile targets and variables."
weight: 200
doctypes: ["installation"]
toc: true
---

{{<custom-styles>}}

## Prerequisites

To get started, you need the following software installed on your machine:

- [Docker v19.03 or higher](https://docs.docker.com/engine/release-notes/19.03/)
- [GNU Make](https://www.gnu.org/software/make/)
- [git](https://git-scm.com/)
- [OpenSSL](https://www.openssl.org/), optionally, if you would like to generate a self-signed certificate and a key for the default server.
- For NGINX Plus users, download the certificate (_nginx-repo.crt_) and key (_nginx-repo.key_) from [MyF5](https://my.f5.com).

Although NGINX Ingress Controller is written in Golang, you don't need to have Golang installed. You can either download the precompiled binary file or build NGINX Ingress Controller in a Docker container.

---

## Prepare the environment {#prepare-environment}

This section helps you get your system ready so you can build and push the NGINX Ingress Controller image.

1. Sign in to your private registry. Replace `<my-docker-registry>` with the path to your own private registry.

    ```shell
    docker login <my-docker-registry>
    ```

2. Clone the NGINX Ingress Controller GitHub repository. Replace `<version_number>` with the version of NGINX Ingress Controller you want.

    ```shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch <version_number>
    cd kubernetes-ingress
    ```

    For instance if you want to clone version v3.2.1, the commands to run would be:

    ```shell
    git clone https://github.com/nginxinc/kubernetes-ingress.git --branch v3.2.1
    cd kubernetes-ingress
    ```

--- 

## Build the NGINX Ingress Controller image {#build-image}

Now that your environment is set up, this section will show you how to build the NGINX Ingress Controller image for either NGINX or NGINX Plus.

{{<note>}}If you have a local Golang environment and want to build the binary yourself, remove `TARGET=download` from the make commands. If you don't have Golang but still want to build the binary, use `TARGET=container`.{{</note>}}

### For NGINX

1. Build the image. Replace `<my-docker-registry>` with your private registry's path.

    - For a Debian-based image:

        ```shell
        make debian-image PREFIX=<my-docker-registry>/nginx-ingress TARGET=download
        ```

    - For an Alpine-based image:

        ```shell
        make alpine-image PREFIX=<my-docker-registry>/nginx-ingress TARGET=download
        ```

    **What to expect**: The image is built and tagged with a version number, which is derived from the `VERSION` variable in the [_Makefile_](#makefile-details). This version number is used for tracking and deployment purposes.

### For NGINX Plus

1. Place your NGINX Plus license files (_nginx-repo.crt_ and _nginx-repo.key_) in the project's root folder. To verify they're in place, run:

    ```shell
    ls nginx-repo.*
    ```

    You should see:

    ```shell
    nginx-repo.crt  nginx-repo.key
    ```
 
2. Build the image. Replace `<my-docker-registry>` with your private registry's path.

    ```shell
    make debian-image-plus PREFIX=<my-docker-registry>/nginx-plus-ingress TARGET=download
    ```

    <br>

     **What to expect**: The image is built and tagged with a version number, which is derived from the `VERSION` variable in the [_Makefile_](#makefile-details). This version number is used for tracking and deployment purposes.

{{<note>}}In the event a patch version of NGINX Plus is released, make sure to rebuild your image to get the latest version. If your system is caching the Docker layers and not updating the packages, add `DOCKER_BUILD_OPTIONS="--pull --no-cache"` to the make command.{{</note>}}

---

## Push the image to your private registry {#push-image}

After building the NGINX Ingress Controller image, this section provides instructions for uploading the image to your own private registry.

### For NGINX

1. Upload the NGINX image. If you're using a custom tag, append `TAG=your-tag` to the command. Replace `<my-docker-registry>` with your private registry's path.

    ```shell
    make push PREFIX=<my-docker-registry>/nginx-ingress
    ```

### For NGINX Plus

1. Upload the NGINX Plus image. Like with the NGINX image, if you're using a custom tag, add `TAG=your-tag` to the end of the command. Replace `<my-docker-registry>` with your private registry's path.

    ```shell
    make push PREFIX=<my-docker-registry>/nginx-plus-ingress
    ```

---

Certainly, let's incorporate those changes to make the section more consistent, readable, and easier to maintain.

---

## Makefile details {#makefile-details}

This section provides comprehensive information on the targets and variables available in the _Makefile_. These targets and variables allow you to customize how you build, tag, and push your NGINX or NGINX Plus images.

### Key Makefile targets

{{<tip>}}To view available _Makefile_ targets, run `make` with no target or type `make help`.{{</tip>}}

Key targets include:

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
| **ubi-image-nap-dos-plus**      | <p>Builds a UBI-based image with NGINX Plus, [NGINX App Protect WAF](/nginx-app-protect/) and the [NGINX App Protect DoS](/nginx-app-protect-dos/) module for [OpenShift](https://www.openshift.com/) clusters.</p> <p> **Important**: Save your RHEL organization and activation keys in a file named _rhel_license_ at the project root.</p> <p> For instance:</p> <pre>RHEL_ORGANIZATION=1111111<br />RHEL_ACTIVATION_KEY=your-key</pre>|
{{</bootstrap-table>}}

### Additional useful targets

{{<bootstrap-table "table table-striped table-bordered">}}
| <div style="width:200px">Target</div> | Description   |
|---------------------------------------|---------------|
| **push**                              | Pushes the built image to the Docker registry. Configures with `PREFIX` and `TAG`.  |
| **all**                               | Runs `test`, `lint`, `verify-codegen`, `update-crds`, and `debian-image`. Stops and reports an error if any of these targets fail.  |
| **test**                              | Runs unit tests.  |
| **certificate-and-key**               | NGINX Ingress Controller requires a certificate and key for the default HTTP/HTTPS server. You have several options: <ul><li>Reference them in a TLS Secret in a command-line argument to NGINX Ingress Controller.</li><li>Add them to the image in in a file in PEM format as `/etc/nginx/secrets/default`.</li><li>Generate a self-signed certificate and key with this target.</li></ul>Note, you must include the `ADD` instruction in your Dockerfile to copy the cert and key to the image. |
{{</bootstrap-table>}}

### Makefile variables you can customize {#makefile-variables}

The _Makefile_ includes several key variables. You have the option to either modify these variables directly in the _Makefile_ or override them when you run the `make` command.

{{<bootstrap-table "table table-striped table-bordered">}}
| <div style="width:200px">Variable</div> | Description   |
|-----------------------------------------|---------------|
| **ARCH**                                | Defines the architecture for the image and binary. The default is `amd64`, but you can also choose from `arm64`, `arm`, `ppc64le`, and `s390x`.   |
| **PREFIX**                              | Gives the image its name. The default is `nginx/nginx-ingress`.  |
| **TAG**                                 | Adds a tag to the image. This is often the version of the NGINX Ingress Controller.   |
| **DOCKER_BUILD_OPTIONS**                | Allows for additional [options](https://docs.docker.com/engine/reference/commandline/build/#options) during the `docker build` process, like `--pull`.  |
| **TARGET**                              | <p>Determines the build environment. NGINX Ingress Controller compiles locally in a Golang environment by default. Ensure the NGINX Ingress Controller repo resides in your `$GOPATH` if you select this option.</p><p>Alternatively, you can set `TARGET=container` to build using a Docker [Golang](https://hub.docker.com/_/golang/) container. To skip compiling the binary if you're on a specific tag or the latest `main` branch commit, set `TARGET=download`.</p>  |
{{</bootstrap-table>}}
