#!/usr/bin/env bash
# Copyright 2022 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

AIOEVENTLET_TMP="/tmp/aioeventlet"
CODEGEN_VERSION="2.2.3"
MAGMA_ROOT="/magma"
PIP_CACHE_HOME="~/.pipcache"
S1AP_TESTER_ROOT="/s1ap-tester"
S1AP_TESTER_SRC="/S1APTester"
SWAGGER_CODEGEN_DIR="/var/tmp/codegen/modules/swagger-codegen-cli/target/"
SWAGGER_CODEGEN_JAR="${SWAGGER_CODEGEN_DIR}swagger-codegen-cli.jar"

mkdir -p ${AIOEVENTLET_TMP}
/magma/third_party/build/bin/aioeventlet_build.sh ${AIOEVENTLET_TMP}
cd /var/tmp/
dpkg -i ${AIOEVENTLET_TMP}/python3-aioeventlet*
rm -rf ${AIOEVENTLET_TMP}
mkdir -p ${SWAGGER_CODEGEN_DIR}
wget -q https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/${CODEGEN_VERSION}/swagger-codegen-cli-${CODEGEN_VERSION}.jar -O ${SWAGGER_CODEGEN_JAR}

cd $MAGMA_ROOT/lte/gateway/python && make
mkdir -p $S1AP_TESTER_ROOT/bin /var/tmp/test_results/

cd integ_tests && make
rm -rf /var/lib/apt/lists/*
$MAGMA_ROOT/lte/gateway/deploy/roles/magma_test/files/clone_s1_tester.sh
$MAGMA_ROOT/lte/gateway/deploy/roles/magma_test/files/build_s1_tester.sh
pip3 install --cache-dir $PIP_CACHE_HOME pyparsing
python3 $MAGMA_ROOT/lte/gateway/deploy/roles/magma_test/files/c_parser.py
