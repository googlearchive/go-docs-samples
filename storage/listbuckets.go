// Copyright 2015 Google, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
)
import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

// ListBuckets returns a slice of all the buckets in the given projectId.
// [START ListBuckets]
func ListBuckets(projectId string) ([]*storage.Bucket, error) {
	// Create the client that uses Application Default Credentials
	client, err := google.DefaultClient(
		oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.read_only")
	if err != nil {
		return nil, err
	}

	// Create the Google Cloud Storage service
	service, err := storage.New(client)
	if err != nil {
		return nil, err
	}

	// Create the request to list buckets for the project id
	request := service.Buckets.List(projectId)

	// Execute the request
	buckets, err := request.Do()
	if err != nil {
		return nil, err
	}

	return buckets.Items, nil
}

// [END ListBuckets]

// main will simply retrieve a list of buckets and print them.
func main() {
	buckets, err := ListBuckets(os.Getenv("TEST_PROJECT_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print out the results
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}
