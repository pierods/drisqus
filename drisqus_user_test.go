// Copyright Piero de Salvia.
// All Rights Reserved

package drisqus

import (
	"testing"

	"github.com/pierods/gisqus"
)

func mockUserURLs() {

	usersMostActiveForumsJSON := readTestFile("usersmostactiveforums.json")
	usersActivitiesJSON := readTestFile("userslistactivity.json")
	usersListPostsJSON := readTestFile("userslistposts.json")
	usersUserDetailsJSON := readTestFile("usersuserdetail.json")
	usersInterestingUsersJSON := readTestFile("usersinterestingusers.json")
	usersActiveForumsJSON := readTestFile("usersactiveforums.json")
	usersFollowersJSON := readTestFile("usersfollowers.json")
	usersFollowingJSON := readTestFile("usersfollowing.json")
	usersForumFollowingJSON := readTestFile("usersfollowingforums.json")

	usersUrls := testGisqus.ReadUsersURLs()

	usersUrls.PostListURL = switchHS(usersUrls.PostListURL, usersListPostsJSON)
	usersUrls.MostActiveForumsURL = switchHS(usersUrls.MostActiveForumsURL, usersMostActiveForumsJSON)
	usersUrls.ActivityURL = switchHS(usersUrls.ActivityURL, usersActivitiesJSON)
	usersUrls.DetailURL = switchHS(usersUrls.DetailURL, usersUserDetailsJSON)
	usersUrls.InterestingIUsersURL = switchHS(usersUrls.InterestingIUsersURL, usersInterestingUsersJSON)
	usersUrls.ActiveForumsURL = switchHS(usersUrls.ActiveForumsURL, usersActiveForumsJSON)
	usersUrls.FollowersURL = switchHS(usersUrls.FollowersURL, usersFollowersJSON)
	usersUrls.FollowingURL = switchHS(usersUrls.FollowingURL, usersFollowingJSON)
	usersUrls.FollowingForumsURL = switchHS(usersUrls.FollowingForumsURL, usersForumFollowingJSON)

	testGisqus.SetUsersURLs(usersUrls)
}

func TestUserPosts(t *testing.T) {

	_, testErr = testDrisqus.UserPostsQuick(testCtx, "", 1)
	if testErr == nil {
		t.Fatal("Should be able to reject a null user")
	}
	posts, err := testDrisqus.UserPostsQuick(testCtx, "79849", 1)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(posts) != 25 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if posts[0].Message != "<p>Looks really good so far....keeping my fingers crossed.</p>" {
		t.Fatal("Should be able to retrieve post message")
	}
	if posts[0].ID != "2978710471" {
		t.Fatal("Should be able to retrieve post id")
	}
	if posts[0].Author.Username != "laross19" {
		t.Fatal("Should be able to retrieve author username")
	}
	if gisqus.ToDisqusTime(posts[0].CreatedAt) != "2016-11-01T02:54:43" {
		t.Fatal("Should be able to retrieve post created at")
	}
	if posts[0].Parent != 2978642690 {
		t.Fatal("Should be able to retrieve post parent")
	}
	if posts[0].Thread != "5268209091" {
		t.Fatal("Should be able to retrieve post thread")
	}
	if posts[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve post forum")
	}
}

func TestMostActiveForums(t *testing.T) {

	_, testErr = testDrisqus.UserMostActiveForums(testCtx, "")
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	forums, err := testDrisqus.UserMostActiveForums(testCtx, "253940813")
	if err != nil {
		t.Fatal(err)
	}
	if gisqus.ToDisqusTimeExact(forums[0].CreatedAt) != "2011-06-02T18:50:59.765656" {
		t.Fatal("Should be able to parse the created at field")
	}
	if forums[0].Founder != "110449899" {
		t.Fatal("Should be able to retrieve founder")
	}

	if forums[0].OrganizationID != 110 {
		t.Fatal("Should be able to retrieve an organization id")
	}
	if !forums[0].Settings.OrganicDiscoveryEnabled {
		t.Log(forums[0].Settings.MustVerifyEmail)
		t.Fatal("Should be able to retrieve organicDiscoveryEnabled")
	}
}

