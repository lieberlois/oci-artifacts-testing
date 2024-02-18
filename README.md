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