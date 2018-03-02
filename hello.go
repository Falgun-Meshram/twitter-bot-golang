package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	fmt.Println("Getting user name now")

	consumerKey := "wT0zMae4oQ0PppULjQcMzO74e"
	consumerSecret := "P7B4NuayiLpOjMfDFWU6lT5Jf6FmRoNjC19M69GOBjPl311fm7"
	accessToken := "935330576361336832-i335w3SfKHnVGLekLYr7xDlpNiyryDh"
	accessSecret := "u9nWww7UrDfEPFGs5iDyXl5hbXlSrK6tkbM3VTt6C40fi"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.Name)

	searchParams := &twitter.SearchTweetParams{
		Query:      "#holi",
		Count:      5,
		ResultType: "recent",
		Lang:       "en",
	}

	searchedTweet, _, _ := client.Search.Tweets(searchParams)

	for _, tweet := range searchedTweet.Statuses {
		tweet_id := tweet.ID
		client.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})
		fmt.Printf("Retweeted %+v\n", tweet.Text)

	}
}
