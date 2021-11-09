# testvidulum

**testvidulum** is a blockchain built using Cosmos SDK and Tendermint.

###  Steps summary

- Recommended OS - Ubuntu 21.04
- 1 **TVDL** == 1,000,000 **uTVDL**
- Download latest release and unzip it
- Move `testvidulumd` to `/usr/local/bin/`
- Download `genesis.json` into `.testvidulum/config/` (create the folders beforehand)
- Start the daemon (u will need peers the first time u start)

---

- [Testnet Explorer](https://testnet-explorer.vidulum.app)

- [Testnet Faucet](https://vidulum.app/testnet-faucet)

- [Vidulum App](https://wallet.vidulum.app)
---

### Steps (Detailed)
**Note:** All the capitalized words will be replaced with personal values
- Download latest release
```sh
wget https://github.com/vidulum/testvidulum/releases/download/v1.0/testvidulum_linux_amd64.tar.gz
```
- Unzip it and move it to right folder
```sh
tar -xf testvidulum_linux_amd64.tar.gz && mv testvidulumd /usr/local/bin/
```
- Download `genesis.json` file directly in the needed folder
```sh
wget -P .testvidulum/config/ https://github.com/vidulum/testvidulum/releases/download/v1.0/genesis.json
```
- Install `tmux` and start a new screen called `testnet`
```sh
apt install tmux -y && tmux new -s testnet
```
- Inside `testnet` sesion start testvidulum daemon with the following arguments
```sh
testvidulumd start --p2p.persistent_peers e7ef78bb156f04f667e4a23a0782e4b1bb673165@216.128.150.25:26656,b9361329891f1acda1f93e55f73642736759e5bb@66.42.124.230:26656
```
- Allow node to fully sync (wait till the current block is the same like in the explorer)
- Press `Ctrl+B` then `D` to close the mux and return to previous session (Use `tmux attach -t testnet` to reconnect to the mux if needed)
- Create a keypair
```sh
testvidulumd keys add EXAMPLE
```
- * Enter a keyring passphrase that is long and unique
- * Re-enter the keyring passphrase
- * Save the keyring passphrase in a secure location
- * Open https://vidulum.app/testnet-faucet
- * Enter your new wallet address to receive 10 TVDL
- Check address balance
```sh
testvidulumd q bank balances ADDRESS
```
- List existing validators (So you can decide where u want to delegate your coins for staking)
```sh
testvidulumd q staking validators
```
- Delegate 6 **TVDL** to the validator (with chosen OPERATOR_ADDRESS)
```sh
testvidulumd tx staking delegate OPERATOR_ADDRESS 6000000utvdl --from EXAMPLE --chain-id testvidulum-1
```
- * Type `y` then press `Enter` to submit delegation transaction

Following command is to see your rewards for the staking amount after a period of time
```sh
testvidulumd q distribution rewards ADDRESS OPERATOR_ADDRESS
```
- * Example Result:
- * rewards:
- - amount: "6649.976987315010000000"
- - denom: utvdl
Following command is to withdraw your rewards from the validator
```sh
testvidulumd tx distribution withdraw-all-rewards --from EXAMPLE --chain-id testvidulum-1
```
- * Type y then press Enter to submit rewards withdraw transaction
- Check address balance to see the received rewards from staking
```sh
testvidulumd q bank balances ADDRESS
```

### Create a validator
If you want to know what each argument does know that all the arguments that are accepted by `testvidulumd` are documented directly in the daemon
```sh
testvidulumd tx staking create-validator --amount "8000000utvdl" --pubkey $(testvidulumd tendermint show-validator) --moniker CoreysValidator --chain-id testvidulum-1 --commission-rate "0.10" --commission-max-rate "0.20" --commission-max-change-rate "0.01" --min-self-delegation "5" --from EXAMPLE --keyring-backend os --gas auto --gas-adjustment "1.5" --gas-prices "0.025utvdl" --broadcast-mode block
```
### Delegate additional (self-delegation) to your validator
```sh
testvidulumd tx staking delegate OPERATOR_ADDRESS 5000utvdl --from EXAMPLE --chain-id testvidulum-1 --keyring-backend os
```

### Edit validator
```sh
testvidulumd tx staking edit-validator --moniker "Corey's Validator" --details "Corey's testnet validator" --website "https://vidulum.app" --from EXAMPLE --chain-id testvidulum-1 --keyring-backend os
```

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```sh
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```sh
curl https://get.starport.network/vidulum/testvidulum@latest! | sudo bash
```

`vidulum/testvidulum` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Vidulum Discord](https://discord.gg/WQAxH4G) [Cosmos Discord](https://discord.gg/cosmosnetwork)
