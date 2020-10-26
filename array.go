package dict

import "fmt"

//Array array
type Array []interface{}

//NewArray new array
func NewArray(values ...interface{}) *Array {
	arr := &Array{}
	if len(values) != 0 {
		arr.Append(values...)
	}

	return arr
}

//Length len
func (self *Array) Length() int {
	return len(*self)
}

//Clear clear
func (self *Array) Clear() *Array {
	*self = (*self)[:0]
	return self
}

func (self *Array) Contains(value interface{}) bool {
	_, found := self.Index(value)
	return found
}

func (self *Array) Index(value interface{}) (index int, found bool) {
	index = 0
	found = false

	for i, v := range *self {
		if v == value {
			index = i
			found = true
			break
		}
	}

	return
}

func (self *Array) Append(values ...interface{}) *Array {
	*self = append(*self, values...)
	return self
}

func (self *Array) Remove(value interface{}) bool {
	index, ok := self.Index(value)

	if ok {
		self.Pop(index)
	}

	return ok
}

func (self *Array) Pop(index int) (interface{}, bool) {
	if index >= len(*self) {
		return nil, false
	}

	value := (*self)[index]
	*self = append((*self)[:index], (*self)[index+1:]...)

	return value, true
}

func (self *Array) Peek(index int) interface{} {
	if index >= len(*self) {
		Raise(fmt.Sprintf("peek out of index: %d", index))
	}

	return (*self)[index]
}
