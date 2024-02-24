package addrs

// TofuAttr is the address of an attribute of the "tofu" object in
// the interpolation scope, like "tofu.workspace".
type TofuAttr struct {
	referenceable
	Name string
}

func (ta TofuAttr) String() string {
	return "tofu." + ta.Name
}

func (ta TofuAttr) UniqueKey() UniqueKey {
	return ta // A TerraformAttr is its own UniqueKey
}

func (ta TofuAttr) uniqueKeySigil() {}
