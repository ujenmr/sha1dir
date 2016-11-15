// Copyright 2016 Ievgen Khmelenko <ujenmr@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var filesList []Path

// Path struct
type Path struct {
	Name        string
	NameLowCase string
	Info        os.FileInfo
}

// ByLowPath implements sort.Interface for []Path based on
// the NameLowCase field.
type ByLowPath []Path

func (a ByLowPath) Len() int {
	return len(a)
}

func (a ByLowPath) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLowPath) Less(i, j int) bool {
	return a[i].NameLowCase < a[j].NameLowCase
}

func sha1sum(s string) string {
	hash := sha1.New()
	io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}

func signWalkFunc(path string, info os.FileInfo, err error) error {
	filesList = append(filesList, Path{path, strings.ToLower(path), info})
	return nil
}

func fileData(path string) (string, error) {
	info, err := os.Stat(path)
	if err == nil {
		if info.Mode().IsRegular() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return "", err
			}
			formatedPath := strings.ToLower(strings.Replace(path, "/", "", -1))
			return string(fileData) + formatedPath, nil
		}
	}
	return "", nil
}

func genBulkFileData(files []Path) (data string, err error) {
	for i := 0; i < len(files); i++ {
		filedata, err := fileData(files[i].Name)
		if err != nil {
			return "", err
		}
		data = data + filedata
	}
	return data, nil
}

func signDir(dir string) (string, error) {
	_, err := os.Stat(dir)
	if err != nil {
		return "-", err
	}

	curDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(curDir)

	filesList = []Path{}
	err = filepath.Walk("./", signWalkFunc)
	if err != nil {
		return "-", err
	}

	sort.Sort(ByLowPath(filesList))

	bulkData, err := genBulkFileData(filesList)
	if err != nil {
		return "-", err
	}

	return sha1sum(bulkData), nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		sha1, _ := signDir("./")
		fmt.Printf("%v - %v\n", ".", sha1)
	}

	for i := 0; i < len(args); i++ {
		sha1, _ := signDir(args[i])
		fmt.Printf("%v - %v\n", args[i], sha1)
	}
}
