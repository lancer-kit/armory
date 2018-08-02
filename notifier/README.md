# notifier
This package is used for sending notifications to NATS.

## Sender
`Sender` is a structure that works with the notification service. The structure contains `*Config`, `*nats.Conn` and `*logrus.Entry`.

#### Method list:
- `Send(*Message) error` – main functional of `Sender`. Publish message to service.
- `SetConfig(*Config) *Sender` – set new configuration.
- `SetLogger(*logrus.Entry) *Sender` – set logger.
- `ErrorLog(error, string)` – write error message into log.
- `IsConnected() error` – check connection to NATS.
- `Disconnect() *Sender` – disconnec.

## SenderI
`SenderI` is an interface for `Sender`.
#### Method list:
- `SetLogger(*logrus.Entry) *Sender`
- `SetConfig(*Config) *Sender`
- `Disconnect() *Sender`
- `Send(*Message) error`

## Config
`Config` is a structure that contains configuration for `Sender`. It has the following fields:  

| Parameter | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| Chanel    | string | Chanel (topic) name. Default = "notifier" |
| Url       | string | NATS connection url                       |
| User      | string | NATS user                                 |
| Password  | string | NATS password                             |
| Token     | string | NATS auth token                           |

## Message
`Message` – main message structure. Used to send notifications to `socket-notifier`.  
`Message` has the following fields:

| Parameter   | Type        | Description                  | JSON name   |
|-------------|-------------|------------------------------|-------------|
| UserId      | int         | User ID for notifications    | userId      |
| Command     | string      | "Command" for front-end      | command     |
| IsBroadcast | bool        | Private or broadcast message | isBroadcast |
| Data        | interface{} | Optional data                | Data        |

`Command` field shows which type of update and which structure is used in `Data` field. Here is the list of possible commands:

- "operation" – update of operation status. Data: WSOperationMessage.
- "payment" – update of payment status. Data: WSPaymentMessage.
- "transaction" – update of transaction status. Data: WSTxMessage.

## WSTxMessage
`WSTxMessage` is a structure used for sending transaction updates to a `socket-notifier`.

#### Fields:
| Parameter | Type         | Description             | JSON name |
|-----------|--------------|-------------------------|-----------|
| TxID      | String       | Transaction ID          | txId      |
| TxType    | txst.TxType  | Transaction type        | txType    |
| TxState   | txst.TxState | Transaction state       | txState   |
| UpdatedAt | int64        | Transaction update time | updatedAt |

#### Usage:

``` go
notifier.Default.Send(&notifier.Message{
			UserId:      userID,
			Command:     "transaction",
			IsBroadcast: false,
			Data: &notifier.WSTxMessage{
				TxID:      tx.TxID,
				TxType:    tx.Type,
				TxState:   tx.State,
				UpdatedAt: tx.UpdatedAt,
			},
		})
```

JSON output:

``` json
{
	"userId": %USER_ID%,
	"command": "transaction",
	"isBroadcast": false,
	"Data": {
		"txId": %TX_ID%,
		"txType": %TX_TYPE%,
		"txState": %TX_STATE%,
		"updatedAt": %UPDATED_AT%
	}
}
```

## WSOperationMessage

`WSOperationMessage` is a structure used for sending operation updates to a `socket-notifier`.

#### Fields:
| Parameter   | Type               | Description           | JSON name     |
|-------------|--------------------|-----------------------|---------------|
| OperationID | string             | Operation ID          | operationId   |
| Type        | txst.OperationType | Operation type        | operationType |
| Status      | string             | Operation status      | status        |
| TxID        | string             | Transaction ID        | txId          |
| UpdatedAt   | int64              | Operation update time | updatedAt     |

#### Usage:

```go
notifier.Default.Send(&notifier.Message{
			UserId:      userID,
			Command:     "operation",
			IsBroadcast: false,
			Data: &notifier.WSTxMessage{
				OperationID: op.OperationID,
				Type:        op.Type,
				Status:      status,
				TxID:        *op.TxID,
				UpdatedAt:   op.UpdatedAt,
			},
		})

```

JSON output:

```json
{
	"userId": %USER_ID%,
	"command": "operation",
	"isBroadcast": false,
	"Data": {
		"operationId": %OPERATION_ID%,
		"operationType": %OPERATION_TYPE%,
		"status": %STATUS%,
		"txId": %TX_ID%,
		"updatedAt": %UPDATED_AT%
	}
}
```

## WSPaymentMessage

`WSPaymentMessage` is a structure used for sending payment updates to a `socket-notifier`.

#### Fields:
| Parameter    | Type              | Description          | JSON name    |
|--------------|-------------------|----------------------|--------------|
| PaymentID    | string            | Payment ID           | paymentId    |
| State        | txst.PaymentState | Payment State        | state        |
| Type         | txst.PaymentType  | Payment Type         | type         |
| Amount       | currency.Coin     | Payment Amount       | amount       |
| FromWalletID | string            | Payer's Wallet ID    | fromWalletId |
| ToWalletID   | string            | Acceptor's Wallet ID | toWalletId   |
| UpdatedAt    | int64             | Payment update time  | updatedAt    |

#### Usage:

```go
notifier.Default.Send(&notifier.Message{
			UserId:      userID,
			Command:     "payment",
			IsBroadcast: false,
			Data: &notifier.WSPaymentMessage{
				PaymentID:    payment.PaymentID,
				State:        payment.State,
				Type:         payment.Type,
				Amount:       payment.Amount,
				FromWalletID: payment.FromWalletID,
				ToWalletID:   payment.ToWalletID,
				UpdatedAt:    payment.UpdatedAt,
			},
		})
```

JSON output:

```json
{
	"userId": %USER_ID%,
	"command": "payment",
	"isBroadcast": false,
	"Data": {
		"paymentId": %PAYMENT_ID%,
		"state": %STATE%,
		"type": %TYPE%,
		"amount": %AMOUNT%,
		"fromWalletId": %FROM_WALLET_ID%,
		"toWalletId": %TO_WALLET_ID%,
		"updatedAt": %UPDATED_AT%
	}
}
```