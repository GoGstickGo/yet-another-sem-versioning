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
	message := os.Args[2]
	//gitVersion := os.Args[3]

	err := checkCommitMsg(message)
	if err != nil {
		fmt.Println("There is no #major, #minor or #patch in commit message.Skipping automatic semversioning.")
		os.Exit(0)
	}

	err = checkVersionFileExist(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		err = createVersionFile(file, message)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		cVersion, vform, err := readInFile(file)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		newVersion := buildVersion(cVersion, message)
		// set version for tagging
		os.Setenv("NEW_VERSION_PATH", newVersion)

		err = modifyVersionFile(newVersion, file, vform)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

const (
	p1 = "#major"
	p2 = "#minor"
	p3 = "#patch"
)

func buildVersion(cV, m string) (nV string) {
	s := strings.Split(cV, ".")
	switch m {
	case p1:
		a, b, c := 0, 1, 2
		mV := buildVersionHelper(s, a)
		nV = mV + "." + s[b] + "." + s[c]
	case p2:
		a, b, c := 1, 0, 2
		mV := buildVersionHelper(s, a)
		nV = s[b] + "." + mV + "." + s[c]
	case p3:
		a, b, c := 2, 0, 1
		mV := buildVersionHelper(s, a)
		nV = s[b] + "." + s[c] + "." + mV
	}
	return nV
}

func readInFile(f string) (cV string, vf bool, err error) {
	currentVersion, err := ioutil.ReadFile(f)
	if err != nil {
		return "", false, fmt.Errorf("issue with reading the file in")
	}
	if strings.HasPrefix(string(currentVersion), "v") {
		return string(currentVersion[1:]), true, nil
	} else {
		return string(currentVersion[0:]), false, nil
	}
}

func createVersionFile(f, m string) (err error) {
	cV := "0.0.0"
	nV := buildVersion(cV, m)
	bs := []byte(nV)
	err = ioutil.WriteFile(f, bs, 0744)
	if err != nil {
		return fmt.Errorf("issue occured when writing to the new version file")
	}
	return nil
}

func modifyVersionFile(nV, f string, vf bool) (err error) {
	var bs []byte
	// remove current file
	err = deleteOldVersionFile(f)
	if err != nil {
		return err
	}
	// create new file & modify
	if vf {
		bs = []byte("v" + nV)
	} else {
		bs = []byte(nV)
	}
	err = ioutil.WriteFile(f, bs, 0755)
	if err != nil {
		return fmt.Errorf("issue occured when writing to the new version file")
	}
	return nil
}

func deleteOldVersionFile(f string) (err error) {
	err = os.Remove(f)
	if err != nil {
		return fmt.Errorf("issue occured when deleting the old version file, error")
	}
	return nil
}

func checkVersionFileExist(f string) (err error) {
	_, err = os.Open(f)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("version file doesn't exist..creating one")
	}
	return nil
}

func checkCommitMsg(c string) (err error) {
	if strings.Contains(c, p1) || strings.Contains(c, p2) || strings.Contains(c, p3) {
		return nil
	}
	return fmt.Errorf("please add major, minor or patch to commit msg")

}

func buildVersionHelper(s []string, a int) (nV string) {
	mV, _ := strconv.Atoi(s[a])
	mV++
	nV = strconv.Itoa(mV)
	return nV
}

/*func compareVersions(gV, cV, f string) (err error) {
	if cV != gV {
		return fmt.Errorf("version mismatch")
	}
	return nil
}*/
