package chains

import (
	"github.com/snowfork/polkadot-ethereum/bridgerelayer/cmd/keybase"
	"github.com/snowfork/polkadot-ethereum/bridgerelayer/cmd/types"
)

// Chain ...
type Chain struct {
	cfg      *ChainConfig // The config of the chain
	core     *Core        // THe chains connection
	streamer *Streamer    // The streamer of this chain
	router   *Router      // The router of the chain
	stop     chan<- int
}

// Core contains important information for each chain, including credentials
type Core interface {
	Keypair() *keybase.Keypair
}

// Streamer streams transactions from a blockchain and passes them to the router
type Streamer interface {
	Start() error
}

// Router packages transaction data as packets and relays them to the bridge
type Router interface {
	BuildPacket(tx map[string]interface{}, block map[string]interface{}) (error, types.Packet)
	SendPacket(types.Packet) error
}
