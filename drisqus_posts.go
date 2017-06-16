// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"context"
	"net/url"
	"sort"
	"time"

	"github.com/pierods/gisqus"
)

// PostsByLikes is used by postSorter to sort Posts
func PostsByLikes(p1, p2 *gisqus.Post) bool {
	return p1.Likes > p2.Likes
}

// PostsByDislikes is used by postSorter to sort Posts
func PostsByDislikes(p1, p2 *gisqus.Post) bool {
	return p1.Dislikes > p2.Dislikes
}

// PostBy is an utility method used to sort Posts
type PostBy func(t1, t2 *gisqus.Post) bool

type postSorter struct {
	posts []*gisqus.Post
	by    func(t1, t2 *gisqus.Post) bool
}

//Sort is the sorting method of PostBy
func (by PostBy) Sort(posts []*gisqus.Post) {
	ps := &postSorter{
		posts: posts,
		by:    by,
	}
	sort.Sort(ps)
}

func (s *postSorter) Len() int {
	return len(s.posts)
}

func (s *postSorter) Swap(i, j int) {
	s.posts[i], s.posts[j] = s.posts[j], s.posts[i]
}

func (s *postSorter) Less(i, j int) bool {
	return s.by(s.posts[i], s.posts[j])
}

/*
PostDetails wraps https://disqus.com/api/docs/posts/details/ (https://disqus.com/api/3.0/posts/details.json)
It does not support the "related" argument.
*/
func (d *Drisqus) PostDetails(ctx context.Context, postID string) (*gisqus.Post, error) {

	values := url.Values{}
	gisqusResponse, err := d.gisqus.PostDetails(ctx, postID, values)
	if err != nil {
		return nil, err
	}
	return gisqusResponse.Response, nil
}

/*
PostListQuick wraps PostList. It includes frequently used parameters, and sets the rest to their zero values
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) PostListQuick(ctx context.Context, threadID string, pages int) ([]*gisqus.Post, error) {
	return d.PostList(ctx, pages, []string{}, []string{threadID}, []string{}, time.Time{}, time.Time{}, "", "")
}

/*
PostList wraps https://disqus.com/api/docs/posts/list/ (https://disqus.com/api/3.0/posts/list.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) PostList(ctx context.Context, pages int, categoryIDs, threadIDs, forumIDs []string, start, end time.Time, sortType, order string) ([]*gisqus.Post, error) {

	values := url.Values{}
	for _, categoryID := range categoryIDs {
		values.Add("category", categoryID)
	}
	for _, threadID := range threadIDs {
		values.Add("thread", threadID)
	}
	for _, forumID := range forumIDs {
		values.Add("forum", forumID)
	}
	if start != (time.Time{}) {
		values.Set("start", gisqus.ToDisqusTime(start))
	}
	if end != (time.Time{}) {
		values.Set("end", gisqus.ToDisqusTime(end))
	}
	if sortType != "" {
		values.Set("sortType", sortType)
	}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	gisqusResponse, err := d.gisqus.PostList(ctx, values)
	if err != nil {
		return nil, err
	}
	var posts []*gisqus.Post

	for page := 0; page < pages || pages == -1; page++ {
		posts = append(posts, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.PostList(ctx, values)
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}

/*
PostPopularQuick wraps PostPopular. It includes frequently used parameters, and sets the rest to their zero values
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) PostPopularQuick(ctx context.Context, threadID string, pages int) ([]*gisqus.Post, error) {
	return d.PostPopular(ctx, pages, []string{}, []string{threadID}, []string{}, "", "", "", "")
}

/*
PostPopular wraps https://disqus.com/api/docs/posts/listPopular/ (https://disqus.com/api/3.0/posts/listPopular.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) PostPopular(ctx context.Context, pages int, forumIDs, threadIDs, includes []string, categoryID, interval, organizationID, order string) ([]*gisqus.Post, error) {

	values := url.Values{}
	for _, threadID := range threadIDs {
		values.Add("thread", threadID)
	}
	for _, forumID := range forumIDs {
		values.Add("forum", forumID)
	}
	for _, include := range includes {
		values.Add("include", include)
	}
	if categoryID != "" {
		values.Set("category", categoryID)
	}
	if interval != "" {
		values.Set("interval", interval)
	}
	if organizationID != "" {
		values.Set("organization", organizationID)
	}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	gisqusResponse, err := d.gisqus.PostPopular(ctx, values)
	if err != nil {
		return nil, err
	}
	return gisqusResponse.Response, nil
}
