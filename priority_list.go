package priority_list

import (
	"errors"
	"sync"
)

type priorityItem[T any] struct {
	data     T
	priority int
}

type PriorityList[T any] struct {
	mu   sync.Mutex
	list []priorityItem[T] //最高优先级的，放在第0位，取的时候，从0开始取
}

func NewPriorityList[T any]() *PriorityList[T] {
	return &PriorityList[T]{
		list: make([]priorityItem[T], 0),
	}
}

func (p *PriorityList[T]) Push(priority int, item T) {
	p.mu.Lock()
	defer p.mu.Unlock()

	l := priorityItem[T]{
		data:     item,
		priority: priority,
	}
	if len(p.list) == 0 {
		p.list = append(p.list, l)
		return
	}
	index := len(p.list)
	for i := 0; i < len(p.list); i++ {
		if l.priority >= p.list[i].priority {
			index = i
			break
		}
	}
	//插入到i之前
	oldList := p.list
	left := oldList[0:index]
	right := oldList[index:]
	p.list = make([]priorityItem[T], 0)
	p.list = append(p.list, left...)
	p.list = append(p.list, l)
	p.list = append(p.list, right...)
}

// 优先级高的先出
func (p *PriorityList[T]) Pop() (T, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.list) == 0 {
		var a T
		return a, errors.New("no item")
	}

	oldList := p.list
	l := oldList[0]
	p.list = oldList[1:]
	return l.data, nil
}
