---
title: Use Case Directory
description: |
  A directory of use cases and where to find documentation and examples.
weight: 1700
doctypes: ["concept"]
toc: true
docs: "DOCS-1220"
---

## NGINX Ingress Controller Use Case Directory

**Simplicity** – horizontally scalable single container architecture ensures safety in failures and simplicity in operation and scale.

**Enterprise grade** – NGINX powered, tried and true, battle tested data plane, performant, highly stable, and reliable.

**Scale** – Fronting 1000s of microservices with 1000s of pods powering the very large

**Wide protocol support** – HTTP/1.1, HTTP/2, gRPC, WebSockets, TCP, UDP (Layer 7 & Layer 4)

**Most capable in the market** – With the full and native power of NGINX / NGINX Plus under the hood the widest breadth of reverse-proxy, API Gateway, and security capabilities is built in.

**API Gateway** – implement various API Gateway use cases at your cluster edge. Method blocking, rate limiting, custom responses, error codes, etc. through the K8s API.

---
| **Capability** | **Resource** | **Enterprise exclusive** | **Documentation** | **Example** | **Lab Exercise** |
| --- | --- | --- | --- | --- | --- |
| **Custom Resources –** provide additional capabilities and granular controls of the configuration to scope aspects of the configuration to different teams, not achievable with Ingress resources. |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/)
[https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/)
[https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/](https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources)
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab9/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab9/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab9/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab9/readme.md)
 |
| **Load balancing** – even distribution of traffic across your service pods including support more multiple methods and advanced routing health checks | **Ingress, CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstream](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstream)
 |
 |
 |
| **wildcard hostname support** | **Ingress, CRD** |
 |
 |
 |
 |
| **Traffic splitting** – canary, blue/green, A/B testing, etc. | **CRD** |
 |
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/traffic-splitting](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/traffic-splitting)
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)

 |
| **Conditional routing** – applying conditional logic to traffic routing decisions based on headers, claims, etc. | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match)
[https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#condition](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#condition)
 |
 |
 |
| **Authentication at the Ingress Controller** – Either basic or OIDC or mTLS is fully supported and baked into the solution at the host or path | **CRD, Ingress** | **\*** | [**h** ttps://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#basicauth](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#basicauth)
https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/advanced-configuration-with-annotations/#auth-and-ssltls
[https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#oidc](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#oidc)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/basic-auth](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/basic-auth)
https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/basic-auth
[https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/oidc](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/oidc)
 |
 |
| **Authentication validation at the ingress controller –** the client presents a token for validation at the Ingress Controller. The token is validated using a local or remote secret | **CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt-using-local-kubernetes-secret](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt-using-local-kubernetes-secret)
[https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt-using-jwks-from-remote-location](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt-using-jwks-from-remote-location)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwt](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwt)
[https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwks](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwks)
 |
 |
| **Authorization –** based on JWT tokens, advanced authorization controls can be implemented at the hostname or path levels | **Snippet, CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt)
 |
 |
 |
| **Rate Limiting** – Supporting your security policy and protecting your APIs against intentional and unintentional DoS attacks. | **CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#ratelimit](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#ratelimit)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/rate-limit](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/rate-limit)
 |
 |
| **Request Queueing –** requests are queued if services are unavailable | **CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamqueue](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamqueue)
 |
 |
 |
| **Method filtering and logic** – easily define your list of allowed methods and response behaviors | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match)
 |
 |
 |
| **WAF Integration** – Best in class integration with the NGINX App Protect Web Application Firewall | **Ingress, CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/app-protect-waf](https://docs.nginx.com/nginx-ingress-controller/app-protect-waf)
 |
 |
 |
| **DoS Prevention Integration** – tight integration with NGINX App Protect Denial of Service | **Ingress, CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/app-protect-dos/](https://docs.nginx.com/nginx-ingress-controller/app-protect-dos/)
 |
 |
 |
| **mTLS encryption to the back end** – a Service Mesh is not required with modern or legacy applications that support SSL | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamtls](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamtls)
 |
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)
 |
| **TLS Termination** – offloading at the edge or in front of your mesh ingress takes the load so your applications don't need to for both Layer 4 and Layer 7. | **Ingress CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertls](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertls)
 |
 |
 |
| **HTTPS Redirection –** Redirect HTTP to HTTPS with a simple setting | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertlsredirect](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertlsredirect)
 |
 |
 |
