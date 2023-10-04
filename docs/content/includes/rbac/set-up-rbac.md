---
docs: 
---

{{<call-out "important" "Admin access required" >}}To complete these steps you need admin access to your cluster. Refer to to your Kubernetes platform's documentation to set up admin access. For Google Kubernetes Engine (GKE), you can refer to their [Role-Based Access Control guide](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control).{{</call-out>}}

1. Create a namespace and a service account:

    ```shell
    kubectl apply -f common/ns-and-sa.yaml
    ```

2. Create a cluster role and binding for the service account:

    ```shell
    kubectl apply -f rbac/rbac.yaml
    ```

3. (App Protect only) Create the *App Protect* role and binding:

    ```shell
    kubectl apply -f rbac/ap-rbac.yaml
    ```

4. (App Protect DoS only) Create the *App Protect DoS* role and binding:

    ```shell
    kubectl apply -f rbac/apdos-rbac.yaml
    ```
