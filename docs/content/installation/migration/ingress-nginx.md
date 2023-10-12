---
title: "Migrating from Ingress-NGINX Controller to NGINX Ingress Controller"
date: 2023-09-29T16:31:21+01:00
description: "This document describes how to migrate from the community-maintained Ingress-NGINX Controller to the F5 NGINX Ingress Controller."
weight: 100
toc: true
tags: [ "docs" ]
docs: "DOCS-000"
categories: ["installation", "platform management"]
doctypes: ["tutorial"]
journeys: ["getting started"]
personas: ["devops"]
authors: ["Jason Williams"]
---

<br>

## Overview

This page explains two different ways to migrate from the community-maintained [Ingress-NGINX Controller](https://github.com/kubernetes/ingress-nginx) project to NGINX Ingress Controller: using NGINX's Ingress Resources or with Kubernetes's built-in Ingress Resources. This is typically because of implementation differences, and to take advantage of features such as [NGINX Plus integration]({{<relref "overview/nginx-plus">}}).

<!-- To understand the differences, you may wish to read [Which Ingress Controller Do I Need?]({{<relref "overview/controller-comparison">}}). -->

The information in this guide is extracted from a free eBook called "_Kubernetes Ingress Controller Deployment and Security with NGINX_", which can be downloaded from the [NGINX Library](https://www.nginx.com/resources/library/kubernetes-ingress-controller-deployment-security-nginx/).

## Before you begin

To complete the instructions in this guide, you need the following:

- A working knowledge of [Ingress Controllers]({{<relref "glossary.md#ingress-controller-ingress-controller">}}).
- An [NGINX Ingress Controller installation]({{<relref "installation/installing-nic">}}) on the same host as an existing Ingress-NGINX Controller.

There are two primary paths for migrating between the community Ingress-NGINX Controller to NGINX Ingress Controller: 

- Using NGINX Ingress Resources
- Using Kubernetes Ingress Resources.

## Migration with NGINX Ingress resources
This path uses Kubernetes Ingress Resources to set root permissions, then NGINX Ingress Resources for configuration using custom resource definitions (CRDs):

* [VirtualServer and VirtualServerRoute]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}})
* [TransportServer]({{<relref "configuration/transportserver-resource">}})
* [GlobalConfiguration]({{<relref "configuration/global-configuration/globalconfiguration-resource">}})
* [Policy]({{<relref "configuration/policy-resource">}})

### Configuring SSL termination and HTTP path-based routing
The following two code examples correspond to a Kubernetes Ingress Resource and an [NGINX VirtualServer Resource]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}}). Although the syntax and indentation is different, they accomplish the same basic Ingress functions, used for SSL termination and Layer 7 path-based routing.

**Kubernetes Ingress Resource**
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx-test
spec:
  tls:
    - hosts:
      - foo.bar.com
      secretName: tls-secret
  rules:
    - host: foo.bar.com
      http:
        paths:
        - path: /login
          backend:
            serviceName: login-svc
            servicePort: 80
        - path: /billing
            serviceName: billing-svc
            servicePort: 80
```

**NGINX VirtualServer Resource**
```yaml
apiVersion: networking.k8s.io/v1
kind: VirtualServer
metadata:
  name: nginx-test
spec:
  host: foo.bar.com
  tls:
    secret: tls-secret
  upstreams:
    - name: login
      service: login-svc
      port: 80
    - name: billing
      service: billing-svc
      port: 80
  routes:
  - path: /login
    action:
      pass: login
  - path: /billing
    action:
      pass: billing
