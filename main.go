package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func measureTime(function func()) time.Duration {
	start := time.Now()
	function()
	end := time.Now()
	return end.Sub(start)
}

func generateImage() {
	// Define the endpoint URL
	const (
		url = "https://api.openai.com/v1/images/generations"
	)
	// Fetch the APIKEY secret
	apiKey := os.Getenv("APIKEY")
	// Creating a list of random prompts for image generation
	var prompts = []string{
		"A robot that can do laundry",
		"A painting of a sunset over a city",
		"A futuristic car",
		"A magical creature",
		"A person playing a musical instrument",
		"A dish from a 5-star restaurant",
		"A scene from a fantasy world",
		"A piece of technology from the year 2050",
		"A garden with exotic flowers",
		"A superhero fighting a villain",
		"A futuristic city skyline at night",
		"A close-up of a lion's face",
		"A group of friends having a picnic in a park",
		"A hot air balloon flying over a mountain range",
		"A robot playing a guitar on stage",
		"A tree with vibrant fall colors",
		"A person scuba diving in crystal clear water",
		"A snowy cabin in the woods",
		"A street market in a bustling city",
		"A giant wave about to crash on a beach",
	}

	// Define the request body data
	rand.Seed(time.Now().UnixNano())
	// Choose a random prompt from the list
	prompt := prompts[rand.Intn(len(prompts))]
	reqBodyData := map[string]interface{}{
		"prompt":     prompt,
		"model":      "image-alpha-001",
		"num_images": 1,
	}
	fmt.Printf("\nGenerating an image of a '%s'\n", prompt)

	// Marshal the request body data to json
	reqBody, _ := json.Marshal(reqBodyData)

	// Create a new http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	// Add headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error while closing the request body: %s", err)
		}
	}(resp.Body)

	// Read the response
	respBody, _ := io.ReadAll(resp.Body)

	// Unmarshal the response
	var data map[string]interface{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		fmt.Printf("Error while reading response data: %s", err)
	}

	// Extract the image URL from the response
	imageURL := data["data"].([]interface{})[0].(map[string]interface{})["url"].(string)
	fmt.Println("Here's the Image URL: ", imageURL)

	// Download the image
	response, err := http.Get(imageURL)
	if err != nil {
		fmt.Printf("Error while downloading image: %s\n", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error while closing the response body: %s", err)
		}
	}(response.Body)

	// Create a file to save the image
	// Get the current time
	//now := time.Now()
	// Format the time as a string
	//timestamp := now.Format("20060102150405")
	// Generate a new image name
	//imageName := "images/image_" + timestamp + ".jpg"
	//file, err := os.Create(fmt.Sprintf("%s", imageName))
	//if err != nil {
	//	fmt.Printf("Error while creating image file: %s\n", err.Error())
	//	return
	//}
	//defer file.Close()

	// Write the image to the file
	//_, err = io.Copy(file, response.Body)
	//if err != nil {
	//	fmt.Printf("Error while saving image: %s\n", err.Error())
	//	return
	//}
	//fmt.Println("Image saved successfully!")
}

func main() {
	var totalTime time.Duration = 0
	for i := 0; i < 7; i++ {
		timeTaken := measureTime(generateImage)
		totalTime += timeTaken
		fmt.Printf("Time taken to generate the image: %s\n", timeTaken)
	}
	fmt.Printf("\nTotal time taken to generate all images: %s\n", totalTime)
}
