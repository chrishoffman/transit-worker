# Transit Worker

Transit Worker is a proof-of-concept for building a horizontally scalable service for doing encryption and decryption operations that are compatible to Hashicorp's Vault transit backend.

## Status

This is a prototype.  Do not use in production.

## Limitations

Vault does not currently provide a way to read the keys it generates from the transit backend.  In order for this model to work, we would need to have access to Vault's encrption keys.  I am also working a modified version of Vault to allow for export of encryption keys.
