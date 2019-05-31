# Simple Go Ethereum - Smart Contract

Interact to smart contract on Ropsten Test Network Ethereum using Golang

Documents: https://goethereumbook.org/en/smart-contracts/

If you're interested in NodeJS connecting to Ethereum, please reference my simple-node-ethereum here: https://github.com/huynhsamha/simple-node-ethereum

## Set up accounts

### Create accounts and deposit ether with MetaMask
+ Install extension MetaMask in Chrome browser.
+ Choose Ropsten Test Network.
+ Create accounts and deposit Ether from Test Faucet
+ Get private keys from accounts (MetaMask requires your password)

### Create an account with Infura
Create account on https://infura.io


## Set up GOPATH
Clone and add this repository path to your GOPATH. 

Example, in `~/.bashrc` or `~/.zshrc`, add the following lines:

```bash
export GOPATH=$HOME/go

export GOPATH=$GOPATH:$HOME/Documents/simple-go-ethereum

export PATH=$PATH:$GOPATH/bin
```


## Install go-ethereum

```bash
go get github.com/ethereum/go-ethereum
```

## Install smart contract compiler

https://goethereumbook.org/en/smart-contract-compile/

### Install solidity compiler (`solc`)

Compile `.sol` file to `.abi` and `.bin` file

```bash
solc --abi Store.sol -o build
solc --bin Store.sol -o build
```

### Install `abigen` tool (`abigen`)

Compile `.bin` and `.abi` files to `.go` file

```bash

Install bla bla ...

# if get error: abigen command not found
sudo ln -s $HOME/go/bin/abigen /usr/local/bin/abigen

# generate go file
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go
```


## Interact to smart contract by Golang

Deploy, load, read and write transaction to smart contract using Go.

https://goethereumbook.org/en/smart-contracts/


### Ropsten Etherscan Urls:

Public address: https://ropsten.etherscan.io/address/0x9338063cd5c8f069334925345fadc94a8e45742f

Smart contract address: https://ropsten.etherscan.io/address/0x71733bc1d86F5de7bCBf0587502bB1BeE176E788

Transaction deploying contract: https://ropsten.etherscan.io/tx/0xa3a15b34944525039873f8a8c30d99802690a5f78ddb3609136d9fe1eae0e2b9

Transaction written to contract: https://ropsten.etherscan.io/tx/0x2366ee7e67753a3f60dfd5917dd0a79c3f56a953922d804b402e4404bd4fbc48
