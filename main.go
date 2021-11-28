package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := os.Args[1]
	change := os.Args[2]

	err := checkCommitMsg(change)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = checkVersionFileExist(file)
	if err != nil {
		fmt.Printf("%v \n", err)
		err = createVersionFile(file, change)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		cVersion, vform, err := readInFile(file)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		newVersion := buildNewVersion(cVersion, change)

		err = modifyVersionFile(newVersion, file, vform)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

const (
	p1 = "major"
	p2 = "minor"
	p3 = "patch"
)

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

func buildNewVersion(cVersion, change string) (nVersion string) {
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

func readInFile(file string) (cVersion string, vform bool, err error) {
	currentVersion, err := ioutil.ReadFile(file)
	if err != nil {
		return "", false, fmt.Errorf("issue with reading the file in")
	}
	if strings.HasPrefix(string(currentVersion), "v") {
		return string(currentVersion[1:]), true, nil
	} else {
		return string(currentVersion[0:]), false, nil
	}
}

func createVersionFile(filePath, change string) (err error) {
	cVersion := "0.0.0"
	nVersion := buildNewVersion(cVersion, change)
	bs := []byte(nVersion)
	err = ioutil.WriteFile(filePath, bs, 0744)
	if err != nil {
		return fmt.Errorf("issue occured when writing to the new version file")
	}
	return nil
}

func modifyVersionFile(nVersion, fileName string, vform bool) (err error) {
	var bs []byte
	// remove current file
	err = deleteOldVersionFile(fileName)
	if err != nil {
		return err
	}
	// create new file & modify
	if vform {
		bs = []byte("v" + nVersion)
	} else {
		bs = []byte(nVersion)
	}
	err = ioutil.WriteFile(fileName, bs, 0755)
	if err != nil {
		return fmt.Errorf("issue occured when writing to the new version file")
	}
	return nil
}

func deleteOldVersionFile(filePath string) (err error) {
	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("issue occured when deleting the old version file, error")
	}
	return nil
}

func checkVersionFileExist(filePath string) (err error) {
	_, err = os.Open(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("version file doesn't exist..creating one")
	}
	return nil
}

func checkCommitMsg(commit string) (err error) {
	if strings.Contains(commit, p1) || strings.Contains(commit, p2) || strings.Contains(commit, p3) {
		return nil
	}
	return fmt.Errorf("please add major, minor or patch to commit msg")

}
