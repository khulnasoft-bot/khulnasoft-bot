// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/airdrop.avdl

package khulnasoft1

type AirdropDetails struct {
	Uid  UID       `codec:"uid" json:"uid"`
	Kid  BinaryKID `codec:"kid" json:"kid"`
	Vid  VID       `codec:"vid" json:"vid"`
	Vers string    `codec:"vers" json:"vers"`
	Time Time      `codec:"time" json:"time"`
}

func (o AirdropDetails) DeepCopy() AirdropDetails {
	return AirdropDetails{
		Uid:  o.Uid.DeepCopy(),
		Kid:  o.Kid.DeepCopy(),
		Vid:  o.Vid.DeepCopy(),
		Vers: o.Vers,
		Time: o.Time.DeepCopy(),
	}
}
