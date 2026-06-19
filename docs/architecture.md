# Architecture

`enoch-luminaries` turns the public surface of a running Enoch federation into a
live map — **the Orrery** — without modifying the protocol. This document
describes the data flow, the collector, the source adapters, and the layout.

## Principles

1. **Zero protocol code.** The collector reads only what `/enoch` already
   exposes. `/enoch` is never modified to feed the visualization. (The rationale
   and the privacy/trust analysis live in
   [observability-tiers.md](observability-tiers.md).)
2. **Manifest-driven.** The federation membership, endpoints and (Phase 2) geo
   placement come from the same `federation_manifest.json` the operators use. To
   point the Orrery at a different federation — local regtest vs. a 5-region AWS
   testnet — you change the manifest, not the code.
3. **Protocol-faithful vocabulary.** Domain types mirror `/enoch` names exactly
   (`operator`, `agent`, `state-root signature`, `anchor`, `FROST round`) so the
   map is legible against the spec and logs. See [naming.md](naming.md).
4. **Source abstraction.** Layers 1+2 (public, mainnet-safe) and Layer 3
   (reconstructed flows, testnet demo) are separate sources behind one
   interface. The frontend is identical; only the active source differs.

## Data flow

```
                         ┌─────────────────────────── /enoch (UNCHANGED) ──────────────────────────┐
                         │                                                                          │
  Bitcoin L1 ───anchors──┤  operator-0  /info  /state_roots/latest  /events (SSE)   ...  operator-N │
                         │                                                                          │
                         └──────────────┬──────────────────────────────┬────────────────────────────┘
                                        │ poll + subscribe (read-only)  │
                            ┌───────────▼───────────┐      ┌────────────▼───────────┐
                            │ source/public         │      │ source/flows           │
                            │  Layer 1 — liveness    │      │  Layer 3 — reconstructed│
                            │  Layer 2 — aggregates  │      │  message choreography  │
                            └───────────┬───────────┘      └────────────┬───────────┘
                                        │  luminaries.Event (one-way)   │
                                        └───────────────┬───────────────┘
                                                ┌───────▼────────┐
                                                │ internal/stream │  merge → timeline
                                                │  + websocket    │
                                                └───────┬────────┘
                                                ┌───────▼────────┐
                                                │  web/ — Orrery  │  geographic map
                                                └────────────────┘
```

## The collector (`cmd/collector`)

A single Go binary. Responsibilities:

- Load `federation_manifest.json` → build the topology (operators, agents,
  endpoints, geo).
- Start the configured **sources** (always `public`; `flows` only when
  `--flows` is set and the network is a test network).
- Merge source output into one time-ordered **timeline** and serve it to the
  Orrery over a websocket, plus a snapshot endpoint for late joiners.

The collector is **out of the federation's trust boundary**: operators never
depend on it, it never writes to them, and its compromise leaks nothing beyond
the already-public data it reads.

## Sources

### `source/public` — Layer 1 + 2 (mainnet-safe, zero protocol code)

Reads existing public surface and emits protocol-faithful `luminaries.Event`s:

| Signal | Source | Emits |
|--------|--------|-------|
| operator up/down, height | `GET /info`, `GET /health` per operator | `Liveness`, height |
| current leader | derived: `(height + view) mod N` from the manifest | `Operator.IsLeader` |
| state-root seals | `GET /state_roots/latest`, `/events` `state_root_signed` | `StateRoot`, Layer-1 `Message`s |
| anchor heartbeat / divergence | Bitcoin L1 + `/events` `anchor_*` | `Anchor`, equivocation alert |
| aggregate throughput | `/events` `tx_applied` / `deposit_*` / `withdrawal_status`, binned | counters (k-anonymized) |
| bonds | manifest / L1 bond outputs | `Operator.BondSats` |

No per-user data leaves this layer except as binned aggregates.

### `source/flows` — Layer 3 (testnet demo bolt-on)

Reconstructs the *implied* inter-node choreography from public lifecycle events
plus the known protocol state machine. When `/events` reports a `deposit_minted`
on operator-2 at height H, the protocol guarantees what happened underneath — a
FROST round among agents, a proposal broadcast, an M-of-N seal — so the source
synthesizes those `Message` arcs (`Tier: 3`) and animates them.

It is a **dramatization driven by real events**, not a packet tap: it requires
no new data surface and no `/enoch` code, and it is gated to test networks. See
[observability-tiers.md](observability-tiers.md) for why this is the chosen
mechanism over a literal RPC tap.

## The Orrery (`web/`)

A geographic world-map frontend. Operators and agents are placed by region;
`Message` arcs animate between them; a timeline/inspector panel drills into any
node or event. Health renders as healthy / degraded / offline. The data source
is pluggable, so the same UI renders a mainnet-safe Layer 1+2 view or the
testnet Layer 1+2+3 demo.

**Two view modes, one Orrery.** The map is a swappable render layer over shared
state — everything else (websocket client, store, top bar, inspector, timeline,
tier toggle) is common to both:

- **Flat map** (`deck.gl` ArcLayer + dark basemap) — the default daily driver:
  all regions legible at once, "mission-control" feel.
- **Globe** (`react-globe.gl`, Three.js) — the reveal: cinematic, planetary,
  makes geo-distribution visceral.

Both consume the identical collector feed; switching is a one-click toggle. Design
references live in [mockups/](mockups/) (`flat-map.html`, `globe.html`).

## Deployment (`deploy/`) — Phase 2

Terraform for one EC2 instance per AWS region, each running the existing
`docker-compose` operator+agent+bitcoind bundle, plus the collector. The work to
get there (real per-entity endpoints + cert SANs from `enoch-fedinit`, secret
distribution) is tracked in [roadmap.md](roadmap.md).