```

### Configuring TCP/UDP load balancing and TLS passthrough
NGINX Ingress Controller exposes TCP and UDP services using [TransportServer]({{<relref "configuration/transportserver-resource">}}) and [GlobalConfiguration]({{<relref "configuration/global-configuration/globalconfiguration-resource">}}) resources. These resources provide a broad range of options for TCP/UDP and TLS Passthrough load balancing. By contrast, the community Ingress-NGINX Controller exposes TCP/UDP services by using a Kubernetes ConfigMap object.

---

### Convert Ingress-NGINX Controller annotations to NGINX Ingress resources
Kubernetes deployments often need to extend basic Ingress rules for advanced use cases such as canary and blue-green deployments, traffic throttling, and ingress-egress traffic manipulation. The community Ingress-NGINX Controller implements many of these using Kubernetes annotations with custom Lua extensions.

These custom Lua extensions are intended for specific NGINX Ingress resource definitions and may not be as granular as required for advanced use cases. The following examples show how to convert these annotations into NGINX Ingress Controller Resources.

---

#### Canary deployments
Canary and blue-green deployments allow you to push code changes to production environments without disrupting existing users. NGINX Ingress Controller runs them on the data plane: to migrate from the community Ingress-NGINX Controller, you must map the latter's annotations to [VirtualServer and VirtualServerRoute resources]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}}).

The Ingress-NGINX Controller evaluates canary annotations in the following order:

1. `nginx.ingress.kubernetes.io/canary-by-header`
1. `nginx.ingress.kubernetes.io/canary-by-cookie`
1. `nginx.ingress.kubernetes.io/canary-by-weight`

For NGINX Ingress Controller to evalute them the same way, they must appear in the same order in the VirtualServer or VirtualServerRoute Manifest.

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/canary: "true"
nginx.ingress.kubernetes.io/canary-by-header: "httpHeader"
```

**NGINX Ingress Controller**
```yaml
matches:
- conditions:
  - header: httpHeader
      value: never
  action:
    pass: echo
  - header: httpHeader
      value: always
  action:
    pass: echo-canary
action:
  pass: echo
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/canary: "true"
nginx.ingress.kubernetes.io/canary-by-header: "httpHeader"
nginx.ingress.kubernetes.io/canary-by-header-value: "my-value"
```

**NGINX Ingress Controller**
```yaml
matches:
- conditions:
  - header: httpHeader
      value: my-value
  action:
    pass: echo-canary
action:
  pass: echo
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/canary: "true"
nginx.ingress.kubernetes.io/canary-by-cookie: "cookieName"
```

**NGINX Ingress Controller**
```yaml
matches:
- conditions:
  - cookie: cookieName
      value: never
  action:
    pass: echo
  - cookie: cookieName
      value: always
  action:
    pass: echo-canary
action:
  pass: echo
```

---

#### Traffic control
Environments using microservices tend to use extensive traffic-control policies to manage ephemeral applications using circuit breaking and rate and connection limiting to prevent error conditions due to unhealthy states or abnormal behavior.

The following examples map Ingress-NGINX Controller annotations to NGINX [VirtualServer and VirtualServerRoute resources]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}}) for rate limiting, custom HTTP errors, custom default backend and URI rewriting.

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/custom-http-errors: "code"

nginx.ingress.kubernetes.io/default-backend: "default-svc"
```

**NGINX Ingress Controller**
```yaml
errorPages:
- codes: [code]
    redirect:
      code: 301
      url: default-svc
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/limit-connections: "number"
```

**NGINX Ingress Controller**
```yaml
http-snippets: |
    limit_conn_zone $binary_remote_addr zone=zone_name:size;
routes:
- path: /path
    location-snippets: |
      limit_conn zone_name number;
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/limit-rate: "number"
nginx.ingress.kubernetes.io/limit-rate-after: "number"
```

**NGINX Ingress Controller**
```yaml
location-snippets: |
    limit_rate number;

    limit_rate_after number;
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/limit-rpm: "number"
nginx.ingress.kubernetes.io/limit-burst-multiplier: "multiplier"
```

**NGINX Ingress Controller**
```yaml
rateLimit:
    rate: numberr/m

    burst: number * multiplier
    key: ${binary_remote_addr}
    zoneSize: size
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/limit-rps: "number"
nginx.ingress.kubernetes.io/limit-burst-multiplier: "multiplier"
```

**NGINX Ingress Controller**
```yaml
rateLimit:
    rate: numberr/s

    burst: number * multiplier
    key: ${binary_remote_addr}
    zoneSize: size
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/limit-whitelist: "CIDR"
```

**NGINX Ingress Controller**
```yaml
http-snippets: |
server-snippets: |
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/rewrite-target: "URI"
```

**NGINX Ingress Controller**
```yaml
rewritePath: "URI"
```

There are four Ingress-NGINX Controller annotations without NGINX Ingress resource fields yet: they must be handled using snippets.

- `nginx.ingress.kubernetes.io/limit-connections`
- `nginx.ingress.kubernetes.io/limit-rate`
- `nginx.ingress.kubernetes.io/limit-rate-after`
- `nginx.ingress.kubernetes.io/limit-whitelist`

---

#### Header manipulation
Manipulating HTTP headers is useful in many cases, as they contain information that is important and relevant to systems involved in HTTP transactions. The community Ingress-NGINX Controller supports enabling and configuring cross-origin resource sharing (CORS) headings used by AJAX applications, where front-end Javascript code interacts with backend applications or web servers.

These code blocks show how the Ingress-NGINX annotations correspond to NGINX Ingress Controller [VirtualServer and VirtualServerRoute resources]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}}).

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/enable-cors: "true"
nginx.ingress.kubernetes.io/cors-allow-credentials: "true"

nginx.ingress.kubernetes.io/cors-allow-headers: "X-Forwarded-For"

nginx.ingress.kubernetes.io/cors-allow-methods: "PUT, GET, POST, OPTIONS"

nginx.ingress.kubernetes.io/cors-allow-origin: "*"

nginx.ingress.kubernetes.io/cors-max-age: "seconds"
```

