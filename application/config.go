package main

import (
	"git.slygods.com/evoplay/wss-go/pkg/helper"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	keyFileDefault    = "key.pem"
	certFileDefault   = "cert.pem"
	keyEnvName        = "KEY_FILE"
	certEnvName       = "CERT_FILE"
	configFileName    = "config.yml"
	defaultListenAddr = ""
	defaultListenPort = "8080"
	addrEnvName       = "APP_WSS_ADDR"
	portEnvName       = "APP_WSS_PORT"
)

type config struct {
	CertFilePath   string
	KeyFilePath    string
	ListenAddr     string
	ListenPort     string
	ExecutablePath string
}

type ParsedConfig struct {
	listenAddr string `yaml:"listen_addr"`
	listenPort string `yaml:"listen_port"`
}

func NewConfig() (*config, error) {
	execPath, err := helper.GetExecPath()
	if err != nil {
		return nil, err
	}

	// Parse confi.yml
	t, err := parseConfig(execPath)
	if err != nil {
		return nil, err
	}

	conf := new(config)
	conf.ListenAddr, conf.ListenPort = getAddrAndPort(t)

	certFilePath, ok := os.LookupEnv(certEnvName)
	if !ok {
		certFilePath = path.Join(execPath, certFileDefault)
	}

	keyFilePath, ok := os.LookupEnv(keyEnvName)
	if !ok {
		keyFilePath = path.Join(execPath, keyFileDefault)
	}

	conf.CertFilePath = certFilePath
	conf.KeyFilePath = keyFilePath
	conf.ExecutablePath = execPath

	return conf, nil
}

func parseConfig(execPath string) (*ParsedConfig, error) {
	configPath, err := filepath.Abs(execPath + string(filepath.Separator) + configFileName)
	if err != nil {
		return nil, err
	}

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	t := &ParsedConfig{}
	err = yaml.Unmarshal(configContent, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getAddrAndPort(t *ParsedConfig) (string, string) {
	addrFromEnv, addrOk := os.LookupEnv(addrEnvName)
	portFromEnv, portOk := os.LookupEnv(portEnvName)

	listenAddr := strings.TrimSpace(t.listenAddr)
	if listenAddr == "" {
		if addrOk {
			listenAddr = strings.TrimSpace(addrFromEnv)
		} else {
			listenAddr = defaultListenAddr
		}
	}

	listenPort := strings.TrimSpace(t.listenPort)
	if listenPort == "" {
		if portOk {
			listenPort = strings.TrimSpace(portFromEnv)
		} else {
			listenPort = defaultListenPort
		}
	}

	return listenAddr, listenPort
}
