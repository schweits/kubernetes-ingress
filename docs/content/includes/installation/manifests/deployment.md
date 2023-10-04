---
docs:
---

When you deploy NGINX Ingress Controller as a deployment, Kubernetes automatically sets up a single NGINX Ingress Controller pod.

- For NGINX, run:

    ```shell
    kubectl apply -f deployment/nginx-ingress.yaml
    ```

- For NGINX Plus, run:

    ```shell
    kubectl apply -f deployment/nginx-plus-ingress.yaml
    ```

    {{<note>}}Update `nginx-plus-ingress.yaml` to include the image you've chosen from the F5 Container registry or the custom container image you've built. {{</note>}}