func TestUserActivities(t *testing.T) {

	_, testErr = testDrisqus.UserActivitiesQuick(testCtx, "", 1)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	activities, err := testDrisqus.UserActivitiesQuick(testCtx, "79849", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(activities.Posts) != 25 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if activities.Posts[0].ID != "3357202237" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if activities.Posts[0].Author.Username != "coachbuzzcut" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if activities.Posts[0].Author.ID != "253940813" {
		t.Fatal("Should be able to retrieve a post's username's id")
	}
	if gisqus.ToDisqusTime(activities.Posts[0].CreatedAt) != "2017-06-13T10:20:29" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if activities.Posts[0].Parent != 3356547778 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if activities.Posts[0].Thread != "5903840168" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if activities.Posts[0].Forum != "tmz" {
		t.Fatal("Should be able to retrieve a post's forum")
	}
}

func TestUserDetails(t *testing.T) {

	_, testErr = testDrisqus.UserDetails(testCtx, "")
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	user, err := testDrisqus.UserDetails(testCtx, "79849")
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != "79849" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if user.Rep != 1.2537909999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if gisqus.ToDisqusTime(user.JoinedAt) != "2008-08-10T02:54:57" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if user.Username != "laross19" {
		t.Fatal("Should be able to retrieve a user's username")
	}
	if user.NumLikesReceived != 14 {
		t.Fatal("Should be able to retrieve a users' likes received")
	}
	if user.NumPosts != 56 {
		t.Fatal("Should be able to retrieve a user's number of posts")
	}
	if user.NumFollowers != 22 {
		t.Fatal("Should be able to retrieve a user's number of followers")
	}
	if user.NumFollowing != 33 {
		t.Fatal("Should be able to retrieve a user's number of following")
	}
	if user.NumForumsFollowing != 44 {
		t.Fatal("Should be able to retrieve a user's number of forums following")
	}
}

func TestUserInteresting(t *testing.T) {

	users, err := testDrisqus.UserInteresting(testCtx, 1)
	if err != nil {
		t.Fatal("Should be able to call the forum list interesting users endpoint - ", err)
	}
	if len(users) != 5 {
		t.Fatal("Should be able to retrieve all users")
	}
	if users[4].User.Username != "anticonsoleshit" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users[4].User.Name != "de ja ful" {
		t.Fatal("Should be able to retrieve a user name")
	}
	if users[4].User.ProfileURL != "https://disqus.com/by/anticonsoleshit/" {
		t.Fatal("Should be able to retrieve a profile url")
	}
	if users[4].User.Reputation != 6.920859999999999 {
		t.Fatal("Should be able to retrieve a user's reputation")
	}
	if gisqus.ToDisqusTime(users[4].User.JoinedAt) != "2015-06-02T14:43:19" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users[4].User.ID != "160076302" {
		t.Fatal("Should be able to retrieve a user's id")
	}

}

func TestUserActiveForums(t *testing.T) {

	_, testErr = testDrisqus.UserActiveForums(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject an empty user id")
	}

	forums, err := testDrisqus.UserActiveForums(testCtx, "46351054", 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(forums) != 25 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums[0].CreatedAt.Format(gisqus.DisqusDateFormatExact) != "2008-04-09T23:30:16.843273" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums[0].Founder != "847" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums[0].ID != "tvnewser" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums[0].Name != "TVNewser" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums[0].OrganizationID != 618 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}

func TestUserFollowers(t *testing.T) {

	_, testErr = testDrisqus.UserFollowers(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testDrisqus.UserFollowers(testCtx, "46351054", 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users[0].ID != "32414357" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users[0].Rep != 0.4153629999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2012-09-18T17:29:47" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users[0].Username != "disqus_IpEgXB3c55" {
		t.Fatal("Should be able to retrieve a user's username")
	}

}

func TestUserFollowing(t *testing.T) {

	_, testErr = testDrisqus.UserFollowing(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testDrisqus.UserFollowing(testCtx, "195792235", 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users[0].ID != "32078576" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users[0].Rep != 1.3459269999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if gisqus.ToDisqusTime(users[0].JoinedAt) != "2012-09-13T05:54:26" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users[0].Username != "flamedance58" {
		t.Fatal("Should be able to retrieve a user's username")
	}

}

func TestUserForumFollowing(t *testing.T) {

	_, testErr = testDrisqus.UserForumFollowing(testCtx, "", 1, "")
	if testErr == nil {
		t.Fatal("Should be able to reject an empty user id")
	}

	forums, err := testDrisqus.UserForumFollowing(testCtx, "46351054", 1, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(forums) != 16 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums[0].CreatedAt.Format(gisqus.DisqusDateFormatExact) != "2015-06-04T17:40:19.641774" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums[0].Founder != "172746617" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums[0].ID != "channel-animeforthepeople" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums[0].Name != "Anime For The People" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums[0].OrganizationID != 3644738 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}
