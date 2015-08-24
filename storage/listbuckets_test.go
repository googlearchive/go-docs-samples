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
	"os"
	"testing"
)

func TestListBuckets(t *testing.T) {
	buckets, err := ListBuckets(os.Getenv("TEST_PROJECT_ID"))
	if err != nil {
		t.Errorf("Error while listing buckets: %s", err)
	}
	if len(buckets) <= 0 {
		t.Error("No bucket returned")
	}

	foundBucket := false
	expectedBucket := os.Getenv("TEST_BUCKET_NAME")
	for _, bucket := range buckets {
		if bucket.Name == expectedBucket {
			foundBucket = true
			break
		}
	}
	if !foundBucket {
		t.Errorf("Expected bucket %s", expectedBucket)
	}

}
