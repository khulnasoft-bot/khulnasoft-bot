// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/apiserver.avdl

package khulnasoft1

type APIRes struct {
	Status     string `codec:"status" json:"status"`
	Body       string `codec:"body" json:"body"`
	HttpStatus int    `codec:"httpStatus" json:"httpStatus"`
	AppStatus  string `codec:"appStatus" json:"appStatus"`
}

func (o APIRes) DeepCopy() APIRes {
	return APIRes{
		Status:     o.Status,
		Body:       o.Body,
		HttpStatus: o.HttpStatus,
		AppStatus:  o.AppStatus,
	}
}
