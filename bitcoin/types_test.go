// Copyright 2021 Rosetta Dogecoin Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Block_UnmarshalJSON(t *testing.T) {
	transactionJSON := loadTypeFixture("transaction.json")

	tests := map[string]struct {
		json        string
		expectedErr error
	}{
		"zero transactions": {
			json:        `{ "tx": [] }`,
			expectedErr: errors.New("expected >= 1 transactions in block, got 0"),
		},
		"one transaction": {
			json: fmt.Sprintf(`{ "tx": [%s] }`, transactionJSON),
		},
		"more than one transaction": {
			json: fmt.Sprintf(`{ "tx": [%s, %s] }`, transactionJSON, transactionJSON),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var block Block
			err := json.Unmarshal([]byte(test.json), &block)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// loadTypeFixture takes a file name and returns the type fixture.
func loadTypeFixture(fileName string) string {
	content, err := ioutil.ReadFile(fmt.Sprintf("type_fixtures/%s", fileName))
	if err != nil {
		log.Fatalf("Failed to load type fixture with file name '%s': %v", fileName, err)
	}
	return string(content)
}
