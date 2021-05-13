package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: map[string]Cloneable{},
	}
}

func (pm *PrototypeManager) Get(name string) Cloneable {
	return pm.prototypes[name]
}

func (pm *PrototypeManager) Set(name string, cloneable Cloneable) {
	pm.prototypes[name] = cloneable
}

type Type1 struct {
	Name string
}

func (t *Type1) Clone() Cloneable {
	t1 := *t
	return &t1
}
