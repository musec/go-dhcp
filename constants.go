package dhcp

const (
	BOOTREQUEST byte = 1
	BOOTREPLY   byte = 2
)

const (
	ETHERNET byte = 1
)

const (
	DISCOVER byte = 1
	OFFER    byte = 2
	REQUEST  byte = 3
	DECLINE  byte = 4
	ACK      byte = 5
	NAK      byte = 6
	RELEASE  byte = 7
	INFORM   byte = 8
)

// Option codes
const ( // Data Length		Meaning
	PAD                               byte = 0  //	None        Pad
	OptSubnetMask                     byte = 1  //	4           Subnet Mask Value
	OptTimeOffset                     byte = 2  //	4           Time Offset in Seconds from UTC (note:deprecated by 100 and 101)
	OptRouter                         byte = 3  //	N           N/4 Router addresses
	OptTimeServer                     byte = 4  //	N           N/4 Timeserver addresses
	OptNameServer                     byte = 5  //	N           N/4 IEN-116 Server addresses
	OptDomainServer                   byte = 6  //	N           N/4 DNS Server addresses
	OptLogServer                      byte = 7  //	N           N/4 Logging Server addresses
	OptQuotesServer                   byte = 8  //	N           N/4 Quotes Server addresses
	OptLPRServer                      byte = 9  //	N           N/4 Printer Server addresses
	OptImpressServer                  byte = 10 //	N           N/4 Impress Server addresses
	OptRLPServer                      byte = 11 //	N           N/4 RLP Server addresses
	OptHostName                       byte = 12 //	N           Hostname string
	OptBootFileSize                   byte = 13 //	2           Size of boot file in 512 byte chunks
	OptMeritDumpFile                  byte = 14 //	N           Client to dump and name the file to dump to
	OptDomainName                     byte = 15 //	N           The DNS domain name of the client
	OptSwapServer                     byte = 16 //	N           Swap Server address
	OptRootPath                       byte = 17 //	N           Path name for root disk
	OptExtensionFile                  byte = 18 //	N           Path name for more BOOTP info
	OptForwardOnOff                   byte = 19 //	1           Enable/Disable IP Forwarding
	OptSrcRteOnOff                    byte = 20 //	1           Enable/Disable Source Routing
	OptPolicyFilter                   byte = 21 //	N           Routing Policy Filters
	OptMaxDGAssembly                  byte = 22 //	2           Max Datagram Reassembly Size
	OptDefaultIPTTL                   byte = 23 //	1           Default IP Time to Live
	OptMTUTImeout                     byte = 24 //	4           Path MTU Aging Timeout
	OptMTUPlateau                     byte = 25 //	N           Path MTU Plateau Table
	OptMTUInterface                   byte = 26 //	2           Interface MTU Size
	OptMTUSubnet                      byte = 27 //	1           All Subnets are Local
	OptBroadcastAddress               byte = 28 //	4           Broadcast Address
	OptMaskDiscovery                  byte = 29 //	1           Perform Mask Discovery
	OptMaskSupplier                   byte = 30 //	1           Provide Mask to Others
	OptRouterDiscovery                byte = 31 //	1           Perform Router Discovery
	OptRouterRequest                  byte = 32 //	4           Router Solicitation Address
	OptStaticRoute                    byte = 33 //	N           Static Routing Table
	OptTrailers                       byte = 34 //	1           Trailer Encapsulation
	OptARPTimeout                     byte = 35 //	4           ARP Cache Timeout
	OptEthernet                       byte = 36 //	1           Ethernet Encapsulation
	OptDefaultTCPTTL                  byte = 37 //	1           Default TCP Time to Live
	OptKeepaliveTime                  byte = 38 //	4           TCP Keepalive Interval
	OptKeppaliveData                  byte = 39 //	1           TCP Keepalive Garbage
	OptNISDomain                      byte = 40 //	N           NIS Domain Name
	OptNISServers                     byte = 41 //	N           NIS Server Addresses
	OptNTPServers                     byte = 42 //	N           NTP Server Addresses
	OptVendorSpecofic                 byte = 43 //	N           Vendor Specific Information
	OptNETBIOSNameSrv                 byte = 44 //	N           NETBIOS Name Servers
	OptNETBIOSDistSrv                 byte = 45 //	N           NETBIOS Datagram Distribution
	OptNETBIOSNodeType                byte = 46 //	1           NETBIOS Node Type
	OptNETBIOSScope                   byte = 47 //	N           NETBIOS Scope
	OptXWindowFont                    byte = 48 //	N           X Window Font Server
	OptXWindowManager                 byte = 49 //	N           X Window Display Manager
	OptAddressRequest                 byte = 50 //	4           Requested IP Address
	OptAddressTime                    byte = 51 //	4           IP Address Lease Time
	OptOverload                       byte = 52 //	1           Overload "sname" or "file"
	OptDHCPMsgType                    byte = 53 //	1           DHCP Message Type
	OptDHCPServerId                   byte = 54 //	4           DHCP Server Identification
	OptParameterList                  byte = 55 //	N           Parameter Request Lis
	OptDHCPMessage                    byte = 56 //	N           DHCP Error Message
	OptDHCPMaxMsgSize                 byte = 57 //	2           DHCP Maximum Message Size
	OptRenewalTime                    byte = 58 //	4           DHCP Renewal (T1) Time
	OptRebindingTime                  byte = 59 //	4           DHCP Rebinding (T2) Time
	OptClassId                        byte = 60 //	N           Class Identifier
	OptClientId                       byte = 61 //	N           Client Identifier
	OptNetWareIPDomain                byte = 62 //	N           NetWare/IP Domain Name
	OptNetWareIPOption                byte = 63 //	N           NetWare/IP sub Options
	OptNISDomainName                  byte = 64 //	N           NIS+ v3 Client Domain Name
	OptNISServerAddr                  byte = 65 //	N           NIS+ v3 Server Addresses
	OptServerName                     byte = 66 //	N           TFTP Server Name
	OptBootfileName                   byte = 67 //	N           Boot File Name
	OptHomeAgentAddrs                 byte = 68 //	N           Home Agent Addresses
	OptSMTPServer                     byte = 69 //	N           Simple Mail Server Addresses
	OptPOP3Server                     byte = 70 //	N           Post Office Server Addresses
	OptNNTPServer                     byte = 71 //	N           Network News Server Addresses
	OptWWWServer                      byte = 72 //	N           WWW Server Addresses
	OptFingerServer                   byte = 73 //	N           Finger Server Addresses
	OptIRCServer                      byte = 74 //	N           Chat Server Addresses
	OptStreetTalkServer               byte = 75 //	N           StreetTalk Server Addresses
	OptSTDAServer                     byte = 76 //	N           ST Directory Assist. Addresses
	OptUserClass                      byte = 77 //	N           User Class Information
	OptDirectoryAgent                 byte = 78 //	N           directory agent information
	OptServiceScope                   byte = 79 //	N           service location agent scope
	OptRapidCommit                    byte = 80 //	0           Rapid Commit
	OptClientFullyQualifiedDomainName byte = 81 //	N           Fully Qualified Domain Name
	OptRelayAgentInformation          byte = 82 //	N           Relay Agent Information
	OptInternetStorageNameService     byte = 83 //	N           Internet Storage Name Service
	//  84				    EMOVED/Unassigned
	OptNDSServers                       byte = 85 //	N           Novell Directory Services
	OptNDSTreeName                      byte = 86 //	N           Novell Directory Services
	OptNDSContext                       byte = 87 //	N           Novell Directory Services
	OptBCMCSControllerDomainNameList    byte = 88 //    		    BCMCS Controller Domain Name list
	OptBCMCSControllerIPv4AddressOption byte = 89 //    		    BCMCS Controller IPv4 address option
	OptAuthentication                   byte = 90 //	N           Authentication
	OptClientLastTransactionTimeOption  byte = 91 //    		    client-last-transaction-time option
	OptAssociatedIPOption               byte = 92 //    		    associated-ip option
	OptClientSystemArchitecture         byte = 93 //	N           Client System Architecture
	OptClientNetworkDeviceInterface     byte = 94 //	N           Client Network Device Interface
	OptLDAP                             byte = 95 //	N           Lightweight Directory Access Protocol
	//  96    			    REMOVED/Unassigned
	OptUUIDGUID      byte = 97  //	N           UUID/GUID-based Client Identifier
	OptUserAuth      byte = 98  //	N           Open Group's User Authentication
	OptGEOCONF_CIVIC byte = 99  //    		    GEOCONF_CIVIC
	OptPCode         byte = 100 //	N           IEEE 1003.1 TZ String
	OptTCode         byte = 101 //	N           Reference to the TZ Database
	//102-107 			    REMOVED/Unassigned
	//  108   			    REMOVED/Unassigned
	//  109   			    Unassigned
	//  110   			    REMOVED/Unassigned
	//  111   			    Unassigned
	OptNetinfoAddress byte = 112 //	N           NetInfo Parent Server Address
	OptNetinfoTag     byte = 113 //	N           NetInfo Parent Server Tag
	OptURL            byte = 114 //	N           URL
	//   115   			    REMOVED/Unassigned
	OptAutoConfig                  byte = 116 //	N           DHCP Auto-Configuration
	OptNameServiceSearch           byte = 117 //	N           Name Service Search
	OptSubnetSelectionOption       byte = 118 //	4           Subnet Selection Option
	OptDomainSearch                byte = 119 //	N           DNS domain search list
	OptSIPServersDHCPOption        byte = 120 //	N           SIP Servers DHCP Option
	OptClasslessStaticRouteOption  byte = 121 //	N           Classless Static Route Option
	OptCCC                         byte = 122 //	N           CableLabs Client Configuration
	OptGeoConfOption               byte = 123 //	16          GeoConf Option
	OptVIVendorClass               byte = 124 //		    Vendor-Identifying Vendor Class
	OptVIVendorSpecificInformation byte = 125 //		    endor-Identifying Vendor-Specific Information
	//   126   			    Removed/Unassigned
	//   127   		 	    Removed/Unassigned
	OptPXE_128                          byte = 128 //		    PXE - undefined (vendor specific)
	OptPXE_129                          byte = 129 //		    PXE - undefined (vendor specific)
	OptPXE_130                          byte = 130 //		    PXE - undefined (vendor specific)
	OptPXE_131                          byte = 131 //		    PXE - undefined (vendor specific)
	OptPXE_132                          byte = 132 //		    PXE - undefined (vendor specific)
	OptPXE_133                          byte = 133 //		    PXE - undefined (vendor specific)
	OptPXE_134                          byte = 134 //		    PXE - undefined (vendor specific)
	OptPXE_135                          byte = 135 //		    PXE - undefined (vendor specific)
	OptOPTIONPANAAGENT                  byte = 136 //		    OPTION_PANA_AGENT
	OptOPTIONV4LOST                     byte = 137 //   		    OPTION_V4_LOST
	OptOPTIONCAPWAPACV4                 byte = 138 //	N           CAPWAP Access Controller addresses
	OptOPTIONIPV4AdressMoS              byte = 139 //	N           a series of suboptions
	OptOPTIONIPV4FQDNMoS                byte = 140 //	N           a series of suboptions
	OptSIPUAConfigurationServiceDomains byte = 141 //	N           List of domain names to search for SIP User Agent Configuration
	OptOPTIONIPv4AddressANDSF           byte = 142 //	N           ANDSF IPv4 Address Option for DHCPv4
	//   143			    Unassigned
	OptGeoLoc                   byte = 144 //	16           Geospatial Location with Uncertainty
	OptFORCERENEW_NONCE_CAPABLE byte = 145 //	1           Forcerenew Nonce Capable
	OptRDNSSSelection           byte = 146 //	N           Information for selecting RDNSS
	// 147-149			    Unassigned
	OptTFTPServerAddress byte = 150 //		    TFTP server address
	//   150			    Etherboot
	//   150   			    GRUB configuration path name
	OptStatusCode           byte = 151 //	N+1         Status code and optional N byte text message describing status.
	OptBaseTime             byte = 152 //	4           Absolute time (seconds since Jan 1, 1970) message was sent.
	OptStartTimeOfState     byte = 153 //	4           Number of seconds in the past when client entered current state.
	OptQueryStartTime       byte = 154 //	4           Absolute time (seconds since Jan 1, 1970) for beginning of query.
	OptQueryEndTime         byte = 155 //	4           Absolute time (seconds since Jan 1, 1970) for end of query.
	OptDHCPState            byte = 156 //	1           State of IP address.
	OptDataSource           byte = 157 //	1           Indicates information came from local or remote server.
	OptOPTION_V4_PCP_SERVER byte = 158 //   		    IP addresses; each list is treated as a separate PCP server.
	OptOPTION_V4_PORTPARAMS byte = 159 //	4           This option is used to configure a set of ports bound to a shared IPv4 address.
	// 160-174			    Unassigned
	OptEtherboot_175 byte = 175 //		    Etherboot (Tentatively Assigned - 2005-06-23)
	OptIPTelephone   byte = 176 //		    IP Telephone (Tentatively Assigned - 2005-06-23)
	//   177   			    Etherboot (Tentatively Assigned - 2005-06-23)
	OptPacketCable byte = 177 //		    PacketCable and CableHome (replaced by 122)
	// 178-207 			    Unassigned
	OptPXELINUXMagic           byte = 208 //	4           magic string = F1:00:74:7E
	OptConfigurationFile       byte = 209 //	N           Configuration file
	OptPathPrefix              byte = 210 //	N           Path Prefix Option
	OptRebootTime              byte = 211 //	4           Reboot Time
	OptOPTION_6RD              byte = 212 //	18 + N      OPTION_6RD with N/4 6rd BR addresses
	OptOPTION_V4_ACCESS_DOMAIN byte = 213 //	N           Access Network Domain Name
	// 214-219			    Unassigned
	OptSubnetAllocationOption byte = 220 //	N           Subnet Allocation Option
	OptVirtualSubnetSelection byte = 221 //		    Virtual Subnet Selection (VSS) Option
	// 222-223			    Unassigned
	// 224-254			    Reserved (Private Use)
	END byte = 255 //	None        END
)
