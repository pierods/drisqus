// Copyright Piero de Salvia.
// All Rights Reserved

/*
t.Skip() the tests based on the flag.
*/

package drisqus

import (
	"testing"

	"github.com/pierods/gisqus"
)

func mockForumURLS() {

	forumInterestingForumsJSON := readTestFile("forumsinterestingforums.json")
	forumDetailsJSON := readTestFile("forumsforumdetails.json")
	forumListCategoriesJSON := readTestFile("forumslistcategories.json")
	forumMostActiveUsersJSON := readTestFile("forumslistmostactive.json")
	forumListUsersJSON := readTestFile("forumslistforumusers.json")
	forumThreadListJSON := readTestFile("forumsforumlistthreads.json")
	forumMostLikedUsersJSON := readTestFile("forumsmostlikedusers.json")
	forumFollowersJSON := readTestFile("forumslistfollowers.json")

	urls := testGisqus.ReadForumsURLs()

	urls.DetailsURL = switchHS(urls.DetailsURL, forumDetailsJSON)
	urls.InterestingForumsURL = switchHS(urls.InterestingForumsURL, forumInterestingForumsJSON)
	urls.CategoriesURL = switchHS(urls.CategoriesURL, forumListCategoriesJSON)
	urls.MostActiveUsersURL = switchHS(urls.MostActiveUsersURL, forumMostActiveUsersJSON)
	urls.ListUsersURL = switchHS(urls.ListUsersURL, forumListUsersJSON)
	urls.ListThreadsURL = switchHS(urls.ListThreadsURL, forumThreadListJSON)
	urls.MostLikedUsersURL = switchHS(urls.MostLikedUsersURL, forumMostLikedUsersJSON)
	urls.ListFollowersURL = switchHS(urls.ListFollowersURL, forumFollowersJSON)

	testGisqus.SetForumsURLs(urls)
}

func TestForumDetails(t *testing.T) {

	details, err := testDrisqus.ForumDetails(testCtx, "mapleleafshotstove")
	if err != nil {
		t.Fatal(err)
	}
	if gisqus.ToDisqusTimeExact(details.CreatedAt) != "2011-04-21T18:47:32.503946" {
		t.Fatal("Should be able to parse the created at field")
	}
	if details.Founder != "9408501" {
		t.Fatal("Should be able to retrieve founder")
	}
	if !details.Settings.AllowAnonPost {
		t.Fatal("Should be able to retrieve Allow Anon Post")
	}
	if details.OrganizationID != 583182 {
		t.Fatal("Should be able to retrieve an organization id")
	}
}

func TestForumInteresting(t *testing.T) {

	interestingForums, err := testDrisqus.ForumInteresting(testCtx, 1)
	if err != nil {
		t.Log(err)
	}
	if len(interestingForums) != 5 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Reason != "7,787 comments this week" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Forum.ID != "mapleleafshotstove" {
		t.Log(interestingForums[0].Forum.ID)
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Forum.Favicon.Permalink != "https://disqus.com/api/forums/favicons/mapleleafshotstove.jpg" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Forum.Favicon.Cache != "https://c.disquscdn.com/uploads/forums/77/598/favicon.png" {
		t.Fatal("Should be able to correctly unmarshal items")
	}

	if interestingForums[0].Forum.CreatedAt.Format(gisqus.DisqusDateFormatExact) != "2011-04-21T18:47:32.503946" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if !interestingForums[0].Forum.Settings.AllowAnonPost {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Forum.Avatar.Small.Permalink != "https://disqus.com/api/forums/avatars/mapleleafshotstove.jpg?size=32" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums[0].Forum.Avatar.Small.Cache != "https://c.disquscdn.com/uploads/forums/77/598/avatar32.jpg?1435553857" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
}

func TestForumCategories(t *testing.T) {

	categories, err := testDrisqus.ForumCategories(testCtx, "mapleleafshotstove", -1, "")
	if err != nil {
		t.Fatal(err)
	}
	if categories[0].Title != "General" {
		t.Fatal("Should be able to retrieve a category name")
	}
	if categories[0].Forum != "alloutdoor" {
		t.Fatal("Should be able to retrieve a forum id")
	}
	if categories[0].ID != "2409406" {
		t.Fatal("Should be able to retrieve a category id")
	}

}

func TestForumMostActiveUsers(t *testing.T) {

	_, testErr = testDrisqus.ForumMostActiveUsers(testCtx, "", -1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testDrisqus.ForumMostActiveUsers(testCtx, "mapleleafshotstove", -1, "")
	if err != nil {
		t.Fatal("Should be able to call the forum followers endpoint - ", err)
	}

	if len(users) != 24 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users[0].Username != "icechest" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users[0].Rep != 23.690665 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2015-07-06T22:57:31" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/icechest.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users[0].Avatar.Small.Cache != "https://c.disquscdn.com/uploads/users/16444/4895/avatar32.jpg?1461376631" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}

func TestForumUsers(t *testing.T) {

	_, testErr = testDrisqus.ForumUsers(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testDrisqus.ForumUsers(testCtx, "mapleleafshotstove", 1, "")
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(users) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users[0].Username != "laross19" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users[0].Rep != 1.2537909999999999 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2008-08-10T02:54:57" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/laross19.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users[0].Avatar.Small.Cache != "//a.disquscdn.com/1495487563/images/noavatar32.png" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}

func TestForumThreads(t *testing.T) {

	_, testErr = testDrisqus.ForumThreadsQuick(testCtx, "", 1)
	if testErr == nil {
		t.Fatal("Should check for an empty forum id")
	}
	threads, err := testDrisqus.ForumThreadsQuick(testCtx, "mapleleafshotstove", 1)
	if err != nil {
		t.Fatal(err)
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

func TestForumMostLikedUsers(t *testing.T) {

	_, testErr = testDrisqus.ForumMostLikedUsers(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testDrisqus.ForumMostLikedUsers(testCtx, "mapleleafshotstove", 1, "")
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(users) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users[0].Username != "Burtonboy" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users[0].ID != "9413311" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2011-04-22T02:22:13" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	// rest of user details are tested in user list test
}

func TestForumFollowers(t *testing.T) {

	_, testErr = testDrisqus.ForumFollowers(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testDrisqus.ForumFollowers(testCtx, "mapleleafshotstove", 1, "")
	if err != nil {
		t.Fatal("Should be able to call the forum followers endpoint - ", err)
	}

	if len(users) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users[0].Username != "Johnld778" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users[0].Rep != 1.375799 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2008-02-27T08:05:22" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/Johnld778.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users[0].Avatar.Small.Cache != "https://c.disquscdn.com/uploads/users/12235/avatar32.jpg?1395182401" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}
