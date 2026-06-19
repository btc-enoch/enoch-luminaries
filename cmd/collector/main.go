// Command collector watches an Enoch federation over its public
// surface and serves the Orrery (web/) a live timeline.
//
// It makes zero changes to /enoch and never writes to an operator: it
// polls /info + /state_roots/latest, subscribes to /events, watches
// Bitcoin L1 for anchors, and reads the federation manifest for
// topology. Layer 1+2 are always on (mainnet-safe); Layer 3 flow
// reconstruction is enabled with --flows and refused on mainnet.
//
// See docs/architecture.md and docs/observability-tiers.md.
package main

import (
	"flag"
	"log"
)

func main() {
	var (
		manifestPath = flag.String("manifest", "../enoch/.local/federation/federation_manifest.json",
			"path to federation_manifest.json (topology + endpoints)")
		listen  = flag.String("listen", ":8090", "address to serve the Orrery websocket/API on")
		network = flag.String("network", "regtest", "network of the target federation (regtest|signet|testnet|mainnet)")
		flows   = flag.Bool("flows", false, "enable Layer 3 flow reconstruction (testnet only; refused on mainnet)")
	)
	flag.Parse()

	if *flows && *network == "mainnet" {
		log.Fatalf("refusing --flows on mainnet: Layer 3 (individual flows) is testnet-only " +
			"(see docs/observability-tiers.md)")
	}

	log.SetPrefix("[collector] ")
	log.Printf("manifest=%s listen=%s network=%s flows=%v", *manifestPath, *listen, *network, *flows)

	// TODO(phase1): load manifest → topology; start sources/public
	// (and sources/flows when *flows); merge via internal/stream; serve
	// the websocket + snapshot on *listen for web/ (the Orrery).
	log.Printf("scaffold only — see docs/roadmap.md Phase 1")
}
