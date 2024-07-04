package main

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	for _, tC := range parseArgsTestcases {
		_ = tC
		config, err := ParseArgs(tC.args)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("config    : %#+v\n", *config)
		t.Logf("tC.config : %#+v\n", tC.config)
		if !config.Equals(tC.config) {
			t.Fail()
		}
	}
}

var parseArgsTestcases = []struct {
	desc   string
	args   []string
	config Config
}{
	// {
	// 	"with empty args, should return error",
	// 	[]string{},
	// 	Config{},
	// },
	{
		"with empty args, should return error",
		[]string{"-r"},
		Config{DisplayMethod: "bar", Recurse: true},
	},
	{
		"with empty args, should return error",
		[]string{"--recurse"},
		Config{DisplayMethod: "bar", Recurse: true},
	},
	{
		"with empty args, should return error",
		[]string{"-d", "box"},
		Config{DisplayMethod: "box"},
	},
	{
		"with empty args, should return error",
		[]string{"-d", "line"},
		Config{DisplayMethod: "line"},
	},
	{
		"with empty args, should return error",
		[]string{"-d", "line-cli"},
		Config{DisplayMethod: "line-cli"},
	},
	{
		"with empty args, should return error",
		[]string{"-d", "stdout"},
		Config{DisplayMethod: "stdout"},
	},
	{
		"with empty args, should return error",
		[]string{"-d", "stdout"},
		Config{DisplayMethod: "stdout"},
	},
}
