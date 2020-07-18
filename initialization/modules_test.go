package initialization

import (
	"errors"
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

func Example() {
	moduleList := Modules{
		Module{
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

		Module{
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

		Module{
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
