package natswrap

type Message struct {
	EventID string      `json:"eventId"`
	Result  string      `json:"result"`
	Msg     string      `json:"msg"`
	Details interface{} `json:"details"`
}

const (
	TopicTxDepositNew     = "tx.deposit.new"
	TopicTxDepositResults = "tx.deposit.results"

	TopicTxPaymentNew     = "tx.payment.new"
	TopicTxPaymentResults = "tx.payment.results"

	TopicTxCommissionClearingNew     = "tx.commission_clearing.new"
	TopicTxCommissionClearingResults = "tx.commission_clearing.results"

	TopicTxRepaymentNew     = "tx.repayment.new"
	TopicTxRepaymentResults = "tx.repayment.results"
)
