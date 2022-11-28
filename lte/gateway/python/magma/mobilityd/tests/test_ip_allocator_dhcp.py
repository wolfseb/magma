"""
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import time
from datetime import datetime, timedelta
from unittest.mock import patch, MagicMock
import unittest

from magma.mobilityd.dhcp_desc import DHCPDescriptor, DHCPState
from magma.mobilityd.ip_allocator_dhcp import IPAllocatorDHCP, DHCP_CLI_HELPER_PATH

MAC = "01:23:45:67:89:ab"
IP = "1.2.3.4"
VLAN = "0"
LEASE_EXPIRATION_TIME = 10


class TestIPAllocator(unittest.TestCase):
    def setUp(self) -> None:
        self.dhcp_desc = DHCPDescriptor(
            mac=MAC,
            ip=IP,
            vlan=VLAN,
            state=DHCPState.ACK,
            state_requested=DHCPState.REQUEST,
            lease_expiration_time=4,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        )
        # self.dhcp_desc_1 = DHCPDescriptor(
        #     mac=MAC,
        #     ip=IP,
        #     vlan=VLAN,
        #     state=DHCPState.ACK,
        #     state_requested=DHCPState.REQUEST,
        #     lease_expiration_time=8,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        # )
        # self.dhcp_desc_2 = DHCPDescriptor(
        #     mac=MAC,
        #     ip=IP,
        #     vlan=VLAN,
        #     state=DHCPState.ACK,
        #     state_requested=DHCPState.REQUEST,
        #     lease_expiration_time=12,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        # )
        self.store = MagicMock()
        self.store.dhcp_store = MagicMock()
        self.store.dhcp_store.values.return_value = [self.dhcp_desc] #, self.dhcp_desc_1, self.dhcp_desc_2]
        self.ip_alloc_dhcp = IPAllocatorDHCP(
            store=self.store,
            lease_renew_wait_min=4,
        )

    def tearDown(self) -> None:
        del self.dhcp_desc
        del self.store
        self.ip_alloc_dhcp._monitor_thread_event.set()
        self.ip_alloc_dhcp._monitor_thread.join()
        del self.ip_alloc_dhcp

    # def test_dhcp_store(self):
    #     assert self.ip_alloc_dhcp._store.dhcp_store.values() == [self.dhcp_desc]

    @patch("subprocess.run")
    def test_no_renewal_of_ip(self, mock_run):
        ret = MagicMock()
        ret.returncode = 0
        ret.stdout = """{"lease_expiration_time": 4}"""
        mock_run.return_value = ret
        time.sleep(1.0)
        mock_run.assert_not_called()

    @patch("subprocess.run")
    def test_renewal_of_ip(self, mock_run):
        ret = MagicMock()
        ret.returncode = 0
        ret.stdout = """{"lease_expiration_time": 4}"""
        mock_run.return_value = ret
        time.sleep(3.0)
        mock_run.assert_called_once()
        mock_run.assert_called_with([
            DHCP_CLI_HELPER_PATH,
            "--mac", str(self.dhcp_desc.mac),
            "--vlan", str(self.dhcp_desc.vlan),
            "--interface", self.ip_alloc_dhcp._iface,
            "--json",
            "renew",
            "--ip", str(self.dhcp_desc.ip),
            "--server-ip", str(self.dhcp_desc.server_ip),
        ],
            capture_output=True
        )

    @patch("subprocess.run")
    def test_allocate_ip_after_expiry(self, mock_run):
        ret = MagicMock()
        ret.returncode = 0
        ret.stdout = """{"lease_expiration_time": 4}"""
        mock_run.return_value = ret
        time.sleep(7.0)
        assert mock_run.call_count == 2

        mock_run.assert_called_with([
            DHCP_CLI_HELPER_PATH,
            "--mac", str(self.dhcp_desc.mac),
            "--vlan", str(self.dhcp_desc.vlan),
            "--interface", self.ip_alloc_dhcp._iface,
            "--json",
            "allocate",
        ],
            capture_output=True
        )

    # def test_renew_lease_after_renew_deadline(self):
    #     """
    #     Test that lease is renewed if lease is expired.
    #     """
    #     self.dhcp_desc.lease_expiration_time = 0
    #     with patch.object(self.dhcp_desc, 'renew_lease') as mock_renew_lease:
    #         self.dhcp_desc.update_lease()
    #         mock_renew_lease.assert_called_once()


