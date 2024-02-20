package params

// App Parameters
const (
	// HumanCoinUnit is human readable representation of the coin name
	HumanCoinUnit = "furya"
	// BaseCoinUnit is the actual name of coin used in transaction
	BaseCoinUnit = "ufury"
	// FURYAExponent is the exponential digits of the coin
	FURYAExponent = 6
	// DefaultBondDenom is the default staking denom of application
	DefaultBondDenom = BaseCoinUnit
)

// Simulation parameter constants
const (
	StakePerAccount           = "stake_per_account"
	InitiallyBondedValidators = "initially_bonded_validators"
)
