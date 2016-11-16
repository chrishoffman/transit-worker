# Transit Worker

Transit Worker is a proof-of-concept for building a horizontally scalable service for doing encryption and decryption operations that are compatible to Hashicorp's Vault transit backend.

Most workloads are pretty predictable in Vault.  The one that does not follow a predictable workload is the transit backend.  Since you could receive payloads of varying sizes and frequency, this makes it hard to scale Vault for these types of workloads.  By moving to an agent that only deals with the crypto operations, you get to easily horizontally scale transit.

## Status

This is a prototype.  Do not use in production.

## Limitations

Vault does not currently provide a way to read the keys it generates from the transit backend.  In order for this model to work, we would need to have access to Vault's encrption key ring.  A modified version of Vault is being developed to allow for access to the encryption key ring.