| **Header manipulation** – add, remove, modify headers and header values in both requests and responses | **Ingress, CRD, snippet** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionproxyrequestheaders](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionproxyrequestheaders)
[https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionproxyrequestheaderssetheader](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionproxyrequestheaderssetheader)
[https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#addheader](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#addheader)
 |
 |
 |
| **Body manipulation** – request and response modification | **snippet** |
 |
 |
 |
 |
| **HSTS** – | **CRD, snippet** |
 | [**h** ttps://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#auth-and-ssltls)
 |
 |
 |
| **CORS** – | **snippet** |
 |
 |
 |
 |
| **Dual stack / IPv6 support –** for front end and / or upstream IPv6 services | **Native** |
 |
 |
 |
 |
| **Dynamic backend configuration** – K8s services are dynamic, pods are constantly changing, the backend service pods are reconfigured with no traffic impact or reload | **Native, N+** | **\*** |
 |
 |
 |
| **Reduced SSL handshake when path is not matched** – when certificates or services are defined in configuration but not available, NGINX SSL reject handshake is enabled to protect your services from malicious discovery through immediately interrupting the handshake process. This also improves the security posture through avoiding default certificates. | **Native** |
 |
 |
 |
 |
| **Access control policies** – allow or deny based on IP or CIDR in a stackable way | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#accesscontrol](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#accesscontrol)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/access-control](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/access-control)
 |
 |
| **Advanced routing rules** – variables, cookies, query parameters, regex path | **CRD, ingress, snippet** |
 |
 |
 |
 |
| **Layer 4 routing** – TCP and UDC routing support | **CRD, ingress** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/](https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/basic-tcp-udp](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/basic-tcp-udp)[https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/tcp-udp](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/tcp-udp)
 |
 |
| **Cert-manger integration** – with the custom VirtualServer resource, in addition to the Ingress resource | **CRD, ingress** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertlscertmanager](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualservertlscertmanager)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/certmanager](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/certmanager)
 |
 |
| **External-dns integration** – with the custom VirtualServer resource, in addition to the Ingress resource | **CRD, ingress** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualserverexternaldns](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#virtualserverexternaldns)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/external-dns](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/external-dns)
 |
 |
| **Cross namespace routing** – routing traffic to backends across different namespaces, supporting full namespace security isolation. | **CRD, ingress** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/cross-namespace-configuration/](https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/cross-namespace-configuration/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/cross-namespace-configuration](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/cross-namespace-configuration)
[https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/mergeable-ingress-types](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/mergeable-ingress-types)
 |
 |
| **mTLS to external services or applications** – TLS encryption and authentication to services within and outside the cluster | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#egressmtls](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#egressmtls)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/egress-mtls](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/egress-mtls)
 |
 |
| **Externalname service support** – forwarding of HTTP or TCP/UDP traffic to services or applications external to the cluster | **CRD** |
 |
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/externalname-services](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/externalname-services)
 |
 |
| **Custom (active) health checks** – granting advanced controls for health checks that drive traffic routing behavior to pods, this is in addition to the basic (passive) health checks which are always active | **CRD** |
 |
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/health-checks](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/health-checks)
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)
 |
| **Deep Service Insight –** Get the proxy's view of the health status of backend service pods expressed in an easy to query and consume endpoint. | **CRD** | **\*** | [https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/service-insight/](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/service-insight/) |
 |
 |
| **Ingress mTLS** – TLS authentication and encryption of incoming traffic to the edge of your cluster, and CRL support | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#ingressmtls](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#ingressmtls)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/ingress-mtls](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/ingress-mtls)
 |
 |
| **JWT**** Authorization **– support for JWT and variations to JWT for authorization and routing decisions |** Snippet, CRD **|** \*** | [https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt](https://docs.nginx.com/nginx-ingress-controller/configuration/policy-resource/#jwt)
[https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#match)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwt](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/jwt)
 |
 |
| **Path rewriting** – giving you the full controls to adapt to service changes at the ingress layer | **CRD** |
 |
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/rewrites](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/rewrites)
 |
 |
| **Path Redirect –** redirect a request to a different URL | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionredirect](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionredirect)
 |
 |
 |
| **Custom Response –** a pre-configured response for a request | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionreturn](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#actionreturn)
 |
 |
 |
| **TLS Passthrough** – TCP Passthrough / proxying of encrypted TCP traffic | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/](https://docs.nginx.com/nginx-ingress-controller/configuration/transportserver-resource/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/tls-passthrough](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/tls-passthrough)
 |
 |
| **TLS Re-encryption –** re-encrypt HTTP and TCP traffic to the backend or external service without requiring a Service Mesh. | **Ingress, CRD** |
 |
 |
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)
 |
