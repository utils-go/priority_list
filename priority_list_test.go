package priority_list

import (
	"testing"
)

func TestNewPriorityList(t *testing.T) {

	list := NewPriorityList[Person]()
	list.Push(2, Person{Name: "2"})
	list.Push(3, Person{Name: "3"})
	list.Push(1, Person{Name: "1"})

	for {
		p, err := list.Pop()
		if err != nil {
			break
		}
		t.Log(p)
	}
}

type Person struct {
	Name string
}
