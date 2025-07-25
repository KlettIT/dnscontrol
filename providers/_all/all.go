// Package all is simply a container to reference all known provider implementations for easy import into other packages
package all

import (
	// Define all known providers here. They should each register themselves with the providers package via init function.
	_ "github.com/StackExchange/dnscontrol/v4/providers/adguardhome"
	_ "github.com/StackExchange/dnscontrol/v4/providers/akamaiedgedns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/autodns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/axfrddns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/azuredns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/azureprivatedns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/bind"
	_ "github.com/StackExchange/dnscontrol/v4/providers/bunnydns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/cloudflare"
	_ "github.com/StackExchange/dnscontrol/v4/providers/cloudns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/cnr"
	_ "github.com/StackExchange/dnscontrol/v4/providers/cscglobal"
	_ "github.com/StackExchange/dnscontrol/v4/providers/desec"
	_ "github.com/StackExchange/dnscontrol/v4/providers/digitalocean"
	_ "github.com/StackExchange/dnscontrol/v4/providers/dnsimple"
	_ "github.com/StackExchange/dnscontrol/v4/providers/dnsmadeeasy"
	_ "github.com/StackExchange/dnscontrol/v4/providers/doh"
	_ "github.com/StackExchange/dnscontrol/v4/providers/domainnameshop"
	_ "github.com/StackExchange/dnscontrol/v4/providers/dynadot"
	_ "github.com/StackExchange/dnscontrol/v4/providers/easyname"
	_ "github.com/StackExchange/dnscontrol/v4/providers/exoscale"
	_ "github.com/StackExchange/dnscontrol/v4/providers/fortigate"
	_ "github.com/StackExchange/dnscontrol/v4/providers/gandiv5"
	_ "github.com/StackExchange/dnscontrol/v4/providers/gcloud"
	_ "github.com/StackExchange/dnscontrol/v4/providers/gcore"
	_ "github.com/StackExchange/dnscontrol/v4/providers/hedns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/hetzner"
	_ "github.com/StackExchange/dnscontrol/v4/providers/hexonet"
	_ "github.com/StackExchange/dnscontrol/v4/providers/hostingde"
	_ "github.com/StackExchange/dnscontrol/v4/providers/huaweicloud"
	_ "github.com/StackExchange/dnscontrol/v4/providers/internetbs"
	_ "github.com/StackExchange/dnscontrol/v4/providers/inwx"
	_ "github.com/StackExchange/dnscontrol/v4/providers/linode"
	_ "github.com/StackExchange/dnscontrol/v4/providers/loopia"
	_ "github.com/StackExchange/dnscontrol/v4/providers/luadns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/mythicbeasts"
	_ "github.com/StackExchange/dnscontrol/v4/providers/namecheap"
	_ "github.com/StackExchange/dnscontrol/v4/providers/namedotcom"
	_ "github.com/StackExchange/dnscontrol/v4/providers/netcup"
	_ "github.com/StackExchange/dnscontrol/v4/providers/netlify"
	_ "github.com/StackExchange/dnscontrol/v4/providers/ns1"
	_ "github.com/StackExchange/dnscontrol/v4/providers/opensrs"
	_ "github.com/StackExchange/dnscontrol/v4/providers/oracle"
	_ "github.com/StackExchange/dnscontrol/v4/providers/ovh"
	_ "github.com/StackExchange/dnscontrol/v4/providers/packetframe"
	_ "github.com/StackExchange/dnscontrol/v4/providers/porkbun"
	_ "github.com/StackExchange/dnscontrol/v4/providers/powerdns"
	_ "github.com/StackExchange/dnscontrol/v4/providers/realtimeregister"
	_ "github.com/StackExchange/dnscontrol/v4/providers/route53"
	_ "github.com/StackExchange/dnscontrol/v4/providers/rwth"
	_ "github.com/StackExchange/dnscontrol/v4/providers/sakuracloud"
	_ "github.com/StackExchange/dnscontrol/v4/providers/softlayer"
	_ "github.com/StackExchange/dnscontrol/v4/providers/transip"
	_ "github.com/StackExchange/dnscontrol/v4/providers/vultr"
)
