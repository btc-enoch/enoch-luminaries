// Package flows is the Layer 3 demo bolt-on: it reconstructs the
// implied inter-node message choreography from real public /events plus
// the known protocol state machine, and emits luminaries.Message arcs
// with Tier 3. It implements luminaries.Source.
//
// It is a dramatization driven by public events, NOT a packet tap: it
// observes no raw RPCs, sig shares, nonces, amounts or tx bodies, adds
// no /enoch code, and is gated to test networks (--flows, refused on
// mainnet). Messaging flows are never reconstructed.
//
// Example: a /events `deposit_minted` on operator-2 at height H implies
// a FROST round among agents, a proposal broadcast, and an M-of-N seal
// among operators at H — this source synthesizes those arcs.
//
// See docs/observability-tiers.md.
package flows

// TODO(phase1b): subscribe to the same /events stream as source/public
// and, per lifecycle event, synthesize the protocol-accurate sequence
// of luminaries.Message{Tier: 3} arcs the event implies.
