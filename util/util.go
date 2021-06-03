package util

import (
	"io/ioutil"
	"os"

	"github.com/TheLazarusNetwork/LazarusTunnel/model"
	log "github.com/sirupsen/logrus"
)

// Version Build Version
var Version = "1.0"

// StandardFields for logger
var StandardFields = log.Fields{
	"hostname": "HostServer",
	"appname":  "TunnelAPI",
}

// ReadFile file content
func ReadFile(path string) (bytes []byte, err error) {
	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// WriteFile content to file
func WriteFile(path string, bytes []byte) (err error) {
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// FileExists check if file exists
func FileExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CheckError for checking any errors
func CheckError(message string, err error) {
	if err != nil {
		log.WithFields(StandardFields).Fatalf("%s %+v", message, err)
	}
}

// LogErrors for checking any errors
func LogError(message string, err error) {
	if err != nil {
		log.WithFields(StandardFields).Warnf("%s %+v", message, err)
	}
}

// Message Return Response as map
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// MessageByte Return Response as byte array
func MessageTunnel(status int, message model.Tunnel) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// MessageByte Return Response as byte array
func MessageTunnels(status int, message []model.Tunnel) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
