// Package stream merges the active sources into one time-ordered
// timeline and serves it to the Orrery over a websocket, with a
// snapshot endpoint so late-joining clients get current state before
// the live feed.
//
// The stream is strictly outbound: it never writes back to any
// operator. The collector is outside the federation trust boundary.
package stream

// TODO(phase1): fan-in source channels → ordered timeline; websocket
// hub with per-client buffering (lossy, like the operator event bus);
// GET /snapshot for current node/topology state.
