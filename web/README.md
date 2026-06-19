# web/ — the Orrery

The live map view: a geographic world-map frontend that renders the federation
as it runs. Operators and agents are placed by region; message arcs animate
between them; an inspector panel drills into any node or event. Health renders as
healthy / degraded / offline.

The data source is **pluggable** — the same UI renders a mainnet-safe Layer 1+2
view or the testnet Layer 1+2+3 demo, depending only on what the collector
feeds it over the websocket.

> **Status:** placeholder. Scaffolded in Phase 1 (see [../docs/roadmap.md](../docs/roadmap.md)).

Planned stack: React + a world-map/great-arc layer + a graph overlay. Built and
served via the containerized toolchain (no host Node) — see the repo `Makefile`.

Vocabulary on screen is protocol-faithful (operator, agent, state-root sig,
anchor, FROST round) per [../docs/naming.md](../docs/naming.md). Only the view
itself is branded — "the Orrery."