**NGINX Ingress Controller**
```yaml
responseHeaders:
  add:
    - name: Access-Control-Allow-Credentials
      value: "true"
    - name: Access-Control-Allow-Headers
      value: "X-Forwarded-For"
    - name: Access-Control-Allow-Methods
      value: "PUT, GET, POST, OPTIONS"
    - name: Access-Control-Allow-Origin
      value: "*"
    - name: Access-Control-Max-Age
      value: "seconds"
```

---

#### Proxying and load balancing
NGINX Ingress Controller has multiple proxy and load balancing functionalities you may want to configure based on the use case, such as configuring the load balancing algorithm and the timeout and buffering settings for proxied connections.

This table shows how Ingress-NGINX Controller annotations map to statements in the upstream field for [VirtualServer and VirtualServerRoute resources]({{<relref "configuration/virtualserver-and-virtualserverroute-resources">}}), covering load balancing, proxy timeout, proxy buffering and connection routing for a services' ClusterIP address and port.

{{< bootstrap-table "table table-bordered table-striped table-responsive" >}}
| Ingress-NGINX Controller | NGINX Ingress Controller |
| ------------------------ | ------------------------ |
| nginx.ingress.kubernetes.io/load-balance | lb-method |
| nginx.ingress.kubernetes.io/proxy-buffering | buffering |
| nginx.ingress.kubernetes.io/proxy-buffers-number | buffers |
| nginx.ingress.kubernetes.io/proxy-buffer-size| buffers |
| nginx.ingress.kubernetes.io/proxy-connect-timeout | connect-timeout |
| nginx.ingress.kubernetes.io/proxy-next-upstream | next-upstream |
| nginx.ingress.kubernetes.io/proxy-next-upstream-timeout | next-upstream-timeout |
| nginx.ingress.kubernetes.io/proxy-read-timeout | read-timeout |
| nginx.ingress.kubernetes.io/proxy-send-timeout | send-timeout |
| nginx.ingress.kubernetes.io/service-upstream | use-cluster-ip |
{{% /bootstrap-table %}}

#### mTLS authentication

mTLS authentication is a way of enforcing mutual authentication on traffic entering and exiting a cluster (north-sourth traffic). This secure form of communication is common within a service mesh, commonly used in strict zero-trust environments.

