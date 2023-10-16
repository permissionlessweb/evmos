package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Default init params
var (
	// these are default intervals _in epochs_ NOT in blocks
	DefaultIBCSolidityCallInterval   uint64 = 1
	// DefaultDelegateInterval      	 uint64 = 1
	// DefaultReinvestInterval      	 uint64 = 1
	// DefaultRewardsInterval       	 uint64 = 1
	// DefaultRedemptionRateInterval uint64 = 1
	// you apparently cannot safely encode floats, so we make commission / 100
	DefaultEvmosCommission             uint64 = 5
	DefaultICATimeoutNanos              uint64 = 600000000000
	DefaultBufferSize                   uint64 = 5             // 1/5=20% of the epoch
	DefaultIbcTimeoutBlocks             uint64 = 300           // 300 blocks ~= 30 minutes
	DefaultFeeTransferTimeoutNanos      uint64 = 1800000000000 // 30 minutes
	// DefaultMinRedemptionRateThreshold   uint64 = 90            // divide by 100, so 90 = 0.9
	// DefaultMaxRedemptionRateThreshold   uint64 = 150           // divide by 100, so 150 = 1.5
	DefaultMaxSolidityICACallsPerEpoch     uint64 = 100
	DefaultIBCTransferTimeoutNanos      uint64 = 1800000000000 // 30 minutes
	// DefaultValidatorSlashQueryThreshold uint64 = 1             // denominated in percentage of TVL (1 => 1%)

	// KeyDepositInterval is store's key for the DepositInterval option
	KeyIBCSolidityCallInterval                   = []byte("IBCSolidityCallInterval")
	// KeyDelegateInterval                  = []byte("DelegateInterval")
	// KeyReinvestInterval                  = []byte("ReinvestInterval")
	KeyRewardsInterval                   = []byte("RewardsInterval")
	// KeyRedemptionRateInterval            = []byte("RedemptionRateInterval")
	KeyEvmosCommission                  = []byte("EvmosCommission")
	KeyICATimeoutNanos                   = []byte("ICATimeoutNanos")
	KeyFeeTransferTimeoutNanos           = []byte("FeeTransferTimeoutNanos")
	KeyBufferSize                        = []byte("BufferSize")
	KeyIbcTimeoutBlocks                  = []byte("IBCTimeoutBlocks")
	// KeyDefaultMinRedemptionRateThreshold = []byte("DefaultMinRedemptionRateThreshold")
	// KeyDefaultMaxRedemptionRateThreshold = []byte("DefaultMaxRedemptionRateThreshold")
	KeyMaxIBCSolidityICACallsPerEpoch          = []byte("MaxIBCSolidityICACallsPerEpoch")
	KeyIBCTransferTimeoutNanos           = []byte("IBCTransferTimeoutNanos")
	// KeyMaxRedemptionRates                = []byte("MaxRedemptionRates")
	// KeyMinRedemptionRates                = []byte("MinRedemptionRates")
	// KeyValidatorSlashQueryThreshold      = []byte("ValidatorSlashQueryThreshold")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	ibcSolidityCallInterval uint64,
	// delegateInterval uint64,
	rewardsInterval uint64,
	// redemptionRateInterval uint64,
	evmosCommission uint64,
	// reinvestInterval uint64,
	icaTimeoutNanos uint64,
	bufferSize uint64,
	ibcTimeoutBlocks uint64,
	feeTransferTimeoutNanos uint64,
	maxIBCSolidityIcaCallsPerEpoch uint64,
	// defaultMinRedemptionRateThreshold uint64,
	// defaultMaxRedemptionRateThreshold uint64,
	ibcTransferTimeoutNanos uint64,
	// validatorSlashQueryInterval uint64,
) Params {
	return Params{
		IBCSolidityCallInterval:           ibcSolidityInterval,
		// DelegateInterval:                  delegateInterval,
		RewardsInterval:                   rewardsInterval,
		// RedemptionRateInterval:            redemptionRateInterval,
		EvmosCommission:                  evmosCommission,
		// ReinvestInterval:                  reinvestInterval,
		IcaTimeoutNanos:                   icaTimeoutNanos,
		BufferSize:                        bufferSize,
		IbcTimeoutBlocks:                  ibcTimeoutBlocks,
		FeeTransferTimeoutNanos:           feeTransferTimeoutNanos,
		MaxIBCSolidityIcaCallsPerEpoch:          maxIBCSolidityIcaCallsPerEpoch,
		// DefaultMinRedemptionRateThreshold: defaultMinRedemptionRateThreshold,
		// DefaultMaxRedemptionRateThreshold: defaultMaxRedemptionRateThreshold,
		IbcTransferTimeoutNanos:           ibcTransferTimeoutNanos,
		// ValidatorSlashQueryThreshold:      validatorSlashQueryInterval,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultIBCSolidityCallInterval,
		// DefaultDelegateInterval,
		DefaultRewardsInterval,
		// DefaultRedemptionRateInterval,
		DefaultEvmosCommission,
		// DefaultReinvestInterval,
		DefaultICATimeoutNanos,
		DefaultBufferSize,
		DefaultIbcTimeoutBlocks,
		DefaultFeeTransferTimeoutNanos,
		DefaultMaxIBCSolidityICACallsPerEpoch,
		// DefaultMinRedemptionRateThreshold,
		// DefaultMaxRedemptionRateThreshold,
		DefaultIBCTransferTimeoutNanos,
		// DefaultValidatorSlashQueryThreshold,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIBCSolidityCallInterval, &p.IBCSolidityCallInterval, isPositive),
		// paramtypes.NewParamSetPair(KeyDelegateInterval, &p.DelegateInterval, isPositive),
		paramtypes.NewParamSetPair(KeyRewardsInterval, &p.RewardsInterval, isPositive),
		// paramtypes.NewParamSetPair(KeyRedemptionRateInterval, &p.RedemptionRateInterval, isPositive),
		paramtypes.NewParamSetPair(KeyEvmosCommission, &p.EvmosCommission, isCommission),
		// paramtypes.NewParamSetPair(KeyReinvestInterval, &p.ReinvestInterval, isPositive),
		paramtypes.NewParamSetPair(KeyICATimeoutNanos, &p.IcaTimeoutNanos, isPositive),
		paramtypes.NewParamSetPair(KeyBufferSize, &p.BufferSize, isPositive),
		paramtypes.NewParamSetPair(KeyIbcTimeoutBlocks, &p.IbcTimeoutBlocks, isPositive),
		paramtypes.NewParamSetPair(KeyFeeTransferTimeoutNanos, &p.FeeTransferTimeoutNanos, validTimeoutNanos),
		paramtypes.NewParamSetPair(KeyMaxIBCSolidityICACallsPerEpoch, &p.MaxIBCSolidityIcaCallsPerEpoch, isPositive),
		paramtypes.NewParamSetPair(KeyDefaultMinRedemptionRateThreshold, &p.DefaultMinRedemptionRateThreshold, validMinRedemptionRateThreshold),
		paramtypes.NewParamSetPair(KeyDefaultMaxRedemptionRateThreshold, &p.DefaultMaxRedemptionRateThreshold, validMaxRedemptionRateThreshold),
		paramtypes.NewParamSetPair(KeyIBCTransferTimeoutNanos, &p.IbcTransferTimeoutNanos, validTimeoutNanos),
		paramtypes.NewParamSetPair(KeyValidatorSlashQueryThreshold, &p.ValidatorSlashQueryThreshold, isPositive),
	}
}

