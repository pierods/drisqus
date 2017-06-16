// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"strconv"

	"github.com/pierods/gisqus"
)

// AuthorID aliases string to signal intent of type
type AuthorID string

// ReplyCount aliases int to signal intent of type
type ReplyCount int

// PostCount aliases int to signal intent of type
type PostCount int

// LikeCount aliases int to signal intent of type
type LikeCount int

//DislikeCount aliases int to signal intent of type
type DislikeCount int

// PostID aliases string to signal intent of type
type PostID string

// AuthorUsername aliases string to signal intent of type
type AuthorUsername string

// AuthorStats contains statistics for a user as far as posting to a single thread goes
type AuthorStats struct {
	ID             AuthorID       `json:"id"`
	Username       AuthorUsername `json:"username"`
	Posts          PostCount      `json:"posts"`
	Replies        ReplyCount     `json:"replies"`
	RepliesWritten ReplyCount     `json:"repliesWritten"`
	Likes          LikeCount      `json:"likes"`
	Dislikes       DislikeCount   `json:"dislikes"`
}

// ThreadStats contains statistics for a single forum thread
type ThreadStats struct {
	AuthorStatsMap *map[AuthorID]*AuthorStats `json:"authorStatsMap"`
	PostMap        *map[PostID]*gisqus.Post   `json:"postMap"`
	OrphanPosts    int                        `json:"orphanPosts"`
	RootPosts      int                        `json:"rootPosts"`
	TotalLikes     LikeCount                  `json:"totalLikes"`
	Totaldislikes  DislikeCount               `json:"totalDislikes"`
	TotalReplies   ReplyCount                 `json:"totalReplies"`
}

// MakeThreadStats produces an instance of ThreadStats given a sampling of posts.
func (d *Drisqus) MakeThreadStats(posts []*gisqus.Post) *ThreadStats {

	postMap := make(map[PostID]*gisqus.Post)
	authorMap := make(map[AuthorID]*AuthorStats)
	var likes, dislikes, replies int

	for _, post := range posts {
		postMap[PostID(post.ID)] = post
	}

	var orphanCount, roots int

	for _, post := range postMap {
		var author *AuthorStats
		var exists bool
		if author, exists = authorMap[AuthorID(post.Author.ID)]; !exists {
			author = &AuthorStats{
				ID:       AuthorID(post.Author.ID),
				Username: AuthorUsername(post.Author.Username),
			}
			authorMap[AuthorID(post.Author.ID)] = author
		}
		author.Posts++
		if post.Likes > 0 {
			author.Likes += LikeCount(post.Likes)
			likes += post.Likes
		}
		if post.Dislikes > 0 {
			author.Dislikes += DislikeCount(post.Dislikes)
			dislikes += post.Dislikes
		}
		if post.Parent != 0 { // let's count replies. roots don't have parents (are not replies) and are not orphans
			parentID := strconv.Itoa(post.Parent)
			parentPost, exists := postMap[PostID(parentID)]
			if !exists {
				// it's an orphan post - cannot attribute reply to parent
				orphanCount++
			} else {
				var parentAuthor *AuthorStats
				if parentAuthor, exists = authorMap[AuthorID(parentPost.Author.ID)]; !exists {
					parentAuthor = &AuthorStats{
						ID:       AuthorID(parentPost.Author.ID),
						Username: AuthorUsername(parentPost.Author.Username),
					}
					authorMap[AuthorID(parentPost.Author.ID)] = parentAuthor
				}
				parentAuthor.Replies++
				author.RepliesWritten++
				replies++
			}
		} else {
			roots++
		}
	}

	return &ThreadStats{
		AuthorStatsMap: &authorMap,
		PostMap:        &postMap,
		OrphanPosts:    orphanCount,
		RootPosts:      roots,
		TotalLikes:     LikeCount(likes),
		Totaldislikes:  DislikeCount(dislikes),
		TotalReplies:   ReplyCount(replies),
	}
}
