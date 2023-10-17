# The IBC-Solidity Module

The IBC-Solidity Module contains the logic handling cross chain calls from cosmos-sdk & cosmwasm compatible blockchains to Evmos:

- it exposes core evm entry points to the user (call and query)
- it executes automated beginBlocker and endBlocker logic to interact with contracts on relevant host zones using Interchain Accounts
- it handles registering new host zones
- it defines all the callbacks used when issuing Interchain Account logic

ICAs allow accounts on Zone A to be controlled by Zone B. ICAs communicate with one another using Interchain Queries (ICQs), which involve Zone A querying Zone B for relevant information.

Two Zones communicate via a connection and channel. All communications between the Controller Zone (the chain that is querying) and the Host Zone (the chain that is being queried) is done through a dedicated IBC channel between the two chains, which is opened the first time the two chains interact.

For context, ICS standards define that each channel is associated with a particular connection, and a connection may have any number of associated channels.

## Params
```
BufferSize (default uint64 = 5)
SolidityCallInterval (default uint64 = 1)
MaxSolidityICACallsPerEpoch (default uint64 = 100)
FeeTransferTimeoutNanos (default uint64 = 1800000000000)
IbcTimeoutBlocks (default uint64 = 300)
ICATimeoutNanos(default uint64 = 600000000000)
IBCTransferTimeoutNanos (default uint64 = 1800000000000)
```

## Keeper functions

- `SolidityCall()`
- `SolidityQuery()`
- `RegisterHostZone()`
- `RestoreInterchainAccount()`

## State

Callbacks

HostZone

Misc

Governance


## Queries
- `QueryInterchainAccountFromAddress`
- `QueryParams`
- `QueryGetHostZone`
- `QueryAllHostZone`
- `QueryModuleAddress`
- `QueryGetEpochTracker`
- `QueryAllEpochTracker`
- `QueryGetNextPacketSequence`

## Events

`ibcsolidity` module emits the following events:

Type: Attribute Key &rarr; Attribute Value
--------------------------------------------------
registerHostZone: module &rarr; ibcsolidity
registerHostZone: connectionId &rarr; connectionId
registerHostZone: chainId &rarr; chainId
onAckPacket (IBC): module &rarr;  moduleName
onAckPacket (IBC): ack &rarr; ackInfo