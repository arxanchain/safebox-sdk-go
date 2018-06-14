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
	"github.com/arxanchain/sdk-go-common/structs/did"
	"github.com/arxanchain/sdk-go-common/structs/safebox"
)

// UpdateAssistCode is used to update assist code.
//
// API-Key must set to header.
func (s *SafeboxClient) UpdateAssistCode(header http.Header, body *safebox.UpdateSecurityCodeRequestBody) error {
	if body == nil {
		err := fmt.Errorf("request payload is null")
		return err
	}

	// Build http request
	r := s.c.NewRequest("POST", "/v1/code/update")
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
		return fmt.Errorf("update code error")
	}

	return nil
}

// RecoverAssistCode is used to recover assist code when user has forgot.
//
// API-Key must set to header.
func (s *SafeboxClient) RecoverAssistCode(header http.Header, id did.Identifier) (result *safebox.CodeInfoReply, err error) {
	if id == "" {
		err = fmt.Errorf("request information is empty")
		return
	}

	// Build http request
	r := s.c.NewRequest("GET", "/v1/code")
	r.SetHeaders(header)
	r.SetParam("user_did", string(id))

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