| **Session persistence** – sticky cookie is built in | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamsessioncookie](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#upstreamsessioncookie)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/session-persistence](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/custom-resources/session-persistence)
 |
 |
| **Full traffic programmability support** through [NGINX Java Script](https://nginx.org/en/docs/njs/) | **CRD, ConfigMap** |
 |
 |
 |
 |
| **Mergeable Ingress** - spread the Ingress configuration for a common host across multiple Ingress resources | **ingress** |
 |
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/mergeable-ingress-types](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/ingress-resources/mergeable-ingress-types)
 |
 |
| **Regex hostname –** provide hostname matches based on regex patterns | **snippet** |
 |
 |
 |
 |
| **FASTCGI -** | **snippet** |
 |
 |
 |
 |
| **Custom Error Page Response –** define your own error page and response from redirect to return | **CRD** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#errorpage](https://docs.nginx.com/nginx-ingress-controller/configuration/virtualserver-and-virtualserverroute-resources/#errorpage)
 |
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)
 |
| **Built-in State sharing** – most useful for authentication use cases, the deployment can be configured to automatically synchronize state objects such as authentication tokens |
 | **\*** |
 |
 |
 |
| **API Key authorization** – support for using API keys to perform service authorization | **snippet** |
 |
 |
 |
 |
| **Tight integration with NGINX Service Mesh –** Get a better and seamless experience with the NGINX service Mesh, giving you better controls from the edge to the backend. | **N+** |
 |
 |
 |
 |
| **Support for 3**** rd **** party Service Meshes **– Service Meshes like Istio and Linkerd expect the Mesh to perform the pod load balancing as opposed to the Ingress/Gateway layer. Through a special setting backend loadbalancing can be disabled to support these Meshes. |** Istio, Linkerd** |
 | [https://docs.nginx.com/nginx-ingress-controller/tutorials/nginx-ingress-istio/](https://docs.nginx.com/nginx-ingress-controller/tutorials/nginx-ingress-istio/)
 |
 |
 |
| **Tracing integration** – the OpenTracing module is currently built-in and will soon be moving to OpenTelemetry |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/third-party-modules/opentracing/](https://docs.nginx.com/nginx-ingress-controller/third-party-modules/opentracing/)
 |
 |
 |
| **Request logging** – full NGINX access logging that can be integrated into your management systems for analysis and malicious pattern detection |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/logging/#nginx-logs](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/logging/#nginx-logs)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-log-format](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-log-format)
 |
 |
| **Log customization –** customize the logging output |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/logging/](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/logging/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-log-format](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-log-format)
 |
 |
| **Additional NGINX module support** – built-in support for any NGINX module |
 |
 |
 |
 |
 |
| **Built-in Prometheus exporter** – with extensive metrics from N+, basic metrics from N OSS, and some additional calculated measures specific to ingress |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/prometheus/](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/prometheus/)
 |
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab8/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab8/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab8/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab8/readme.md)
 |
| **Proxy Protocol –** receiving and forwarding the real client IP and headers |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#listeners](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#listeners)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/proxy-protocol](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/proxy-protocol)
 |
 |
| **NGINX Snippets** – to-the-metal injection of NGINX configuration, unlocking the full capabilities of NGINX at the cluster edge. |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/configmap-resource/#snippets-and-custom-templates)
 |
 |
 |
| **Custom Templates –** Flexibility to modify the configuration generating templates to control specific behavior | **CRD, Ingress** |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/custom-templates/](https://docs.nginx.com/nginx-ingress-controller/configuration/global-configuration/custom-templates/)
 | [https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-templates](https://github.com/nginxinc/kubernetes-ingress/tree/main/examples/shared-examples/custom-templates)
 |
 |
| **Native NGINX Data plane** – wide tool and integration support, tried and trusted for years |
 |
 |
 |
 |
 |
| **Tight API security –** only monitoring those namespaces that are required, not global |
 |
 | [https://docs.nginx.com/nginx-ingress-controller/configuration/security/](https://docs.nginx.com/nginx-ingress-controller/configuration/security/)
 |
 |
 |
| **Rich and deep metrics and backend insights –** from detailed request and response times to response code awareness, over 100 detailed metrics |
 | **\*** | [https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/status-page/](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/status-page/)
 |
 |
 |
| **NTLM authentication to backends** | **CRD** |
 |
 |
 |
 |
| **readOnlyRootFilesystem support** |
 |
 |
 |
 |
 |
| **FIPS Inside** |
 |
 |
 |
 |
 |
| **Cache Support** – optimize your K8s applications using dynamic caching capabilities of NGINX |
 |
 |
 |
 | [https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/OSS/labs/lab10/readme.md)
[https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md](https://github.com/nginxinc/nginx-ingress-workshops/blob/main/Plus/labs/lab10/readme.md)
 |
---
