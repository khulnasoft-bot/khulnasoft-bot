// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/merkle_store.avdl

package khulnasoft1

type MerkleStoreSupportedVersion int

func (o MerkleStoreSupportedVersion) DeepCopy() MerkleStoreSupportedVersion {
	return o
}

type MerkleStoreKitHash string

func (o MerkleStoreKitHash) DeepCopy() MerkleStoreKitHash {
	return o
}

type MerkleStoreKit string

func (o MerkleStoreKit) DeepCopy() MerkleStoreKit {
	return o
}

type MerkleStoreEntryString string

func (o MerkleStoreEntryString) DeepCopy() MerkleStoreEntryString {
	return o
}

type MerkleStoreEntry struct {
	Hash  MerkleStoreKitHash     `codec:"hash" json:"hash"`
	Entry MerkleStoreEntryString `codec:"entry" json:"entry"`
}

func (o MerkleStoreEntry) DeepCopy() MerkleStoreEntry {
	return MerkleStoreEntry{
		Hash:  o.Hash.DeepCopy(),
		Entry: o.Entry.DeepCopy(),
	}
}
