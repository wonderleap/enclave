package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	shell "github.com/ipfs/go-ipfs-api"
)

var (
	ipfsHost = "localhost"
	ipfsPort = "5001"
	password = []byte("hunter2")
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getIPFSURL() string {
	return strings.Join([]string{ipfsHost, ipfsPort}, ":")
}

func encryptData(s string) string {
	armor, err := helper.EncryptMessageWithPassword(password, s)
	check(err)
	return armor
}

func decryptData(armor string) string {
	message, err := helper.DecryptMessageWithPassword(password, armor)
	check(err)
	return message
}

func addToIPFS(sh *shell.Shell, armor string) string {
	cid, err := sh.Add(strings.NewReader(armor))
	check(err)
	return cid
}

func catFromIPFS(sh *shell.Shell, hash string) io.ReadCloser {
	cat, err := sh.Cat(hash)
	check(err)
	return cat
}

func getFromIPFS(sh *shell.Shell, hash string) {
	err := sh.Get(hash, "file.tmp")
	check(err)
}

func main() {
	sh := shell.NewShell(getIPFSURL())
	fmt.Printf("daemon is up: %s\n", strconv.FormatBool(sh.IsUp()))

	// encrypt and add to IPFS
	armor := encryptData("test")
	cid := addToIPFS(sh, armor)

	// cat from IPFS
	reader := catFromIPFS(sh, cid)
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	// decrypt data
	text := decryptData(buf.String())
	fmt.Println(text)

	// fmt.Printf("added %s", cid)
	// fmt.Println(decryptData(armor))
}
