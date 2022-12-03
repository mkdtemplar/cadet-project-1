package saml_handler

import (
	"cadet-project/configurations"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"log"
	"net/http"
	"net/url"
)

func AuthorizationRequest() *samlsp.Middleware {

	keyPair, err := tls.LoadX509KeyPair(configurations.Config.Crt, configurations.Config.Key)
	if err != nil {
		log.Println(err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Println(err)

	}
	microsoftURL := fmt.Sprintf(configurations.Config.MSUrl, configurations.Config.TenantID)

	idpMetadataURL, err := url.Parse(microsoftURL)
	if err != nil {
		log.Fatalln(err)

	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Println(err)

	}

	rootURL, err := url.Parse(configurations.Config.RootUrl)
	if err != nil {
		log.Println(err)

	}

	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:    "spn:" + configurations.Config.AppId,
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	return samlSP
}
