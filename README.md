# `enclave`

`enclave` is a p2p data sharing tool built on top of IPFS and OpenPGP. The tool aims to make it easier to encrypt and distribute private data through the Internet. We believe that humans should be able to:

- own their data
- have portability of their data
- know when their data is being accessed
- request that their data become inaccessible at any time

This library does not aim to accomplish all of the above, but take small steps.

## Getting started

You first need to go and get IPFS, then run the daemon.

```
ipfs daemon
```

Then:

```
go install .
enclave
```

## How it works

### Authentication and encryption

1. The data owner first creates a password, which becomes the passphrase of their private key generated with their public key.
2. The data owner uses their public key to encrypt their data before uploading this to IPFS.
3. After uploading the data to IPFS, data can be decrypted using the user's private key and password.

### Public key storage

The public key can be stored in one of two ways:

- In a smart contract, which stores a key-value pair of the intended recipient's Ethereum wallet address, mapped to the hash of the file in IPFS that contains the message recipient's public key.
- In a centralized key-value DB, which stores all data owners' Ethereum wallets mapped to the data owner's public key.

The benefit of storing the public key in a smart contract is that it allows for public, immutable auditability of the data accesses. However, it will cost gas to be able to trigger the lookup on the smart contract and to change or modify the key-value store.

The benefit of the centralized key-value DB is that the lookup and all changes do not cost anything to perform. However it's centralized and may become a point of failure.

### Private key storage

No private keys should ever be in the public, whether that's on the blockchain or in IPFS.
