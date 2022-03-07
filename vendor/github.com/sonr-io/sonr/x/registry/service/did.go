package service

import (
	"context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/x/registry/types"
	v1 "github.com/sonr-io/sonr/x/registry/types"
)

// DID is the interface for managing the APIv1 Did's
type DID interface {
	// Create creates a new DID from an existing DID
	Create(options ...Option) (DID, error)

	// GetBase return the did:method:network:identifier string of the DID
	GetBase() string

	// HasNetwork returns true if the DID has the provided network
	HasNetwork() bool

	// HasPath returns true if the DID has the provided path
	HasPath() bool

	// HasQuery returns true if the DID has the provided query
	HasQuery() bool

	// HasFragment returns true if the DID has the provided fragment
	HasFragment() bool

	// IsValid returns true if the DID is valid
	IsValid() bool

	// PublicKey returns the public key of the DID
	PublicKey() crypto.PubKey

	// PublicKeyBase58 returns the base58 encoded public key of the DID
	PublicKeyBase58() (string, error)

	// PublicKeyBuffer returns the public key of the DID as a byte array
	PublicKeyBuffer() ([]byte, error)

	// Resolve resolves the DID and
	Resolve(ctx context.Context, resolverFunc ResolverFunc) (*v1.DidDocument, error)

	// ToProto converts the DID to a proto representation v1.Did
	ToProto() *v1.Did

	// ToRegistryProto converts the DID to a proto representation sonrio.sonr.registry.Did
	ToRegistryProto() *types.Did

	// ToString returns the string representation of the DID
	ToString() string
}

// Create builds a new DID from the provided id and options
func Create(id string, options ...v1.Option) (DID, error) {
	did, err := v1.CreateDid(id, options...)
	if err != nil {
		return nil, err
	}

	return &didImpl{
		Did: did,
	}, nil
}

// New creates a new DID with a provided public key and options
func New(pubKey crypto.PubKey, options ...Option) (DID, error) {
	// Marshal the public key into bytes
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	// Convert Options
	v1Opts := make([]v1.Option, len(options))
	for i, opt := range options {
		v1Opts[i] = v1.Option(opt)
	}

	pubStr := base58.Encode(pubBuf)
	d, err := v1.CreateDid(pubStr)
	if err != nil {
		return nil, err
	}

	return &didImpl{
		Did:    d,
		pubKey: pubKey,
	}, nil
}

// FromString parses a DID string into a DID struct
func FromString(s string) (DID, error) {
	di, err := v1.ParseDid(s)
	if err != nil {
		return nil, err
	}

	return &didImpl{
		Did: di,
	}, nil
}

// didImpl is the implementation of the DID interface
type didImpl struct {
	DID
	*v1.Did
	pubKey crypto.PubKey
}

// Generate creates a new DID from an existing DID
func (di *didImpl) Generate(options ...Option) (DID, error) {
	id := di.GetId()
	if id == "" {
		return nil, errors.New("Current DID does not have an Identifier")
	}

	// Create a new DID
	did, err := v1.CreateDid(id, options...)
	if err != nil {
		return nil, err
	}

	return &didImpl{
		Did: did,
	}, nil
}

// GetBase returns the did:method:network:identifier string of the DID
func (di *didImpl) GetBase() string {
	return di.Did.GetBase()
}

// HasNetwork returns true if the DID has the provided network
func (di *didImpl) HasNetwork() bool {
	return di.Did.HasNetwork()
}

// HasFragment returns true if the DID has the provided fragment
func (di *didImpl) HasFragment() bool {
	return di.Did.HasFragment()
}

// HasPath returns true if the DID has the provided path
func (di *didImpl) HasPath() bool {
	return di.Did.HasPath()
}

// HasQuery returns true if the DID has the provided query
func (di *didImpl) HasQuery() bool {
	return di.Did.HasQuery()
}

// IsValid returns true if the DID is valid
func (di *didImpl) IsValid() bool {
	return di.Did.IsValid()
}

// ToProto converts the DID to a proto representation v1.Did
func (di *didImpl) ToProto() *v1.Did {
	return di.Did
}

// ToString returns the string representation of the DID
func (di *didImpl) ToString() string {
	return di.Did.ToString()
}

// PublicKey returns the public key of the DID
func (di *didImpl) PublicKey() crypto.PubKey {
	return di.pubKey
}

// PublicKeyBuffer returns the public key of the DID as a byte array
func (di *didImpl) PublicKeyBuffer() ([]byte, error) {
	return crypto.MarshalPublicKey(di.pubKey)
}

// PublicKeyBase58 returns the base58 encoded public key of the DID
func (di *didImpl) PublicKeyBase58() (string, error) {
	pubBuf, err := di.PublicKeyBuffer()
	if err != nil {
		return "", errors.Wrap(err, "Could not create base58 encoded public key")
	}
	return base58.Encode(pubBuf), nil
}

// TODO: Resolve resolves the DID and returns the DID document
func (di *didImpl) Resolve(ctx context.Context, resFunc ResolverFunc) (*v1.DidDocument, error) {
	return nil, nil
}

// ToRegistryProto converts the DID to a proto representation sonrio.sonr.registry.Did
func (di *didImpl) ToRegistryProto() *types.Did {
	return &types.Did{
		Method:   di.GetMethod(),
		Network:  di.GetNetwork(),
		Id:       di.GetId(),
		Paths:    di.GetPaths(),
		Query:    di.GetQuery(),
		Fragment: di.GetFragment(),
	}
}
