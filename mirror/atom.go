package mirror

type Atom interface {
	Type() string
	String() string
	Tomap() map[string]interface{}
}
