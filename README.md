# drisqus - a high level client of Disqus' API, written in Go

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![](https://godoc.org/github.com/pierods/drisqus?status.svg)](http://godoc.org/github.com/pierods/drisqus)
[![Go Report Card](https://goreportcard.com/badge/github.com/pierods/drisqus)](https://goreportcard.com/report/github.com/pierods/drisqus)
[![Build Status](https://travis-ci.org/pierods/drisqus.svg?branch=master)](https://travis-ci.org/pierods/drisqus)

Drisqus is a high level client of Disqus' API ((https://disqus.com/api/docs/)). Its underlying library is Gisqus (https://github.com/pierods/gisqus).
Drisqus aims to cover endpoints that read data (GET method) for reporting, statistical and data drilling purposes.

As in Gisqus, Drisqus has the following limitations:
* it only supports authentication in the form of "Authenticating as the Account Owner" (https://disqus.com/api/docs/auth/)
* endpoints that require entity IDs (thread ID, forum ID etc) but where they can be provided implicitly by authentication have their wrappers 
  requiring those parameters explicitly in the method signature

The "related" parameter in many Disqus endpoints is not supported, since data returned through it can always be gotten with a direct call to the 
respective api. In this sense, Drisqus covers the complete hierarchy of Disqus' object model.

### Usage
After having obtained an API key from Disqus (you must create an app for that on Disqus' control panel), one must obtain an instance of Gisqus and
then pass it to Drisqus:

```Go
    import  "context"
    import  "net/url"
    import  "github.com/pierods/gisqus"
    import  "github.com/pierods/drisqus"
    ...
    g := NewGisqus("api key")
    d := NewDrisqus(g)
    values := url.Values{}
    ctx, cancel := context.WithCancel(context.TODO())
```

One can then proceed to make calls against Disqus' endpoints. Calls do not support timeouts, but they are cancellable (https://golang.org/pkg/context/).

```Go
    
    posts, err := d.PostList(ctx, values)
    if err != nil {
        ...
    }
    fmt.Println(posts[0].ID)
```
Methods that support the "pages" parameter will retrieve pages of size 100. If pages is set to -1, all pages are retrieved.

### Data drilling and statistics
Drisqus makes it really is to drill down in Disqus' API and calculate statistics.

As an example, let's pick the latest thread from a forum and check out which comment authors don't have any replies to their posts.

```Go
    import "fmt"
    import  "context"
    import  "net/url"
    import  "github.com/pierods/gisqus"
    import  "github.com/pierods/drisqus"
    ...
    g := NewGisqus("api key")
    d := NewDrisqus(g)
    
    ctx, cancel := context.WithCancel(context.TODO())
    forums, err := d.ForumInteresting(ctx, 1)
    if err != nil {
        ...
    }
    forumID := forums[0].Forum.ID
    
     
```

### Notes
All calls are cancellable, so they won't catastrophically block on a call chain.

The complete Disqus hierarchy is modeled:


![hierarchy](assets/chart-api-relationships.png)

except for Categories, since the vast majority of forums use only one category for all posts.

### Endpoints covered
https://disqus.com/api/docs/
##### Forums
* details
* interestingForums
* listCategories
* listFollowers 
* listMostActiveUsers
* listMostLikedUsers
* listThreads
* listUsers

##### Threads
* details
* list
* listHot 
* listPopular 
* listPosts
* listUsersVotedThread
* set

##### Posts
* details
* getContext 
* list
* listPopular

##### Users
* details 
* interestingUsers
* listActiveForums
* listActivity 
* listFollowers 
* listFollowing 
* listFollowingForums 
* listMostActiveForums 
* listPosts
