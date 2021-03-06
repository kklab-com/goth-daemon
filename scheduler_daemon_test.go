package kkdaemon

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedulerDaemon(t *testing.T) {
	daemon := &testSchedulerDaemon{}
	assert.Nil(t, RegisterDaemon(daemon))
	testSchedulerPerFiveMinuteDaemon := &testSchedulerPerFiveMinuteDaemon{}
	assert.Nil(t, RegisterDaemon(testSchedulerPerFiveMinuteDaemon))
	testSchedulerPerFiveSecondDaemon := &testSchedulerPerFiveSecondDaemon{}
	assert.Nil(t, RegisterDaemon(testSchedulerPerFiveSecondDaemon))
	assert.Nil(t, Start())
	assert.Equal(t, GetDaemon(daemon.Name()).Next.Format("2006-01-02 15:04:05"), time.Now().Truncate(time.Minute).Add(time.Minute).Format("2006-01-02 15:04:05"))
	assert.Equal(t, GetDaemon(testSchedulerPerFiveMinuteDaemon.Name()).Next.Format("2006-01-02 15:04:05"), time.Now().Truncate(time.Minute*5).Add(time.Minute*5).Format("2006-01-02 15:04:05"))
	assert.Equal(t, GetDaemon(testSchedulerPerFiveSecondDaemon.Name()).Next.Format("2006-01-02 15:04:05"), time.Now().Truncate(time.Second*5).Add(time.Second*5).Format("2006-01-02 15:04:05"))
	assert.Equal(t, 1, daemon.start)
	testSchedulerPerSecondDaemon := &testSchedulerPerSecondDaemon{t: t}
	assert.Nil(t, RegisterDaemon(testSchedulerPerSecondDaemon))
	assert.Nil(t, StartDaemon(testSchedulerPerSecondDaemon.Name()))
	<-time.After(time.Second * 6)
	assert.Equal(t, 0, testSchedulerPerFiveMinuteDaemon.loop)
	assert.True(t, testSchedulerPerFiveSecondDaemon.loop > 0)
	assert.True(t, testSchedulerPerSecondDaemon.loop > 5)
	assert.Nil(t, Stop(syscall.SIGKILL))
	assert.Equal(t, "testSchedulerDaemon", daemon.Name())
	assert.Equal(t, "testSchedulerPerFiveMinuteDaemon", testSchedulerPerFiveMinuteDaemon.Name())
	assert.Equal(t, "testSchedulerPerFiveSecondDaemon", testSchedulerPerFiveSecondDaemon.Name())
	assert.Equal(t, 1, daemon.stop)
}

type testSchedulerDaemon struct {
	DefaultSchedulerDaemon
	start int
	loop  int
	stop  int
}

func (d *testSchedulerDaemon) When() CronSyntax {
	return "* * * * *"
}

func (d *testSchedulerDaemon) Start() {
	d.start = 1
}

func (d *testSchedulerDaemon) Loop() error {
	d.loop++
	return nil
}

func (d *testSchedulerDaemon) Stop(sig os.Signal) {
	d.stop = 1
}

type testSchedulerPerFiveMinuteDaemon struct {
	DefaultSchedulerDaemon
	start int
	loop  int
	stop  int
}

func (d *testSchedulerPerFiveMinuteDaemon) When() CronSyntax {
	return "*/5 * * * *"
}

func (d *testSchedulerPerFiveMinuteDaemon) Start() {
	d.start = 1
}

func (d *testSchedulerPerFiveMinuteDaemon) Loop() error {
	d.loop++
	return nil
}

func (d *testSchedulerPerFiveMinuteDaemon) Stop(sig os.Signal) {
	d.stop = 1
}

type testSchedulerPerFiveSecondDaemon struct {
	DefaultSchedulerDaemon
	start int
	loop  int
	stop  int
}

func (d *testSchedulerPerFiveSecondDaemon) When() CronSyntax {
	return "*/5 * * * * *"
}

func (d *testSchedulerPerFiveSecondDaemon) Start() {
	d.start = 1
}

func (d *testSchedulerPerFiveSecondDaemon) Loop() error {
	d.loop++
	return nil
}

func (d *testSchedulerPerFiveSecondDaemon) Stop(sig os.Signal) {
	d.stop = 1
}

type testSchedulerPerSecondDaemon struct {
	DefaultSchedulerDaemon
	start int
	loop  int
	stop  int
	t     *testing.T
}

func (d *testSchedulerPerSecondDaemon) When() CronSyntax {
	return "* * * * * *"
}

func (d *testSchedulerPerSecondDaemon) Start() {
	assert.Equal(d.t, StateRun, d.State())
	d.start = 1
}

func (d *testSchedulerPerSecondDaemon) Loop() error {
	assert.Equal(d.t, StateRun, d.State())
	d.loop++
	return nil
}

func (d *testSchedulerPerSecondDaemon) Stop(sig os.Signal) {
	assert.Equal(d.t, StateStop, d.State())
	d.stop = 1
}
