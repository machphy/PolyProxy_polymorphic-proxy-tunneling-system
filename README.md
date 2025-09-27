
---

# Fast Reverse Proxy (FRP) – Documentation

##  Overview

**FRP (Fast Reverse Proxy)** is a **high-performance reverse proxy application**.
It allows users to expose **local services** (running on private networks or behind NAT/firewalls) to the **public internet** through a remote FRP server.

###  Core Idea

* Run **`frpc` (client)** on your **local machine**.
* Run **`frps` (server)** on a **remote server** with public access.
* `frpc` forwards your **local TCP/UDP services** → `frps` → makes them **publicly accessible**.

---

##  Features

### 🔹 Port Range Mapping (v0.56.0+)

FRP supports **port range mapping** using Go template’s `parseNumberRangePair`.
This allows you to create **multiple proxy configs at once** instead of writing them individually.

**Example:**

```toml
{{- range $_, $v := parseNumberRangePair "6000-6006,6007" "6000-6006,6007" }}
[[proxies]]
name = "tcp-{{ $v.First }}"
type = "tcp"
localPort = {{ $v.First }}
remotePort = {{ $v.Second }}
{{- end }}
```

 This configuration creates **8 proxies**:
`tcp-6000` → `tcp-6007`

---

### 🔹 Client Plugins

By default, `frpc` forwards only **TCP/UDP ports**.
Plugins extend `frpc` with **additional features**.

#### Built-in Plugins

* `unix_domain_socket`
* `http_proxy`
* `socks5`
* `static_file`
* `http2https`
* `https2http`
* `https2https`

**Example – HTTP Proxy Plugin:**

```toml
[[proxies]]
name = "http_proxy"
type = "tcp"
remotePort = 6000

[proxies.plugin]
type = "http_proxy"
httpUser = "abc"
httpPassword = "abc"
```

Here, `httpUser` and `httpPassword` are **authentication credentials**.

---

### 🔹 Server Manage Plugins

The FRP server (`frps`) also supports **server-side plugins** to extend functionality.
Additional plugins are available in the **FRP plugin repository**, built using FRP’s **extension mechanism**.

---

### 🔹 SSH Tunnel Gateway (v0.53.0+)

FRP supports listening on an **SSH port** on the `frps` side.
This allows **TCP proxying via ssh -R protocol**, without requiring `frpc`.

**Example Config:**

```toml
# frps.toml
sshTunnelGateway.bindPort = 2200
```

When started, `frps` generates a **private key** (`.autogen_ssh_key`) in the working directory.

**Usage Command:**

```bash
ssh -R :80:127.0.0.1:8080 v0@{frp_address} -p 2200 \
    tcp --proxy_name "test-tcp" --remote_port 9090
```

Equivalent to running:

```bash
frpc tcp --proxy_name "test-tcp" \
         --local_ip 127.0.0.1 \
         --local_port 8080 \
         --remote_port 9090
```

---

### 🔹 Virtual Network (VirtualNet) – Alpha (v0.62.0+)

FRP now supports **virtual networking** via a **TUN interface**, extending beyond port forwarding to **full IP-level routing (like a VPN)**.

**Enable in config:**

```toml
featureGates = { VirtualNet = true }
```

---

##  Feature Gates

Feature gates control **experimental features** in FRP.

| Feature    | Stage | Default | Description                            |
| ---------- | ----- | ------- | -------------------------------------- |
| VirtualNet | ALPHA | false   | Enables virtual networking (TUN-based) |

**Lifecycle:**

1. **ALPHA** → Disabled by default, may be unstable.
2. **BETA** → More stable, may be enabled by default.
3. **GA (Generally Available)** → Production-ready, enabled by default.

---

##  Related Projects

* **gofrp/plugin** → Additional FRP plugins (via extension mechanism).
* **gofrp/tiny-frpc** → Lightweight FRP client (~3.5 MB) using SSH protocol, ideal for **low-resource devices**.

---

##  Project Status

 Work in Progress – This documentation and project will evolve with more features, research, and security extensions.

 Suggestions or collaboration ideas are welcome → **reach out at:**
 [rajeevsharmamachphy@gmail.com](mailto:rajeevsharmamachphy@gmail.com)

---



