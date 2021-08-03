package daemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	assert.EqualValues(t, nil, RegisterService(&DefaultService{
		name:  "SS",
		order: 1,
	}))

	assert.EqualValues(t, 1, GetService("SS").Order())
	assert.NotNil(t, RegisterService(&DefaultService{
		name:  "SS",
		order: 2,
	}))

	RegisterService(&P1{Service: &DefaultService{
		name:  "P1",
		order: 1,
	}})

	RegisterService(&P2{Service: &DefaultService{
		name:  "P2",
		order: 2,
	}})

	Start()
	Stop()
}

type P1 struct {
	Service
}

func (p *P1) Start() {
	println("start p1")
}

func (p *P1) Stop() {
	println("stop p1")
}

type P2 struct {
	Service
}

func (p *P2) Start() {
	println("start p2")
}

func (p *P2) Stop() {
	println("stop p2")
}