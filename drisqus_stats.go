package drisqus

import (
	"strconv"

	"github.com/pierods/gisqus"
)

type AuthorID string
type ReplyCount int
type PostCount int
type LikeCount int
type DislikeCount int
type PostID string
type AuthorUsername string

type AuthorStats struct {
	ID             AuthorID       `json:"id"`
	Username       AuthorUsername `json:"username"`
	Posts          PostCount      `json:"posts"`
	Replies        ReplyCount     `json:"replies"`
	RepliesWritten ReplyCount     `json:"repliesWritten"`
	Likes          LikeCount      `json:"likes"`
	Dislikes       DislikeCount   `json:"dislikes"`
}

type ThreadStats struct {
	AuthorStatsMap *map[AuthorID]*AuthorStats `json:"authorStatsMap"`
	PostMap        *map[PostID]*gisqus.Post   `json:"postMap"`
	OrphanPosts    int                        `json:"orphanPosts"`
	RootPosts      int                        `json:"rootPosts"`
	TotalLikes     LikeCount                  `json:"totalLikes"`
	Totaldislikes  DislikeCount               `json:"totalDislikes"`
	TotalReplies   ReplyCount                 `json:"totalReplies"`
}

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
			parentId := strconv.Itoa(post.Parent)
			parentPost, exists := postMap[PostID(parentId)]
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

type AuthorReplyCount struct {
	ReplierName AuthorUsername `json:"replierName"`
	AuthorName  AuthorUsername `json:"authorName"`
	Replies     ReplyCount     `json:"replies"`
}

func (arc *AuthorReplyCount) MarshalJSON() ([]byte, error) {
	replies := strconv.Itoa(int(arc.Replies))
	oneEl := `["` + string(arc.ReplierName) + `","` + string(arc.AuthorName) + `",` + replies + `]`

	return []byte(oneEl), nil
}

func (d *Drisqus) MakeReplyGroups(posts []*gisqus.Post) *[]AuthorReplyCount {

	postMap := make(map[PostID]*gisqus.Post)

	for _, post := range posts {
		postMap[PostID(post.ID)] = post
	}

	authorMap := make(map[AuthorUsername]map[AuthorUsername]int)

	for _, post := range postMap {
		// let's count replies.
		if post.Parent != 0 { // roots don't have parents (are not replies)
			parentId := strconv.Itoa(post.Parent)
			// it's not an orphan post, otherwise could not attribute reply to parent
			if parentPost, exists := postMap[PostID(parentId)]; exists {
				var parentAuthorsMap map[AuthorUsername]int
				if parentAuthorsMap, exists = authorMap[AuthorUsername(post.Author.Username)]; !exists {
					parentAuthorsMap = make(map[AuthorUsername]int)
					authorMap[AuthorUsername(post.Author.Username)] = parentAuthorsMap
				}
				parentAuthorsMap[AuthorUsername(parentPost.Author.Username)]++
			}
		}
	}

	var replySlice []AuthorReplyCount

	for replier, parentMap := range authorMap {
		for author, replies := range parentMap {
			replySlice = append(replySlice, AuthorReplyCount{
				ReplierName: replier,
				AuthorName:  author,
				Replies:     ReplyCount(replies),
			})
		}
	}

	return &replySlice
}