NGINX Ingress Controller layer can handle mTLS authentication for end systems through the presentation of valid certificates for external connections. It accomplishes this through [Policy]({{<relref "configuration/policy-resource">}}) resources, which correspond to Ingress-NGINX Controller annotations for [client certificate authentication](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#client-certificate-authentication) and [backend certificate authentication](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#backend-certificate-authentication).

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/auth-tls-secret: secretName
nginx.ingress.kubernetes.io/auth-tls-verify-client: "on"
nginx.ingress.kubernetes.io/auth-tls-verify-depth: "1"
```

**NGINX Ingress Controller**
```yaml
ingressMTLS:
   clientCertSecret: secretName
   verifyClient: "on"

   verifyDepth: 1
```

---

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/proxy-ssl-secret: "secretName"
nginx.ingress.kubernetes.io/proxy-ssl-verify: "on|off"
nginx.ingress.kubernetes.io/proxy-ssl-verify-depth: "1"
nginx.ingress.kubernetes.io/proxy-ssl-protocols: "TLSv1.2"
nginx.ingress.kubernetes.io/proxy-ssl-ciphers: "DEFAULT"
nginx.ingress.kubernetes.io/proxy-ssl-name: "server-name"
nginx.ingress.kubernetes.io/proxy-ssl-server-name: "on|off"
```

**NGINX Ingress Controller**
```yaml
egressMTLS:
   tlsSecret: secretName

   verifyServer: true|false

   verifyDepth: 1

   protocols: TLSv1.2

   ciphers: DEFAULT

   sslName: server-name

   serverName: true|false
```

---

#### Session persistence with NGINX Plus
With [NGINX Plus]({{<relref "overview/nginx-plus">}}), you can use [Policy]({{<relref "configuration/policy-resource">}}) resources for session persistence, which have corresponding annotations for the community Ingress-NGINX Controller.

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/affinity: "cookie"
nginx.ingress.kubernetes.io/session-cookie-name: "cookieName"
nginx.ingress.kubernetes.io/session-cookie-expires: "x"
nginx.ingress.kubernetes.io/session-cookie-path: "/route"
nginx.ingress.kubernetes.io/session-cookie-secure: "true"
```

**NGINX Ingress Controller**
```yaml
sessionCookie:
  enable: true

  name: cookieName

  expires: xh

  path: /route

  secure: true
```

## Migration with Kubernetes Ingress resources
The other option for migrating from the community Ingress-NGINX Controller to NGINX Ingress Controller is using only [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) and [ConfigMaps](https://kubernetes.io/docs/concepts/configuration/configmap/) from standard Kubernetes resources, potentially relying on [mergeable Ingress types](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/mergeable-ingress-types).

This ensures that all configuration is kept in the Ingress object.

{{< warning >}}
Do not alter the `spec` field of the Ingress resource when taking this option.
{{< /warning >}}

### Advanced configuration with annotations
This table maps the Ingress-NGINX Controller annotations to NGINX Ingress Controller's equivalent annotations, and the respective NGINX Directive.

{{< bootstrap-table "table table-bordered table-striped table-responsive" >}}
| Ingress-NGINX Controller | NGINX Ingress Controller | NGINX Directive |
| ------------------------ | ------------------------ | --------------- |
| [`nginx.ingress.kubernetes.io/configuration-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#configuration-snippet) | [`nginx.org/location-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#snippets-and-custom-templates) | N/A |
| [`nginx.ingress.kubernetes.io/load-balance`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#custom-nginx-load-balancing) (1) |  [`nginx.org/lb-method`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#backend-services-upstreams) | [`random two least_conn`](https://nginx.org/en/docs/http/ngx_http_upstream_module.html#random) |
| [`nginx.ingress.kubernetes.io/proxy-buffering`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#proxy-buffering) | [`nginx.org/proxy-buffering`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_buffering`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffering) |
| [`nginx.ingress.kubernetes.io/proxy-buffers-number`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#proxy-buffers-number) | [`nginx.org/proxy-buffers`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_buffers`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffers) |
| [`nginx.ingress.kubernetes.io/proxy-buffer-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#proxy-buffer-size) | [`nginx.org/proxy-buffer-size`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_buffer_size`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffer_size) |
| [`nginx.ingress.kubernetes.io/proxy-connect-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#custom-timeouts) | [`nginx.org/proxy-connect-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_connect_timeout`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_connect_timeout) |
| [`nginx.ingress.kubernetes.io/proxy-read-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#custom-timeouts) | [`nginx.org/proxy-read-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_read_timeout`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_read_timeout) |
| [`nginx.ingress.kubernetes.io/proxy-send-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#custom-timeouts) | [`nginx.org/proxy-send-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#general-customization) | [`proxy_send_timeout`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_send_timeout) |
| [`nginx.ingress.kubernetes.io/rewrite-target`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#rewrite) | [`nginx.org/rewrites`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#request-uriheader-manipulation) | [`rewrite`](https://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite) |
| [`nginx.ingress.kubernetes.io/server-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#server-snippet)| [`nginx.org/server-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#snippets-and-custom-templates) | N/A |
| [`nginx.ingress.kubernetes.io/ssl-redirect`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#server-side-https-enforcement-through-redirect) | [`ingress.kubernetes.io/ssl-redirect`](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#auth-and-ssltls) | N/A (2) |
{{% /bootstrap-table %}}

1. Ingress-NGINX Controller implements some of its load balancing algorithms with Lua, which may not have an equivalent in NGINX Ingress Controller.
1. To redirect HTTP (80) traffic to HTTPS (443), NGINX Ingress Controller uses native NGINX `if` conditions while Ingress-NGINX Controller uses Lua.

The following two snippets outline Ingress-NGINX Controller annotations that correspond to annotations for NGINX Ingress Controller with NGINX Plus.

**Ingress-NGINX Controller**
```yaml
nginx.ingress.kubernetes.io/affinity: "cookie"
nginx.ingress.kubernetes.io/session-cookie-name: "cookie_name"
nginx.ingress.kubernetes.io/session-cookie-expires: "seconds"
nginx.ingress.kubernetes.io/session-cookie-path: "/route"
```

**NGINX Ingress Controller (with NGINX Plus)**
```yaml
nginx.com/sticky-cookie-services: "serviceName=example-svc cookie_name expires=time path=/route"
```

{{< note >}}
NGINX Ingress Controller has additional annotations for features using NGINX Plus that have no Ingress-NGINX Controller equivalent, such as active health checks and authentication using JSON Web Tokens (JWTs).
{{< /note >}}

### Global configuration with ConfigMaps

This table maps the Ingress-NGINX Controller ConfigMap keys to NGINX Ingress Controller's equivalent ConfigMap keys.

<!-- {{< note >}}
Some of the key names are identical, and each Ingress Controller has ConfigMap keys that the other does not (Which are indicated).
{{< /note >}} -->

{{< bootstrap-table "table table-bordered table-striped table-responsive" >}}
| Ingress-NGINX Controller | NGINX Ingress Controller |
| ------------------------ | ------------------------ |
| [`disable-access-log`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#disable-access-log) | [`access-log-off`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#logging) |
| [`error-log-level`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#error-log-level) | [`error-log-level`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#logging) |
| [`hsts`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#hsts) | [`hsts`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`hsts-include-subdomains`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#hsts-include-subdomains) | [`hsts-include-subdomains`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls)       |
| [`hsts-max-age`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#hsts-max-age) | [`hsts-max-age`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`http-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#http-snippet) | [`http-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates) |
| [`keep-alive`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#keep-alive) | [`keepalive-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`keep-alive-requests`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#keep-alive-requests) | [`keepalive-requests`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`load-balance`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#load-balance) | [`lb-method`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#backend-services-upstreams) |
| [`location-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#location-snippet) | [`location-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates) |
| [`log-format-escape-json`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#log-format-escape-json) | [`log-format-escaping: "json"`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#logging) |
| [`log-format-stream`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#log-format-stream) | [`stream-log-format`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#logging) |
| [`log-format-upstream`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#log-format-upstream) | [`log-format`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#logging) |
| [`main-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#main-snippet) | [`main-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates) |
| [`max-worker-connections`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#max-worker-connections) | [`worker-connections`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`max-worker-open-files`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#max-worker-open-files) | [`worker-rlimit-nofile`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-body-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-body-size) | [`client-max-body-size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-buffering`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-buffering) | [`proxy-buffering`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-buffers-number`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-buffers-number) | [`proxy-buffers: number size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-buffer-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-buffer-size) | [`proxy-buffers: number size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-connect-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-connect-timeout) | [`proxy-connect-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-read-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-read-timeout) | [`proxy-read-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-send-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-send-timeout) | [`proxy-send-timeout`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`server-name-hash-bucket-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#server-name-hash-bucket-size) | [`server-names-hash-bucket-size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`proxy-headers-hash-max-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#proxy-headers-hash-max-size) | [`server-names-hash-max-size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`server-snippet`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#server-snippet) | [`server-snippets`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates) |
| [`server-tokens `](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#server-tokens) | [`server-tokens`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`ssl-ciphers`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-ciphers) | [`ssl-ciphers`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`ssl-dh-param`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-dh-param) | [`ssl-dhparam-file`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`ssl-protocols`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-protocols) | [`ssl-protocols`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`ssl-redirect`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-redirect) | [`ssl-redirect`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls) |
| [`upstream-keepalive-connections`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#upstream-keepalive-connections) | [`keepalive`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#backend-services-upstreams) |
| [`use-http2`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#use-http2) | [`http2`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#listeners) |
| [`use-proxy-protocol`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#use-proxy-protocol) | [`proxy-protocol`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#listeners) |
| [`variables-hash-bucket-size`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#variables-hash-bucket-size)     | [`variables-hash-bucket-size`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`worker-cpu-affinity`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#worker-cpu-affinity) | [`worker-cpu-affinity`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`worker-processes`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#worker-processes) | [`worker-processes`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
| [`worker-shutdown-timeout`](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#worker-shutdown-timeout) | [`worker-shutdown-timeole`](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#general-customization) |
{{% /bootstrap-table %}}