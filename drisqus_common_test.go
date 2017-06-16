package drisqus

import (
	"context"
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

func init() {
	testGisqus = gisqus.NewGisqus("secret")
	testDrisqus = NewDrisqus(testGisqus)
	testCtx, _ = context.WithCancel(context.TODO())

	goPath := os.Getenv("GOPATH")
	testDataDir = goPath + "/src/github.com/pierods/drisqus/testdata/"
}

func readTestFile(fileName string) (string, error) {

	f, err := os.Open(testDataDir + fileName)
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

func TestMain(m *testing.M) {

	var live bool

	if len(os.Args) > 1 {
		for _, arg := range os.Args {
			if arg == "live" {
				live = true
				break
			}
		}
	}
	if !live {
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
	}

	retCode := m.Run()
	os.Exit(retCode)
}
