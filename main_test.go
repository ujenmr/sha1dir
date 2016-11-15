package main

import (
	"os"
	"sort"
	"strings"
	"testing"
)

func TestMain_Sha1Sum(t *testing.T) {
	testString := "test"
	origSha1Sum := "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	testSha1Sum := sha1sum(testString)
	if testSha1Sum != origSha1Sum {
		t.Fatalf("origin sha1sum '%v' test sha1sum '%v' is not valid", origSha1Sum, testSha1Sum)
	}
}

func TestMain_FileData(t *testing.T) {
	testDir := "test_dir"
	testFile := "file1"
	origFileData := "DATAFILE1\ntest_dirfile1"

	// info, err := os.Stat(testDir + "/" + testFile)
	// if err != nil {
	// 	t.Error(err.Error())
	// }

	testFileData, err := fileData(testDir + "/" + testFile)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if origFileData != testFileData {
		t.Fatalf("origin FileData '%v' test FileData '%v' is not valid", origFileData, testFileData)
	}
}

func TestMain_GenBulkFileData(t *testing.T) {
	origPath := [2]string{"test_dir/file1", "test_dir1/test_subdir1/test_file1"}
	origBulkFileData := "DATAFILE1\ntest_dirfile1DATAFILE2\ntest_dir1test_subdir1test_file1"

	testPath := []Path{}
	for i := 0; i < len(origPath); i++ {
		info, err := os.Stat(origPath[i])
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		testPath = append(testPath, Path{origPath[i], strings.ToLower(origPath[i]), info})
	}

	testBulkFileData, err := genBulkFileData(testPath)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if origBulkFileData != testBulkFileData {
		t.Fatalf("origin BulkFileData '%v' test BulkFileData '%v' is not valid", origBulkFileData, testBulkFileData)
	}
}

func TestMain_PathSort(t *testing.T) {
	origPath := [2]string{"test_dir2/fileB", "test_dir2/filea"}
	testPath := []Path{}
	for i := 0; i < len(origPath); i++ {
		info, err := os.Stat(origPath[i])
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		testPath = append(testPath, Path{origPath[i], strings.ToLower(origPath[i]), info})
	}
	if testPath[0].Name != "test_dir2/fileB" {
		t.Fatalf("Orig sorting failed")
	}
	sort.Sort(ByLowPath(testPath))
	if testPath[0].Name != "test_dir2/filea" {
		t.Fatalf("Test sorting failed")
	}
}

func TestMain_SignDir(t *testing.T) {
	origSha1Sum := "9892abb1e795eed35e4ef246bfcc1535e5e8da2a"
	testSha1Sum, err := signDir("test_dir3")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if origSha1Sum != testSha1Sum {
		t.Fatalf("SignDir failed")
	}
}
