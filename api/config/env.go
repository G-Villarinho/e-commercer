package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var Env Environment

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	if _, err := env.UnmarshalFromEnviron(&Env); err != nil {
		return fmt.Errorf("init env: %w", err)
	}

	if Env.Key.PrivateKey == "" || Env.Key.PublicKey == "" {
		privateKey, err := LoadKeyFromFile("ecdsa_private.pem")
		if err != nil {
			return fmt.Errorf("load private key: %w", err)
		}

		publicKey, err := LoadKeyFromFile("ecdsa_public.pem")
		if err != nil {
			return fmt.Errorf("load public key: %w", err)
		}

		Env.Key.PrivateKey = privateKey
		Env.Key.PublicKey = publicKey
	}

	return nil
}

func LoadKeyFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return strings.TrimSpace(string(data)), nil
}
