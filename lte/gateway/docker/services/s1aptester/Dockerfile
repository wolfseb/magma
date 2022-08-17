################################################################################
# Copyright 2022 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
ARG OS_DIST=ubuntu
ARG OS_RELEASE=focal


FROM $OS_DIST:$OS_RELEASE

ARG APT_UTILS_VERSION=2.0.9
ARG BINUTILS_VERSION=2.34-6ubuntu1.3
ARG BUILD_ESSENTIAL_VERSION=12.8ubuntu1.1
ARG CODEGEN_VERSION=2.2.3
ARG DEBIAN_FRONTEND=noninteractive
ARG EVENTLET_VERSION=0.25.1-2ubuntu1.1
ARG GIT_VERSION=1:2.25.1-1ubuntu3.5
ARG JDK_VERSION=8u342-b07-0ubuntu1~20.04
ARG LIBIPERF0_VERSION=3.7-3
ARG LIBMYSQL_VERSION=8.0.30-0ubuntu0.20.04.2
ARG LIBPROTOBUF_VERSION=3.6.1.3-2ubuntu5
ARG LIBSYSTEMD_DEV_VERSION=245.4-4ubuntu3.17
ARG LSB_RELEASE_VERSION=11.1.0ubuntu2
ARG MVN_VERSION=3.5.4
ARG PIP_VERSION=20.0.2-5ubuntu1.6
ARG PKG_CONFIG_VERSION=0.29.1-0ubuntu4
ARG PROTOBUF_VERSION=3.6.1.3-2ubuntu5
ARG PYSTEMD_VERSION=0.7.0-2build1
ARG PYTHON_IS_PYTHON3_VERSION=3.8.2-4
ARG RUBY_VERSION=1:2.7+1
ARG RUBY_DEV_VERSION=1:2.7+1
ARG RSYNC_VERSION=3.1.3-8ubuntu0.3
ARG SSH_VERSION=1:8.2p1-4ubuntu0.5
ARG SSHPASS_VERSION=1.06-1
ARG SUDO_VERSION=1.8.31-1ubuntu1.2
ARG VENV_VERSION=3.8.2-0ubuntu2
ARG VIRTUALENV_VERSION=20.0.17-1ubuntu0.4
ARG WGET_VERSION=1.20.3-1ubuntu2

ENV VIRTUAL_ENV=/build/venv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"
ENV MAGMA_ROOT=/magma
ENV S1AP_TESTER_ROOT=/s1ap-tester
ENV S1AP_TESTER_SRC=/S1APTester
ENV PYTHON_BUILD=/build
ENV PIP_CACHE_HOME="~/.pipcache"
ENV TZ=Europe/Paris
ENV SWAGGER_CODEGEN_DIR=/var/tmp/codegen/modules/swagger-codegen-cli/target/
ENV SWAGGER_CODEGEN_JAR=${SWAGGER_CODEGEN_DIR}swagger-codegen-cli.jar
ENV S1AP_TESTER_CFG=$MAGMA_ROOT/lte/gateway/python/integ_tests/data/s1ap_tester_cfg
ENV GRPC_POLL_STRATEGY=poll

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  apt-utils=$APT_UTILS_VERSION \
  binutils=$BINUTILS_VERSION \
  build-essential=$BUILD_ESSENTIAL_VERSION \
  git=$GIT_VERSION \
  libiperf0=$LIBIPERF0_VERSION \
  libmysqlclient-dev=$LIBMYSQL_VERSION \
  libprotobuf-dev=$LIBPROTOBUF_VERSION \
  libsystemd-dev=$LIBSYSTEMD_DEV_VERSION \
  lsb-release=$LSB_RELEASE_VERSION \
  openjdk-8-jdk=$JDK_VERSION \
  pkg-config=$PKG_CONFIG_VERSION \
  python-is-python3=$PYTHON_IS_PYTHON3_VERSION \
  python3-eventlet=$EVENTLET_VERSION \
  python3-pip=$PIP_VERSION \
  python3-protobuf=$PROTOBUF_VERSION \
  python3-pystemd=$PYSTEMD_VERSION \
  python3-venv=$VENV_VERSION \
  rsync=$RSYNC_VERSION \
  ruby=$RUBY_VERSION \
  ruby-dev=$RUBY_DEV_VERSION \
  ssh=$SSH_VERSION \
  sshpass=$SSHPASS_VERSION \
  sudo=$SUDO_VERSION \
  virtualenv=$VIRTUALENV_VERSION \
  wget=$WGET_VERSION && \
  apt-get clean && \
  gem install fpm \
