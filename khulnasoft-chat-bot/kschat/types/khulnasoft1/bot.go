// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/bot.avdl

package khulnasoft1

type BotToken string

func (o BotToken) DeepCopy() BotToken {
	return o
}

type BotTokenInfo struct {
	Token BotToken `codec:"token" json:"bot_token"`
	Ctime Time     `codec:"ctime" json:"ctime"`
}

func (o BotTokenInfo) DeepCopy() BotTokenInfo {
	return BotTokenInfo{
		Token: o.Token.DeepCopy(),
		Ctime: o.Ctime.DeepCopy(),
	}
}
