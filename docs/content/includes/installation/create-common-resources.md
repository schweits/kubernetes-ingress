---
docs: 
---

In this section, you'll create resources common for most NGINX Ingress Controller installations:

1. (Optional) Create a secret containing a TLS certificate and key for the default NGINX server.

    {{<call-out "important" "Optional step" >}}
Complete this step only if you're using the [default server TLS secret]({{< relref "configuration/global-configuration/command-line-arguments#cmdoption-default-server-tls-secret.md" >}}) command-line argument. Otherwise, you can skip it. We recommend using your own certificate. 
{{</call-out>}}

    First, make sure you're in the `kubernetes-ingress/deployment` directory, and then run:

    ```shell
    kubectl apply -f ../examples/shared-examples/default-server-secret/default-server-secret.yaml
    ```

    {{<note>}}By default, the server returns a _404 Not Found_ page for all requests where no ingress rules are set up. Although we've provided a self-signed certificate and key for testing purposes, it's best to use your own. {{</note>}}

2. Create a ConfigMap to customize your NGINX settings:

    ```shell
    kubectl apply -f common/nginx-config.yaml
    ```

3. Create an IngressClass resource:

    ```shell
    kubectl apply -f common/ingress-class.yaml
    ```

    To make this NGINX Ingress Controller instance the default for your cluster, remove the comment from the annotation `ingressclass.kubernetes.io/is-default-class`. Doing this will automatically assign this IngressClass to any new Ingresses that don't specify an `ingressClassName`.

    {{<note>}}NGINX Ingress Controller won't start unless you create an IngressClass resource. {{</note>}}
