# OCI Artifacts

1. Start up Zot as a OCI compliant Artifact Registry (e.g. on a Hetzner Server)

```bash
docker run -d -p 5000:5000 --name oras-quickstart ghcr.io/project-zot/zot-linux-amd64:latest
```

2. Push OCI Artifacts

```bash
go run cmd/push/main.go
```

3. Push an Image Index

```bash
go run cmd/index/main.go
```

4. Test with ORAS (these commands are for the ORAS host, hence the localhost)

```bash
mkdir index && cd index
oras manifest fetch localhost:5000/myindex:latest
oras pull --plain-http localhost:5000/myindex:latest
```

Note: Note: to setup an Ubuntu Machine with Docker and Oras, use the following commands:
```bash
apt update -y && apt install curl vim -y

# Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Oras
VERSION="1.1.0"
curl -LO "https://github.com/oras-project/oras/releases/download/v${VERSION}/oras_${VERSION}_linux_amd64.tar.gz"
mkdir -p oras-install/
tar -zxf oras_${VERSION}_*.tar.gz -C oras-install/
sudo mv oras-install/oras /usr/local/bin/
rm -rf oras_${VERSION}_*.tar.gz oras-install/
```