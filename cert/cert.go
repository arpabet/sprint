/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package cert

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/certpb"
	"go.arpabet.com/store"
	"golang.org/x/crypto/acme/autocert"
	"net"
	"reflect"
	"time"
)


var CertificateManagerClass = reflect.TypeOf((*CertificateManager)(nil)).Elem()

type CertificateManager interface {
	glue.InitializingBean
	glue.DisposableBean

	GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error)

	InvalidateCache(domain string)

	ListActive() map[string]error

	ListRenewal() map[string]time.Time

	ListUnknown() map[string]time.Time

	ExecuteCommand(cmd string, args []string) (string, error)

}

var CertificateRepositoryClass = reflect.TypeOf((*CertificateRepository)(nil)).Elem()

type CertificateRepository interface {
	glue.DisposableBean

	/**
	Self Signer CRUID
	*/
	SaveSelfSigner(self *certpb.SelfSigner) error

	FindSelfSigner(name string) (*certpb.SelfSigner, error)

	ListSelfSigners(prefix string, cb func(*certpb.SelfSigner) bool) error

	DeleteSelfSigner(name string) error

	/**
	Acme Account CRUID
	*/
	SaveAccount(account *certpb.AcmeAccount) error

	FindAccount(email string) (*certpb.AcmeAccount, error)

	ListAccounts(prefix string, cb func(*certpb.AcmeAccount) bool) error

	DeleteAccount(email string) error

	/**
	Domain zone CRUID
	*/

	SaveZone(zone *certpb.Zone) error

	FindZone(zone string) (*certpb.Zone, error)

	ListZones(prefix string, cb func(*certpb.Zone) bool) error

	DeleteZone(zone string) error

	/**
	Watch zone changes
	*/

	Watch(ctx context.Context, cb func(zone, event string) bool) (cancel context.CancelFunc, err error)

	/**
	Gets backend using for storing certificates
	*/

	Backend() store.DataStore

	/**
	Sets backend using for storing certificates
	*/
	SetBackend(storage store.DataStore)
}

var CertificateServiceClass = reflect.TypeOf((*CertificateService)(nil)).Elem()

type AcmeAccount struct {
	Status string `json:"status,omitempty"`
	Contact []string `json:"contact,omitempty"`
	TermsOfServiceAgreed bool `json:"termsOfServiceAgreed,omitempty"`
	Orders string `json:"orders,omitempty"`
	OnlyReturnExisting bool `json:"onlyReturnExisting,omitempty"`
	ExternalAccountBinding json.RawMessage `json:"externalAccountBinding,omitempty"`
}

type AcmeResource struct {
	Body  AcmeAccount `json:"body,omitempty"`
	URI   string       `json:"uri,omitempty"`
}

type AcmeUser struct {
	Email         string
	Registration  *AcmeResource
	PrivateKey    crypto.PrivateKey
}

type CertificateService interface {
	glue.InitializingBean

	CreateAcmeAccount(email string) error

	GetOrCreateAcmeUser(email string) (user *AcmeUser, err error)

	CreateSelfSigner(cn string, withInter bool) error

	RenewCertificate(zone string) error

	ExecuteCommand(cmd string, args []string) (string, error)

	IssueAcmeCertificate(entry *certpb.Zone) (string, error)

	IssueSelfSignedCertificate(entry *certpb.Zone) error

}

var AutocertStorageClass = reflect.TypeOf((*AutocertStorage)(nil)).Elem()

type AutocertStorage interface {

	Cache(serverName string) autocert.Cache

}

var CertificateIssuerClass = reflect.TypeOf((*CertificateIssuer)(nil)).Elem()
var CertificateIssuerServiceClass = reflect.TypeOf((*CertificateIssueService)(nil)).Elem()

type CertificateDesc struct {
	Organization string
	Country      string
	Province     string
	City         string
	Street       string
	Zip          string
}

type IssuedCertificate interface {

	KeyFileContents() []byte

	CertFileContents() []byte

	PrivateKey() crypto.Signer

	Certificate() *x509.Certificate

}

type CertificateIssuer interface {

	Parent() (CertificateIssuer, bool)

	Certificate() IssuedCertificate

	IssueInterCert(cn string) (CertificateIssuer, error)

	IssueClientCert(cn string, password string) (cert IssuedCertificate, pfxData []byte, err error)

	IssueServerCert(cn string, domains []string, ipAddresses []net.IP) (IssuedCertificate, error)

}

type CertificateIssueService interface {

	LoadCertificateDesc() (*CertificateDesc, error)

	CreateIssuer(cn string, info *CertificateDesc) (CertificateIssuer, error)

	LoadIssuer(*certpb.SelfSigner) (CertificateIssuer, error)

	LocalIPAddresses(addLocalhost bool) ([]net.IP, error)

}

