package drisqus

// RepliesInThread  is used by MakeReplySlice
type RepliesInThread struct {
	ID       AuthorID
	UserName AuthorUsername
	Replies  ReplyCount
}

/*
MakeReplySlice is a d3.js-friendly method, makes a slice of author-replies
*/
func (d *Drisqus) MakeReplySlice(authorStatsMap *map[AuthorID]*AuthorStats) []RepliesInThread {

	var replies []RepliesInThread

	for _, authorStats := range *authorStatsMap {
		replies = append(replies, RepliesInThread{
			authorStats.ID,
			authorStats.Username,
			authorStats.Replies,
		})
	}
	return replies
}

/*
MakeReplyWrittenSlice is a d3.js-friendly method, makes a slice of author-replies (written)
*/
func (d *Drisqus) MakeReplyWrittenSlice(authorStatsMap *map[AuthorID]*AuthorStats) []RepliesInThread {

	var replies []RepliesInThread

	for _, authorStats := range *authorStatsMap {
		replies = append(replies, RepliesInThread{
			authorStats.ID,
			authorStats.Username,
			authorStats.RepliesWritten,
		})
	}
	return replies
}

// PostsByAuthor is used by MakePostCountSlice
type PostsByAuthor struct {
	ID       AuthorID
	UserName AuthorUsername
	Posts    PostCount
}

/*
MakePostCountSlice is a d3.js-friendly method, makes a slice of author-post count
*/
func (d *Drisqus) MakePostCountSlice(authorStatsMap *map[AuthorID]*AuthorStats) []PostsByAuthor {

	var posts []PostsByAuthor

	for _, authorStats := range *authorStatsMap {
		posts = append(posts, PostsByAuthor{
			authorStats.ID,
			authorStats.Username,
			authorStats.Posts,
		})
	}
	return posts
}

//LikesByAuthor is used by MakeLikeCountSlice
type LikesByAuthor struct {
	ID       AuthorID
	UserName AuthorUsername
	Likes    LikeCount
	Dislikes DislikeCount
}

/*
MakeLikeCountSlice is a d3.js-friendly method, makes a slice of author-likes-dislikes count
*/
func (d *Drisqus) MakeLikeCountSlice(authorStatsMap *map[AuthorID]*AuthorStats) []LikesByAuthor {

	var likes []LikesByAuthor

	for _, authorStats := range *authorStatsMap {
		likes = append(likes, LikesByAuthor{
			authorStats.ID,
			authorStats.Username,
			authorStats.Likes,
			authorStats.Dislikes,
		})
	}
	return likes
}

// AuthorReplyCount is used by MakeReplyGroups
type AuthorReplyCount struct {
	ReplierName AuthorUsername `json:"replierName"`
	AuthorName  AuthorUsername `json:"authorName"`
	Replies     ReplyCount     `json:"replies"`
}

// MarshalJSON is used to coerce an AuthorReplyCount into a mixed type Javascript array of the form [string, string, int]
func (arc *AuthorReplyCount) MarshalJSON() ([]byte, error) {
	replies := strconv.Itoa(int(arc.Replies))
	oneEl := `["` + string(arc.ReplierName) + `","` + string(arc.AuthorName) + `",` + replies + `]`

	return []byte(oneEl), nil
}

/*
MakeReplyGroups is a d3.js-friendly method, makes a slice of author-to-author reply counts
*/
func (d *Drisqus) MakeReplyGroups(posts []*gisqus.Post) *[]AuthorReplyCount {

	postMap := make(map[PostID]*gisqus.Post)

	for _, post := range posts {
		postMap[PostID(post.ID)] = post
	}

	authorMap := make(map[AuthorUsername]map[AuthorUsername]int)

	for _, post := range postMap {
		// let's count replies.
		if post.Parent != 0 { // roots don't have parents (are not replies)
			parentID := strconv.Itoa(post.Parent)
			// it's not an orphan post, otherwise could not attribute reply to parent
			if parentPost, exists := postMap[PostID(parentID)]; exists {
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
