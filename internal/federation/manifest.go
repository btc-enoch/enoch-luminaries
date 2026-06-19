// Package federation reads the same federation_manifest.json the
// operators use and turns it into the Orrery's topology: the set of
// operators and agents, their endpoints, and (Phase 2) their geo
// placement. Point the Orrery at a different federation by changing the
// manifest, not the code.
package federation

// TODO(phase1): parse federation_manifest.json (operators[], agents[]
// with id/label/grpc_endpoint) into []model.Node. Geo placement
// (region/lat/lon) is sourced from a side map keyed by label until the
// manifest carries it in Phase 2.
