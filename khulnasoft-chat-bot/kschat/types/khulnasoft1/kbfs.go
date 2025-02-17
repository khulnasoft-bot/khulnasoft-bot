// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/kbfs.avdl

package khulnasoft1

type KBFSTeamSettings struct {
	TlfID TLFID `codec:"tlfID" json:"tlfID"`
}

func (o KBFSTeamSettings) DeepCopy() KBFSTeamSettings {
	return KBFSTeamSettings{
		TlfID: o.TlfID.DeepCopy(),
	}
}
