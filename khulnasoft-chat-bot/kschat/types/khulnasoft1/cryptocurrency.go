// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/cryptocurrency.avdl

package khulnasoft1

type RegisterAddressRes struct {
	Type   string `codec:"type" json:"type"`
	Family string `codec:"family" json:"family"`
}

func (o RegisterAddressRes) DeepCopy() RegisterAddressRes {
	return RegisterAddressRes{
		Type:   o.Type,
		Family: o.Family,
	}
}
