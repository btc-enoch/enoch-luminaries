# Roadmap

## Phase 1 — Local live map (Layer 1+2)
The real product, against a local regtest federation. Zero `/enoch` changes.

- [ ] `internal/federation` — read `federation_manifest.json` → topology + geo
- [ ] `internal/model` — protocol-faithful domain types
- [ ] `sources/public` — Layer 1 (liveness/leader/anchors) + Layer 2 (binned aggregates)
- [ ] `internal/stream` — merge → timeline + websocket
- [ ] `web/` — the Orrery: geographic map, message arcs, inspector panel
- [ ] run end-to-end against `../enoch` `make federation-up`

## Phase 1b — Layer 3 demo bolt-on
Reconstructed flow choreography for the investor / operator-contributor reveal.
Testnet-gated (`--flows`). Zero `/enoch` changes.

- [ ] `sources/flows` — reconstruct message choreography from `/events` + the protocol state machine
- [ ] Orrery flow toggle + animation
- [ ] testnet gate enforced in the collector

## Phase 2 — AWS geo-distributed testnet
Lift the same federation across AWS regions so the map shows real cross-region
latency. EC2 per region running the existing `docker-compose` bundle.

- [ ] extend `enoch-fedinit` with real per-entity endpoints + cert SANs
      (today it emits docker hostnames like `operator-0:9090`)
- [ ] secret distribution — each host holds only its own keys (SSM / Secrets Manager)
- [ ] `deploy/` — Terraform: VPC/region, EC2, security groups (gRPC 9090/9190,
      bitcoin P2P 18444, operator HTTP 8080), persistent disk for bitcoind
- [ ] collector reads the multi-region manifest unchanged

## Phase 3 — Depth
- [ ] richer Layer 1 signals (view-change inference, slashing timeline)
- [ ] latency heatmaps across regions
- [ ] scenario replay (record a timeline, replay it offline for demos)
- [ ] mainnet-safe transparency build (Layer 1+2 only, public sources)

## Parked
- Literal RPC tap (build-tagged interceptor in `/enoch`, or deploy-time mesh
  sidecar). Only revisited if Layer 3 reconstruction proves visually
  unconvincing **and** true per-RPC fidelity is required. See
  [observability-tiers.md](observability-tiers.md).
