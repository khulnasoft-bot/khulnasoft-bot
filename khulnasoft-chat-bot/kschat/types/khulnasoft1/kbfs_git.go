// Auto-generated to Go types using avdl-compiler v1.4.10 (https://github.com/khulnasoft/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/khulnasoft1/kbfs_git.avdl

package khulnasoft1

type GcOptions struct {
	MaxLooseRefs         int  `codec:"maxLooseRefs" json:"maxLooseRefs"`
	PruneMinLooseObjects int  `codec:"pruneMinLooseObjects" json:"pruneMinLooseObjects"`
	PruneExpireTime      Time `codec:"pruneExpireTime" json:"pruneExpireTime"`
	MaxObjectPacks       int  `codec:"maxObjectPacks" json:"maxObjectPacks"`
}

func (o GcOptions) DeepCopy() GcOptions {
	return GcOptions{
		MaxLooseRefs:         o.MaxLooseRefs,
		PruneMinLooseObjects: o.PruneMinLooseObjects,
		PruneExpireTime:      o.PruneExpireTime.DeepCopy(),
		MaxObjectPacks:       o.MaxObjectPacks,
	}
}
