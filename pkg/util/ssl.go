/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package util

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
	"github.com/arpabet/sprint/pkg/app"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
	"github.com/arpabet/preferences"
)


const (

	MAX_SSL_IDLE_CONNS = 1024

	SSL_PREFIX = "ssl:"

	SERVER_CRT = "server.crt"
	SERVER_KEY = "server.key"
	CA_CRT = "ca.crt"

	CLIENT_CRT = "client.crt"
	CLIENT_KEY = "client.key"

)

var (

	ERR_CA_DECODE = errors.New("Can't decode client certificate authority")
	ERR_CA_PARSE = errors.New("Can't parse client certificate authority")
	ERR_CA_NO_SIGN = errors.New("Can't find signature of client certificate authority")

)

func ImportCertificates(storage app.Storage, sslDir string) error {

	for _, certName := range []string{CA_CRT, SERVER_CRT, SERVER_KEY} {
		if content, err := ioutil.ReadFile(filepath.Join(sslDir, certName)); err != nil {
			return err
		} else if err := storage.Put([]byte(SSL_PREFIX + certName), []byte(content)); err != nil {
			return err
		}
	}

	appDir, err := preferences.MakeAppDataDir(app.ApplicationName)
	if err != nil {
		return err
	}

	for _, certName := range []string{CLIENT_CRT, CLIENT_KEY} {
		srcFile := filepath.Join(sslDir, certName)
		dstFile := filepath.Join(appDir, certName)
		if err = CopyFile(srcFile, dstFile, 0644); err != nil {
			return errors.Errorf("copy file error from %s to %s, %v", srcFile, dstFile, err)
		}
	}

	return nil

}

func PromptCertificates(storage app.Storage) error {

	content := PromptPassword("Enter ca.crt content:")
	if err := storage.Put([]byte(SSL_PREFIX + CA_CRT), []byte(content)); err != nil {
		return err
	}

	content = PromptPassword("Enter server.crt content:")
	if err := storage.Put([]byte(SSL_PREFIX + SERVER_CRT), []byte(content)); err != nil {
		return err
	}

	content = PromptPassword("Enter server.key content:")
	if err := storage.Put([]byte(SSL_PREFIX + SERVER_KEY), []byte(content)); err != nil {
		return err
	}

	return nil

}


func CAKeyId(storage app.Storage) (string, error) {

	caPEMBlock, err := storage.Get([]byte(SSL_PREFIX + CA_CRT), true)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(caPEMBlock)
	if block == nil {
		return "", ERR_CA_DECODE
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	signature := cert.Signature
	if signature == nil {
		return "", ERR_CA_NO_SIGN
	}

	return app.Encoding.EncodeToString(signature)[:8], nil

}


func LoadServerConfig(storage app.Storage) (*tls.Config, error) {

	certPEMBlock, err := storage.Get([]byte(SSL_PREFIX + SERVER_CRT), true)
	if err != nil {
		return nil, err
	}

	keyPEMBlock, err := storage.Get([]byte(SSL_PREFIX + SERVER_KEY), true)
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	certpool := x509.NewCertPool()
	caPEMBlock, err := storage.Get([]byte(SSL_PREFIX + CA_CRT), true)
	if err != nil {
		return nil, err
	}
	if !certpool.AppendCertsFromPEM(caPEMBlock) {
		return nil, ERR_CA_PARSE
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certpool,
		Rand: 		  rand.Reader,
	}, nil

}

func LoadClientConfig() (*tls.Config, error) {

	appDir := preferences.AppDataDir(app.ApplicationName)

	certFile := filepath.Join(appDir, CLIENT_CRT)
	keyFile := filepath.Join(appDir, CLIENT_KEY)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, errors.Errorf("LoadX509KeyPair for client SSL from %s and %s failed, %v", certFile, keyFile, err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}, nil

}

func NewClient() (*http.Client, error) {

	config, err := LoadClientConfig()
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: config,
		MaxIdleConnsPerHost: MAX_SSL_IDLE_CONNS,
		TLSHandshakeTimeout: 0 * time.Second,
	}

	return &http.Client{Transport: tr}, nil

}

func ReadAll(resp *http.Response, err error) (string, error) {

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return "", err
	}

	if resp.StatusCode == 200 {
		return string(content), nil
	} else {
		return "", errors.New(string(content))
	}

}

func PostResp(resp *http.Response) (string, error) {

	if resp.StatusCode != 200 {
		return "", errors.Errorf("invalid response code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(content), nil
}