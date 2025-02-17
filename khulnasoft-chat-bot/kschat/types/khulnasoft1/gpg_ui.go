// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/gpg_ui.avdl

package khulnasoft1

type SelectKeyRes struct {
	KeyID        string `codec:"keyID" json:"keyID"`
	DoSecretPush bool   `codec:"doSecretPush" json:"doSecretPush"`
}

func (o SelectKeyRes) DeepCopy() SelectKeyRes {
	return SelectKeyRes{
		KeyID:        o.KeyID,
		DoSecretPush: o.DoSecretPush,
	}
}
