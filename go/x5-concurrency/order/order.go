package order

import "sync"

type Order struct {
	id     int
	status string
	mutex  sync.RWMutex
}

func New(id int, status string) Order {
	return Order{
		id:     id,
		status: status,
	}
}

func (o *Order) ID() int {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return o.id
}

func (o *Order) Status() string {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return o.status
}

func (o *Order) SetStatus(status string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.status = status
}
