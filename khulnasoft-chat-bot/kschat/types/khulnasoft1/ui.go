// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/ui.avdl

package khulnasoft1

import (
	"fmt"
)

type PromptDefault int

const (
	PromptDefault_NONE PromptDefault = 0
	PromptDefault_YES  PromptDefault = 1
	PromptDefault_NO   PromptDefault = 2
)

func (o PromptDefault) DeepCopy() PromptDefault { return o }

var PromptDefaultMap = map[string]PromptDefault{
	"NONE": 0,
	"YES":  1,
	"NO":   2,
}

var PromptDefaultRevMap = map[PromptDefault]string{
	0: "NONE",
	1: "YES",
	2: "NO",
}

func (e PromptDefault) String() string {
	if v, ok := PromptDefaultRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}