func validTimeoutNanos(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	tenMin := uint64(600000000000)
	oneHour := uint64(600000000000 * 6)

	if ival < tenMin {
		return fmt.Errorf("parameter must be g.t. 600000000000ns: %d", ival)
	}
	if ival > oneHour {
		return fmt.Errorf("parameter must be less than %dns: %d", oneHour, ival)
	}
	return nil
}

// func validMaxRedemptionRateThreshold(i interface{}) error {
// 	ival, ok := i.(uint64)
// 	if !ok {
// 		return fmt.Errorf("parameter not accepted: %T", i)
// 	}

// 	maxVal := uint64(1000) // divide by 100, so 1000 => 10

// 	if ival > maxVal {
// 		return fmt.Errorf("parameter must be l.t. 1000: %d", ival)
// 	}

// 	return nil
// }

// func validMinRedemptionRateThreshold(i interface{}) error {
// 	ival, ok := i.(uint64)
// 	if !ok {
// 		return fmt.Errorf("parameter not accepted: %T", i)
// 	}

// 	minVal := uint64(75) // divide by 100, so 75 => 0.75

// 	if ival < minVal {
// 		return fmt.Errorf("parameter must be g.t. 75: %d", ival)
// 	}

// 	return nil
// }

func isPositive(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	if ival <= 0 {
		return fmt.Errorf("parameter must be positive: %d", ival)
	}
	return nil
}

func isCommission(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("commission not accepted: %T", i)
	}

	if ival > 100 {
		return fmt.Errorf("commission must be less than 100: %d", ival)
	}
	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := isPositive(p.IBCSolidityCallInterval); err != nil {
		return err
	}
	// if err := isPositive(p.DelegateInterval); err != nil {
	// 	return err
	// }
	if err := isPositive(p.RewardsInterval); err != nil {
		return err
	}
	// if err := isPositive(p.RedemptionRateInterval); err != nil {
	// 	return err
	// }
	if err := isCommission(p.EvmosCommission); err != nil {
		return err
	}
	// if err := isPositive(p.ReinvestInterval); err != nil {
	// 	return err
	// }
	if err := isPositive(p.IcaTimeoutNanos); err != nil {
		return err
	}
	if err := isPositive(p.BufferSize); err != nil {
		return err
	}
	if err := isPositive(p.IbcTimeoutBlocks); err != nil {
		return err
	}
	if err := validTimeoutNanos(p.FeeTransferTimeoutNanos); err != nil {
		return err
	}
	if err := isPositive(p.MaxIBCSolidityIcaCallsPerEpoch); err != nil {
		return err
	}
	// if err := validMinRedemptionRateThreshold(p.DefaultMinRedemptionRateThreshold); err != nil {
	// 	return err
	// }
	// if err := validMaxRedemptionRateThreshold(p.DefaultMaxRedemptionRateThreshold); err != nil {
	// 	return err
	// }
	if err := validTimeoutNanos(p.IbcTransferTimeoutNanos); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
