// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/metadata_update.avdl

package khulnasoft1

type RekeyRequest struct {
	FolderID string `codec:"folderID" json:"folderID"`
	Revision int64  `codec:"revision" json:"revision"`
}

func (o RekeyRequest) DeepCopy() RekeyRequest {
	return RekeyRequest{
		FolderID: o.FolderID,
		Revision: o.Revision,
	}
}
