package processor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rohit-jaisinghani/gin-gorm-rest/models"
)

// SearchForArtist is used to fetch the details from spotify api based on irfc
func SearchForArtist(token string, artistName string) []models.Artist {
	urll := "https://api.spotify.com/v1/search"
	query := fmt.Sprintf("?q=%s&type=track&limit=1", artistName)

	queryURL := urll + query
	authOptions := url.Values{}

	req, err := http.NewRequest("GET", queryURL, strings.NewReader(authOptions.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	var body models.TrackDetail
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fmt.Println("value of body---", body)
	var artistDetails []models.Artist

	for _, v := range body.Tracks.Items {
		var artistDetail models.Artist
		artistDetail.ISRC = v.ExternalIds.Isrc

		for _, ll := range v.Artists {

			artistDetail.ArtistName = ll.Name
			artistDetail.URI = ll.URI
			artistDetails = append(artistDetails, artistDetail)
		}

	}
	return artistDetails
}

// GetToken func is used to get the token from spotify client credential to use the apis.
func GetToken() string {
	godotenv.Load(".env")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	authOptions := url.Values{}
	authOptions.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(authOptions.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	var body struct {
		AccessToken string `json:"access_token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return ""
	}

	token := body.AccessToken
	return token
}
