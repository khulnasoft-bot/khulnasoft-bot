// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/quota.avdl

package khulnasoft1

type VerifySessionRes struct {
	Uid       UID    `codec:"uid" json:"uid"`
	Sid       string `codec:"sid" json:"sid"`
	Generated int    `codec:"generated" json:"generated"`
	Lifetime  int    `codec:"lifetime" json:"lifetime"`
}

func (o VerifySessionRes) DeepCopy() VerifySessionRes {
	return VerifySessionRes{
		Uid:       o.Uid.DeepCopy(),
		Sid:       o.Sid,
		Generated: o.Generated,
		Lifetime:  o.Lifetime,
	}
}
