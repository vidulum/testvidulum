# testvidulum

**testvidulum** is a blockchain built using Cosmos SDK and Tendermint.

##  Quick Start

```
Download release
mv testvidulumd /usr/local/bin/testvidulumd
copy genesis.json to .testvidulum/config/genesis.json
testvidulumd start --p2p.persistent_peers e7ef78bb156f04f667e4a23a0782e4b1bb673165@216.128.150.25:26656,b9361329891f1acda1f93e55f73642736759e5bb@66.42.124.230:26656

1,000,000 uTVDL == 1 TVDL

```

- [Testnet Explorer](https://testnet-explorer.vidulum.app)


### Notes - Work in Progress
```
    - * Ubuntu 21.04
    - wget https://github.com/vidulum/testvidulum/releases/download/v1.0/testvidulum_linux_amd64.tar.gz
    - tar -xf testvidulum_linux_amd64.tar.gz && mv testvidulumd /usr/local/bin/
    - mkdir .testvidulum && mkdir .testvidulum/config && cd .testvidulum/config
    - wget -P .testvidulum/config/ https://github.com/vidulum/testvidulum/releases/download/v1.0/genesis.json
    - tmux new -s testnet
    - testvidulumd start
    - * Allow node to fully sync
    - * Press control+b then d to close the mux and return to previous session
    - * Use tmux attach -t testnet to reconnect to the mux if needed
    - testvidulumd keys add EXAMPLE
    - * Enter a keyring passphrase that is long and unique
    - * Re-enter the keyring passphrase
    - * Save the keyring passphrase in a secure location
    - * Open https://vidulum.app/testnet-faucet
    - * Enter your new wallet address to receive 10 TVDL
    - testvidulumd q bank balances ADDRESS
    - testvidulumd q staking validators
    - testvidulumd tx staking delegate OPERATOR_ADDRESS 1000000utvdl --from EXAMPLE --chain-id testvidulum-1
    - * Type y then press Enter to submit delegation transaction
    - testvidulumd q distribution rewards ADDRESS OPERATOR_ADDRESS
    - * Example Result:
    - * rewards:
    - - amount: "6649.976987315010000000"
    -   denom: utvdl
    - testvidulumd tx distribution withdraw-all-rewards --from EXAMPLE --chain-id testvidulum-1
    - * Type y then press Enter to submit rewards withdraw transaction
    - testvidulumd q bank balances ADDRESS

```

### Create a validator
```
testvidulumd tx staking create-validator \
--amount "8000000utvdl" \
--pubkey $(testvidulumd tendermint show-validator) \
--moniker CoreysValidator \
--chain-id testvidulum-1 \
--commission-rate "0.10" \
--commission-max-rate "0.20" \
--commission-max-change-rate "0.01" \
--min-self-delegation "5" \
--from EXAMPLE \
--keyring-backend os \
--gas auto \
--gas-adjustment "1.5" \
--gas-prices "0.025utvdl" \
--broadcast-mode block
```

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/vidulum/testvidulum@latest! | sudo bash
```

`vidulum/testvidulum` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
