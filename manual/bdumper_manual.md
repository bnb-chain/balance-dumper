# User manual of Balance Dumper

## Introduce
**Balance Dumper** is a CLI designed to dump a snapshot of balance on a particular token at certain height by starting a `fullnode` on client side. It stops the `fullnode` at the given height, and analyzes the database to find all the accounts of the specified token, exports them to a CSV file in your specified directory.

## Installation
Download the binary, place it into any folder you like. Open a terminal in that folder, check its availability
```
$ bdumper --help
Balance Dumper

Usage:
  bdumper [flags]

Flags:
      --asset string    query asset 
      --height int      query height 
  -h, --help            help for bdumper
      --home string     directory for config and data (default "${HOME}/.bdumper")
  -o, --output string   directory for storing the csv file of balance result (default "${HOME}/.bdumper")
```

## Usage
You can use the `height` flag to specify the height of the snapshot to dump, and the `asset` flag to specify the asset for the account balance to list. Assuming that we want to list the account balance of *`BNB`* at the height of *`56503900`*, simply enter the following command
```
$ bdumper --height 56503900 --asset BNB --home ~/myhome -o ~/myoutput &
```
We recommend that it run as a daemon, as it takes time in most cases.
It doesn't matter if we do not specify the `home` and `output`, because the CLI will take *${HOME}/.bdumper* as the default home, and HOME as the default output directory.

## Log
When the user executes the command, the process log is printed on the console, and output to a log file named `dumper.log` under the `HOME` folder. The log shows like the following
```
$ bdumper --height 56503900 --asset BNB --home ~/myhome -o ~/myoutput &
===>got the latest breathe block height = 56503807
===>start node,home = /Users/fletcher/.bdumper, stopAt = 56503900, StateSyncHeight = 56503807
===>node started from height = 0
===>syncing......
```
Notice that the log may stop at *`'syncing......'`* for a long time, because it would take long time to download the block data from other peers. You can `curl` localhost:27147/stauts to check if the process is running.
```
$ curl localhost:27147/status
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
      },
      "id": "086d8e31b7c149a47c772cf747d39b1ef6c23576",
      "listen_addr": "tcp://0.0.0.0:27146",
      "network": "Binance-Chain-Tigris",
      "version": "0.32.3",
      "channels": "36402021222330380041",
      "moniker": "fletcherfandeMacBook-Pro.local",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://0.0.0.0:27147"
      }
    },
    "sync_info": {
      "latest_block_hash": "7516C780F5700804A99F83984900FFED07B6FD7794D18B71F58471EF36BA10CB",
      "latest_app_hash": "FEF3B44D86F714B989C92D2134E9920F87936F2B06CA4EDC7B75D03EA1D68E2A",
      "latest_block_height": "56503807", // the latest block height downloaded from peers
      "latest_block_time": "2019-12-22T00:00:00.414320569Z",
      "catching_up": true,
      "index_height": "0"
    },
    "validator_info": {
      "address": "E52D4ABB8A467467DA7BA92B0E503C46023156B7",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "CvLiuDs81Z5N1B3Z39xHu3+/R4wdauyHj33Milrkjv8="
      },
      "voting_power": "0"
    }
  }
}
```
If the data similar to the above of `json` format is returned, it indicates that the process is still running well. The value of `latest_block_height` represents the latest block height you've downloaded from other peers. It is likely that this value will initially remain at the height of the breath block for some time, meaning that the process is synchronizing the data in *`state sync`* mode. Then this value increases, which means it has switched to fast sync mode. After it finishes synchronizing, the log shows like the below
```
===>node catches up the target height 56503900, terminal the node
```
It stops the synchronization， and starts to fetch accounts at this height.
```
===>start to fetch at height = 56503900
===>finish fetching,got 107677 matched account
```
Finally, those account balances will be ouput to a CSV file, named by asset and height, as *BNB_56503900.csv*. The following is the example of a result file

![BNB_56503900.csv](/manual/result_shot.png?raw=true "BNB_56503900.csv")

>**Note that for the balance, we take the last 8 digits as the decimal place, meaning the balance of *bnb1edrs5cukhx060e02u98v9j8spum7vhuqg9ctxd* in the above example is actually *32888888BNB*.**

## Notice
- If a folder serves as a *home* of a fullnode that you started ever, then you should be careful to use it as your home directory for this executive tool, since the historic block data could be removed by this tool.
- If you has launched a fullnode that is keeping synced with the Block Chain. You can do a quick search by using the BNCHOME as the home dir of this tool.The premise is to stop the whole node for a moment.
- Once it starts, it will take a long time to download data from other peers. For current experience, the speed of *`fast sync`* is about 20 blocks one second. The longer the height is from the *breathe block* of the day, the longer it will take.


  











