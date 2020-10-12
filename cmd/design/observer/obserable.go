package observer

import "github.com/rs/zerolog/log"

// 主题(被观察者)
type Observable interface {
	Attach(observer ...Observer) Observable
	Detach(observer ...Observer) Observable
	Notify() error
}

// Observer订阅(观察者)
type Observer interface {
	Do(observable Observable) error
}

// OderObserver订阅者实现
type OderObserver struct {
}

// Do 订阅者具体业务实现
func (oder *OderObserver) Do(observable Observable) error {
	log.Info().Msg("order status")
	return nil
}
