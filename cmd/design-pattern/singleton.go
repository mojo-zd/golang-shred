package design_pattern

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var (
	singleton *Singleton
	once      sync.Once
)

// GetInstance 获取单例
func GetInstance() *Singleton {
	once.Do(func() {
		if singleton == nil {
			singleton = &Singleton{}
		}
	})
	return singleton
}
