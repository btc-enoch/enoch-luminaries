// Package public is the Layer 1+2 source: it reads only the
// federation's existing public surface and emits protocol-faithful
// model objects. It is mainnet-safe and makes zero /enoch changes.
//
// Layer 1 (liveness): GET /info + /health per operator; current leader
// derived as (height+view) mod N from the manifest; state-root seals
// from /state_roots/latest and the /events `state_root_signed` event;
// anchor heartbeat + divergence from Bitcoin L1 and the /events
// `anchor_*` events.
//
// Layer 2 (aggregates): /events `tx_applied` / `deposit_*` /
// `withdrawal_status`, binned into k-anonymized counts. No per-user
// data leaves this layer.
//
// See docs/observability-tiers.md.
package public

// TODO(phase1): implement the poller (/info,/health,/state_roots/latest),
// the SSE subscriber (/events), and the L1 anchor watcher; emit
// model.Node / model.StateRoot / model.Anchor / model.Aggregate updates.
