package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	uuid "github.com/hashicorp/go-uuid"
	"github.com/hashicorp/vault/helper/keysutil"
	"github.com/julienschmidt/httprouter"
)

type EncryptRequest struct {
	Plaintext           string `json:"plaintext"`
	Context             string `json:"context"`
	Nonce               string `json:"nonce"`
	Type                string `json:"type"`
	ConvergentEncyption bool   `json:"convergent_encryption"`
}

type EncryptResponse struct {
	Data EncryptResponseData `json:"data"`
}

type EncryptResponseData struct {
	Ciphertext string `json:"ciphertext"`
}

func EncryptEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var req EncryptRequest
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Plaintext == "" {
		http.Error(w, "missing plaintext to encrypt", http.StatusBadRequest)
		return
	}

	var context []byte
	if req.Context != "" {
		context, err = base64.StdEncoding.DecodeString(req.Context)
		if err != nil {
			http.Error(w, "failed to base64-decode context", http.StatusBadRequest)
			return
		}
	}

	var nonce []byte
	if req.Nonce != "" {
		nonce, err = base64.StdEncoding.DecodeString(req.Nonce)
		if err != nil {
			http.Error(w, "failed to base64-decode nonce", http.StatusBadRequest)
			return
		}
	}

	// Generate a 256bit key
	newKey, err := uuid.GenerateRandomBytes(32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "empty ciphertext returned", http.StatusInternalServerError)
		return
	}

	rsp := EncryptResponse{
		Data: EncryptResponseData{
			Ciphertext: ciphertext,
		},
	}
	json.NewEncoder(w).Encode(rsp)
}
