package drisqus

// d3.js-friendly struct
type RepliesInThread struct {
	ID       AuthorID
	UserName AuthorUsername
	Replies  ReplyCount
}

/*
d3.js-friendly method, makes a slice of author-replies
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
d3.js-friendly method, makes a slice of author-replies (written)
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

// d3.js-friendly struct
type PostsByAuthor struct {
	ID       AuthorID
	UserName AuthorUsername
	Posts    PostCount
}

/*
d3.js-friendly method, makes a slice of author-post count
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

type LikesByAuthor struct {
	ID       AuthorID
	UserName AuthorUsername
	Likes    LikeCount
	Dislikes DislikeCount
}

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
