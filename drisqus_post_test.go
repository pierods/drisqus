// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"testing"

	"github.com/pierods/gisqus"
)

func mockPostURLS() {

	postPopularJSON := readTestFile("postspostpopular.json")
	postDetailsJSON := readTestFile("postspostdetails.json")
	postListJSON := readTestFile("postspostlist.json")

	postsURLs := testGisqus.ReadPostURLs()

	postsURLs.PostDetailsURL = switchHS(postsURLs.PostDetailsURL, postDetailsJSON)
	postsURLs.PostListURL = switchHS(postsURLs.PostListURL, postListJSON)
	postsURLs.PostPopularURL = switchHS(postsURLs.PostPopularURL, postPopularJSON)

	testGisqus.SetPostsURLs(postsURLs)
}

func TestPostDetails(t *testing.T) {

	_, testErr = testDrisqus.PostDetails(testCtx, "")
	if testErr == nil {
		t.Fatal("Should check for an empty post id")
	}
	details, err := testDrisqus.PostDetails(testCtx, "3320987826")
	if err != nil {
		t.Fatal(err)
	}
	if details.ID != "3320987826" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if details.Author.Username != "royalewithcrowne" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if details.Author.ID != "209257634" {
		t.Fatal("Should be able to retrieve a post's username's id")
	}
	if gisqus.ToDisqusTime(details.CreatedAt) != "2017-05-23T17:57:41" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if details.Parent != 3320975377 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if details.Thread != "5843656825" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if details.Forum != "wrestlinginc" {
		t.Fatal("Should be able to retrieve a post's forum")
	}
}

func TestPostList(t *testing.T) {

	posts, err := testDrisqus.PostListQuick(testCtx, "3324481803", 1)
	if err != nil {
		t.Fatal("Should be able to call the post list endpoint - ", err)
	}
	if len(posts) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts[0].ID != "3324481803" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts[0].Author.Username != "bautista8190p" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts[0].Author.ID != "242978772" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if gisqus.ToDisqusTime(posts[0].CreatedAt) != "2017-05-25T17:43:47" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts[1].Parent != 3324406982 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if posts[0].Thread != "4978714775" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts[0].Forum != "pregunta2" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}

func TestPostPopular(t *testing.T) {

	posts, err := testDrisqus.PostPopularQuick(testCtx, "3357275751", 1)
	if err != nil {
		t.Fatal("Should be able to call the post popular endpoint - ", err)
	}
	if len(posts) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts[0].ID != "3357275751" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts[0].Author.Username != "mychive-3683e7511cad5234db651099216183d0" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts[0].Author.ID != "252976252" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if gisqus.ToDisqusTime(posts[0].CreatedAt) != "2017-06-13T11:40:12" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts[0].Thread != "5906159869" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts[0].Forum != "thechiverules" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}
