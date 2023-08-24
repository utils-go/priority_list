package priority_list

import (
	"testing"
)

func TestNewPriorityList(t *testing.T) {

	list := NewPriorityList[Person]()
	list.Push(2, Person{Name: "2"})
	list.Push(3, Person{Name: "3"})
	list.Push(1, Person{Name: "1"})
	list.Push(5, Person{Name: "5"})
	list.Push(4, Person{Name: "4"})
	list.Push(100, Person{Name: "100"})
	list.Push(101, Person{Name: "101"})
	list.Push(50, Person{Name: "50"})
	list.Push(101, Person{Name: "101+"})
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
