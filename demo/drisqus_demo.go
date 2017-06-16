// Copyright Piero de Salvia.
// All Rights Reserved

package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pierods/drisqus"
	"github.com/pierods/gisqus"
)

var (
	drisq drisqus.Drisqus
)

func readFile(name string) ([]byte, error) {

	goPath := os.Getenv("GOPATH")
	filePath := goPath + "/src/github.com/pierods/drisqus/demo/" + name

	f, fErr := os.Open(filePath)
	defer f.Close()

	if fErr != nil {
		return nil, fErr
	}
	bytes, fErr := ioutil.ReadAll(f)

	if fErr != nil {
		return nil, fErr
	}
	return bytes, nil
}

func handler(rw http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if strings.HasPrefix(path, "/interestingforums") {
		ctx, _ := context.WithCancel(context.TODO())
		iForums, err := drisq.ForumInteresting(ctx, -1)

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		json.NewEncoder(rw).Encode(&iForums)

		return
	}
	if strings.HasPrefix(path, "/trends") {

		ctx, _ := context.WithCancel(context.TODO())
		trends, err := drisq.ThreadTrending(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw, err)
			fmt.Println(err)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(&err)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&trends)

		return
	}
	if strings.HasPrefix(path, "/threadlatest") {
		query := r.URL.Query()
		forum := query.Get("forum")

		ctx, _ := context.WithCancel(context.TODO())
		threads, err := drisq.ThreadListQuick(ctx, forum, 1)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw, err)
			fmt.Println(err)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(&err)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&threads)

		return
	}
	if strings.HasPrefix(path, "/popularthreads") {
		query := r.URL.Query()
		forum := query.Get("forum")

		ctx, _ := context.WithCancel(context.TODO())
		threads, err := drisq.ThreadPopular(ctx, "", forum, gisqus.Interval1h, true)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw, err)
			fmt.Println(err)
			return
		}
		rw.Header().Set("Content-Type", "application/json")

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(&err)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&threads)

		return
	}
	if strings.HasPrefix(path, "/hotthreads") {
		query := r.URL.Query()
		forum := query.Get("forum")

		ctx, _ := context.WithCancel(context.TODO())
		threads, err := drisq.ThreadHot(ctx, []string{}, []string{forum}, []string{}, []string{})
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw, err)
			fmt.Println(err)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&threads)

		return
	}
	if strings.HasPrefix(path, "/threadposts") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 1)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		json.NewEncoder(rw).Encode(&posts)

		return
	}
	if strings.HasPrefix(path, "/makestats") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		json.NewEncoder(rw).Encode(&stats)

		return
	}
	if strings.HasPrefix(path, "/makepostcountslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)
		slice := drisq.MakePostCountSlice(stats.AuthorStatsMap)
		rw.Header().Set("Content-Type", "text/csv")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		records := [][]string{{"id", "value"}}
		for _, author := range slice {
			records = append(records, []string{string(author.UserName), strconv.Itoa(int(author.Posts))})
		}
		w := csv.NewWriter(rw)
		w.WriteAll(records)
		if err := w.Error(); err != nil {
			fmt.Println(err)
		}
		return
	}
	if strings.HasPrefix(path, "/makereplygroupslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		slice := drisq.MakeReplyGroups(posts)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		json.NewEncoder(rw).Encode(&slice)
		return
	}
	if strings.HasPrefix(path, "/makereplyhadslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)
		slice := drisq.MakeReplySlice(stats.AuthorStatsMap)
		rw.Header().Set("Content-Type", "text/csv")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		records := [][]string{{"id", "value"}}
		for _, author := range slice {
			records = append(records, []string{string(author.UserName), strconv.Itoa(int(author.Replies))})
		}
		w := csv.NewWriter(rw)
		w.WriteAll(records)
		if err := w.Error(); err != nil {
			fmt.Println(err)
		}
		return
	}
	if strings.HasPrefix(path, "/makereplywrittenslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)
		slice := drisq.MakeReplyWrittenSlice(stats.AuthorStatsMap)
		rw.Header().Set("Content-Type", "text/csv")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		records := [][]string{{"id", "value"}}
		for _, author := range slice {
			records = append(records, []string{string(author.UserName), strconv.Itoa(int(author.Replies))})
		}
		w := csv.NewWriter(rw)
		w.WriteAll(records)
		if err := w.Error(); err != nil {
			fmt.Println(err)
		}
		return
	}
	if strings.HasPrefix(path, "/makelikecountslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)
		slice := drisq.MakeLikeCountSlice(stats.AuthorStatsMap)
		rw.Header().Set("Content-Type", "text/csv")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		records := [][]string{{"id", "value"}}
		for _, author := range slice {
			records = append(records, []string{string(author.UserName), strconv.Itoa(int(author.Likes))})
		}
		w := csv.NewWriter(rw)
		w.WriteAll(records)
		if err := w.Error(); err != nil {
			fmt.Println(err)
		}
		return
	}

	if strings.HasPrefix(path, "/makedislikecountslice") {
		query := r.URL.Query()
		thread := query.Get("thread")

		ctx, _ := context.WithCancel(context.TODO())
		posts, err := drisq.ThreadPostsQuick(ctx, thread, 2)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			fmt.Println(err)
			return
		}
		stats := drisq.MakeThreadStats(posts)
		slice := drisq.MakeLikeCountSlice(stats.AuthorStatsMap)
		rw.Header().Set("Content-Type", "text/csv")
		rw.WriteHeader(200)
		if err != nil {
			json.NewEncoder(rw).Encode(&err)
			return
		}
		records := [][]string{{"id", "value"}}
		for _, author := range slice {
			records = append(records, []string{string(author.UserName), strconv.Itoa(int(author.Dislikes))})
		}
		w := csv.NewWriter(rw)
		w.WriteAll(records)
		if err := w.Error(); err != nil {
			fmt.Println(err)
		}
		return
	}

	if path == "/" {
		path = "/index.html"
	}
	page, err := readFile(path)
	rw.Header().Set("Content/Type", "text/html")
	if err != nil {
		rw.WriteHeader(404)
		fmt.Fprint(rw, err)
	}
	rw.WriteHeader(200)
	rw.Write(page)
}

func main() {

	keyBytes, err := readFile("/disqus_secret_key.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	key := string(keyBytes)

	http.HandleFunc("/", handler)

	g := gisqus.NewGisqus(key)
	drisq = drisqus.NewDrisqus(g)
	fmt.Println("Drisqus demo : Started : Listening on port 30000")
	http.ListenAndServe(":30000", nil)
}
