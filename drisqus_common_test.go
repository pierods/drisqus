// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/pierods/gisqus"
	"github.com/pierods/gisqus/mock"
)

var mockServer *mock.Server
var testErr error
var testGisqus gisqus.Gisqus
var testDrisqus Drisqus
var testCtx context.Context
var testValues url.Values
var testDataDir string

var runLiveTests = flag.Bool("live", false, "whether to run tests requiring a connection to the internet")

func init() {
	testCtx, _ = context.WithCancel(context.TODO())

	goPath := os.Getenv("GOPATH")
	testDataDir = goPath + "/src/github.com/pierods/drisqus/testdata/"
}

func TestMain(m *testing.M) {

	if !*runLiveTests {
		testGisqus = gisqus.NewGisqus("secret")
		testDrisqus = NewDrisqus(testGisqus)

		mockServer = mock.NewMockServer()
		defer mockServer.Close()

		mockForumURLS()
		mockThreadURLS()
		mockPostURLS()
		mockUserURLs()
	} else {
		key, err := readKeyFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		testGisqus = gisqus.NewGisqus(key)
		testDrisqus = NewDrisqus(testGisqus)
		// TODO live tests
	}

	retCode := m.Run()
	os.Exit(retCode)
}

func switchHS(URL, JSON string) string {
	result, err := mockServer.SwitchHostAndScheme(URL, JSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return result
}

func readTestFile(fileName string) string {

	f, err := os.Open(testDataDir + fileName)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return string(bytes)
}

func readKeyFile() (string, error) {
	f, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/pierods/drisqus/demo/disqus_secret_key.txt")
	defer f.Close()

	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
