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

// ------------------------test TrusteeKeyPair---------------------------
func TestTrusteeKeypairSucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	payload := &safebox.SaveKeyPairReply{
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
		Post(trusteeURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.SaveKeyPairRequetBody{
		UserDid:    "did:anx:00001",
		PrivateKey: "privatekey",
		PublicKey:  "publckey",
	}

	resp, err := safeboxClient.TrusteeKeyPair(header, req)
	if err != nil {
		t.Fatalf("trutee key pair error, %v", err)
	}
	if resp == nil {
		t.Fatalf("trutee response is nil")
	}
	if resp.Code != "我是中国人" {
		t.Fatalf("trutee key pair return code error")
	}
}

func TestTrusteeKeypairFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoIsExist,
		ErrMessage: "user exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Post(trusteeURLPath).
		Reply(int(errors.UserInfoIsExist)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.SaveKeyPairRequetBody{
		UserDid:    "did:anx:00001",
		PrivateKey: "privatekey",
		PublicKey:  "publckey",
	}

	resp, err := safeboxClient.TrusteeKeyPair(header, req)
	if err == nil {
		t.Fatalf("trutee key pair error, %v", err)
	}
	if resp != nil {
		t.Fatalf("trutee response is error")
	}
}

func TestTrusteeKeypairPayloadErr(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
		Payload: 1,
	}
	//mock http response
	gock.New(safeboxURL).
		Post(trusteeURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.SaveKeyPairRequetBody{
		UserDid:    "did:anx:00001",
		PrivateKey: "privatekey",
		PublicKey:  "publckey",
	}

	resp, err := safeboxClient.TrusteeKeyPair(header, req)
	if err == nil {
		t.Fatalf("trutee key pair error, %v", err)
	}
	if resp != nil {
		t.Fatalf("trutee response is nil")
	}
}

func TestTrusteeKeypairBodyNil(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.TrusteeKeyPair(header, nil)
	if err == nil {
		t.Fatalf("trutee key pair error, %v", err)
	}
	if resp != nil {
		t.Fatalf("trutee response is error")
	}
}

// ------------------------test QueryPrivateKey---------------------------
func TestQueryPrivateKeySucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	payload := &safebox.PrivateKeyReply{
		PrivateKey: "privatekey",
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
		Get(privateURLPath).
		MatchParam("user_did", "did:anx:00001").
		MatchParam("code", "我是中国人").
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}
	resp, err := safeboxClient.QueryPrivateKey(header, req)
	if err != nil {
		t.Fatalf("get private error, %v", err)
	}
	if resp == nil {
		t.Fatalf("get private response is nil")
	}
	if resp.PrivateKey != "privatekey" {
		t.Fatalf("get private return code error")
	}
}

func TestQueryPrivateKeyFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoNotExit,
		ErrMessage: "user does not exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Get(privateURLPath).
		Reply(int(errors.UserInfoNotExit)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	resp, err := safeboxClient.QueryPrivateKey(header, req)
	if err == nil {
		t.Fatalf("query private key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query private response is error")
	}
}

func TestQueryPrivateKeyPayloadErr(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
		Payload: 1,
	}
	//mock http response
	gock.New(safeboxURL).
		Get(privateURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	resp, err := safeboxClient.QueryPrivateKey(header, req)
	if err == nil {
		t.Fatalf("query private key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query private response is nil")
	}
}

func TestQueryPrivateKeyBodyNil(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.QueryPrivateKey(header, nil)
	if err == nil {
		t.Fatalf("query private key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query private response is error")
	}
}

// ------------------------test QueryPublicKey---------------------------
func TestQueryPublicKeySucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	payload := &safebox.PublicKeyReply{
		PublicKey: "publickey",
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
		Get(publicURLPath).
		MatchParam("user_did", "did:anx:00001").
		MatchParam("code", "我是中国人").
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}
	resp, err := safeboxClient.QueryPublicKey(header, req)
	if err != nil {
		t.Fatalf("get public error, %v", err)
	}
	if resp == nil {
		t.Fatalf("get public response is nil")
	}
	if resp.PublicKey != "publickey" {
		t.Fatalf("get public return code error")
	}
}

func TestQueryPublicKeyFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoNotExit,
		ErrMessage: "user does not exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Get(publicURLPath).
		Reply(int(errors.UserInfoNotExit)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	resp, err := safeboxClient.QueryPublicKey(header, req)
	if err == nil {
		t.Fatalf("query public key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query public response is error")
	}
}

func TestQueryPublicKeyPayloadErr(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
		Payload: 1,
	}
	//mock http response
	gock.New(safeboxURL).
		Get(publicURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	resp, err := safeboxClient.QueryPublicKey(header, req)
	if err == nil {
		t.Fatalf("query public key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query public response is nil")
	}
}

func TestQueryPublicKeyBodyNil(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	resp, err := safeboxClient.QueryPublicKey(header, nil)
	if err == nil {
		t.Fatalf("query public key error, %v", err)
	}
	if resp != nil {
		t.Fatalf("query public response is error")
	}
}

// ------------------------test DeleteKeyPair---------------------------
func TestDeleteKeyPairSucc(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode: 0,
	}
	//mock http response
	gock.New(safeboxURL).
		Post(deleteURLPath).
		Reply(http.StatusOK).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	err := safeboxClient.DeleteKeyPair(header, req)
	if err != nil {
		t.Fatalf("delete key pair error, %v", err)
	}
}

func TestDeleteKeyPairFail(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	respBody := &rtstructs.Response{
		ErrCode:    errors.UserInfoNotExit,
		ErrMessage: "user does not exist",
	}
	//mock http response
	gock.New(safeboxURL).
		Post(deleteURLPath).
		Reply(int(errors.UserInfoNotExit)).
		JSON(respBody)

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	req := &safebox.OperateKeyInfo{
		UserDid: "did:anx:00001",
		Code:    "我是中国人",
	}

	err := safeboxClient.DeleteKeyPair(header, req)
	if err == nil {
		t.Fatalf("delete key pair error, %v", err)
	}
}

func TestDeleteKeyPairBodyNil(t *testing.T) {
	initTestSafeboxClient(t)
	defer gock.Off()

	header := http.Header{}
	header.Set(structs.APIKeyHeader, apiKey)

	err := safeboxClient.DeleteKeyPair(header, nil)
	if err == nil {
		t.Fatalf("delete key pair error, %v", err)
	}
}
