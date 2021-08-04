package outoftree

import "fmt"

var registry = make(Registry)

type CollectFactory interface {
	Initializer() error
}

type Registry map[string]CollectFactory

type Options func(registry Registry)

// Unregister remove an existing plugin from the registry. if no plugin with the provided name exists, it returns an error
func (r Registry) Unregister(name string) error {
	if _, ok := r[name]; ok {
		delete(r, name)
	}
	return fmt.Errorf("not found plugin %v", name)
}

// Register add a new plugin to the registry, if a plugin with the same name exists, it returns an error.
func (r Registry) Register(name string, factory CollectFactory) error {
	if _, ok := r[name]; ok {
		return fmt.Errorf("plugin[%s] has exist, can't repeat registry", name)
	}
	r[name] = factory
	return nil
}

// Get return collect factory with name
func (r Registry) Get(name string) (CollectFactory, bool) {
	factory, ok := r[name]
	return factory, ok
}

// Merge merge the give registry to the current
func (r Registry) Merge(in Registry) error {
	for name, factory := range in {
		if err := r.Register(name, factory); err != nil {
			return err
		}
	}
	return nil
}

// WithPlugin registry factory to registry, with this method we can implement a "out of tree" plugin
func WithPlugin(name string, factory CollectFactory) Options {
	return func(registry Registry) {
		registry.Register(name, factory)
	}
}
