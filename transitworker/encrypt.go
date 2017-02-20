package transitworker

import (
	"errors"
	"fmt"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/hashicorp/vault/helper/keysutil"
)

func Encrypt() error {

	// TEMP Data
	var context, nonce []byte
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
		Keys:          map[int]keysutil.KeyEntry{1: entry},
	}

	ciphertext, err := policy.Encrypt(context, nonce, "temp")
	if err != nil {
		return err
	}

	if ciphertext == "" {
		return errors.New("empty ciphertext returned")
	}

	fmt.Println(ciphertext)
	return nil
}
