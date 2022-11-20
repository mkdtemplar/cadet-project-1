package saml_handler

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/IvanMarkovskiSF/cadet-project/configurations"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func SamlRequest() *samlsp.Middleware {
	config, err := configurations.LoadConfig(".configurations")
	if err != nil {
		log.Fatalln("cannot load configurations")
	}

	keyPair, err := tls.LoadX509KeyPair("saml_handler/myservice.crt", "saml_handler/myservice.key")
	if err != nil {
		log.Println(err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Println(err)

	}
	microsoftURL := fmt.Sprintf("https://login.microsoftonline.com/%s/federationmetadata/2007-06/federationmetadata.xml",
		config.TenantID)

	idpMetadataURL, err := url.Parse(microsoftURL)
	if err != nil {
		log.Println(err)

	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Println(err)

	}

	rootURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Println(err)

	}

	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:    "spn:" + config.AppId,
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	return samlSP
}
