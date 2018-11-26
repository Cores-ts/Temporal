package net

import (
	"context"
	"io"

	ic "gx/ipfs/QmNiJiXwWE3kRhZrC5ej3kSjWHm337pYfhjLGSCDNKJP2s/go-libp2p-crypto"
	pstore "gx/ipfs/QmQAGG1zxfePqj2t7bLxyN8AFccZ889DDR9Gn8kVLDrGZo/go-libp2p-peerstore"
	ma "gx/ipfs/QmRKLtwMw131aK7ugC3G7ybpumMz78YrJe5dzneyindvG1/go-multiaddr"
	"gx/ipfs/QmSF8fPo3jgVBAy8fpdjjYqgG87dkJgUprRBHRd2tmfgpP/goprocess"
	smux "gx/ipfs/QmY9JXR3FupnYAYJWK9aMr9bCpqWKcToQ1tz8DVGTrHpHw/go-stream-muxer"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	peer "gx/ipfs/QmcqU6QUDSXprb1518vYDGczrTJTyGwLG9eUa5iNX4xUtS/go-libp2p-peer"
)

// MessageSizeMax is a soft (recommended) maximum for network messages.
// One can write more, as the interface is a stream. But it is useful
// to bunch it up into multiple read/writes when the whole message is
// a single, large serialized object.
const MessageSizeMax = 1 << 22 // 4 MB

// Stream represents a bidirectional channel between two agents in
// the IPFS network. "agent" is as granular as desired, potentially
// being a "request -> reply" pair, or whole protocols.
// Streams are backed by stream-muxer streams underneath the hood.
type Stream interface {
	smux.Stream

	Protocol() protocol.ID
	SetProtocol(protocol.ID)

	// Stat returns metadata pertaining to this stream.
	Stat() Stat

	// Conn returns the connection this stream is part of.
	Conn() Conn
}

// Direction represents which peer in a stream initiated a connection.
type Direction int

const (
	// DirUnknown is the default direction.
	DirUnknown Direction = iota
	// DirInbound is for when the remote peer initiated a connection.
	DirInbound
	// DirOutbound is for when the local peer initiated a connection.
	DirOutbound
)

// Stat stores metadata pertaining to a given Stream/Conn.
type Stat struct {
	Direction Direction
	Extra     map[interface{}]interface{}
}

// StreamHandler is the type of function used to listen for
// streams opened by the remote side.
type StreamHandler func(Stream)

// ConnSecurity is the interface that one can mix into a connection interface to
// give it the security methods.
type ConnSecurity interface {
	// LocalPeer returns our peer ID
	LocalPeer() peer.ID

	// LocalPrivateKey returns our private key
	LocalPrivateKey() ic.PrivKey

	// RemotePeer returns the peer ID of the remote peer.
	RemotePeer() peer.ID

	// RemotePublicKey returns the public key of the remote peer.
	RemotePublicKey() ic.PubKey
}

// ConnMultiaddrs is an interface mixin for connection types that provide multiaddr
// addresses for the endpoints.
type ConnMultiaddrs interface {
	// LocalMultiaddr returns the local Multiaddr associated
	// with this connection
	LocalMultiaddr() ma.Multiaddr

	// RemoteMultiaddr returns the remote Multiaddr associated
	// with this connection
	RemoteMultiaddr() ma.Multiaddr
}

// Conn is a connection to a remote peer. It multiplexes streams.
// Usually there is no need to use a Conn directly, but it may
// be useful to get information about the peer on the other side:
//  stream.Conn().RemotePeer()
type Conn interface {
	io.Closer

	ConnSecurity
	ConnMultiaddrs

	// NewStream constructs a new Stream over this conn.
	NewStream() (Stream, error)

	// GetStreams returns all open streams over this conn.
	GetStreams() []Stream

	// Stat stores metadata pertaining to this conn.
	Stat() Stat
}

// ConnHandler is the type of function used to listen for
// connections opened by the remote side.
type ConnHandler func(Conn)

// Network is the interface used to connect to the outside world.
// It dials and listens for connections. it uses a Swarm to pool
// connnections (see swarm pkg, and peerstream.Swarm). Connections
// are encrypted with a TLS-like protocol.
type Network interface {
	Dialer
	io.Closer

	// SetStreamHandler sets the handler for new streams opened by the
	// remote side. This operation is threadsafe.
	SetStreamHandler(StreamHandler)

	// SetConnHandler sets the handler for new connections opened by the
	// remote side. This operation is threadsafe.
	SetConnHandler(ConnHandler)

	// NewStream returns a new stream to given peer p.
	// If there is no connection to p, attempts to create one.
	NewStream(context.Context, peer.ID) (Stream, error)

	// Listen tells the network to start listening on given multiaddrs.
	Listen(...ma.Multiaddr) error

	// ListenAddresses returns a list of addresses at which this network listens.
	ListenAddresses() []ma.Multiaddr

	// InterfaceListenAddresses returns a list of addresses at which this network
	// listens. It expands "any interface" addresses (/ip4/0.0.0.0, /ip6/::) to
	// use the known local interfaces.
	InterfaceListenAddresses() ([]ma.Multiaddr, error)

	// Process returns the network's Process
	Process() goprocess.Process
}

// Dialer represents a service that can dial out to peers
// (this is usually just a Network, but other services may not need the whole
// stack, and thus it becomes easier to mock)
type Dialer interface {

	// Peerstore returns the internal peerstore
	// This is useful to tell the dialer about a new address for a peer.
	// Or use one of the public keys found out over the network.
	Peerstore() pstore.Peerstore

	// LocalPeer returns the local peer associated with this network
	LocalPeer() peer.ID

	// DialPeer establishes a connection to a given peer
	DialPeer(context.Context, peer.ID) (Conn, error)

	// ClosePeer closes the connection to a given peer
	ClosePeer(peer.ID) error

	// Connectedness returns a state signaling connection capabilities
	Connectedness(peer.ID) Connectedness

	// Peers returns the peers connected
	Peers() []peer.ID

	// Conns returns the connections in this Netowrk
	Conns() []Conn

	// ConnsToPeer returns the connections in this Netowrk for given peer.
	ConnsToPeer(p peer.ID) []Conn

	// Notify/StopNotify register and unregister a notifiee for signals
	Notify(Notifiee)
	StopNotify(Notifiee)
}

// Connectedness signals the capacity for a connection with a given node.
// It is used to signal to services and other peers whether a node is reachable.
type Connectedness int

const (
	// NotConnected means no connection to peer, and no extra information (default)
	NotConnected Connectedness = iota

	// Connected means has an open, live connection to peer
	Connected

	// CanConnect means recently connected to peer, terminated gracefully
	CanConnect

	// CannotConnect means recently attempted connecting but failed to connect.
	// (should signal "made effort, failed")
	CannotConnect
)

// Notifiee is an interface for an object wishing to receive
// notifications from a Network.
type Notifiee interface {
	Listen(Network, ma.Multiaddr)      // called when network starts listening on an addr
	ListenClose(Network, ma.Multiaddr) // called when network stops listening on an addr
	Connected(Network, Conn)           // called when a connection opened
	Disconnected(Network, Conn)        // called when a connection closed
	OpenedStream(Network, Stream)      // called when a stream opened
	ClosedStream(Network, Stream)      // called when a stream closed

	// TODO
	// PeerConnected(Network, peer.ID)    // called when a peer connected
	// PeerDisconnected(Network, peer.ID) // called when a peer disconnected
}
