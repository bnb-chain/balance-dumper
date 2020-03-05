Binance Airdrop
-------------
Binance Airdrop is a tool designed to provide users with the simplest way to distribute tokens to multiple accounts.

### Usage
Note that we have two binary files to distinguish testNet or production network. `bairdrop` is used for production and `tbairdrop` for testNet. 

Now, let us do a demo with `tbairdrop`. First of all, place the binary file in the specified folder, enter this folder, and execute it to check the usage of it
```
$ tbairdrop
Testnet BlockChain Airdrop

Usage:
  tbairdrop [command]

Available Commands:
  keys        Add or view local private keys
  run         Run airdrop
  help        Help about any command

Flags:
  -h, --help            help for bairdrop
      --home string     directory for config and data (default "/Users/fletcher/.tbairdrop")
  -o, --output string   Output format (text|json) (default "text")

Use "tbairdrop [command] --help" for more information about a command.
```
You can specify a home directory for this operation by `--home`. Otherwise, `${HOME}/.bairdrop` will serve as the default `home`.
### keys
Before distributing, you have to configure the source account for this distribution . And command `keys` help you with that.

- Add account
  ```
  Usage:
    tbairdrop keys add <name> [flags]
  ```
  example
  ```
  $ tbairdrop keys add fromAcc --recover 
  Enter a passphrase for your key:
  Repeat the passphrase:
  > Enter your recovery seed phrase:
  {***your seed phrase***}
  NAME:	TYPE:	ADDRESS:						PUBKEY:
  fromAcc	local	tbnb1m38ds8d69kwd8a4uaz5fm3hmvh94wk5gfeszxn	bnbp1addwnpepqdqlls9gxnqujgkdpty6nluxtc6cuurqe7fhe8jmp87exwzq7s5vkfxcvxk
  ```
- List accounts
  ```
  $ tbairdrop keys list
  NAME:	TYPE:	ADDRESS:						PUBKEY:
  fromAcc	local	tbnb1m38ds8d69kwd8a4uaz5fm3hmvh94wk5gfeszxn	bnbp1addwnpepqdqlls9gxnqujgkdpty6nluxtc6cuurqe7fhe8jmp87exwzq7s5vkfxcvxk
  ```
Moreover, you can check other usages of the command `keys`

### File Prepare
After initializing the account, we need to prepare the CSV file of transfer details, including target accounts and the amount for the individual. It has its specific format which looks like the below:

![transfer_details.csv](/manual/transfer_example.png?raw=true "example")

From left to right are `target address`,`amount`,`asset`. For `amount`, we take the last 8 bits as the decimal place. For instance,`addr1,100000000,BNB` means transfer addr1 1BNB.

Save it as CSV format

![save](/manual/transfer_save.png?raw=true "save")


### Run
Then we can start distributing with the essential command `run`.
```
$ tbairdrop run --help
Run airdrop

Usage:
  tbairdrop run [flags]

Flags:
      --chain-id string   Chain ID of tendermint node
  -f, --file string       File of Transfer details
      --from string       Name or address of private key with which to sign
  -h, --help              help for run
      --memo string       Memo to send along with transaction
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --source int        Source of tx
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
      --home string     directory for config and data (default "/Users/fletcher/.tbairdrop")
  -o, --output string   Output format (text|json) (default "text")
```

You can execute the `airdrop` with the transfer file specified.
```
$ tbairdrop run --file transfer_details.csv --chain-id Binance-Chain-Nile  --node data-seed-pre-0-s1.binance.org:80 --from fromAcc
```

Following is the log of this execution

```
1.  ==>Start to run with file: transfer_details.csv
2.  ==>Start batch 1(from tbnb1rtzy6szuyzcj4amfn6uarvne8a5epxrdc28nhr to tbnb1tyrc4usqp52ne60y2qnta4jk997e79tzcmvlcm)
3.  Password to sign with 'fromAcc':
4.  ==>Transaction hash: BDB452AE09AB9961FD77109DB6DB36559C64415D37F7CEDEF7027FEC43D1130B, sending...
5.  ==>Sending completed, committed at block 69726953 (tx hash: BDB452AE09AB9961FD77109DB6DB36559C64415D37F7CEDEF7027FEC43D1130B)

6.  ==>Start batch 2(from tbnb1vze2xyajsl3dpumkewuz0jdschstnyr6wtyctz to tbnb15xemc2fk9cvxewa87d0js3ypq7aawrku24l7px)
7.  ==>Transaction hash: C66DFFCF6DE1146468CE0947AB5E554E7FF1C87942969F20167692228F8B37A1, sending...
8.  ==>Sending completed, committed at block 69726961 (tx hash: C66DFFCF6DE1146468CE0947AB5E554E7FF1C87942969F20167692228F8B37A1)

9.  ==>Start batch 3(from tbnb17jql0796cjuxzxwmjfdm779gd2306lpexemx27 to tbnb1qgtdt7y062mk66vgu37e0lgamngwgj5asce74f)
10. ==>Transaction hash: F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725, sending...
11. ==>Sending completed, committed at block 69726969 (tx hash: F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725)

12. ==>Start batch 4(from tbnb1zgczyzlxyk34rwqdzrpxez35k52qy7uyl7lp8q to tbnb1zgczyzlxyk34rwqdzrpxez35k52qy7uyl7lp8q)
13. ==>Transaction hash: DD1C8F9B6A7C0F526CFA41A6D221D28AC26DD196D5AC109B7F72655B0E726E63, sending...
14. ==>Sending completed, committed at block 69726979 (tx hash: DD1C8F9B6A7C0F526CFA41A6D221D28AC26DD196D5AC109B7F72655B0E726E63)`
```

Like the log shows above, this tool processes the transfer in batches if the number of target accounts exceeds 300 for performance concerns. It treats each 300 rows as a batch, with an interval of 2s between each batch.

### Error Handle
It is worth noting that if for any reason it fails during a batch execution, you need to double confirm that whether this transfer is on chain based on tx hash the console prints. And you should remove the records of the successful transfer from the CSV to prevent repeated transfer, and execute it again.
 
For example, if it crashes after the console prints line:
+ 8/9：meaning batch1 and batch2 are transferred successfully. We should remove the related rows from the CSV, and re-run this file with the command.
+ 10: meaning batch1 and batch2 are transferred successfully. But it is not sure if batch3 is successfully executed, we should take a look with the transaction hash *F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725* offline. For instance, you can check it on the explorer `https://testnet-explorer.binance.org/tx/F474F3ABF5E770EE1291821C07234236403028767BD912011F5CAB42228D8725`. 




