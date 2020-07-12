# Natswrap

`natsx` is a  simple wrapper for [nats.io](nats-io/go-nats) client. 

## Usage 

To start use the `natsx` package add import:

``` go
...
import (
  "github.com/lancer-kit/armory/natsx"
)
...
```

- Fill config structure:

| Field | Type | Required |
| ----- | ---- | ---- |
| Host | string | + |
| Port | int | + |
| User | string |   |
| Password | string |

- Set config: 

``` go
natsx.SetCfg(cfg)
```

Connect will be initialized at first try to push or subscribe a message
``` go
err := PublishMessage(topic, obj)
```

#### Example

``` go
package main

import (
    "fmt"
    "github.com/lancer-kit/armory/natsx"
)

var config = natsx.Config{
    Host: "127.0.0.1",
    Port: 4222,
    User: "user",
    Password: "password",  
}

err := config.Validate()
if err != nil {
    fmt.Println('invalid nats configuration')
}

natsx.SetConfig(&config)

testMsg := []string {"1", "2"}
err := natsx.PublishMessage("Topic", testMsg)
if err != nil {
    log.Get().Error(err)
}

```

