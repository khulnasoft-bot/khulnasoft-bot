// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/session.avdl

package khulnasoft1

type Session struct {
	Uid             UID    `codec:"uid" json:"uid"`
	Username        string `codec:"username" json:"username"`
	Token           string `codec:"token" json:"token"`
	DeviceSubkeyKid KID    `codec:"deviceSubkeyKid" json:"deviceSubkeyKid"`
	DeviceSibkeyKid KID    `codec:"deviceSibkeyKid" json:"deviceSibkeyKid"`
}

func (o Session) DeepCopy() Session {
	return Session{
		Uid:             o.Uid.DeepCopy(),
		Username:        o.Username,
		Token:           o.Token,
		DeviceSubkeyKid: o.DeviceSubkeyKid.DeepCopy(),
		DeviceSibkeyKid: o.DeviceSibkeyKid.DeepCopy(),
	}
}
