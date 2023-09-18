---
title: Ingress
doctypes: ["concept"]
---

## What is the Ingress?

The Ingress is a Kubernetes resource that lets you configure an HTTP load balancer for applications running on Kubernetes, represented by one or more [Services](https://kubernetes.io/docs/concepts/services-networking/service/). Such a load balancer is necessary to deliver those applications to clients outside of the Kubernetes cluster.

The Ingress resource supports the following features:

- **Content-based routing**:
  - *Host-based routing*. For example, routing requests with the host header `foo.example.com` to one group of services and the host header `bar.example.com` to another group.
  - *Path-based routing*. For example, routing requests with the URI that starts with `/serviceA` to service A and requests with the URI that starts with `/serviceB` to service B.
- **TLS/SSL termination** for each hostname, such as `foo.example.com`.

See the [Ingress Reference Documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/) to learn more about the Ingress resource.