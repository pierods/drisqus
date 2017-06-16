package drisqus

import (
	"context"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/pierods/gisqus"
)

/*
ThreadTrending wraps https://disqus.com/api/docs/trends/listThreads/ (https://disqus.com/api/3.0/trends/listThreads.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (d *Drisqus) ThreadTrending(ctx context.Context) ([]*gisqus.Trend, error) {

	values := url.Values{}
	values.Set("limit", "10")
	trendsResponse, err := d.gisqus.ThreadTrending(ctx, values)
	if err != nil {
		return nil, err
	}
	return trendsResponse.Response, nil
}

/*
ThreadPostsQuick wraps ThreadPosts. It includes frequently used parameters, and sets the rest to their zero values
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ThreadPostsQuick(ctx context.Context, threadID string, pages int) ([]*gisqus.Post, error) {
	return d.ThreadPosts(ctx, threadID, pages, []string{}, []string{}, "", time.Time{}, "")
}

/*
ThreadPosts wraps https://disqus.com/api/docs/threads/listPosts/ (https://disqus.com/api/3.0/threads/listPosts.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ThreadPosts(ctx context.Context, threadID string, pages int, filters, includes []string, forumID string, since time.Time, order string) ([]*gisqus.Post, error) {

	values := url.Values{}
	for _, filter := range filters {
		values.Add("filter", filter)
	}
	for _, include := range includes {
		values.Add("include", include)
	}
	if forumID != "" {
		values.Set("forum", forumID)
	}
	if since != (time.Time{}) {
		values.Set("since", gisqus.ToDisqusTime(since))
	}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")

	var posts []*gisqus.Post

	postResponse, err := d.gisqus.ThreadPosts(ctx, threadID, values)
	if err != nil {
		return nil, err
	}

	for page := 0; page < pages || pages == -1; page++ {
		posts = append(posts, postResponse.Response...)
		if !postResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", postResponse.Cursor.Next)
		postResponse, err = d.gisqus.ThreadPosts(ctx, threadID, values)
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}

/*
ThreadListQuick wraps ThreadList. It includes frequently used parameters, and sets the rest to their zero values
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ThreadListQuick(ctx context.Context, forumID string, pages int) ([]*gisqus.Thread, error) {
	return d.ThreadList(ctx, []string{}, []string{forumID}, []string{}, []string{}, []string{}, pages, time.Time{}, "")
}

/*
ThreadList wraps https://disqus.com/api/docs/threads/list/ (https://disqus.com/api/3.0/threads/list.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ThreadList(ctx context.Context, categoryIDs, forumIDs, threadIDs, authorIDs, includes []string, pages int, since time.Time, order string) ([]*gisqus.Thread, error) {

	values := url.Values{}
	for _, categoryID := range categoryIDs {
		values.Add("category", categoryID)
	}
	for _, forumID := range categoryIDs {
		values.Add("forum", forumID)
	}
	for _, threadID := range threadIDs {
		values.Add("thread", threadID)
	}
	for _, authorID := range authorIDs {
		values.Add("author", authorID)
	}
	for _, include := range includes {
		values.Add("include", include)
	}
	if since != (time.Time{}) {
		values.Set("since", gisqus.ToDisqusTime(since))
	}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	for _, forumID := range forumIDs {
		values.Add("forum", forumID)
	}
	threadResponse, err := d.gisqus.ThreadList(ctx, values)
	if err != nil {
		return nil, err
	}
	var threads []*gisqus.Thread

	for page := 0; page < pages || pages == -1; page++ {
		threads = append(threads, threadResponse.Response...)
		if !threadResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", threadResponse.Cursor.Next)
		threadResponse, err = d.gisqus.ThreadList(ctx, values)
		if err != nil {
			return nil, err
		}
	}
	return threads, nil
}

/*
ThreadHotQuick wraps ThreadHot. It includes frequently used parameters, and sets the rest to their zero values
*/
func (d *Drisqus) ThreadHotQuick(ctx context.Context) ([]*gisqus.Thread, error) {
	return d.ThreadHot(ctx, []string{}, []string{}, []string{}, []string{})
}

