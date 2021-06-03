package core

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/util"
)

//Init initializes json file for caddy and nginx
func Init() {
	//caddy.json path
	path := filepath.Join(os.Getenv("APP_CONF_DIR"), "caddy.json")

	//check if exists
	if !util.FileExists(path) {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		util.CheckError("caddy.json error: ", err)

		_, err = file.Write([]byte("[]"))
		util.CheckError("caddy.json error: ", err)
	}

	//nginx.json path
	path = filepath.Join(os.Getenv("APP_CONF_DIR"), "nginx.json")

	//check if exists
	if !util.FileExists(path) {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		util.CheckError("nginx.json error: ", err)

		_, err = file.Write([]byte("[]"))
		util.CheckError("nginx.json error: ", err)
	}
}

// Writefile appends data to file
func Writefile(path string, bytes []byte) (err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		util.LogError("File Open error: ", err)
		return err
	}

	defer file.Close()

	_, err = file.WriteString(string(bytes))
	if err != nil {
		util.LogError("File write error: ", err)
		return err
	}

	return nil
}

//ScanPort checks avilability of port
func ScanPort(port int) (string, error) {
	ip := os.Getenv("SERVER")
	timer := 500 * time.Millisecond

	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timer)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timer)
			ScanPort(port)
		} else {
			return "inactive", nil
		}
		return "", err
	}

	conn.Close()
	return "active", nil
}

//GetPort returns available port based on random generation
func GetPort(max, min int) (int, error) {
	port := rand.Intn(max-min) + min

	status, err := ScanPort(port)
	if err != nil {
		util.LogError("Scan Port error: ", err)
		return -1, err
	}

	if status == "inactive" {
		return port, nil
	} else if status == "active" {
		GetPort(max, min)
	}

	return -1, nil
}
