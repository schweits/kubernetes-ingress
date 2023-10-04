---
docs:
---

When you deploy NGINX Ingress Controller as a DaemonSet, Kubernetes creates an Ingress Controller pod on every node in the cluster.

{{<note>}}
For guidance on how to limit NGINX Ingress Controller to specific nodes rather than running NGINX Ingress Controller on every node in the cluster, refer to the Kubernetes [DaemonSet docs](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/).
{{</note>}}

- For NGINX, run:

    ```shell
    kubectl apply -f daemon-set/nginx-ingress.yaml
    ```

- For NGINX Plus, run:

    ```shell
    kubectl apply -f daemon-set/nginx-plus-ingress.yaml
    ```

    {{<note>}}Update `nginx-plus-ingress.yaml` to include the image you've chosen from the F5 Container registry or the custom container image you've built. {{</note>}}
