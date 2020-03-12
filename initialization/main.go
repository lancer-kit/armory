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

func (modules Modules) Add(m *Module) Modules {
	if m == nil {
		return modules
	}
	return append(modules, *m)
}

func (modules Modules) InitAll() {
	if err := modules.validate(); err != nil {
		log.Default.Fatal(err)
	}
	wg := &sync.WaitGroup{}

	locks := make(map[string]chan bool)
	for i := range modules {
		locks[modules[i].Name] = make(chan bool, 1)
	}

	for i := range modules {
		wg.Add(1)
		go modules[i].Run(wg, locks)
	}

	wg.Wait()
	for i := range modules {
		close(locks[modules[i].Name])
	}
	return
}

func (modules Modules) validate() error {
	if d, ok := modules.duplicates(); ok {
		return errors.New("found duplicated modules: " + d)
	}

	if c, ok := modules.cycle(); ok {
		return errors.New("found dependency cycle: " + c)
	}

	return nil
}

func (modules Modules) duplicates() (string, bool) {
	duplicatesMap := make(map[string]int, len(modules))
	for i := range modules {
		duplicatesMap[modules[i].Name]++
	}
	for k, v := range duplicatesMap {
		if v > 1 {
			return k, true
		}
	}
	return "", false
}

func (modules Modules) cycle() (string, bool) {
	depsMap := make(map[string]string, len(modules))
	for i := range modules {
		depsMap[modules[i].Name] = modules[i].DependsOn
	}
	for i := range modules {
		current := modules[i].DependsOn
		cycle := modules[i].Name
		for len(current) != 0 {
			cycle = cycle + " -> " + current
			if current == modules[i].Name {
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

func (mod *Module) Run(wg *sync.WaitGroup, locks map[string]chan bool) {
	if err := mod.validate(); err != nil {
		log.Default.Fatal(err)
	}
	defer wg.Done()

	for {
		if len(mod.DependsOn) == 0 {
			break
		}
		if ok := <-locks[mod.DependsOn]; ok {
			locks[mod.DependsOn] <- true
			break
		}
	}

	ok := tools.RetryIncrementallyUntil(
		mod.InitInterval,
		mod.Timeout,
		mod.initCall,
	)
	if !ok {
		log.Default.Fatalf("Can't init %s", mod.Name)
	}

	locks[mod.Name] <- true
}

func (mod *Module) initCall() bool {
	err := mod.Init(log.Default.WithField("init", mod.Name))
	if err != nil {
		log.Default.WithError(err).Error("Can't init " + mod.Name)
	}
	return err == nil
}

func (mod *Module) validate() error {
	var err error
	baseErr := errors.New("invalid module init")

	const notEmpty = "%s should not be empty"
	if len(mod.Name) == 0 {
		err = errors.Wrapf(baseErr, notEmpty, "Name")
	}

	if mod.Timeout.Nanoseconds() == 0 {
		err = errors.Wrapf(baseErr, notEmpty, "Timeout")
	}

	if mod.InitInterval.Nanoseconds() == 0 {
		err = errors.Wrapf(baseErr, notEmpty, "Interval")
	}

	if mod.Init == nil {
		err = errors.Wrapf(baseErr, notEmpty, "Init")
	}

	if err != nil {
		return err
	}

	return nil
}
