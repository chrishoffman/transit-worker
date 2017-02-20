package transitworker

import (
	"errors"
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/hashicorp/vault/helper/keysutil"
)

func Encrypt() error {
	// Generate a 256bit key (TEMP)
	newKey, err := uuid.GenerateRandomBytes(32)
	if err != nil {
		return errors.New("unable to generate random key")
	}

	entry := keysutil.KeyEntry{
		AESKey: newKey,
	}

	policy := &keysutil.Policy{
		Type:          keysutil.KeyType_AES256_GCM96,
		LatestVersion: 1,
	}
	policy.Keys[1] = entry

	ciphertext, err := policy.Encrypt(context, nonce, req.Plaintext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ciphertext == "" {
		return errors.New("empty ciphertext returned")
	}

}
