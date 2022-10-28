/*
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package modules

import (
	"context"
	"errors"
	"fmt"

	"layeh.com/radius"
)

func WaitForRadiusServerToBeReady(secret []byte, port string) (err error) {
	MaxRetries := 10
	for r := 0; r < MaxRetries; r++ {
		_, err = radius.Exchange(context.Background(), radius.New(radius.CodeStatusServer, secret), port)
		if err == nil {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("radius server failed to be ready after %d retries: %v", MaxRetries, err))
}
