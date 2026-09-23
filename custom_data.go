package inttegro

import "github.com/zebodotdev/inttegro-sdk-go/v10/customdata"

const (
	MaxCustomDataKeyBytes = customdata.MaxKeyBytes
	MaxCustomDataBytes    = customdata.MaxBytes
)

var (
	ErrCustomDataKeyTooLong = customdata.ErrKeyTooLong
	ErrCustomDataTooLarge   = customdata.ErrTooLarge
)

// CustomData is merchant-defined string data attached to an API resource.
type CustomData = customdata.Data

// CustomDataInput is open-ended custom data accepted by create and replace requests.
type CustomDataInput = customdata.Input

// CustomDataPatch records custom-data merge operations.
type CustomDataPatch = customdata.Patch

func NewCustomData(values ...map[string]string) (*CustomData, error) {
	return customdata.New(values...)
}

func NewCustomDataInput(values ...map[string]any) (*CustomDataInput, error) {
	return customdata.NewInput(values...)
}

func NewCustomDataPatch() *CustomDataPatch {
	return customdata.NewPatch()
}
