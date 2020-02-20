Balance Dumper
-------------
Balance Dumper is a tool to get a snapshot of balance on a particular token at certain height by starting a fullnode on client side.It will stop the `fullnode` at the given height, and analyses on the database to find out all the accounts of the specified token , then export them to a CSV file in your specified directory.


## Prepare

If you do not have golang yet, please [install it](https://golang.org/dl) or use brew on macOS: `brew install go` and `brew install dep`.

Clone the project from github

```
$ git clone http://github.com/binance-chain/balance-dumper
```

## Install

```
$ cd balance-dumper
$ go install ./bdumper
```

To test that installation worked, try to run the cli tool:
```
$ bdumper -h
Balance Dumper

Usage:
  bnbaccr [flags]

Flags:
      --asset string    query asset 
      --height int      query height 
  -h, --help            help for bdumper
      --home string     directory for config and data (default "${HOME}/.bdumper")
  -o, --output string   directory for storing the csv file of balance result (default "${HOME}/.bdumper")
```

## Execute

```
$ bdumper --home {home} --asset {asset} --height {height} -o {outputDir} &
```

The progress can be checked in the logfile named `bdumper.log` under users `home` directory.

## Strategy

User can use this tool to perform multiple query statistics. The node data is stored in the user-specified home directory. Of course, the full node does not have to launched every time and this tool is always trying to find a most efficient way to get the snapshot of a particular height. The following shows how does this tool work in different situations.

Assuming that `queryH` as the height user want to query,`breatheH` as the breathe block height of the day of `queryH`,`latestH` as the latest height in DB of user's local environment.

- If `queryH` > `latestH` which means user's query height is greater than the latest height user's `fullnode` synchronizes to, then the reporter will compare the `latestH` to `breatheH`:
  
  - If `breatheH` > `latestH` + 15000, delete the existing `fullnode` data that has been synchronized, and restart by state sync mode with ``stateSyncHeight = breatheH`` set.
  - Otherwise, start the `fullnode` by fast sync mode from `latestH`
  
  The `fullnode` will be stopped immediately once it catches up the `queryH`. And then the tool will analyze the accounts in DB, and export the result to the User-specified directory. 
  
- If `queryH` <= `latestH`, then the tool will query the related data in DB first. If the data exists, export the result.If not, it will delete the existing `fullnode` data that has been synchronized, and restart by state sync mode with `stateSyncHeight = breatheH` set.  

## Notice

- Data in DB can only be queried when the `fullnode` is terminated.
- If a folder serves as a home directory of a `fullnode` that you started ever, then you should be careful to use it as your home dir for this executive tool, since the historic block data could be removed by this tool.
- If user has launched a `fullnode` that is keeping synced with the Block Chain. You can do a quick search by using the `BNCHOME` as the home dir of this tool.The premise is to stop the whole node for a moment.