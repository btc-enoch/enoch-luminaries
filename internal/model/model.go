// Package model holds the domain types the Orrery renders.
//
// Vocabulary policy (docs/naming.md): these types mirror /enoch's own
// names exactly — operator, agent, state-root signature, anchor, FROST
// round — so what shows on the map maps one-to-one onto the protocol
// spec, logs and code. The only invented names here are pure-viz
// concepts /enoch does not name (health states, the "message" arc).
package model

// Role distinguishes the two federation node types. Both are real
// /enoch entities; the Orrery places and styles them differently.
type Role string

const (
	RoleOperator Role = "operator"
	RoleAgent    Role = "agent"
)

// Liveness is a node's health as the collector currently sees it.
// Pure-viz concept (the protocol does not name these); kept literal.
type Liveness string

const (
	Healthy  Liveness = "healthy"  // responding, height advancing
	Degraded Liveness = "degraded" // reachable but lagging / partial
	Offline  Liveness = "offline"  // unreachable
)

// Node is one operator or agent in the federation, as placed on the map.
type Node struct {
	ID       int      `json:"id"`
	Role     Role     `json:"role"`
	Label    string   `json:"label"`            // manifest label, e.g. "operator-0"
	Region   string   `json:"region,omitempty"` // AWS region (Phase 2 geo)
	Lat      float64  `json:"lat,omitempty"`
	Lon      float64  `json:"lon,omitempty"`
	Height   uint64   `json:"height"`    // last known L2 height (operators)
	IsLeader bool     `json:"is_leader"` // derived: (height+view) mod N
	Liveness Liveness `json:"liveness"`
	BondSats uint64   `json:"bond_sats,omitempty"`
}

// MessageKind names a message arc by the /enoch RPC or event it
// represents — never a celestial alias.
type MessageKind string

const (
	MsgStateRootSig  MessageKind = "state_root_sig" // operator → operator partial sig
	MsgProposal      MessageKind = "proposal"       // leader → peers apply-on-seal proposal
	MsgFrostRound    MessageKind = "frost_round"    // agent ↔ operator FROST nonce/share
	MsgDepositGossip MessageKind = "deposit_gossip" // operator → peers deposit detected
	MsgViewChange    MessageKind = "view_change"    // operator → peers view-change endorsement
	MsgAnchor        MessageKind = "anchor"         // operator → Bitcoin L1
)

// Tier records which observability layer produced a Message
// (docs/observability-tiers.md). Tiers 1–2 are public/mainnet-safe;
// tier 3 is reconstructed flow choreography, testnet-only.
type Tier int

// Message is one animated arc between nodes (or a node and L1).
type Message struct {
	From     int         `json:"from"` // node id; -1 for L1/external
	To       int         `json:"to"`   // node id; -1 for L1/external
	Kind     MessageKind `json:"kind"`
	AtUnixNs int64       `json:"at_unix_ns"`
	Tier     Tier        `json:"tier"`
}

// StateRoot is a quorum-sealed state root — the federation's heartbeat.
// Mirrors /enoch's StateRootSignedData wire shape.
type StateRoot struct {
	Height   uint64 `json:"height"`
	Root     string `json:"root"`     // 32-byte hex
	NumSigs  int    `json:"num_sigs"` // M of N
	AtUnixNs int64  `json:"at_unix_ns"`
}

// Anchor is a state root committed to Bitcoin L1. Diverged flags an
// A6 equivocation candidate (/enoch anchor_divergence_detected).
type Anchor struct {
	L2Height uint64 `json:"l2_height"`
	TxID     string `json:"txid"`
	L1Height uint64 `json:"l1_height"`
	Diverged bool   `json:"diverged"`
	AtUnixNs int64  `json:"at_unix_ns"`
}

// Aggregate is a binned activity count for a time bucket (Layer 2).
// k-anonymized: never carries per-user identifiers.
type Aggregate struct {
	BucketUnixNs int64  `json:"bucket_unix_ns"`
	Kind         string `json:"kind"` // "tx" | "deposit" | "withdrawal"
	Count        int    `json:"count"`
}
