package payout

import "github.com/inttegro/inttegro-sdk-go/v10/customdata"

// CustomData contains merchant-defined payout metadata.
//
// Keys are chosen by the merchant; values are always strings on the wire.
type CustomData = customdata.Data
