// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/fs.avdl

package khulnasoft1

type File struct {
	Path string `codec:"path" json:"path"`
}

func (o File) DeepCopy() File {
	return File{
		Path: o.Path,
	}
}

type ListResult struct {
	Files []File `codec:"files" json:"files"`
}

func (o ListResult) DeepCopy() ListResult {
	return ListResult{
		Files: (func(x []File) []File {
			if x == nil {
				return nil
			}
			ret := make([]File, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Files),
	}
}
