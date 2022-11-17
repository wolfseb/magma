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
import binascii
import ipaddress
import logging
import socket

import dpkt
import netifaces
from lte.protos.mobilityd_pb2 import IPAddress
from magma.pipelined.app.packet_parser import ParseSocketPacket
from magma.pipelined.ifaces import get_mac_address_from_iface
from scapy.arch import get_if_addr
from scapy.data import ETH_P_ALL, ETHER_BROADCAST
from scapy.error import Scapy_Exception
from scapy.layers.inet6 import getmacbyip6
from scapy.layers.l2 import ARP, Dot1Q, Ether
from scapy.sendrecv import srp1


def get_gw_mac_address(ip: IPAddress, vlan: str, non_nat_arp_egress_port: str) -> str:
    gw_ip = str(ipaddress.ip_address(ip.address))
    if ip.version == IPAddress.IPV4:
        return _get_gw_mac_address_v4(gw_ip, vlan, non_nat_arp_egress_port)
    elif ip.version == IPAddress.IPV6:
        if vlan == "NO_VLAN":
            return _get_gw_mac_address_v6(gw_ip)
        logging.error("Not supported: GW IPv6: %s over vlan %d", gw_ip, vlan)
    return ""


def get_iface_by_ip4(target_ip4):
    for iface in netifaces.interfaces():
        iface_ip4 = netifaces.ifaddresses(iface)[netifaces.AF_INET][0]['addr']
        netmask = netifaces.ifaddresses(iface)[netifaces.AF_INET][0]['netmask']
        res_iface = ipaddress.ip_network(f'{iface_ip4}/{netmask}', strict=False)
        res_target_ip4 = ipaddress.ip_network(f'{target_ip4}/{netmask}', strict=False)
        if res_iface == res_target_ip4:
            return iface
    return None


def get_mac_by_ip4(target_ip4):
    iface = get_iface_by_ip4(target_ip4)
    if iface:
        return _get_gw_mac_address_v4(
            gw_ip=target_ip4,
            vlan="NO_VLAN",
            non_nat_arp_egress_port=iface,
        )


def _get_gw_mac_address_v4(gw_ip: IPAddress, vlan: str, non_nat_arp_egress_port: str) -> str:
    try:
        logging.debug(
            "sending arp via egress: %s",
            non_nat_arp_egress_port,
        )
        eth_mac_src, psrc = _get_addresses(non_nat_arp_egress_port)
        pkt = _make_arp_packet(eth_mac_src, psrc, gw_ip, vlan)
        logging.debug("ARP Req pkt:\n%s", pkt.pprint())

        res = _send_packet_and_receive_response(pkt, non_nat_arp_egress_port, vlan)
        if res is None:
            logging.debug("Got Null response")
            return ""

        parsed = ParseSocketPacket(res)
        logging.debug("ARP Res pkt %s", str(parsed))
        if str(parsed.arp.psrc) != str(gw_ip):
            logging.warning(
                f"Unexpected IP in ARP response. expected: {str(gw_ip)} pkt: {str(parsed)}",
            )
            return ""
        if vlan.isdigit():
            if parsed.dot1q is not None and str(parsed.dot1q.vlan) == vlan:
                mac = parsed.arp.hwsrc
            else:
                logging.warning(
                    "Unexpected vlan in ARP response. expected: %s pkt: %s",
                    vlan,
                    str(parsed),
                )
                return ""
        else:
            mac = parsed.arp.hwsrc
        return mac.mac_address

    except ValueError:
        logging.warning(
            "Invalid GW Ip address: [%s] or vlan %s",
            gw_ip, vlan,
        )
        return ""


def _get_addresses(non_nat_arp_egress_port):
    eth_mac_src = get_mac_address_from_iface(non_nat_arp_egress_port)
    eth_mac_src = binascii.unhexlify(eth_mac_src.replace(':', ''))
    psrc = "0.0.0.0"
    egress_port_ip = netifaces.ifaddresses(non_nat_arp_egress_port)
    if netifaces.AF_INET in egress_port_ip:
        psrc = egress_port_ip[netifaces.AF_INET][0]['addr']
    return eth_mac_src, psrc


def _make_arp_packet(eth_mac_src, psrc, gw_ip, vlan):
    pkt = dpkt.arp.ARP(
        sha=eth_mac_src,
        spa=socket.inet_aton(psrc),
        tha=b'\x00' * 6,
        tpa=socket.inet_aton(str(gw_ip)),
        op=dpkt.arp.ARP_OP_REQUEST,
    )
    if vlan.isdigit():
        pkt = dpkt.ethernet.VLANtag8021Q(
            id=int(vlan), data=bytes(pkt), type=dpkt.ethernet.ETH_TYPE_ARP,
        )
        t = dpkt.ethernet.ETH_TYPE_8021Q
    else:
        t = dpkt.ethernet.ETH_TYPE_ARP
    pkt = dpkt.ethernet.Ethernet(
        dst=b'\xff' * 6, src=eth_mac_src, data=bytes(pkt), type=t,
    )
    return pkt


def _send_packet_and_receive_response(pkt, non_nat_arp_egress_port, vlan):
    buffsize = 2 ** 16
    sol_packet = 263
    packet_aux_data = 8
    with socket.socket(socket.AF_PACKET, socket.SOCK_RAW, socket.ntohs(0x0003)) as s:
        s.setsockopt(socket.SOL_SOCKET, socket.SO_RCVBUF, buffsize)
        if vlan.isdigit():
            s.setsockopt(sol_packet, packet_aux_data, 1)
            s.setsockopt(socket.SOL_SOCKET, socket.SO_MARK, 1)
        s.bind((non_nat_arp_egress_port, 0x0003))
        s.send(bytes(pkt))
        if vlan.isdigit():
            res, aux, _, _ = s.recvmsg(0xffff, socket.CMSG_LEN(4096))
            for cmsg_level, cmsg_type, cmsg_data in aux:
                if cmsg_level == sol_packet and cmsg_type == packet_aux_data:
                    # add VLAN tag after ethernet header
                    res = res[:12] + cmsg_data[-1:-5:-1] + res[12:]
        else:
            res = s.recv(0xffff)
    return res


def _get_gw_mac_address_v6(gw_ip: IPAddress) -> str:
    try:
        mac = getmacbyip6(gw_ip)
        logging.debug("Got mac %s for IP: %s", mac, gw_ip)
        return mac
    except Scapy_Exception as ex:
        logging.warning("Error in probing Mac address: err %s", ex)
        return ""
    except ValueError:
        logging.warning(
            "Invalid GW Ip address: [%s]",
            gw_ip,
        )
        return ""
