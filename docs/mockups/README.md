# Orrery — design mockups

Two interactive mockups of the same federation, same dark UI chrome — **only the
map differs**, so you can compare the two hero visuals for "the Orrery". Open
either in a browser (they fetch their map assets from a CDN, so you need a
connection the first time).

| File | Option | Stack it previews |
|------|--------|-------------------|
| [globe.html](globe.html) | **A — 3D globe** | `react-globe.gl` (Three.js / WebGL) |
| [flat-map.html](flat-map.html) | **B — flat dark map** | `deck.gl` ArcLayer + dark basemap |

> These are **static design mocks**, not wired to real data. The federation
> (5 operators across AWS regions + 3 agents), the metrics, and the message arcs
> are illustrative of the Phase-2 geo-distributed testnet. The mockups use
> `globe.gl` / `d3-geo` directly (no framework, no build step) purely so each is
> a single openable file; the real Orrery uses the React stack in
> [../architecture.md](../architecture.md).

Both cycle through the same four "moments": a state-root seal (leader fan-out),
a FROST signing round, deposit gossip, and an L1 anchor — colored per message
kind, with node health and a leader badge shown live.
