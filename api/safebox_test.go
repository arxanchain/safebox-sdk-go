/*
Copyright ArxanFintech Technology Ltd. 2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"net/http"
	"testing"

	"github.com/arxanchain/sdk-go-common/rest/api"
	gock "gopkg.in/h2non/gock.v1"
)

const (
	safeboxURL = "http://127.0.0.1:8014"
	apiKey     = "1234567890"

	trusteeURLPath = "/v1/keypair/save"
	privateURLPath = "/v1/keypair/private"
	publicURLPath  = "/v1/keypair/public"
	deleteURLPath  = "/v1/keypair/delete"

	updateCodeURLPath  = "/v1/code/update"
	recoverCodeURLPath = "/v1/code"
)

var (
	safeboxClient *SafeboxClient
)

func initTestSafeboxClient(t *testing.T) {
	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)
	var err error
	safeboxClient, err = NewSafeboxClient(&api.Config{Address: safeboxURL, HttpClient: client})
	if err != nil {
		t.Fatalf("New safebox client fail: %v", err)
	}
}

func TestNewSafeboxClientSucc(t *testing.T) {
	initTestSafeboxClient(t)
}

func TestNewSafeboxNilConf(t *testing.T) {
	var err error
	safeboxClient, err = NewSafeboxClient(nil)
	if err == nil {
		t.Fatalf("New safebox client fail: %v", err)
	}
}
