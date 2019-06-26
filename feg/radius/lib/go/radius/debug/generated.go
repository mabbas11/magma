/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Code generated by generate_main.go. DO NOT EDIT.

package debug

import "fbc/lib/go/radius/dictionary"

var IncludedDictionary = &dictionary.Dictionary{Attributes: []*dictionary.Attribute{&dictionary.Attribute{Name: "User-Name", OID: "1", Type: 1}, &dictionary.Attribute{Name: "User-Password", OID: "2", Type: 1, FlagEncrypt: dictionary.Int(1)}, &dictionary.Attribute{Name: "CHAP-Password", OID: "3", Type: 2}, &dictionary.Attribute{Name: "NAS-IP-Address", OID: "4", Type: 3}, &dictionary.Attribute{Name: "NAS-Port", OID: "5", Type: 5}, &dictionary.Attribute{Name: "Service-Type", OID: "6", Type: 5}, &dictionary.Attribute{Name: "Framed-Protocol", OID: "7", Type: 5}, &dictionary.Attribute{Name: "Framed-IP-Address", OID: "8", Type: 3}, &dictionary.Attribute{Name: "Framed-IP-Netmask", OID: "9", Type: 3}, &dictionary.Attribute{Name: "Framed-Routing", OID: "10", Type: 5}, &dictionary.Attribute{Name: "Filter-Id", OID: "11", Type: 1}, &dictionary.Attribute{Name: "Framed-MTU", OID: "12", Type: 5}, &dictionary.Attribute{Name: "Framed-Compression", OID: "13", Type: 5}, &dictionary.Attribute{Name: "Login-IP-Host", OID: "14", Type: 3}, &dictionary.Attribute{Name: "Login-Service", OID: "15", Type: 5}, &dictionary.Attribute{Name: "Login-TCP-Port", OID: "16", Type: 5}, &dictionary.Attribute{Name: "Reply-Message", OID: "18", Type: 1}, &dictionary.Attribute{Name: "Callback-Number", OID: "19", Type: 1}, &dictionary.Attribute{Name: "Callback-Id", OID: "20", Type: 1}, &dictionary.Attribute{Name: "Framed-Route", OID: "22", Type: 1}, &dictionary.Attribute{Name: "Framed-IPX-Network", OID: "23", Type: 3}, &dictionary.Attribute{Name: "State", OID: "24", Type: 2}, &dictionary.Attribute{Name: "Class", OID: "25", Type: 2}, &dictionary.Attribute{Name: "Vendor-Specific", OID: "26", Type: 10}, &dictionary.Attribute{Name: "Session-Timeout", OID: "27", Type: 5}, &dictionary.Attribute{Name: "Idle-Timeout", OID: "28", Type: 5}, &dictionary.Attribute{Name: "Termination-Action", OID: "29", Type: 5}, &dictionary.Attribute{Name: "Called-Station-Id", OID: "30", Type: 1}, &dictionary.Attribute{Name: "Calling-Station-Id", OID: "31", Type: 1}, &dictionary.Attribute{Name: "NAS-Identifier", OID: "32", Type: 1}, &dictionary.Attribute{Name: "Proxy-State", OID: "33", Type: 2}, &dictionary.Attribute{Name: "Login-LAT-Service", OID: "34", Type: 1}, &dictionary.Attribute{Name: "Login-LAT-Node", OID: "35", Type: 1}, &dictionary.Attribute{Name: "Login-LAT-Group", OID: "36", Type: 2}, &dictionary.Attribute{Name: "Framed-AppleTalk-Link", OID: "37", Type: 5}, &dictionary.Attribute{Name: "Framed-AppleTalk-Network", OID: "38", Type: 5}, &dictionary.Attribute{Name: "Framed-AppleTalk-Zone", OID: "39", Type: 1}, &dictionary.Attribute{Name: "CHAP-Challenge", OID: "60", Type: 2}, &dictionary.Attribute{Name: "NAS-Port-Type", OID: "61", Type: 5}, &dictionary.Attribute{Name: "Port-Limit", OID: "62", Type: 5}, &dictionary.Attribute{Name: "Login-LAT-Port", OID: "63", Type: 1}, &dictionary.Attribute{Name: "Acct-Status-Type", OID: "40", Type: 5}, &dictionary.Attribute{Name: "Acct-Delay-Time", OID: "41", Type: 5}, &dictionary.Attribute{Name: "Acct-Input-Octets", OID: "42", Type: 5}, &dictionary.Attribute{Name: "Acct-Output-Octets", OID: "43", Type: 5}, &dictionary.Attribute{Name: "Acct-Session-Id", OID: "44", Type: 1}, &dictionary.Attribute{Name: "Acct-Authentic", OID: "45", Type: 5}, &dictionary.Attribute{Name: "Acct-Session-Time", OID: "46", Type: 5}, &dictionary.Attribute{Name: "Acct-Input-Packets", OID: "47", Type: 5}, &dictionary.Attribute{Name: "Acct-Output-Packets", OID: "48", Type: 5}, &dictionary.Attribute{Name: "Acct-Terminate-Cause", OID: "49", Type: 5}, &dictionary.Attribute{Name: "Acct-Multi-Session-Id", OID: "50", Type: 1}, &dictionary.Attribute{Name: "Acct-Link-Count", OID: "51", Type: 5}, &dictionary.Attribute{Name: "Acct-Tunnel-Connection", OID: "68", Type: 1}, &dictionary.Attribute{Name: "Acct-Tunnel-Packets-Lost", OID: "86", Type: 5}, &dictionary.Attribute{Name: "Acct-Input-Gigawords", OID: "52", Type: 5}, &dictionary.Attribute{Name: "Acct-Output-Gigawords", OID: "53", Type: 5}, &dictionary.Attribute{Name: "Event-Timestamp", OID: "55", Type: 4}, &dictionary.Attribute{Name: "ARAP-Password", OID: "70", Type: 2, Size: dictionary.Int(16)}, &dictionary.Attribute{Name: "ARAP-Features", OID: "71", Type: 2, Size: dictionary.Int(14)}, &dictionary.Attribute{Name: "ARAP-Zone-Access", OID: "72", Type: 5}, &dictionary.Attribute{Name: "ARAP-Security", OID: "73", Type: 5}, &dictionary.Attribute{Name: "ARAP-Security-Data", OID: "74", Type: 1}, &dictionary.Attribute{Name: "Password-Retry", OID: "75", Type: 5}, &dictionary.Attribute{Name: "Prompt", OID: "76", Type: 5}, &dictionary.Attribute{Name: "Connect-Info", OID: "77", Type: 1}, &dictionary.Attribute{Name: "Configuration-Token", OID: "78", Type: 1}, &dictionary.Attribute{Name: "EAP-Message", OID: "79", Type: 2, FlagConcat: dictionary.Bool(true)}, &dictionary.Attribute{Name: "Message-Authenticator", OID: "80", Type: 2}, &dictionary.Attribute{Name: "ARAP-Challenge-Response", OID: "84", Type: 2, Size: dictionary.Int(8)}, &dictionary.Attribute{Name: "Acct-Interim-Interval", OID: "85", Type: 5}, &dictionary.Attribute{Name: "NAS-Port-Id", OID: "87", Type: 1}, &dictionary.Attribute{Name: "Framed-Pool", OID: "88", Type: 1}, &dictionary.Attribute{Name: "NAS-IPv6-Address", OID: "95", Type: 6}, &dictionary.Attribute{Name: "Framed-Interface-Id", OID: "96", Type: 8}, &dictionary.Attribute{Name: "Framed-IPv6-Prefix", OID: "97", Type: 7}, &dictionary.Attribute{Name: "Login-IPv6-Host", OID: "98", Type: 6}, &dictionary.Attribute{Name: "Framed-IPv6-Route", OID: "99", Type: 1}, &dictionary.Attribute{Name: "Framed-IPv6-Pool", OID: "100", Type: 1}, &dictionary.Attribute{Name: "Error-Cause", OID: "101", Type: 5}}, Values: []*dictionary.Value{&dictionary.Value{Attribute: "Service-Type", Name: "Login-User", Number: 1}, &dictionary.Value{Attribute: "Service-Type", Name: "Framed-User", Number: 2}, &dictionary.Value{Attribute: "Service-Type", Name: "Callback-Login-User", Number: 3}, &dictionary.Value{Attribute: "Service-Type", Name: "Callback-Framed-User", Number: 4}, &dictionary.Value{Attribute: "Service-Type", Name: "Outbound-User", Number: 5}, &dictionary.Value{Attribute: "Service-Type", Name: "Administrative-User", Number: 6}, &dictionary.Value{Attribute: "Service-Type", Name: "NAS-Prompt-User", Number: 7}, &dictionary.Value{Attribute: "Service-Type", Name: "Authenticate-Only", Number: 8}, &dictionary.Value{Attribute: "Service-Type", Name: "Callback-NAS-Prompt", Number: 9}, &dictionary.Value{Attribute: "Service-Type", Name: "Call-Check", Number: 10}, &dictionary.Value{Attribute: "Service-Type", Name: "Callback-Administrative", Number: 11}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "PPP", Number: 1}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "SLIP", Number: 2}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "ARAP", Number: 3}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "Gandalf-SLML", Number: 4}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "Xylogics-IPX-SLIP", Number: 5}, &dictionary.Value{Attribute: "Framed-Protocol", Name: "X.75-Synchronous", Number: 6}, &dictionary.Value{Attribute: "Framed-Routing", Name: "None", Number: 0}, &dictionary.Value{Attribute: "Framed-Routing", Name: "Broadcast", Number: 1}, &dictionary.Value{Attribute: "Framed-Routing", Name: "Listen", Number: 2}, &dictionary.Value{Attribute: "Framed-Routing", Name: "Broadcast-Listen", Number: 3}, &dictionary.Value{Attribute: "Framed-Compression", Name: "None", Number: 0}, &dictionary.Value{Attribute: "Framed-Compression", Name: "Van-Jacobson-TCP-IP", Number: 1}, &dictionary.Value{Attribute: "Framed-Compression", Name: "IPX-Header-Compression", Number: 2}, &dictionary.Value{Attribute: "Framed-Compression", Name: "Stac-LZS", Number: 3}, &dictionary.Value{Attribute: "Login-Service", Name: "Telnet", Number: 0}, &dictionary.Value{Attribute: "Login-Service", Name: "Rlogin", Number: 1}, &dictionary.Value{Attribute: "Login-Service", Name: "TCP-Clear", Number: 2}, &dictionary.Value{Attribute: "Login-Service", Name: "PortMaster", Number: 3}, &dictionary.Value{Attribute: "Login-Service", Name: "LAT", Number: 4}, &dictionary.Value{Attribute: "Login-Service", Name: "X25-PAD", Number: 5}, &dictionary.Value{Attribute: "Login-Service", Name: "X25-T3POS", Number: 6}, &dictionary.Value{Attribute: "Login-Service", Name: "TCP-Clear-Quiet", Number: 8}, &dictionary.Value{Attribute: "Login-TCP-Port", Name: "Telnet", Number: 23}, &dictionary.Value{Attribute: "Login-TCP-Port", Name: "Rlogin", Number: 513}, &dictionary.Value{Attribute: "Login-TCP-Port", Name: "Rsh", Number: 514}, &dictionary.Value{Attribute: "Termination-Action", Name: "Default", Number: 0}, &dictionary.Value{Attribute: "Termination-Action", Name: "RADIUS-Request", Number: 1}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Async", Number: 0}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Sync", Number: 1}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "ISDN", Number: 2}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "ISDN-V120", Number: 3}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "ISDN-V110", Number: 4}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Virtual", Number: 5}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "PIAFS", Number: 6}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "HDLC-Clear-Channel", Number: 7}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "X.25", Number: 8}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "X.75", Number: 9}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "G.3-Fax", Number: 10}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "SDSL", Number: 11}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "ADSL-CAP", Number: 12}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "ADSL-DMT", Number: 13}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "IDSL", Number: 14}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Ethernet", Number: 15}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "xDSL", Number: 16}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Cable", Number: 17}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Wireless-Other", Number: 18}, &dictionary.Value{Attribute: "NAS-Port-Type", Name: "Wireless-802.11", Number: 19}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Start", Number: 1}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Stop", Number: 2}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Alive", Number: 3}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Interim-Update", Number: 3}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Accounting-On", Number: 7}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Accounting-Off", Number: 8}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Failed", Number: 15}, &dictionary.Value{Attribute: "Acct-Authentic", Name: "RADIUS", Number: 1}, &dictionary.Value{Attribute: "Acct-Authentic", Name: "Local", Number: 2}, &dictionary.Value{Attribute: "Acct-Authentic", Name: "Remote", Number: 3}, &dictionary.Value{Attribute: "Acct-Authentic", Name: "Diameter", Number: 4}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "User-Request", Number: 1}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Lost-Carrier", Number: 2}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Lost-Service", Number: 3}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Idle-Timeout", Number: 4}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Session-Timeout", Number: 5}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Admin-Reset", Number: 6}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Admin-Reboot", Number: 7}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Port-Error", Number: 8}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "NAS-Error", Number: 9}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "NAS-Request", Number: 10}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "NAS-Reboot", Number: 11}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Port-Unneeded", Number: 12}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Port-Preempted", Number: 13}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Port-Suspended", Number: 14}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Service-Unavailable", Number: 15}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Callback", Number: 16}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "User-Error", Number: 17}, &dictionary.Value{Attribute: "Acct-Terminate-Cause", Name: "Host-Request", Number: 18}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Start", Number: 9}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Stop", Number: 10}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Reject", Number: 11}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Link-Start", Number: 12}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Link-Stop", Number: 13}, &dictionary.Value{Attribute: "Acct-Status-Type", Name: "Tunnel-Link-Reject", Number: 14}, &dictionary.Value{Attribute: "ARAP-Zone-Access", Name: "Default-Zone", Number: 1}, &dictionary.Value{Attribute: "ARAP-Zone-Access", Name: "Zone-Filter-Inclusive", Number: 2}, &dictionary.Value{Attribute: "ARAP-Zone-Access", Name: "Zone-Filter-Exclusive", Number: 4}, &dictionary.Value{Attribute: "Prompt", Name: "No-Echo", Number: 0}, &dictionary.Value{Attribute: "Prompt", Name: "Echo", Number: 1}, &dictionary.Value{Attribute: "Service-Type", Name: "Authorize-Only", Number: 17}, &dictionary.Value{Attribute: "Error-Cause", Name: "Residual-Context-Removed", Number: 201}, &dictionary.Value{Attribute: "Error-Cause", Name: "Invalid-EAP-Packet", Number: 202}, &dictionary.Value{Attribute: "Error-Cause", Name: "Unsupported-Attribute", Number: 401}, &dictionary.Value{Attribute: "Error-Cause", Name: "Missing-Attribute", Number: 402}, &dictionary.Value{Attribute: "Error-Cause", Name: "NAS-Identification-Mismatch", Number: 403}, &dictionary.Value{Attribute: "Error-Cause", Name: "Invalid-Request", Number: 404}, &dictionary.Value{Attribute: "Error-Cause", Name: "Unsupported-Service", Number: 405}, &dictionary.Value{Attribute: "Error-Cause", Name: "Unsupported-Extension", Number: 406}, &dictionary.Value{Attribute: "Error-Cause", Name: "Administratively-Prohibited", Number: 501}, &dictionary.Value{Attribute: "Error-Cause", Name: "Proxy-Request-Not-Routable", Number: 502}, &dictionary.Value{Attribute: "Error-Cause", Name: "Session-Context-Not-Found", Number: 503}, &dictionary.Value{Attribute: "Error-Cause", Name: "Session-Context-Not-Removable", Number: 504}, &dictionary.Value{Attribute: "Error-Cause", Name: "Proxy-Processing-Error", Number: 505}, &dictionary.Value{Attribute: "Error-Cause", Name: "Resources-Unavailable", Number: 506}, &dictionary.Value{Attribute: "Error-Cause", Name: "Request-Initiated", Number: 507}, &dictionary.Value{Attribute: "Error-Cause", Name: "Invalid-Attribute-Value", Number: 407}, &dictionary.Value{Attribute: "Error-Cause", Name: "Multiple-Session-Selection-Unsupported", Number: 508}}}
