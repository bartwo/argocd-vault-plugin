package backends

import (
	"encoding/json"
	"fmt"
	"log"

	thycoticsecretserver "github.com/thycotic/tss-sdk-go/server"
)

// NewThycoticSecretServerBackend initializes a new Thycotic Secret Server backend
func NewThycoticSecretServerBackend(client IBMSecretsManagerClient) *IBMSecretsManager {
	ibmSecretsManager := &IBMSecretsManager{
		Client: client,
	}
	return ibmSecretsManager
}

func main() {
	// fmt.Println("vim-go")
	tss, _ := thycoticsecretserver.New(thycoticsecretserver.Configuration{
		Credentials: thycoticsecretserver.UserCredential{
			Username: "supersecretusername",
			Password: "supersecretpassword",
		},
		ServerURL: "https://supersecretserver/SecretServer",
	})

	if err != nil {
		log.Fatal("Error calling thycoticsecretserver.Secret", err)
	}

	if pw, ok := s.Field("password"); ok {
		fmt.Println("the password is", pw)
		es, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(string(es))
	}
}
