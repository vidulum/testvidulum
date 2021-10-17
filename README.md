# testvidulum

**testvidulum** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
Download release
copy genesis.json to .testvidulum/config/
testvidulumd start --p2p.persistent_peers e7ef78bb156f04f667e4a23a0782e4b1bb673165@216.128.150.25:26656,b9361329891f1acda1f93e55f73642736759e5bb@66.42.124.230:26656
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
