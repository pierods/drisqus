// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"context"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pierods/gisqus"
)

/*
ForumMostActiveUsers wraps https://disqus.com/api/docs/forums/listMostActiveUsers/ (https://disqus.com/api/3.0/forums/listMostActiveUsers.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumMostActiveUsers(ctx context.Context, forumID string, pages int, order string) ([]*gisqus.User, error) {
	values := url.Values{}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.ForumMostActiveUsers(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumMostActiveUsers(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil

}

// InterestingForum models the response to a call to Disqus' Interesting Forums (https://disqus.com/api/docs/forums/interestingForums/)
type InterestingForum struct {
	Reason string        `json:"reason"`
	Forum  *gisqus.Forum `json:"forum"`
}

type forumsByNumberOfPosts []InterestingForum

func (a forumsByNumberOfPosts) Len() int      { return len(a) }
func (a forumsByNumberOfPosts) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a forumsByNumberOfPosts) Less(i, j int) bool {

	ris := strings.Replace(strings.Split(a[i].Reason, " ")[0], ",", "", -1)
	rjs := strings.Replace(strings.Split(a[j].Reason, " ")[0], ",", "", -1)
	ri, _ := strconv.Atoi(ris)
	rj, _ := strconv.Atoi(rjs)
	return ri > rj
}

/*
ForumInteresting wraps https://disqus.com/api/docs/forums/interestingForums/ (https://disqus.com/api/3.0/forums/interestingForums.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumInteresting(ctx context.Context, pages int) ([]InterestingForum, error) {

	var interestingForums []InterestingForum
	values := url.Values{}
	values.Set("limit", "100")
	gisqusResponse, err := d.gisqus.ForumInteresting(ctx, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		for _, item := range gisqusResponse.Response.Items {
			interestingForum := InterestingForum{}
			interestingForum.Reason = item.Reason
			interestingForum.Forum = gisqusResponse.Response.Objects[item.ID]
			interestingForums = append(interestingForums, interestingForum)
		}
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumInteresting(ctx, values)
		if err != nil {
			return nil, err
		}
	}

	sort.Sort(forumsByNumberOfPosts(interestingForums))
	return interestingForums, nil
}

/*
ForumCategories wraps https://disqus.com/api/docs/forums/listCategories/ (https://disqus.com/api/3.0/forums/listCategories.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumCategories(ctx context.Context, forumID string, pages int, order string) ([]*gisqus.Category, error) {

	values := url.Values{}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var categories []*gisqus.Category

	gisqusResponse, err := d.gisqus.ForumCategories(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		categories = append(categories, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumCategories(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return categories, nil
}

/*
ForumDetails wraps https://disqus.com/api/docs/forums/details/ (https://disqus.com/api/3.0/forums/details.json)
It does not support the "related" url parameter (other funcs can be used for drilldown)
*/
func (d *Drisqus) ForumDetails(ctx context.Context, forumID string) (*gisqus.Forum, error) {

	values := url.Values{}
	forum, err := d.gisqus.ForumDetails(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	return forum.Response, nil
}

/*
ForumFollowers wraps https://disqus.com/api/docs/forums/listFollowers/ (https://disqus.com/api/3.0/forums/listFollowers.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumFollowers(ctx context.Context, forumID string, pages int, order string) ([]*gisqus.User, error) {

	values := url.Values{}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.ForumFollowers(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumFollowers(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

/*
ForumUsers wraps https://disqus.com/api/3.0/forums/listUsers.json (https://disqus.com/api/docs/forums/listUsers/)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumUsers(ctx context.Context, forumID string, pages int, order string) ([]*gisqus.User, error) {

	values := url.Values{}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.ForumUsers(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumUsers(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

/*
ForumThreadsQuick wraps ForumThreads. It includes frequently used parameters, and sets the rest to their zero values
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumThreadsQuick(ctx context.Context, forumID string, pages int) ([]*gisqus.Thread, error) {
	return d.ForumThreads(ctx, forumID, []string{}, []string{}, pages, time.Time{}, "")
}

/*
ForumThreads wraps https://disqus.com/api/docs/forums/listThreads/ (https://disqus.com/api/3.0/forums/listThreads.json)
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumThreads(ctx context.Context, forumID string, threadIDs, includes []string, pages int, since time.Time, order string) ([]*gisqus.Thread, error) {

	values := url.Values{}
	for _, threadID := range threadIDs {
		values.Add("thread", threadID)
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
	var threads []*gisqus.Thread

	gisqusResponse, err := d.gisqus.ForumThreads(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		threads = append(threads, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumThreads(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return threads, nil
}

/*
ForumMostLikedUsers wraps https://disqus.com/api/docs/forums/listMostLikedUsers/ (https://disqus.com/api/3.0/forums/listMostLikedUsers.json)
Disqus does not return the # of likes with this call.
When pages is -1, all pages are retrieved.
*/
func (d *Drisqus) ForumMostLikedUsers(ctx context.Context, forumID string, pages int, order string) ([]*gisqus.User, error) {

	values := url.Values{}
	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.ForumMostLikedUsers(ctx, forumID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.ForumMostLikedUsers(ctx, forumID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}
