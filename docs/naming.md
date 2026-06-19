# Naming — poetic identity, protocol-faithful internals

## Decision

> The poetry is skin-deep. It lives in the **brand** (`enoch-luminaries`) and the
> **name of the view** ("the Orrery"). Everything with semantic content uses
> `/enoch`'s vocabulary verbatim.

| Layer | Name |
|-------|------|
| Repo / brand | **`enoch-luminaries`** |
| The live map view (UI) | **"the Orrery"** |
| Collector binary | **`collector`** (plain — `/enoch` has no "scribe") |
| Node | **operator** / **agent** |
| Membership / topology | **federation** (from the manifest) |
| L1 backdrop | **anchors / Bitcoin L1** |
| Message arcs | labeled by their RPC — **state-root sig, proposal, FROST round, deposit gossip, anchor** |
| State stream | **timeline / events** |
| Health | **healthy / degraded / offline** |

The only invented labels are pure-visualization concepts `/enoch` does not name
(health states, "message arc"), and even those stay literal.

## Why

- **One vocabulary, zero translation tax.** An operator-contributor or an
  investor doing diligence sees the *same words* on screen as in the spec, the
  logs, and the code, and maps what they are watching straight onto the protocol.
- **It fits the honesty culture.** A trust-minimized verification instrument
  should call things exactly what they are. Aliasing an "operator" to a
  "Luminary" in the data model would be a small dishonesty the diligence
  audience has to see through.
- **The evocative naming costs nothing where it lives.** The repo *is*
  "Luminaries"; the map *is* "the Orrery" — memorable and on-theme — but the
  instant you are looking at data, it is protocol-faithful.

## On the name "Luminaries"

*Enoch* (Hebrew *Ḥanokh*) means "dedicated / initiated." In the Book of Enoch he
is taken on a tour of the heavens and shown the hidden order of the cosmos,
returning as the **scribe and witness** who records it. 1 Enoch 72–82 is
specifically the **Book of the Heavenly Luminaries** — Enoch's chart of the
courses of the sun, moon and stars. A tool that charts the courses of the
federation's bodies inherits that name honestly.

*An orrery* is a working model that visualizes celestial bodies in motion — not
the place you observe *from* (an observatory) but the visualization *itself*.
Hence: the repo is named for its subject (the luminaries); the live view is
named for what it is (an orrery).

## Considered and not chosen

- `enoch-observatory` — clear, but names the building, not the artifact.
- `enoch-orrery` (as the **repo**) — beautiful and precise, but "orrery" is a
  word much of the public cannot define or pronounce on sight, a real cost for a
  public/investor-facing brand. Preserved instead as the name of the map view,
  where its precision is exactly apt.
- `enoch-specula` (Latin *watchtower/observatory*, root of *speculate*),
  `enoch-vigil`, `enoch-witness` — strong, but "witness" collides with Bitcoin
  segwit terminology and the others are less legible than "luminaries."
- Celestial aliases in the data model (`Luminary`, `Ray`, `Ephemeris`,
  `Firmament`, `Constellation`) — rejected for internals per the decision above;
  they would impose a second vocabulary on a tool whose value is legibility.
