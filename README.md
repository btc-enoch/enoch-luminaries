# enoch-luminaries

A live, trust-minimized visualization of the [Enoch](../enoch) federation — the
protocol working and behaving as it should, rendered as it happens.

The map view is **the Orrery**: a live model of the federation's bodies in
motion. Each operator and agent is a node; state-root signatures, anchors, FROST
rounds and deposit/withdrawal lifecycle events animate between them. Locally it
watches a regtest federation; in [Phase 2](docs/roadmap.md) it watches a
geo-distributed testnet across AWS regions, so cross-region latency becomes
visible on the map.

> **Status:** scaffold. Phase 1 (Layer 1+2 collector + Orrery) in progress.

## What it is — and what it deliberately is not

`enoch-luminaries` reads only what the federation **already exposes publicly**
(Bitcoin L1 anchors, the operator HTTP API, the `/events` stream, the
manifest). It makes **zero changes to the `/enoch` protocol repo.** That is a
hard design constraint, not an accident — see
[docs/observability-tiers.md](docs/observability-tiers.md).

The visualization is split into three layers by privacy sensitivity:

| Layer | Shows | Privacy cost | Where it runs |
|------:|-------|--------------|---------------|
| **1 — Liveness** | operator health, leader, height, anchor heartbeat, bonds, equivocation/slashing alerts | none (already public / on-chain) | mainnet-safe |
| **2 — Aggregate** | binned throughput, deposit/withdrawal/tx counts | none (k-anonymized) | mainnet-safe |
| **3 — Flows** | reconstructed per-event message choreography (FROST rounds, seals) | demo-only | **testnet only** |

Layers 1+2 are the real product: a verifiable, on-thesis view of a
"federated, not trustless" system proving its own honesty. Layer 3 is a
**bolt-on for the investor / operator-contributor reveal** — reconstructed from
public events, gated to test networks, never pointed at a real federation.

## Naming

Poetic on the outside, protocol-faithful on the inside
([docs/naming.md](docs/naming.md)):

- **`enoch-luminaries`** — the repo / brand. *Luminaries* is canonical: 1 Enoch
  72–82 is the *Book of the Heavenly Luminaries*, Enoch's chart of the courses
  of the sun, moon and stars.
- **the Orrery** — the live map view.
- Everything with semantic content uses `/enoch`'s own words — **operator**,
  **agent**, **state-root signature**, **anchor**, **FROST round** — so what you
  see on screen maps one-to-one onto the spec, the logs, and the code.

## Layout

```
cmd/collector/        the collector binary — watches the federation, serves the Orrery
internal/
  federation/         read federation_manifest.json → topology + geo placement
  sources/
    public/           Layer 1+2: /info, /state_roots/latest, /events, Bitcoin L1
    flows/            Layer 3 bolt-on: reconstruct message choreography from /events
  model/              protocol-faithful domain types (Operator, Agent, Message, ...)
  stream/             merge sources → timeline, serve websocket to the Orrery
web/                  the Orrery — geographic world-map frontend
deploy/               Terraform for the AWS geo-distributed testnet (Phase 2)
docs/                 architecture, observability tiers, naming, roadmap
```

## Quickstart

> Toolchain is containerized — no host Go/Node required.

```
make collector       # build + run the collector against a local regtest federation
make orrery          # serve the Orrery web UI
```

(Bring up the federation first from `../enoch`: `make federation-up`.)

See [docs/architecture.md](docs/architecture.md) for the full design.
