package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dghubble/oauth1"
)

const (
	consumerKey    = "rmCcKa3arY3dQA7QbSXSQhB6E"
	consumerSecret = "w9sy8hOKcrap0BTDhDV6HIePgUHDhC9vcxbL2TVrO688XJzFae"
	accessToken    = "1831479382457696257-aPD22FxGILAStd1acHNrGKRL4jNDHI"
	accessSecret   = "GpDGVP5LolMUthXgrjyCBzyG8EDGjePFXbdJ5WkK30YNP"
)

func main() {

	tweetID, err := postTweet("Hello from Twitter API!")
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}
	fmt.Printf("Posted tweet with ID: %s\n", tweetID)

	time.Sleep(15 * time.Second)

	err = deleteTweet(tweetID)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)
	}
	fmt.Printf("Deleted tweet with ID: %s\n", tweetID)
}

func postTweet(content string) (string, error) {

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	tweetData := map[string]interface{}{
		"text": content,
	}
	tweetJSON, err := json.Marshal(tweetData)
	if err != nil {
		return "", fmt.Errorf("error marshaling tweet content: %w", err)
	}

	response, err := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetJSON))
	if err != nil {
		return "", fmt.Errorf("failed to post tweet: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		var result map[string]interface{}
		body, _ := io.ReadAll(response.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %w", err)
		}

		return result["data"].(map[string]interface{})["id"].(string), nil
	}

	body, _ := io.ReadAll(response.Body)
	return "", fmt.Errorf("failed to post tweet: %s, response: %s", response.Status, string(body))
}

func deleteTweet(tweetID string) error {

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute delete request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("failed to delete tweet: %s, response: %s", response.Status, string(body))
	}

	return nil
}
