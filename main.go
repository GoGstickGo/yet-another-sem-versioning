package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cVersion, err := readInFile("VERSION")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(cVersion)
	change := "patch"
	newVersion := createNewVersion(cVersion, change)
	fmt.Println(newVersion)
	err = modifyVersionFile(newVersion)
	if err != nil {
		log.Fatal(err)
	}

}

func readInFile(file string) (cVersion string, err error) {
	currentVersion, err := ioutil.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("issue with reading the file")
	}
	if strings.HasPrefix(string(currentVersion), "v") == true {
		return string(currentVersion[1:]), nil
	} else {
		return string(currentVersion[0:]), nil
	}
}
func modifyVersion(a, b, c int, s []string) (mVersion string) {
	mV, _ := strconv.Atoi(s[a])
	mV++
	nV := strconv.Itoa(mV)
	switch a {
	case 0:
		mVersion = nV + "." + s[b] + "." + s[c]
	case 1:
		mVersion = s[b] + "." + nV + "." + s[c]
	case 2:
		mVersion = s[b] + "." + s[c] + "." + nV
	}
	return mVersion
}

func createNewVersion(cVersion, change string) (nVersion string) {
	v := strings.Split(cVersion, ".")
	switch change {
	case "major":
		a, b, c := 0, 1, 2
		nVersion = modifyVersion(a, b, c, v)
	case "minor":
		a, b, c := 1, 0, 2
		nVersion = modifyVersion(a, b, c, v)
	case "patch":
		a, b, c := 2, 0, 1
		nVersion = modifyVersion(a, b, c, v)
	}
	return nVersion
}

func modifyVersionFile(nVersion string) (err error) {
	// remove current file
	err = os.Remove("VERSION")
	if err != nil {
		return fmt.Errorf("issue occured when deleting the old version file, error: %v", err)
	}
	// create new file & modify
	bs := []byte("v" + nVersion)
	err = ioutil.WriteFile("VERSION", bs, 0744)
	if err != nil {
		return fmt.Errorf("issue occured with new version file error: %v", err)
	}
	return nil
}
