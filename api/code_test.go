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
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arxanchain/sdk-go-common/errors"
	rtstructs "github.com/arxanchain/sdk-go-common/rest/structs"
	"github.com/arxanchain/sdk-go-common/structs"
	"github.com/arxanchain/sdk-go-common/structs/safebox"
	gock "gopkg.in/h2non/gock.v1"
)

// ------------------------test UpdateAssistCode---------------------------
func TestUpdateAssistCodeSucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
	}
	//mock http response
	gock.New(safeboxURL).
		Post(updateCodeURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.UpdateSecurityCodeRequestBody{
		UserDid:      "did:anx:00001",
		OriginalCode: "我是中国人",
		NewCode:      "我爱你中国",
	}

	err := safeboxClient.UpdateAssistCode(header, req)
	if err != nil {
		t.Fatalf("update code error, %v", err)
	}
}

func TestUpdateAssistCodeFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoNotExit,
		ErrMessage: "user does not exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Post(updateCodeURLPath).
		Reply(int(errors.UserInfoNotExit)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.UpdateSecurityCodeRequestBody{
		UserDid:      "did:anx:00001",
		OriginalCode: "我是中国人",
		NewCode:      "我爱你中国",
	}

	err := safeboxClient.UpdateAssistCode(header, req)
	if err == nil {
		t.Fatalf("update code error, %v", err)
	}
}

func TestUpdateAssistCodeBodyNil(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	err := safeboxClient.UpdateAssistCode(header, nil)
	if err == nil {
		t.Fatalf("update code error, %v", err)
	}
}

// ------------------------test RecoverCode---------------------------
func TestRecoverCodeSucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	payload := &safebox.CodeInfoReply{
		Code: "我是中国人",
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}
	respBody := &rtstructs.Response{
		ErrCode: 0,
		Payload: string(byPayload),
	}
	//mock http response
	gock.New(safeboxURL).
		Get(recoverCodeURLPath).
		MatchParam("user_did", "did:anx:00001").
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.RecoverAssistCode(header, "did:anx:00001")
	if err != nil {
		t.Fatalf("recover code error, %v", err)
	}
	if resp == nil {
		t.Fatalf("recover code response is nil")
	}
	if resp.Code != "我是中国人" {
		t.Fatalf("recover code return code error")
	}
}

func TestRecoverCodeIdEmpty(t *testing.T) {
	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.RecoverAssistCode(header, "")
	if err == nil {
		t.Fatalf("recover code error, %v", err)
	}
	if resp != nil {
		t.Fatalf("recover code response error")
	}
}

func TestRecoverAssistCodeFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoNotExit,
		ErrMessage: "user does not exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Get(recoverCodeURLPath).
		Reply(int(errors.UserInfoNotExit)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.RecoverAssistCode(header, "did:anx:00001")
	if err == nil {
		t.Fatalf("recover code error, %v", err)
	}
	if resp != nil {
		t.Fatalf("recover code response is error")
	}
}

func TestRecoverAssistCodePayloadErr(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
		Payload: 1,
	}
	//mock http response
	gock.New(safeboxURL).
		Get(recoverCodeURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.RecoverAssistCode(header, "did:anx:00001")
	if err == nil {
		t.Fatalf("recover code error, %v", err)
	}
	if resp != nil {
		t.Fatalf("recover code response is nil")
	}
}
