package drisqus

import (
	"context"
	"net/url"
	"time"

	"github.com/pierods/gisqus"
)

// UserActivities models the response of the user activity list endpoint
type UserActivities struct {
	Posts []*gisqus.Post
}

/*
UserActivitiesQuick wraps UserActivities. It includes frequently used parameters, and sets the rest to their zero values
*/
func (d *Drisqus) UserActivitiesQuick(ctx context.Context, userID string, pages int) (*UserActivities, error) {
	return d.UserActivities(ctx, userID, pages, time.Time{}, []string{})
}

/*
UserActivities wraps https://disqus.com/api/docs/users/listActivity/ (https://disqus.com/api/3.0/users/listActivity.json)
*/
func (d *Drisqus) UserActivities(ctx context.Context, userID string, pages int, since time.Time, includes []string) (*UserActivities, error) {

	values := url.Values{}

	if since != (time.Time{}) {
		values.Set("since", gisqus.ToDisqusTime(since))
	}
	for _, include := range includes {
		values.Add("include", include)
	}
	values.Set("limit", "100")
	var activities UserActivities

	gisqusResponse, err := d.gisqus.UserActivities(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		activities.Posts = append(activities.Posts, gisqusResponse.Posts...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserActivities(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return &activities, nil
}

/*
UserMostActiveForums wraps https://disqus.com/api/docs/users/listMostActiveForums/ (https://disqus.com/api/3.0/users/listMostActiveForums.json)
*/
func (d *Drisqus) UserMostActiveForums(ctx context.Context, userID string) ([]*gisqus.Forum, error) {

	values := url.Values{}
	values.Set("limit", "100")
	forums, err := d.gisqus.UserMostActiveForums(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	return forums.Response, nil
}

/*
UserPostsQuick wraps UserPosts. It includes frequently used parameters, and sets the rest to their zero values
*/
func (d *Drisqus) UserPostsQuick(ctx context.Context, userID string, pages int) ([]*gisqus.Post, error) {
	return d.UserPosts(ctx, userID, []string{}, pages, time.Time{}, "")
}

/*
UserPosts wraps https://disqus.com/api/docs/users/listPosts/ (https://disqus.com/api/3.0/users/listPosts.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (d *Drisqus) UserPosts(ctx context.Context, userID string, includes []string, pages int, since time.Time, order string) ([]*gisqus.Post, error) {

	values := url.Values{}
	if since != (time.Time{}) {
		values.Set("since", gisqus.ToDisqusTime(since))
	}
	if order != "" {
		values.Set("order", order)
	}
	for _, include := range includes {
		values.Add("include", include)
	}
	values.Set("limit", "100")
	var posts []*gisqus.Post

	gisqusResponse, err := d.gisqus.UserPosts(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		posts = append(posts, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserPosts(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}

/*
UserDetails wraps https://disqus.com/api/docs/users/details/ (https://disqus.com/api/3.0/users/details.json)
*/
func (d *Drisqus) UserDetails(ctx context.Context, userID string) (*gisqus.User, error) {

	values := url.Values{}
	details, err := d.gisqus.UserDetails(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	return details.Response, nil
}

// InterestingUser models the response of the interesting users endpoint.
type InterestingUser struct {
	Reason string
	User   *gisqus.User
}

/*
UserInteresting wraps https://disqus.com/api/docs/users/interestingUsers/ (https://disqus.com/api/3.0/users/interestingUsers.json)
*/
func (d *Drisqus) UserInteresting(ctx context.Context, pages int) ([]*InterestingUser, error) {

	values := url.Values{}
	values.Set("limit", "100")

	var interestingUsers []*InterestingUser

	gisqusResponse, err := d.gisqus.UserInteresting(ctx, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		for _, item := range gisqusResponse.Response.Items {
			interestingUser := InterestingUser{
				Reason: item.Reason,
				User:   gisqusResponse.Response.Objects[item.ID],
			}
			interestingUsers = append(interestingUsers, &interestingUser)
		}

		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserInteresting(ctx, values)
		if err != nil {
			return nil, err
		}
	}

	return interestingUsers, nil
}

/*
UserActiveForums wraps https://disqus.com/api/docs/users/listActiveForums/ (https://disqus.com/api/3.0/users/listActiveForums.json)
*/
func (d *Drisqus) UserActiveForums(ctx context.Context, userID string, pages int, order string) ([]*gisqus.Forum, error) {

	values := url.Values{}

	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var forums []*gisqus.Forum

	gisqusResponse, err := d.gisqus.UserActiveForums(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		forums = append(forums, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserActiveForums(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return forums, nil
}

/*
UserFollowers wraps https://disqus.com/api/docs/users/listFollowers/ (https://disqus.com/api/3.0/users/listFollowers.json)
Numlikes, NumPosts, NumFollowers are not returned by Disqus' API
*/
func (d *Drisqus) UserFollowers(ctx context.Context, userID string, pages int, order string) ([]*gisqus.User, error) {
	values := url.Values{}

	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.UserFollowers(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserFollowers(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

/*
UserFollowing wraps https://disqus.com/api/docs/users/listFollowing/ (https://disqus.com/api/3.0/users/listFollowing.json)
Numlikes, NumPosts, NumFollowers are not returned by Disqus' API
*/
func (d *Drisqus) UserFollowing(ctx context.Context, userID string, pages int, order string) ([]*gisqus.User, error) {
	values := url.Values{}

	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var users []*gisqus.User

	gisqusResponse, err := d.gisqus.UserFollowing(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		users = append(users, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserFollowing(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

/*
UserForumFollowing wraps https://disqus.com/api/docs/users/listFollowingForums/ (https://disqus.com/api/3.0/users/listFollowingForums.json)
*/
func (d *Drisqus) UserForumFollowing(ctx context.Context, userID string, pages int, order string) ([]*gisqus.Forum, error) {
	values := url.Values{}

	if order != "" {
		values.Set("order", order)
	}
	values.Set("limit", "100")
	var forums []*gisqus.Forum

	gisqusResponse, err := d.gisqus.UserForumFollowing(ctx, userID, values)
	if err != nil {
		return nil, err
	}
	for page := 0; page < pages || pages == -1; page++ {
		forums = append(forums, gisqusResponse.Response...)
		if !gisqusResponse.Cursor.HasNext {
			break
		}
		values.Set("cursor", gisqusResponse.Cursor.Next)
		gisqusResponse, err = d.gisqus.UserForumFollowing(ctx, userID, values)
		if err != nil {
			return nil, err
		}
	}

	return forums, nil
}
