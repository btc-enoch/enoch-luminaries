// Package manifest reads the same federation_manifest.json the operators
// use and turns it into the Orrery's topology: the set of operators and
// agents, their endpoints, and (Phase 2) their geo placement. Point the
// Orrery at a different federation by changing the manifest, not the code.
package manifest

// Load parses federation_manifest.json at path into the federation
// topology as []luminaries.Node.
//
// TODO(phase1): parse operators[]/agents[] (id/label/grpc_endpoint);
// geo placement (region/lat/lon) comes from a side map keyed by label
// until the manifest carries it in Phase 2.
//
//	func Load(path string) ([]luminaries.Node, error)
