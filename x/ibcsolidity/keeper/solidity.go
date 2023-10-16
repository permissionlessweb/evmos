


var (
	// A valid IBC path for the LSM token must only consist of 1 channel hop along a transfer channel
	// (e.g. "transfer/channel-0")
	IsValidIBCPath = regexp.MustCompile(fmt.Sprintf(`^%s/(%s[0-9]{1,20})$`, transfertypes.PortID, channeltypes.ChannelPrefix)).MatchString

	// // Timeout for the validator slash query that occurs at periodic deposit intervals
	// // LSMSlashQueryTimeout = time.Minute * 5 // 5 minutes

	// // Time for the detokenization ICA
	// DetokenizationTimeout = time.Hour * 24 // 1 day
)

// This function returns the associated host zone along with the initial contract record
func (k Keeper) ValidateIBCSolidityCall(ctx sdk.Context, msg types.MsgIBCSolidityCall) (types.IBCSolidityCall, error) {

}

// Generates a unique ID for the IBC Solidity contract call so that,
// the query callback can be joined back with this tx
// The key in the store for an IBCSolidityCall is chainId + solidityAddress (meaning, there
// can only be 1 IBCSolidityCall in progress per tokenization)
func GetIBCSolidityContractID(blockHeight int64, chainID, callerAddress, contract string){

}

// Loops through all active host zones, grabs queued IBCSolidtyCalls for that host
// that are in status TRANSFER_QUEUE, and submits the IBC Transfer to the host
func (k Keeper) SendAllIBCSolidityCalls(ctx sdk.Context){

}