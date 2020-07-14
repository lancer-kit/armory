package natsx

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

var (
	cfg      *Config    // nolint:gochecknoglobals
	natsConn *nats.Conn // nolint:gochecknoglobals
)

// SetConfig sets NATS configuration singleton.
func SetConfig(config *Config) {
	cfg = config
}

// GetConn returns nats.Conn singleton.
func GetConn() (*nats.Conn, error) {
	if natsConn != nil {
		return natsConn, nil
	}

	if cfg == nil {
		return nil, errors.New("Nats config didn't set")
	}

	var err error
	natsConn, err = nats.Connect(cfg.ToURL(), nats.UserInfo(cfg.User, cfg.Password))
	return natsConn, err
}

// PublishMessage serializes the message in JSON and sends to the topic.
func PublishMessage(topic string, msg interface{}) error {
	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err,
			"unable to marshal message for the topic - "+topic)
	}

	_, err = GetConn()
	if err != nil {
		return errors.Wrap(err, "unable to establish connection")
	}

	err = natsConn.Publish(topic, rawMsg)
	if err != nil {
		return errors.Wrap(err, "unable to publish into the topic "+topic)
	}

	return nil
}

// Subscribe initiates chan subscription for topic.
func Subscribe(topic string, msgs chan *nats.Msg) (*nats.Subscription, error) {
	_, err := GetConn()
	if err != nil {
		return nil, errors.Wrap(err, "unable to establish connection")
	}

	subscription, err := natsConn.ChanSubscribe(topic, msgs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to subscribe into the topic "+topic)
	}
	return subscription, nil
}
