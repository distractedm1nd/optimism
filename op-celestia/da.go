package celestia

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/rollkit/celestia-openrpc/types/blob"
)

// DerivationVersionCelestia is a byte marker for celestia references submitted
// to the batch inbox address as calldata.
// Mnemonic 0xce = celestia
// version 0xce references are encoded as:
// [8]byte block height ++ [32]byte commitment
// in little-endian encoding.
// see: https://github.com/rollkit/celestia-da/blob/1f2df375fd2fcc59e425a50f7eb950daa5382ef0/celestia.go#L141-L160
const DerivationVersionCelestia = 0xce

// heightLen is a length (in bytes) of serialized height.
//
// This is 8 as uint64 consist of 8 bytes.
const heightLen = 8

// ID is a unique identifier of a Blob posted on Celestia, containing the height
// and blob commitment.
type ID []byte

// NewID constructs a new ID from the provided height and blob.
func NewID(height uint64, blob *blob.Blob) ID {
	id := make([]byte, heightLen+len(blob.Commitment)+1)
	id[0] = DerivationVersionCelestia
	binary.LittleEndian.PutUint64(id, height)
	copy(id[heightLen+1:], blob.Commitment)
	return id
}

// Split splits the ID into height and commitment.
func (id ID) Split() (height uint64, commitment blob.Commitment) {
	if len(id) <= heightLen {
		return 0, nil
	}
	commitment = blob.Commitment(id[1+heightLen:])
	return binary.LittleEndian.Uint64(id[1:heightLen]), commitment
}

func (id ID) String() string {
	return hex.EncodeToString(id)
}
