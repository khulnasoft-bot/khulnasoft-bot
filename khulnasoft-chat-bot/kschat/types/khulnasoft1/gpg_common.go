// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/gpg_common.avdl

package khulnasoft1

type GPGKey struct {
	Algorithm  string        `codec:"algorithm" json:"algorithm"`
	KeyID      string        `codec:"keyID" json:"keyID"`
	Creation   string        `codec:"creation" json:"creation"`
	Expiration string        `codec:"expiration" json:"expiration"`
	Identities []PGPIdentity `codec:"identities" json:"identities"`
}

func (o GPGKey) DeepCopy() GPGKey {
	return GPGKey{
		Algorithm:  o.Algorithm,
		KeyID:      o.KeyID,
		Creation:   o.Creation,
		Expiration: o.Expiration,
		Identities: (func(x []PGPIdentity) []PGPIdentity {
			if x == nil {
				return nil
			}
			ret := make([]PGPIdentity, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Identities),
	}
}
