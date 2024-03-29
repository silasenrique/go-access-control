package helper

import "fmt"

type helper struct {
	sentinel error
	internal error
}

func NewHelper(sentinel error) *helper {
	return &helper{sentinel: sentinel}
}

func (p helper) Error() string {
	return fmt.Sprintf("%s: %s", p.sentinel, p.internal)
}

func (p *helper) AddIntenal(err error) *helper {
	p.internal = err

	return p
}
