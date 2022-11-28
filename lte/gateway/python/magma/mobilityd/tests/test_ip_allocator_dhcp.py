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
from ipaddress import IPv4Network
from unittest.mock import patch, MagicMock
import unittest

from magma.mobilityd.dhcp_desc import DHCPDescriptor, DHCPState
from magma.mobilityd.ip_allocator_dhcp import IPAllocatorDHCP, DHCP_CLI_HELPER_PATH
from magma.mobilityd.ip_descriptor import IPState
from magma.mobilityd.ip_descriptor_map import IpDescriptorMap


MAC = "01:23:45:67:89:ab"
MAC2 = "01:23:45:67:89:cd"
IP = "1.2.3.4"
IP2 = "1.2.3.5"
VLAN = "0"
LEASE_EXPIRATION_TIME = 10


class TestMonitorDhcpState(unittest.TestCase):
    def setUp(self) -> None:
        self.dhcp_desc = DHCPDescriptor(
            mac=MAC,
            ip=IP,
            vlan=VLAN,
            state=DHCPState.ACK,
            state_requested=DHCPState.REQUEST,
            lease_expiration_time=4,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        )
        self.store = MagicMock()
        self.store.dhcp_store = MagicMock()
        self.store.dhcp_store.values.return_value = [self.dhcp_desc]
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

    # @patch("subprocess.run")
    # def test_no_renewal_of_ip(self, mock_run):
    #     ret = MagicMock()
    #     ret.returncode = 0
    #     ret.stdout = """{"lease_expiration_time": 4}"""
    #     mock_run.return_value = ret
    #     time.sleep(1.0)
    #     mock_run.assert_not_called()
    #
    # @patch("subprocess.run")
    # def test_renewal_of_ip(self, mock_run):
    #     ret = MagicMock()
    #     ret.returncode = 0
    #     ret.stdout = """{"lease_expiration_time": 4}"""
    #     mock_run.return_value = ret
    #     time.sleep(3.0)
    #     mock_run.assert_called_once()
    #     mock_run.assert_called_with([
    #         DHCP_CLI_HELPER_PATH,
    #         "--mac", str(self.dhcp_desc.mac),
    #         "--vlan", str(self.dhcp_desc.vlan),
    #         "--interface", self.ip_alloc_dhcp._iface,
    #         "--json",
    #         "renew",
    #         "--ip", str(self.dhcp_desc.ip),
    #         "--server-ip", str(self.dhcp_desc.server_ip),
    #     ],
    #         capture_output=True
    #     )
    #
    # @patch("subprocess.run")
    # def test_allocate_ip_after_expiry(self, mock_run):
    #     ret = MagicMock()
    #     ret.returncode = 0
    #     ret.stdout = """{"lease_expiration_time": 4}"""
    #     mock_run.return_value = ret
    #     time.sleep(7.0)
    #     assert mock_run.call_count == 2
    #
    #     mock_run.assert_called_with([
    #         DHCP_CLI_HELPER_PATH,
    #         "--mac", str(self.dhcp_desc.mac),
    #         "--vlan", str(self.dhcp_desc.vlan),
    #         "--interface", self.ip_alloc_dhcp._iface,
    #         "--json",
    #         "allocate",
    #     ],
    #         capture_output=True
    #     )


class TestRemoveIpBlocks(unittest.TestCase):
    def setUp(self) -> None:
        self.dhcp_desc_expired = DHCPDescriptor(
            mac=MAC,
            ip=IP,
            vlan=VLAN,
            state=DHCPState.ACK,
            state_requested=DHCPState.REQUEST,
            lease_expiration_time=-2,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        )
        self.dhcp_desc_used = DHCPDescriptor(
            mac=MAC2,
            ip=IP2,
            vlan=VLAN,
            state=DHCPState.ACK,
            state_requested=DHCPState.REQUEST,
            lease_expiration_time=100,  # datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        )
        self.store = MagicMock()
        self.store.dhcp_store = MagicMock()
        self.store.dhcp_store.values.return_value = [
            self.dhcp_desc_expired,
            self.dhcp_desc_used
        ]
        self.ip_alloc_dhcp = IPAllocatorDHCP(
            store=self.store,
            lease_renew_wait_min=4,
        )

    def tearDown(self) -> None:
        del self.dhcp_desc_expired
        del self.dhcp_desc_used
        del self.store
        self.ip_alloc_dhcp._monitor_thread_event.set()
        self.ip_alloc_dhcp._monitor_thread.join()
        del self.ip_alloc_dhcp

    @patch("subprocess.run")
    def test_remove_one_ip_block(self, mock_run):
        ret = MagicMock()
        ret.returncode = 0
        mock_run.return_value = ret
        self.store.assigned_ip_blocks = {IP, IP2}
        self.store.ip_state_map = IpDescriptorMap({IP: IPState.ALLOCATED, IP2: IPState.ALLOCATED})

        removed_block = self.ip_alloc_dhcp.remove_ip_blocks(IPv4Network(IP))
        assert removed_block == IPv4Network(IP)
        assert self.store.assigned_ip_blocks == {IP2}
