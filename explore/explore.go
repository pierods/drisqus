package main

/*

forum : site (takismag, returnofkings)
thread: article
post: comment

https://0value.com/Let-the-Doer-Do-it
 reentrancy
*/

import (
	"context"
	"fmt"
	"net/url"
	"os"
	//"github.com/davecgh/go-spew/spew"
	"encoding/json"

	"github.com/pierods/drisqus"
	"github.com/pierods/gisqus"
)

func exploreTrendingThreads(ctx context.Context) {

	values := url.Values{}
	values.Set("forum", "rawstory")
	values.Set("interval", "1h")
	trends, err := g.ThreadTrending(values, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, trend := range trends.Response {
		fmt.Println(trend.TrendingThread.Id, ", ", trend.TrendingThread.CleanTitle, trend.TrendingThread.Posts)
	}

}

func exploreInterestingForums(ctx context.Context) {

	iForums, err := d.ForumInteresting(1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for _, iForum := range iForums {
		fmt.Println(iForum.Reason, ", ", iForum.Forum.Id, ", ", iForum.Forum.Name)
	}
	fmt.Println("*******************************************************", len(iForums))
}

func exploreAllInterestingForums(ctx context.Context) {

	iForums, err := d.ForumInteresting(-1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for _, iForum := range iForums {
		fmt.Println(iForum.Reason, ", ", iForum.Forum.Id, ", ", iForum.Forum.Name)
	}
	fmt.Println("*******************************************************", len(iForums))
}

func explorePopularThreads(ctx context.Context) {

	values := url.Values{}
	values.Set("forum", "rawstory")
	values.Set("interval", "1h")
	threads, err := g.ThreadPopular(values, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, thread := range threads.Response {
		fmt.Println(thread.Id, ", ", thread.CleanTitle, thread.Posts)
	}

}

func threadDetails(threadId string, ctx context.Context) {
	values := url.Values{}
	thread, err := g.ThreadDetails(threadId, values, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(thread.Response)
}

func forumDetails(ctx context.Context) {
	values := url.Values{}
	forumId := "returnofkings"
	fDetails, err := g.ForumDetails(forumId, values, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%+v/n", fDetails.Response)
}

func listThreadsByPostCount(forumId string, ctx context.Context) {
	f := drisqus.ThreadsByPostCount

	threads, err := d.ThreadList(forumId, 7, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	drisqus.ThreadBy(f).Sort(threads)

	for _, thread := range threads {
		fmt.Println(thread.Posts, ", ", thread.Id, ", ", thread.Title, ", ", thread.Likes)
	}
	fmt.Println("*******************************************************", len(threads))
}

func listThreadsByLikes(forumId string, ctx context.Context) {
	f := drisqus.ThreadsByLikes

	threads, err := d.ThreadList(forumId, 5, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	drisqus.ThreadBy(f).Sort(threads)
	for _, thread := range threads {
		fmt.Println(thread.Posts, ", ", thread.Id, ", ", thread.Title, ", ", thread.Likes)
	}
	fmt.Println("*******************************************************", len(threads))
}

var g gisqus.Gisqus
var d drisqus.Drisqus

func main() {

	//spew.Dump(users)

	g = gisqus.NewGisqus("KpVAypnhCxG27eRLRbXad0i1xfbyUsHPE7E8on5wbFJkbQcIzjB0pkJ4kMOfTRmx")
	d = drisqus.NewDrisqus(g)
	ctx, _ := context.WithCancel(context.TODO())

	//exploreAllInterestingForums(ctx)
	//exploreInterestingForums(ctx)
	//explorePopularThreads(ctx)
	//exploreTrendingThreads(ctx)

	//listThreadsByPostCount("wccftech", ctx)
	//listThreadsByPostCount("rawstory", ctx)
	//listThreadsByPostCount("returnofkings", ctx)
	//listThreadsByPostCount("tmz", ctx)

	//listThreadsByLikes("wccftech", ctx)
	//listThreadsByLikes("returnofkings", ctx)
	//listThreadsByLikes("rawstory", ctx)
	//listThreadsByLikes("tmz", ctx)

	//listAllPosts("5859394484", ctx) // rok - pochi
	//listPosts("5549194138", ctx) // rok
	//listPostsByLikes("5549194138", ctx) // rok
	//listPostsByDislikes("5549194138", ctx) // rok

	//listPostsByLikes("5834414038", ctx) // wccftech
	//listPostsByDislikes("5834414038", ctx) // wccftech

	//listPostsByLikes("5860794546", ctx) // tmz
	//listPostsByDislikes("5860794546", ctx) // tmz

	//makePostTree("5549194138", ctx) // rok
	//makePostTree("5860794546", ctx) // tmz
	//makePostTree("5834414038", ctx) // wccftech
	//makePostTree("5854266958", ctx) // rawstory

	//makeCompletePostTree("5783374349", ctx) // rok

	exploreReplyGroups("5549194138", ctx)
	fmt.Println(g.Limits())
}

func exploreReplyGroups(thread string, ctx context.Context) {
	posts, err := d.ThreadPosts(thread, 1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	groups := d.MakeReplyGroups(posts)
	for _, group := range *groups {
		fmt.Println(group.ReplierName, " replied to  ", group.AuthorName, " ", group.Replies, " times")
	}
	json, err := json.Marshal(groups)
	fmt.Println(string(json), err)
}

func dumpPostMap(thread string, ctx context.Context) {
	posts, err := d.ThreadPosts(thread, 1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	stats := d.MakeThreadStats(posts)
	for _, post := range *stats.PostMap {
		fmt.Println(post.Id, ", ", post.Parent)
	}
}

func makePostTree(thread string, ctx context.Context) {
	posts, err := d.ThreadPosts(thread, 1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var orphanRate float32
	stats := d.MakeThreadStats(posts)
	orphanRate = float32(stats.OrphanPosts) / float32(len(*stats.PostMap))
	for _, authStats := range *stats.AuthorStatsMap {
		fmt.Println(authStats.Id, ", ", authStats.Replies)
	}
	fmt.Println("len(postMap)=", len(*stats.PostMap), " roots=", stats.RootPosts, " orphans=", stats.OrphanPosts,
		" orphanRate=", orphanRate, " totalLikes=", stats.TotalLikes, ", ", " totalDislikes=", stats.Totaldislikes, " totalReplies=", stats.TotalReplies)
}

func makeCompletePostTree(thread string, ctx context.Context) {
	posts, err := d.ThreadPosts(thread, -1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var orphanRate float32
	stats := d.MakeThreadStats(posts)
	orphanRate = float32(stats.OrphanPosts) / float32(len(*stats.PostMap))
	for _, authStats := range *stats.AuthorStatsMap {
		fmt.Println(authStats.Id, ", ", authStats.Replies)
	}
	fmt.Println("len(postMap)=", len(*stats.PostMap), " roots=", stats.RootPosts, " orphans=", stats.OrphanPosts,
		" orphanRate=", orphanRate, " totalLikes=", stats.TotalLikes, ", ", " totalDislikes=", stats.Totaldislikes, " totalReplies=", stats.TotalReplies)
}

func listAllPosts(thread string, ctx context.Context) {

	posts, err := d.ThreadPosts(thread, -1, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, post := range posts {
		fmt.Println(post.Likes, ", ", post.Id, ", ", post.DisqusTimeCreatedAt, ", ", post.Message)
	}
	fmt.Println("*******************************************************", len(posts))
}
func listPosts(thread string, ctx context.Context) {

	posts, err := d.ThreadPosts(thread, 5, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, post := range posts {
		fmt.Println(post.Likes, ", ", post.Id, ", ", post.DisqusTimeCreatedAt, ", ", post.Message)
	}
	fmt.Println("*******************************************************", len(posts))
}

func listPostsByLikes(thread string, ctx context.Context) {

	posts, err := d.ThreadPosts(thread, 7, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	f := drisqus.PostsByLikes
	drisqus.PostBy(f).Sort(posts)

	for _, post := range posts {
		fmt.Println(post.Likes, ", ", post.Id, ", ", post.DisqusTimeCreatedAt, ", ", post.Message)
	}
	fmt.Println("*******************************************************", len(posts))
}

func listPostsByDislikes(thread string, ctx context.Context) {

	posts, err := d.ThreadPosts(thread, 7, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	f := drisqus.PostsByDislikes
	drisqus.PostBy(f).Sort(posts)

	for _, post := range posts {
		fmt.Println(post.Dislikes, ", ", post.Id, ", ", post.DisqusTimeCreatedAt, ", ", post.Message)
	}
	fmt.Println("*******************************************************", len(posts))
}
