package kkdaemon

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	assert.EqualValues(t, nil, RegisterDaemon("SS", 1, &DefaultDaemon{}))

	assert.EqualValues(t, 1, GetService("SS").Order)
	assert.NotNil(t, RegisterDaemon("SS", 2, &DefaultDaemon{}))

	RegisterDaemon("P1", 1, &P1{Daemon: &DefaultDaemon{}})

	RegisterDaemon("P2", 2, &P2{Daemon: &DefaultDaemon{}})

	RegisterServiceInline("P3", 3, func() {
		println("start p3")
	}, func(sig os.Signal) {
		println("stop p3")
	})

	Start()
	Stop(nil)
}

type P1 struct {
	Daemon
}

func (p *P1) Start() {
	println("start p1")
}

func (p *P1) Stop(sig os.Signal) {
	println("stop p1")
}

type P2 struct {
	Daemon
}

func (p *P2) Start() {
	println("start p2")
}

func (p *P2) Stop(sig os.Signal) {
	println("stop p2")
}
