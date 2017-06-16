// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"fmt"
	"os"
	"testing"

	"github.com/pierods/gisqus"
)

func mockThreadURLS() {

	var err error

	threadPostsJSON, err := readTestFile("threadsthreadposts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListHotJSON, err := readTestFile("threadshotlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListPopularJSON, err := readTestFile("threadspopular.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListJSON, err := readTestFile("threadsthreadlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListTrendingJSON, err := readTestFile("threadstrending.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadSetJSON, err := readTestFile("threadsset.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadDetailsJSON, err := readTestFile("threadsthreaddetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadUsersVotedJSON, err := readTestFile("threadsusersvoted.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	urls := testGisqus.ReadThreadsURLs()

	urls.ThreadPostsURL, err = mockServer.SwitchHostAndScheme(urls.ThreadPostsURL, threadPostsJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadHotURL, err = mockServer.SwitchHostAndScheme(urls.ThreadHotURL, threadListHotJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadPopularURL, err = mockServer.SwitchHostAndScheme(urls.ThreadPopularURL, threadListPopularJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadListURL, err = mockServer.SwitchHostAndScheme(urls.ThreadListURL, threadListJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadTrendingURL, err = mockServer.SwitchHostAndScheme(urls.ThreadTrendingURL, threadListTrendingJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadDetailURL, err = mockServer.SwitchHostAndScheme(urls.ThreadDetailURL, threadDetailsJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadSetURL, err = mockServer.SwitchHostAndScheme(urls.ThreadSetURL, threadSetJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	urls.ThreadUsersVotedURL, err = mockServer.SwitchHostAndScheme(urls.ThreadUsersVotedURL, threadUsersVotedJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	testGisqus.SetThreadsURLs(urls)
}

func TestThreadStats(t *testing.T) {

	posts, err := testDrisqus.ThreadPostsQuick(testCtx, "5846923796", 1)
	if err != nil {
		t.Fatal(err)
	}

	stats := testDrisqus.MakeThreadStats(posts)
	for _, post := range *stats.PostMap {
		t.Log(post.ID, ", ", post.Parent)
	}
}

func TestThreadPosts(t *testing.T) {

	_, testErr = testDrisqus.ThreadPostsQuick(testCtx, "", 1)
	if testErr == nil {
		t.Fatal("Should check for empty thread id")
	}
	posts, err := testDrisqus.ThreadPostsQuick(testCtx, "5846923796", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts[0].ID != "3325943139" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts[0].Author.Username != "loovtrain" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts[0].Author.ID != "163477624" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if gisqus.ToDisqusTime(posts[0].CreatedAt) != "2017-05-26T15:12:18" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts[0].Parent != 3325896546 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if posts[0].Thread != "5846923796" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}

func TestThreadHot(t *testing.T) {

	threads, err := testDrisqus.ThreadHotQuick(testCtx)
	if err != nil {
		t.Fatal("Should be able to call the thread list endpoint - ", err)
	}
	if len(threads) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads[0].Feed != "https://mapleleafshotstove.disqus.com/leafs_links_bob_mckenzie_discusses_kyle_dubas_report_shoots_down_fictitious_william_nylander_trade_r/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads[0].ID != "5846923796" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads[0].Category != "783882" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads[0].Author != "9408501" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(threads[0].CreatedAt) != "2017-05-24T16:41:44" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads[0].Title != "Leafs Links: Bob McKenzie discusses Kyle Dubas report, shoots down fictitious William Nylander trade rumours; Sheldon Keefe on Carl Grundstrom, Kasperi Kapanen and more" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadPopular(t *testing.T) {

	threads, err := testDrisqus.ThreadPopularQuick(testCtx)
	if err != nil {
		t.Fatal("Should be able to call the thread list popular endpoint - ", err)
	}
	if len(threads) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads[0].Feed != "https://alloutdoor.disqus.com/sig_sauer_sued_by_new_jersey_state_police/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads[0].ID != "5829486853" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads[0].Category != "2409406" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads[0].Author != "37536641" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(threads[0].CreatedAt) != "2017-05-18T19:23:54" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads[0].Forum != "alloutdoor" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads[0].Title != "Sig Sauer Sued By New Jersey State Police" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadTrending(t *testing.T) {

	trends, err := testDrisqus.ThreadTrending(testCtx)
	if err != nil {
		t.Fatal("Should be able to call the thread trending endpoint - ", err)
	}
	if len(trends) != 10 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if trends[2].PostLikes != 1665 {
		t.Fatal("Should be able to retrieve a trend's postlikes")
	}
	if trends[2].Posts != 62 {
		t.Fatal("Should be able to retrieve a trend's posts")
	}
	if trends[2].Score != 1.497732426303855 {
		t.Fatal("Should be able to retrieve a trends's score")
	}
	if trends[2].Likes != 90 {
		t.Fatal("Should be able to retrieve a trends' likes")
	}
	if trends[2].TrendingThread.Feed != "https://kissanime.disqus.com/berserk_2017_anime_watch_berserk_2017_anime_online_in_high_quality/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if trends[2].TrendingThread.ID != "5592902940" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if trends[2].TrendingThread.Category != "3204063" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if trends[2].TrendingThread.Author != "100108732" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(trends[2].TrendingThread.CreatedAt) != "2017-03-01T01:42:44" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if trends[2].TrendingThread.Forum != "kissanime" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if trends[2].TrendingThread.Title != "Berserk (2017) anime | Watch Berserk (2017) anime online in high quality" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
	if trends[2].TrendingThread.HighlightedPost.ID != "3316658778" {
		t.Fatal("Should be able to retrieve a trend's highlighted post id")
	}
	if gisqus.ToDisqusTime(trends[2].TrendingThread.HighlightedPost.CreatedAt) != "2017-05-20T23:15:11" {
		t.Fatal("Should be able to retrieve a trend's highlighted post created at")
	}
	if trends[2].TrendingThread.HighlightedPost.Author.Username != "Umbrielle" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's author")
	}
	if gisqus.ToDisqusTime(trends[2].TrendingThread.HighlightedPost.Author.JoinedAt) != "2015-02-06T14:28:51" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's joined at")
	}
	if trends[2].TrendingThread.HighlightedPost.Author.ID != "143213885" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's author")
	}
}

func TestThreadList(t *testing.T) {

	threads, err := testDrisqus.ThreadListQuick(testCtx, "tmz", 1)
	if err != nil {
		t.Fatal("Should be able to call the thread list endpoint - ", err)
	}
	if len(threads) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads[0].Feed != "https://babbel-magazine.disqus.com/personalidades_multilingues_ao_longo_da_historia_babbelcom_087/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads[0].ID != "5850192558" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads[0].Category != "3261556" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads[0].Author != "121561733" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(threads[0].CreatedAt) != "2017-05-25T18:16:19" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads[0].Forum != "babbel-magazine" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads[0].Title != "Personalidades multilíngues ao longo da História - Babbel.com" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadSet(t *testing.T) {

	_, testErr = testDrisqus.ThreadSet(testCtx, []string{})
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	_, testErr = testDrisqus.ThreadSet(testCtx, nil)
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	threads, err := testDrisqus.ThreadSet(testCtx, []string{"5903840168", "5850192558"})
	if err != nil {
		t.Fatal(err)
	}
	if len(threads) != 2 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads[0].Feed != "https://tmz.disqus.com/039bachelor_in_paradise039_star_corinne_olympios_says_she_didn039t_consent_to_sexual_contact_with_de/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads[0].ID != "5903840168" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads[0].Category != "3341905" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads[0].Author != "116162885" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(threads[0].CreatedAt) != "2017-06-12T17:48:04" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads[0].Forum != "tmz" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads[0].Title != "&#039;Bachelor in Paradise&#039; Star Corinne Olympios Says She Didn&#039;t Consent to Sexual Contact with DeMario Jackson" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadDetails(t *testing.T) {

	_, testErr = testDrisqus.ThreadDetails(testCtx, "")
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	details, err := testDrisqus.ThreadDetails(testCtx, "5846923796")
	if err != nil {
		t.Fatal(err)
	}
	if details.ID != "5846923796" {
		t.Fatal("Should be able to retrieve a thread id")
	}
	if details.Category != "783882" {
		t.Fatal("Should be able to retrieve a thread id")
	}
	if details.Author != "9408501" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if gisqus.ToDisqusTime(details.CreatedAt) != "2017-05-24T16:41:44" {
		t.Fatal("Should be able to parse a thread's created at")
	}
	if details.Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if details.Posts != 1927 {
		t.Fatal("Should be able to retrieve a thread's number of posts")
	}
}

func TestThreadUsersVoted(t *testing.T) {

	_, testErr = testDrisqus.ThreadUsersVoted(testCtx, "")
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	users, err := testDrisqus.ThreadUsersVoted(testCtx, "5846923796")
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 5 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users[0].ID != "19365741" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users[0].IsPowerContributor {
		t.Fatal("Should be able to retrieve a user's power contributor")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2011-11-22T10:43:15" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users[0].Username != "bigboss400" {
		t.Fatal("Should be able to retrieve a user's username")
	}
}
