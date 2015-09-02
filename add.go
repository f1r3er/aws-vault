package main

import (
	"github.com/99designs/aws-vault/keyring"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/bgentry/speakeasy"
)

type AddCommandInput struct {
	Profile string
	Keyring keyring.Keyring
}

func AddCommand(ui Ui, input AddCommandInput) {
	accessKeyId, err := prompt("Enter Access Key ID: ")
	if err != nil {
		ui.Error.Fatal(err)
	}

	secretKey, err := speakeasy.Ask("Enter Secret Access Key: ")
	if err != nil {
		ui.Error.Fatal(err)
	}

	creds := credentials.Value{AccessKeyID: accessKeyId, SecretAccessKey: secretKey}
	provider := &KeyringProvider{Keyring: input.Keyring, Profile: input.Profile}

	if err := provider.Store(creds); err != nil {
		ui.Error.Fatal(err)
	}

	ui.Printf("Added credentials to profile %q in vault", input.Profile)
}