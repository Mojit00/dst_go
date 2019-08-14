package readbook

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool //过滤函数
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(out time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     out,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscriberTopic(nil)
}

//添加一个订阅者，订阅过滤器筛选后的主题

func (p *Publisher) SubscriberTopic(topicFunc2 topicFunc) chan interface{} {
	ch := make(chan interface{})
	p.m.Lock()
	p.subscribers[ch] = topicFunc2
	p.m.Unlock()
	return ch
}

//推出订阅
func (p *Publisher) Exitpub(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

//发一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup

	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

}

func (p *Publisher) sendTopic(subscribers subscriber, t topicFunc, v interface{}, group *sync.WaitGroup) {
	defer group.Done()

	if t != nil && !t(v) {
		return
	}

	select {
	case subscribers <- v:
	case <-time.After(p.timeout):
	}
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers,sub)
		close(sub)
	}
}