/*
ThreadHot wraps https://disqus.com/api/docs/threads/listHot/ (https://disqus.com/api/3.0/threads/listHot.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (d *Drisqus) ThreadHot(ctx context.Context, categoryIDs, forumIDs, authorIDs, includes []string) ([]*gisqus.Thread, error) {

	values := url.Values{}
	values.Set("limit", "100")
	for _, forum := range forumIDs {
		values.Add("forum", forum)
	}
	for _, author := range authorIDs {
		values.Add("author", author)
	}

	threadResponse, err := d.gisqus.ThreadHot(ctx, values)
	if err != nil {
		return nil, err
	}
	var threads []*gisqus.Thread

	for _, thread := range threadResponse.Response {
		threads = append(threads, thread)
	}
	return threads, nil
}

/*
ThreadPopularQuick wraps ThreadPopular. It includes frequently used parameters, and sets the rest to their zero values
*/
func (d *Drisqus) ThreadPopularQuick(ctx context.Context) ([]*gisqus.Thread, error) {
	return d.ThreadPopular(ctx, "", "", "", false)
}

/*
ThreadPopular wraps https://disqus.com/api/docs/threads/listPopular/ (https://disqus.com/api/3.0/threads/listPopular.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (d *Drisqus) ThreadPopular(ctx context.Context, categoryID, forumID, interval string, withTopPost bool) ([]*gisqus.Thread, error) {

	values := url.Values{}
	values.Set("limit", "100")
	if forumID != "" {
		values.Set("forum", forumID)
	}
	if categoryID != "" {
		values.Add("category", categoryID)
	}
	if interval != "" {
		values.Set("interval", interval)
	}
	if withTopPost {
		values.Set("with_top_post", strconv.FormatBool(withTopPost))
	}
	threadResponse, err := d.gisqus.ThreadPopular(ctx, values)
	if err != nil {
		return nil, err
	}
	var threads []*gisqus.Thread

	for _, thread := range threadResponse.Response {
		threads = append(threads, thread)
	}
	return threads, nil
}

/*
ThreadSet wraps https://disqus.com/api/docs/threads/set/ (https://disqus.com/api/3.0/threads/set.json)
*/
func (d *Drisqus) ThreadSet(ctx context.Context, threadIDs []string) ([]*gisqus.Thread, error) {
	values := url.Values{}

	gisqusResponse, err := d.gisqus.ThreadSet(ctx, threadIDs, values)
	if err != nil {
		return nil, err
	}
	return gisqusResponse.Response, nil
}

/*
ThreadDetails wraps https://disqus.com/api/docs/threads/details/ (https://disqus.com/api/3.0/threads/details.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (d *Drisqus) ThreadDetails(ctx context.Context, threadID string) (*gisqus.ThreadDetail, error) {
	values := url.Values{}

	gisqusResponse, err := d.gisqus.ThreadDetails(ctx, threadID, values)
	if err != nil {
		return nil, err
	}
	return gisqusResponse.Response, nil

}

/*
ThreadUsersVoted wraps https://disqus.com/api/docs/threads/listUsersVotedThread/ (https://disqus.com/api/3.0/threads/listUsersVotedThread.json)
Complete users are not returned by Disqus on this call
*/
func (d *Drisqus) ThreadUsersVoted(ctx context.Context, threadID string) ([]*gisqus.User, error) {
	values := url.Values{}
	values.Set("limit", "100")

	gisqusResponse, err := d.gisqus.ThreadUsersVoted(ctx, threadID, values)
	if err != nil {
		return nil, err
	}
	return gisqusResponse.Response, nil
}

// ThreadsByPostCount is used by threadSorter to sort Threads
func ThreadsByPostCount(p1, p2 *gisqus.Thread) bool {
	return p1.Posts > p2.Posts
}

// ThreadsByLikes is used by threadSorter to sort Threads
func ThreadsByLikes(p1, p2 *gisqus.Thread) bool {
	return p1.Likes > p2.Likes
}

// ThreadBy is an utility type used to sort Threads
type ThreadBy func(t1, t2 *gisqus.Thread) bool

type threadSorter struct {
	threads []*gisqus.Thread
	by      func(t1, t2 *gisqus.Thread) bool
}

// Sort is the sorting method of ThreadBy
func (by ThreadBy) Sort(threads []*gisqus.Thread) {
	ts := &threadSorter{
		threads: threads,
		by:      by,
	}
	sort.Sort(ts)
}

func (s *threadSorter) Len() int {
	return len(s.threads)
}

func (s *threadSorter) Swap(i, j int) {
	s.threads[i], s.threads[j] = s.threads[j], s.threads[i]
}

func (s *threadSorter) Less(i, j int) bool {
	return s.by(s.threads[i], s.threads[j])
}
