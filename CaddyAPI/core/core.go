package core

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/util"
)

func Init() {
	path := filepath.Join(os.Getenv("CADDY_CONF_DIR"), "server.json")

	if !util.FileExists(path) {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		util.CheckError("Server.json error: ", err)

		_, err = file.Write([]byte("[]"))
		util.CheckError("Server.json error: ", err)
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

func GetPort() (int, error) {
	max, err := strconv.Atoi(os.Getenv("UpperRange"))
	if err != nil {
		util.LogError("String Conversion error: ", err)
		return -1, err
	}

	min, err := strconv.Atoi(os.Getenv("LowerRange"))
	if err != nil {
		util.LogError("String Conversion error: ", err)
		return -1, err
	}

	port := rand.Intn(max-min) + min

	status, err := ScanPort(port)
	if err != nil {
		util.LogError("Scan Port error: ", err)
		return -1, err
	}

	if status == "inactive" {
		return port, nil
	} else if status == "active" {
		GetPort()
	}

	return -1, nil
}
