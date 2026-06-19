# Observability tiers — security, confidentiality & trust

This is the load-bearing design document for `enoch-luminaries`. It defines
**what may be visualized, on which networks, and why** — so the boundary between
a legitimate trust-minimized instrument and a surveillance oracle is a written
commitment, not a chat decision.

Enoch is **"federated, not trustless — a checking account, not a vault,"** with
an explicit roadmap toward an *untraceable* payments and messaging layer. A
visualization of the federation must not quietly erode that. The tiering below
is how we get a beautiful live map **without** security or confidentiality
compromise.

## The map is three layers, not one

The privacy cost of "showing the federation" depends entirely on *what* you
show. Decomposed:

### Layer 1 — Liveness / federation health
Operator up/down, current leader, current height, anchor heartbeat, bond status,
view-changes, equivocation/slashing alerts.

- **This is infrastructure status, not user metadata.** Most of it is *already
  public or protocol-public*: state-root anchors land on Bitcoin L1; the leader
  schedule is deterministic (`(h+v) mod N` from the manifest); bonds are
  on-chain.
- **Privacy cost: none.** Publishable on mainnet.
- **On-thesis.** For a "federated, not trustless" system, letting users *verify*
  that M-of-N is live, bonded, anchoring on cadence, and not equivocating is a
  **trust-minimization feature** — the visual form of the honesty the project
  already commits to in prose.

### Layer 2 — Aggregate activity
Binned throughput, total volume, deposit/withdrawal/tx counts over coarse time
buckets.

- With k-anonymity thresholds this reveals nothing about any individual — it is
  what every L1 explorer and Lightning dashboard already shows.
- **Privacy cost: none.** Publishable on mainnet.

### Layer 3 — Individual flows
Per-event, real-time message choreography: *this* deposit → *that* operator,
*this* FROST session, sender→recipient arcs.

- **This is the firehose.** Per-event and correlatable, it is precisely the
  metadata the privacy roadmap exists to protect.
- **Privacy cost: real on a live network.** Enables leader prediction for
  targeted DoS, deposit/withdrawal timing correlation, and deanonymizing which
  operator handled what.
- **Messaging flows (who messages whom) must NEVER appear, at any
  aggregation** — the mixnet threat model treats traffic *shape* as
  adversary-visible-but-unlinkable. Aggregate message *volume* may be shown;
  flows may not.

## The rule

> **Layers 1+2 (infrastructure + aggregates) are on-thesis and mainnet-safe.
> Layer 3 (individual flows) is a testnet-only demo instrument. Messaging flows
> are never visualized.**

The beautiful "pulsing arcs between geo-distributed operators" is achievable on
mainnet using Layer 1+2 data — each pulse is "a state root sealed" / "an anchor
posted" / "N payments settled this minute," **not** "Alice paid Bob."

## How Layer 3 is built — and why it stays clean

Layer 3 is **reconstructed**, not tapped. The collector synthesizes the implied
choreography from real public `/events` plus the known protocol state machine
(see [architecture.md](architecture.md) → `sources/flows`). Consequences:

- **Zero `/enoch` code.** No interceptor, no new endpoint, no build flag in the
  protocol repo. The protocol repo stays pristine — which is itself part of the
  trust story.
- **No new data surface.** It reads the same public events Layer 1+2 reads; it
  does not observe raw RPCs, sig shares, nonces, amounts, or tx bodies.
- **The committed-nonce slashing primitive is untouched** — nonces never transit
  any telemetry, because there is no telemetry.
- **Testnet-only by gate** (`--flows`, refused unless the network is a test
  network) — and because it is a *dramatization of public events*, even when
  enabled it discloses nothing a `/events` subscriber couldn't already see.

### Why not a literal RPC tap?
A literal tap — a build-tagged gRPC interceptor inside `/enoch`, or a
TLS-terminating mesh sidecar at deploy time — was considered and **parked**. It
would add fidelity (exact per-RPC timing, verified caller identity, FROST
session ids) at the cost of either protocol-repo code or operationally invasive
cert handling. Reconstruction gives a *more legible* demo (real RPC timing is
microsecond-bursty and illegible on screen) with none of that cost. The tap
remains a documented fallback to revisit only if reconstruction proves visually
unconvincing **and** true fidelity is required — and even then, build-tag-gated
so a mainnet binary contains zero tap code.

## Consensus safety

Independent of layer: the collector is **read-only and out-of-band**. It polls
and subscribes; it never writes to an operator, never sits on the consensus
path, and its failure or compromise cannot affect funds, safety, or liveness. It
is explicitly **outside the federation trust boundary**.

## Mainnet posture (summary)

| | Mainnet | Test networks |
|--|---------|---------------|
| Layer 1 — liveness | ✅ ships (on-thesis) | ✅ |
| Layer 2 — aggregates | ✅ ships (k-anonymized) | ✅ |
| Layer 3 — individual flows | ❌ never | ✅ demo (`--flows`) |
| Messaging flows | ❌ never | ❌ never |
| `/enoch` protocol changes | none | none |

## Where this is honestly bounded

A live, federation-wide map of a **real** network with independent operators and
real funds **cannot** be made fully leak-free at Layer 3 — the communication
graph *is* the sensitive thing, and rendering it live is the point. We do not
build that. The only genuinely compromise-free per-flow view on a real network
is a per-operator *self*-view (an operator visualizing only its own traffic,
which it already knows) — which is not the federation-wide map this tool
provides. The Layer 1+2 / Layer 3 split, and the testnet gate, are what keep us
from drifting across that line by accident.
