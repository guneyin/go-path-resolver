package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type TestItem struct {
	FilePath string
	Want     int
}

var Tests = []TestItem{
	{
		FilePath: "test_files/test1.json",
		Want:     18,
	},
	{
		FilePath: "test_files/test2.json",
		Want:     4,
	},
	{
		FilePath: "test_files/test3.json",
		Want:     228,
	},
}

func (i TestItem) Test(t *testing.T) {
	data, err := ioutil.ReadFile(i.FilePath)
	if err != nil {
		t.Error(err.Error())
	}

	var r Request
	_ = json.Unmarshal(data, &r)

	r.Init()

	want := i.Want
	got := r.Response.PathSum

	name := filepath.Base(i.FilePath)

	if want != got {
		t.Errorf("Test %v failed: want = %v, got = %v", name, want, got)
	} else {
		t.Logf("Test %v success", name)
	}
}

func TestAll(t *testing.T) {
	args := os.Args

	fmt.Println(args)

	for _, test := range Tests {
		test.Test(t)
	}
}
