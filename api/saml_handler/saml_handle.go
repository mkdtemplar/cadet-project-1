package saml_handler

import (
	"cadet-project/config"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func AuthorizationRequest() *samlsp.Middleware {

	keyPair, err := tls.LoadX509KeyPair(config.Config.Crt, config.Config.Key)
	if err != nil {
		log.Println(err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Println(err)

	}
	microsoftURL := fmt.Sprintf(config.Config.MSUrl, config.Config.TenantID)

	idpMetadataURL, err := url.Parse(microsoftURL)
	if err != nil {
		log.Fatalln(err)

	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Println(err)

	}

	rootURL, err := url.Parse(config.Config.RootUrl)
	if err != nil {
		log.Println(err)

	}

	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:    "spn:" + config.Config.AppId,
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	return samlSP
}
