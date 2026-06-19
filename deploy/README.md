# deploy/ — AWS geo-distributed testnet (Phase 2)

Terraform to lift the federation across AWS regions so the Orrery shows real
cross-region latency. **EC2 per region**, each running the existing
`../enoch` `docker-compose` operator+agent+bitcoind bundle, plus the collector.

> **Status:** placeholder. Phase 2 (see [../docs/roadmap.md](../docs/roadmap.md)).

Prerequisites tracked in the roadmap, the load-bearing one being:

- **`enoch-fedinit` real endpoints + cert SANs.** Today it emits docker
  hostnames (`operator-0:9090`); multi-region needs routable per-entity
  `host:port` and matching cert SANs. This is the single biggest change, and it
  is in `/enoch` tooling, not the protocol path.
- **Secret distribution** — each host holds only its own keys (SSM / Secrets
  Manager), never the whole `.local/` bundle.
- **Networking** — security groups for gRPC (9090/9190), bitcoin P2P (18444),
  operator HTTP (8080); persistent disk for bitcoind state.

The collector reads the resulting multi-region manifest unchanged.
