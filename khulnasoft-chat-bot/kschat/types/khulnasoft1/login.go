// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/login.avdl

package khulnasoft1

type ConfiguredAccount struct {
	Username        string   `codec:"username" json:"username"`
	Fullname        FullName `codec:"fullname" json:"fullname"`
	HasStoredSecret bool     `codec:"hasStoredSecret" json:"hasStoredSecret"`
	IsCurrent       bool     `codec:"isCurrent" json:"isCurrent"`
}

func (o ConfiguredAccount) DeepCopy() ConfiguredAccount {
	return ConfiguredAccount{
		Username:        o.Username,
		Fullname:        o.Fullname.DeepCopy(),
		HasStoredSecret: o.HasStoredSecret,
		IsCurrent:       o.IsCurrent,
	}
}
