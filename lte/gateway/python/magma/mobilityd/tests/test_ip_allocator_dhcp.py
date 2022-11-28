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
from datetime import datetime, timedelta
from unittest.mock import patch, MagicMock
import unittest

from magma.mobilityd.dhcp_desc import DHCPDescriptor, DHCPState
from magma.mobilityd.ip_allocator_dhcp import IPAllocatorDHCP

MAC = "01:23:45:67:89:ab"
IP = "1.2.3.4"
VLAN = "0"
LEASE_EXPIRATION_TIME = 10


class TestDHCPDescriptor(unittest.TestCase):
    def setUp(self) -> None:
        self.dhcp_desc = DHCPDescriptor(
            mac=MAC,
            ip=IP,
            vlan=VLAN,
            state=DHCPState.ACK,
            state_requested=DHCPState.REQUEST,
            lease_expiration_time=datetime.now() + timedelta(seconds=LEASE_EXPIRATION_TIME),
        )
        self.store = MagicMock()
        self.store.dhcp_store = MagicMock()
        self.store.dhcp_store.values.return_value = [self.dhcp_desc]
        self.ip_alloc_dhcp = IPAllocatorDHCP(
            store=self.store,
        )

    def tearDown(self) -> None:
        del self.dhcp_desc
        del self.store
        del self.ip_alloc_dhcp

    def test_dhcp_store(self):
        assert self.ip_alloc_dhcp._store.dhcp_store.values() == [self.dhcp_desc]

    def test_no_renewal_of_ip(self):
        assert self.ip_alloc_dhcp.

    # def test_renew_lease_after_renew_deadline(self):
    #     """
    #     Test that lease is renewed if lease is expired.
    #     """
    #     self.dhcp_desc.lease_expiration_time = 0
    #     with patch.object(self.dhcp_desc, 'renew_lease') as mock_renew_lease:
    #         self.dhcp_desc.update_lease()
    #         mock_renew_lease.assert_called_once()


