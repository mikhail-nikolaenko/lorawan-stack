// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package ttjsv2

// ClaimData contains information about the claim.
type ClaimData struct {
	HomeNetID string  `json:"homeNetID"`
	HomeNSID  *string `json:"homeNSID,omitempty"`
	Locked    bool    `json:"locked"`
}

// KEK contains the key encryption key.
type KEK struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

// ClaimRequest is the claim request message.
type ClaimRequest struct {
	OwnerToken           string  `json:"ownerToken"`
	HomeNetID            string  `json:"homeNetID"`
	HomeNSID             *string `json:"homeNSID,omitempty"`
	ASID                 string  `json:"asID"`
	RegenerateOwnerToken *bool   `json:"regenerateOwnerToken,omitempty"`
	Lock                 *bool   `json:"lock,omitempty"`
	KEK                  *KEK    `json:"kek,omitempty"`
}

// ErrorResponse is a message that may be returned by The Things Join Server in case of an error.
type ErrorResponse struct {
	Message string `json:"message"`
}

func boolValue(b bool) *bool {
	return &b
}

func stringValue(s string) *string {
	return &s
}
