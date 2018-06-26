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
	"fmt"
	"net/http"
	"reflect"

	"github.com/arxanchain/sdk-go-common/errors"
	"github.com/arxanchain/sdk-go-common/rest"
	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	reststruct "github.com/arxanchain/sdk-go-common/rest/structs"
	safebox "github.com/arxanchain/sdk-go-common/structs/safebox"
)

// TrusteeKeyPair is used to trutee keypair.
//
// API-Key must set to header.
func (s *SafeboxClient) TrusteeKeyPair(header http.Header, body *safebox.SaveKeyPairRequetBody) (result *safebox.SaveKeyPairReply, err error) {
	if body == nil {
		err = fmt.Errorf("request payload is null")
		return
	}

	// Build http request
	r := s.c.NewRequest("POST", "/v1/keypair/save")
	r.SetHeaders(header)
	r.SetBody(body)

	// Do http request
	_, resp, err := restapi.RequireOK(s.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Parse http response
	var respBody reststruct.Response
	if err = restapi.DecodeBody(resp, &respBody); err != nil {
		return
	}

	if respBody.ErrCode != errors.SuccCode {
		err = rest.CodedError(respBody.ErrCode, respBody.ErrMessage)
		return
	}

	payload, ok := respBody.Payload.(string)
	if !ok {
		err = fmt.Errorf("response payload type invalid: %v", reflect.TypeOf(respBody.Payload))
		return
	}

	err = json.Unmarshal([]byte(payload), &result)

	return
}

// QueryPrivateKey is used to query private key.
//
// API-Key must set to header.
func (s *SafeboxClient) QueryPrivateKey(header http.Header, info *safebox.OperateKeyInfo) (result *safebox.PrivateKeyReply, err error) {
	if info == nil {
		err = fmt.Errorf("request information is nil")
		return
	}

	// Build http request
	r := s.c.NewRequest("GET", "/v1/keypair/private")
	r.SetHeaders(header)
	r.SetParam("user_did", info.UserDid)
	r.SetParam("code", info.Code)

	// Do http request
	_, resp, err := restapi.RequireOK(s.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Parse http response
	var respBody reststruct.Response
	if err = restapi.DecodeBody(resp, &respBody); err != nil {
		return
	}

	if respBody.ErrCode != errors.SuccCode {
		err = rest.CodedError(respBody.ErrCode, respBody.ErrMessage)
		return
	}

	payload, ok := respBody.Payload.(string)
	if !ok {
		err = fmt.Errorf("response payload type invalid: %v", reflect.TypeOf(respBody.Payload))
		return
	}

	err = json.Unmarshal([]byte(payload), &result)
	return
}

// QueryPublicKey is used to query public key.
//
// API-Key must set to header.
func (s *SafeboxClient) QueryPublicKey(header http.Header, info *safebox.OperateKeyInfo) (result *safebox.PublicKeyReply, err error) {
	if info == nil {
		err = fmt.Errorf("request information is nil")
		return
	}

	// Build http request
	r := s.c.NewRequest("GET", "/v1/keypair/public")
	r.SetHeaders(header)
	r.SetParam("user_did", info.UserDid)
	r.SetParam("code", info.Code)

	// Do http request
	_, resp, err := restapi.RequireOK(s.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Parse http response
	var respBody reststruct.Response
	if err = restapi.DecodeBody(resp, &respBody); err != nil {
		return
	}

	if respBody.ErrCode != errors.SuccCode {
		err = rest.CodedError(respBody.ErrCode, respBody.ErrMessage)
		return
	}

	payload, ok := respBody.Payload.(string)
	if !ok {
		err = fmt.Errorf("response payload type invalid: %v", reflect.TypeOf(respBody.Payload))
		return
	}

	err = json.Unmarshal([]byte(payload), &result)
	return
}

// DeleteKeyPair is used to delete keypair.
//
// API-Key must set to header.
func (s *SafeboxClient) DeleteKeyPair(header http.Header, body *safebox.OperateKeyInfo) error {
	if body == nil {
		err := fmt.Errorf("request payload is nil")
		return err
	}

	// Build http request
	r := s.c.NewRequest("POST", "/v1/keypair/delete")
	r.SetHeaders(header)
	r.SetBody(body)

	// Do http request
	_, resp, err := restapi.RequireOK(s.c.DoRequest(r))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse http response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete key pair error")
	}

	return nil
}
