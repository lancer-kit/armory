[![GoDoc](https://godoc.org/github.com/lancer-kit/armory/initialization?status.png)](https://godoc.org/github.com/lancer-kit/armory/initialization)

# initialization

Is a package for initialization of the application modules. Ex. database connectors, service providers, etc.


## Usage

```go
package main

import (
	"errors"
	"log"
	"time"

	"github.com/lancer-kit/armory/initialization"
	"github.com/sirupsen/logrus"
)

func main() {
	moduleList := initialization.Modules{
		initialization.Module{
			Name:         "database_conn",
			DependsOn:    "",
			Timeout:      15 * time.Second,
			InitInterval: time.Second,
			Init: func(entry *logrus.Entry) error {
				entry.WithField("module", "database_conn").
					Info("try to initialize module")
				time.Sleep(500 * time.Millisecond)
				return errors.New("ha-ha-ha")
			},
		},

		initialization.Module{
			Name:         "rabbit",
			DependsOn:    "",
			Timeout:      15 * time.Second,
			InitInterval: time.Second,
			Init: func(entry *logrus.Entry) error {
				entry.WithField("module", "rabbit").
					Info("try to initialize module")
				time.Sleep(500 * time.Millisecond)
				return nil
			},
		},

		initialization.Module{
			Name:         "listener",
			DependsOn:    "database_conn",
			Timeout:      15 * time.Second,
			InitInterval: time.Second,
			Init: func(entry *logrus.Entry) error {
				entry.WithField("module", "listener").
					Info("try to initialize module")
				time.Sleep(500 * time.Millisecond)
				return nil
			},
		},
	}

	err := moduleList.InitAll()
	if err != nil {
		log.Fatal("Init failed with error: ", err.Error())
	}
}
```
