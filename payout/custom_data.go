package payout

// CustomData contains merchant-defined payout metadata.
//
// Keys are chosen by the merchant; values are always strings on the wire.
type CustomData map[string]string
