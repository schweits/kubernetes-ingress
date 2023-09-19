---
title: Ingress
doctypes: ["concept"]
---

_Ingress_ refers to an _Ingress Resource_, a Kubernetes API object which allows access to [Services](https://kubernetes.io/docs/concepts/services-networking/service/) within a cluster. They are managed by an [Ingress Controller]({{< relref "glossary/ingress-controller">}}).

_Ingress_ resources enable the following functionality:

- **Load balancing**, extended through the use of Services
- **Content-based routing**, using hosts and paths
- **TLS/SSL termination**, based on hostnames

For additional information, please read the official [Kubernetes Ingress Documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/).