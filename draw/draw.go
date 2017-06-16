package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"

	"github.com/pierods/drisqus"
	"github.com/pierods/gisqus"
	"github.com/pierods/gisqus/mock"
)

var mockServer *httptest.Server
var err error
var testCtx context.Context
var testGisqus gisqus.Gisqus
var ms mock.MockServer

func init() {
	testGisqus = gisqus.NewGisqus("secret")
	testCtx, _ = context.WithCancel(context.TODO())

	goPath := os.Getenv("GOPATH")
	testDataDir := goPath + "/src/github.com/pierods/drisqus/testdata/"

	ms = mock.NewMockServer(testDataDir)
}

func main() {

	urls := testGisqus.ReadThreadsURLs()
	mockServer = ms.NewServer()
	defer mockServer.Close()

	urls.Thread_posts, err = mock.SwitchHostAndScheme(urls.Thread_posts, mockServer.URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	testGisqus.SetThreadsURLs(urls)

	drisq := drisqus.NewDrisqus(testGisqus)

	posts, err := drisq.ThreadPosts("5894030679", 1, testCtx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	stats := drisq.MakeThreadStats(posts)

	for _, author := range *stats.AuthorStatsMap {
		fmt.Println(author.Username, ", ", author.Id)
	}

	fmt.Println(stats.OrphanPosts)
	fmt.Println(stats.RootPosts)
}
