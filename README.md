---

NeuroTunnel — Advanced Tunneling & Virtual Networking (FRP Based)

Tagline: Exploring intelligent tunneling, plugin-driven proxies, and VPN-like virtual networks.


---

Overview

NeuroTunnel is a research-focused project built on Fast Reverse Proxy (FRP), a high-performance reverse proxy application. FRP allows exposing local services (behind NAT/firewalls) to the public internet through a remote server.

The main architecture:

frpc (client): Runs on your local machine and forwards local TCP/UDP services.

frps (server): Runs on a publicly accessible server and receives forwarded traffic.


This project explores advanced features of FRP and extends it with research-focused experiments for secure, scalable networking.


---

Key Features

1. Port Range Mapping (v0.56.0)

Automates creation of multiple proxy configurations using Go templates and parseNumberRangePair.
Example:

{{- range $_, $v := parseNumberRangePair "6000-6006,6007" "6000-6006,6007" }}
[[proxies]]
name = "tcp-{{ $v.First }}"
type = "tcp"
localPort = {{ $v.First }}
remotePort = {{ $v.Second }}
{{- end }}

Creates 8 proxies (tcp-6000 → tcp-6007), each mapping the remote port to the corresponding local port.


---

2. Client Plugins

Extend FRPC functionality beyond simple TCP/UDP forwarding.

Built-in Plugins:

unix_domain_socket

http_proxy

socks5

static_file

http2https

https2http

https2https


HTTP Proxy Example:

[[proxies]]
name = "http_proxy"
type = "tcp"
remotePort = 6000

[proxies.plugin]
type = "http_proxy"
httpUser = "abc"
httpPassword = "abc"

httpUser and httpPassword are authentication credentials.



---

3. Server Manage Plugins

FRPS supports server-side plugins for extended functionality.

Additional plugins are implemented using the FRP extension mechanism.

Plugins can customize proxy behavior and protocol handling.



---

4. SSH Tunnel Gateway (v0.53.0)

Enables TCP proxying via SSH -R protocol, without requiring frpc.

Configuration:

# frps.toml
sshTunnelGateway.bindPort = 2200

Example Command:

ssh -R :80:127.0.0.1:8080 v0@<frp_address> -p 2200 \
    tcp --proxy_name "test-tcp" --remote_port 9090

Forwards local port 8080 → remote 9090

Equivalent frpc command:


frpc tcp --proxy_name "test-tcp" \
         --local_ip 127.0.0.1 \
         --local_port 8080 \
         --remote_port 9090


---

5. Virtual Network (VirtualNet) — Alpha Feature (v0.62.0)

Creates and manages TUN-based virtual networks for full IP-level routing (VPN-like).

Enable VirtualNet:

featureGates = { VirtualNet = true }


---

6. Feature Gates

Allow enabling/disabling experimental FRP features.

Lifecycle:

1. ALPHA: Disabled by default, may be unstable


2. BETA: More stable, may be enabled by default


3. GA: Production-ready, enabled by default





---

Related Projects

gofrp/plugin: Repository for additional FRP plugins.

gofrp/tiny-frpc: Lightweight (~3.5MB) FRP client using SSH protocol.



---

Why NeuroTunnel?

This project pushes FRP beyond simple port forwarding into a next-generation tunneling and networking framework:

Automates proxy creation via port range mapping.

Extends functionality with plugins.

Enables SSH-based dynamic proxy creation.

Explores VPN-like virtual networks using TUN interfaces.


Suitable for developers, sysadmins, and cybersecurity researchers exploring adaptive, secure, and polymorphic networking solutions.


---

Contact / Collaboration

Project is continuously evolving. Feedback, ideas, or contributions are welcome.
Reach out: rajeevsharmamachphy@gmail.com


---

