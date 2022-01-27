package initialization

import (
	"time"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/tools"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// Modules is a set of modules for initialization.
type Modules []Module

// Add push new module to set.
func (modules Modules) Add(m *Module) Modules {
	if m == nil {
		return modules
	}
	return append(modules, *m)
}

// InitAll checks status and run initialization of all modules for the fist to last.
func (modules Modules) InitAll() error {
	if err := modules.validate(); err != nil {
		return errors.Wrap(err, "set is invalid")
	}

	locks := make(map[string]chan bool)
	for i := range modules {
		locks[modules[i].Name] = make(chan bool, 1)
	}

	eg := errgroup.Group{}
	for i := range modules {
		m := modules[i]
		eg.Go(func() error { return m.Run(locks) })
	}

	err := eg.Wait()
	if err != nil {
		return errors.Wrap(err, "set is invalid")
	}

	for i := range modules {
		close(locks[modules[i].Name])
	}
	return nil
}

// validate checks that there are no duplicates and dependency cycle is set.
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

// Module is unit for initialization.
type Module struct {
	// Name unique identifier for module.
	Name string
	// DependsOn is a name of other module which should be initialized before this.
	DependsOn string
	// InitInterval is initial timeout before second attempt to perform initCall.
	InitInterval time.Duration
	// Timeout is a deadline for the initialization attempts.
	Timeout time.Duration
	// Init is a module init function.
	Init func(*logrus.Entry) error
}

// Run tries to initialize the module, if the attempt failed, it will be repeated,
// but with an incremental delay. So it will be until the maximum initialization timeout is reached.
func (mod *Module) Run(locks map[string]chan bool) error {
	if err := mod.validate(); err != nil {
		return errors.New("can't validate " + mod.Name)
	}

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
		return errors.New("can't init " + mod.Name)
	}

	locks[mod.Name] <- true
	return nil
}

func (mod *Module) initCall() bool {
	err := mod.Init(log.Get().WithField("init", mod.Name))
	if err != nil {
		log.Get().WithError(err).Error("Can't init " + mod.Name)
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
