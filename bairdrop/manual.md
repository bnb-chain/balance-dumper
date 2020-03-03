Binance Airdrop
-------------
Binance Airdrop is a tool designed to provide users with the simplest way to distribute tokens to multiple accounts.

### Usage
Place the binary file in the specified folder, enter this folder, and execute it to check the usage of it
```
$ bairdrop
BlockChain Airdrop

Usage:
  bairdrop [command]

Available Commands:
  keys        Add or view local private keys
  run         Run airdrop
  help        Help about any command

Flags:
  -h, --help            help for bairdrop
      --home string     directory for config and data (default "/Users/fletcher/.bairdrop")
  -o, --output string   Output format (text|json) (default "text")

Use "bairdrop [command] --help" for more information about a command.
```
You can specify a home directory for this operation by `--home`. Otherwise, `${HOME}/.bairdrop` will serve as the default `home`.
### keys
Before distributing, you have to configure the source account for this distribution . And command `keys` help you with that.

- Add account
  ```
  Usage:
    bairdrop keys add <name> [flags]
  ```
  example
  ```
  $ bairdrop keys add fromAcc --recover 
  Enter a passphrase for your key:
  Repeat the passphrase:
  > Enter your recovery seed phrase:
  {***your seed phrase***}
  NAME:	TYPE:	ADDRESS:						PUBKEY:
  fromAcc	local	tbnb1m38ds8d69kwd8a4uaz5fm3hmvh94wk5gfeszxn	bnbp1addwnpepqdqlls9gxnqujgkdpty6nluxtc6cuurqe7fhe8jmp87exwzq7s5vkfxcvxk
  ```
- List accounts
  ```
  $ bairdrop keys list
  NAME:	TYPE:	ADDRESS:						PUBKEY:
  fromAcc	local	tbnb1m38ds8d69kwd8a4uaz5fm3hmvh94wk5gfeszxn	bnbp1addwnpepqdqlls9gxnqujgkdpty6nluxtc6cuurqe7fhe8jmp87exwzq7s5vkfxcvxk
  ```
Moreover, you can check other usages of the command `keys`

### Run
After initializing the account, we need to prepare the CSV file of transfer details, including target accounts and the amount for the individual. It has its specific format which looks like the below:
```
tbnb1rtzy6szuyzcj4amfn6uarvne8a5epxrdc28nhr,10000000,BNB
tbnb1zj8qnx774kjf5qsra88qf6mzu7fpqtj6m7qwvs,10000000,BNB
tbnb1tyrc4usqp52ne60y2qnta4jk997e79tzcmvlcm,20000000,BNB
tbnb1vze2xyajsl3dpumkewuz0jdschstnyr6wtyctz,60000000,BNB
tbnb1mpvfz9x93f6gparvg532ha7kcm9kc35xprq2fl,70000000,BNB
tbnb15xemc2fk9cvxewa87d0js3ypq7aawrku24l7px,90000000,BNB
tbnb17jql0796cjuxzxwmjfdm779gd2306lpexemx27,10000000,BNB
tbnb16h277la79h06fhxeykd7chzdhysgww0rsada2f,8050000,BNB
tbnb1qgtdt7y062mk66vgu37e0lgamngwgj5asce74f,65000000,BNB
tbnb1zgczyzlxyk34rwqdzrpxez35k52qy7uyl7lp8q,90000000,BNB
......
```
From left to right are `target address`,`amount`,`asset`. For `amount`, we take the last 8 bits as the decimal place. For instance,`addr1,100000000,BNB` means transfer addr1 1BNB.

Then we can start distributing with the essential command `run`.
```
$ bairdrop run --help
Run airdrop

Usage:
  bairdrop run [flags]

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
$ bairdrop run --from fromAcc --file transfer_details.csv &
```

For the performance concerns, this tool processes the transfer in batches if the number of target accounts exceeds 300.It treats each 300 rows as a batch, with an interval of 2s between each batch.

It is worth noting that if for any reason it fails during a batch execution, you need to double confirm that whether this transfer is on chain based on tx hash the console prints. And you should remove the records of the successful transfer from the CSV to prevent repeated transfer, and execute it again.
 




