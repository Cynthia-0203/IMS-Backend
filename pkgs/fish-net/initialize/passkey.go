package initialize

import (
	"IMS-Backend/pkgs/fish-net/glb"

	"os"

	"github.com/go-webauthn/webauthn/webauthn"
)

func initAuth() {
	Auth, err := webauthn.New(&webauthn.Config{
		RPDisplayName: "Go Webauthn",                                              // Display Name for your site
		RPID:          "localhost",                                                // Generally the FQDN for your site
		RPOrigins:     []string{"http://localhost:3000", "http://localhost:3001"}, // The origin URLs allowed for WebAuthn requests
		RPIcon:        "https://go-webauthn.local/logo.png",                       // Optional icon URL for your site
	})
	if err != nil {
		glb.LOG.Error(err.Error())
		os.Exit(0)
	}
	glb.Auth = Auth
}
