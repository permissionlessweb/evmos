# Records

Handles record keeping and accounting for the Evmos blockchain.

IBC Solidity Call Records
- `GetIBCSolidityCallCount()`
- `SetIBCSolidityCallCount()`
- `AppendIBCSolidityCallRecord()`
- `SetIBCSolidityCall()`
- `GetIBCSolidityCall()`
- `RemoveIBCSolidityCall()`
- `GetAllIBCSolidityCallRecord()`
- `GetTransferIBCSolidityRecordByEpochAndChain()`

Epoch Records

- `SetEpochRecord()`
- `GetEpochRecord()`
- `RemoveEpochRecord()`
- `GetAllEpochRecord()`
- `GetAllPreviousEpochRecords()`
- `GetHostZoneByChainID()` 
- `AddHostZoneToEpochRecord()`
- `SetHostZoneEpochs()`


## State

Callbacks

- `TransferCallback`

Genesis

- `Params`
- `RecordsPacketData`
- `NoData`
- `IBCSolidityCallRecord`
- `HostZoneEpochs`
- `EpochsRecord`
- `GenesisState`

## Queries 

- `Params`
- `GetIBCSolidityCallRecord`
- `AllIBCSolidityCallRecord`
- `GetEpochRecord`
- `AllEpochRecord`

## Events 

The `records` module emits does not currently emit any events.
