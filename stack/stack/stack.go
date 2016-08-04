package stack

import (
	"fmt"
	"errors" 
)

type Element struct {
	next *Element
	prev *Element
	value interface{}
}


type Stack struct {
	size int
	top *Element
}

func (e *Element) GetValue() interface{} {
	return e.value
}

func (e *Element) Update(newElem Element) {
	e = &newElem
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (s *Stack) Push (value interface{}) {

	newElem := &Element{ nil,  s.top, value }

	if s.top != nil {
		old := s.top;
		s.top = newElem
		old.next = s.top
	} else {
		s.top = newElem
	}

	s.size++
}

func (s *Stack) Get(index int) (*Element, error) {
	k := s.top
	i := 0
	for k != nil {
		if i == index {
			return k, nil
		}
		k = k.Prev()
		i++
	}
	return nil, errors.New("Unknown index")
}

func (s *Stack) Remove(index int) {
	if s.size > 0 {
		elem, err := s.Get(index)
		if err == nil{
			prev := elem.Prev()
			next := elem.Next()

			prev.next = next;
			next.prev = prev;

			s.size--;
		}
	}
}

func (s *Stack) Print() {
	top := s.top

	if top != nil {
		for top != nil {
			fmt.Println(top.value)
			top = top.Prev()
		}
	}
}

func (s *Stack) Len() int {
	return s.size
}
