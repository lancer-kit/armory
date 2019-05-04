package initialization

import (
	"sync"
	"time"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/tools"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Modules []Module

func (mm Modules) Add(m *Module) Modules {
	if m == nil {
		return mm
	}
	return append(mm, *m)
}

func (mm Modules) InitAll() {
	if err := mm.validate(); err != nil {
		log.Default.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	locks := make(map[string]chan bool)
	for i := range mm {
		locks[mm[i].Name] = make(chan bool, 1)
	}
	for i := range mm {
		wg.Add(1)
		go mm[i].Run(wg, locks)
	}
	wg.Wait()
	for i := range mm {
		close(locks[mm[i].Name])
	}
	return
}

func (mm Modules) validate() error {
	if d, ok := mm.duplicates(); ok {
		return errors.New("found duplicated modules: " + d)
	}

	if c, ok := mm.cycle(); ok {
		return errors.New("found dependency cycle: " + c)
	}

	return nil
}

func (mm Modules) duplicates() (string, bool) {
	duplicatesMap := make(map[string]int, len(mm))
	for i := range mm {
		duplicatesMap[mm[i].Name]++
	}
	for k, v := range duplicatesMap {
		if v > 1 {
			return k, true
		}
	}
	return "", false
}

func (mm Modules) cycle() (string, bool) {
	depsMap := make(map[string]string, len(mm))
	for i := range mm {
		depsMap[mm[i].Name] = mm[i].DependsOn
	}
	for i := range mm {
		current := mm[i].DependsOn
		cycle := mm[i].Name
		for len(current) != 0 {
			cycle = cycle + " -> " + current
			if current == mm[i].Name {
				return cycle, true
			}
			current = depsMap[current]

		}
	}
	return "", false
}

type Module struct {
	Name         string
	DependsOn    string
	Timeout      time.Duration
	InitInterval time.Duration
	Init         func(*logrus.Entry) error
}

func (m *Module) Run(wg *sync.WaitGroup, locks map[string]chan bool) {
	if err := m.validate(); err != nil {
		log.Default.Fatal(err)
	}
	defer wg.Done()
	for {
		if len(m.DependsOn) == 0 {
			break
		}
		if ok := <-locks[m.DependsOn]; ok {
			locks[m.DependsOn] <- true
			break
		}
	}

	ok := tools.RetryIncrementallyUntil(
		m.InitInterval,
		m.Timeout,
		m.initCall,
	)
	if !ok {
		log.Default.Fatalf("Can't init %s", m.Name)
	}
	locks[m.Name] <- true
}

func (m *Module) initCall() bool {
	err := m.Init(log.Default.WithField("init", m.Name))
	if err != nil {
		log.Default.WithError(err).Error("Can't init " + m.Name)
	}
	return err == nil
}

func (m *Module) validate() error {
	err := errors.New("invalid module init")
	ok := true
	const notEmpty = "%s should not be empty"
	if len(m.Name) == 0 {
		ok = false
		err = errors.Wrapf(err, notEmpty, "Name")
	}
	if m.Timeout.Nanoseconds() == 0 {
		ok = false
		err = errors.Wrapf(err, notEmpty, "Timeout")
	}
	if m.InitInterval.Nanoseconds() == 0 {
		ok = false
		err = errors.Wrapf(err, notEmpty, "Interval")
	}
	if m.Init == nil {
		ok = false
		err = errors.Wrapf(err, notEmpty, "Init")
	}
	if !ok {
		return err
	}
	return nil
}
