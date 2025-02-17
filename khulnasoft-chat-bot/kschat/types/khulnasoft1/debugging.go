// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/debugging.avdl

package khulnasoft1

type FirstStepResult struct {
	ValPlusTwo int `codec:"valPlusTwo" json:"valPlusTwo"`
}

func (o FirstStepResult) DeepCopy() FirstStepResult {
	return FirstStepResult{
		ValPlusTwo: o.ValPlusTwo,
	}
}
