package fortigate

import (
	"fmt"
	"golang.org/x/net/idna"
	"net"
	"strings"

	"github.com/StackExchange/dnscontrol/v4/models"
)

/*
	--------------------------------------------------------------------------
	  nativeToRecord – convert an fgDNSRecord coming from FortiGate into a
	  *models.RecordConfig that dnscontrol understands

---------------------------------------------------------------------------
*/
func nativeToRecord(domain string, n fgDNSRecord) (*models.RecordConfig, error) {
	rc := &models.RecordConfig{}
	rc.Type = strings.ToUpper(n.Type)
	rc.Original = n

	// --- Label / Name -----------------------------------------------------
	label := strings.TrimSuffix(n.Hostname, ".")
	if label == "@" {
		label = ""
	}
	rc.SetLabel(label, domain)

	// --- TTL --------------------------------------------------------------
	rc.TTL = n.TTL
	if rc.TTL == 0 {
		rc.TTL = 300 // oder ggf. dc.DefaultTTL
	}

	// --- Status → Metadata ------------------------------------------------
	if strings.ToLower(n.Status) != "enable" {
		if rc.Metadata == nil {
			rc.Metadata = map[string]string{}
		}
		rc.Metadata["fortigate_status"] = "disable"
	}

	// --- Type-specific fields ---------------------------------------------
	switch rc.Type {
	case "A":
		ip := net.ParseIP(n.IP)
		if ip == nil || ip.To4() == nil {
			return nil, fmt.Errorf("invalid IPv4 address %q in %+v", n.IP, n)
		}
		rc.SetTargetIP(ip)

	case "AAAA":
		ip := net.ParseIP(n.IPv6)
		if ip == nil || ip.To16() == nil || ip.To4() != nil {
			return nil, fmt.Errorf("invalid IPv6 address %q in %+v", n.IPv6, n)
		}
		rc.SetTargetIP(ip)

	case "CNAME":
		if n.CanonicalName == "" {
			return nil, fmt.Errorf("CNAME record without canonical-name (id=%d)", n.ID)
		}
		if err := rc.SetTarget(ensureDot(n.CanonicalName)); err != nil {
			return nil, err
		}

	case "PTR", "PTR_V6":
		var ip net.IP
		if n.Type == "PTR" {
			ip = net.ParseIP(n.IP)
		} else {
			ip = net.ParseIP(n.IPv6)
		}
		if ip == nil {
			return nil, fmt.Errorf("%s record without valid IP (id=%d)", n.Type, n.ID)
		}
		rc.Type = "PTR"
		rc.SetTargetIP(ip)

	default:
		// NS is not supported due to FortiGate limitations
		return nil, fmt.Errorf("record type %q is not supported by fortigate provider", rc.Type)
	}

	fmt.Printf("FINAL RC: type=%s fqdn=%s target=%s ttl=%d\n",
		rc.Type,
		rc.GetLabelFQDN(),
		rc.GetTargetField(),
		rc.TTL,
	)

	return rc, nil
}

func recordsToNative(recs, existing models.Records) ([]*fgDNSRecord, []error) {
	var resourceRecords []*fgDNSRecord
	var errors []error

	id := 1

	for _, record := range recs {

		n := &fgDNSRecord{
			Status: "enable",
			Type:   strings.ToUpper(record.Type),
			TTL:    record.TTL, // 0 = inherit
		}

		// --- Wildcard support -------------------------------------------------
		if strings.Contains(record.GetLabelFQDN(), "*") {
			errors = append(errors, fmt.Errorf("wildcard records are not supported by FortiGate: %s", record.GetLabelFQDN()))
			continue
		}

		// --- Status from Metadata ---------------------------------------------
		if v, ok := record.Metadata["fortigate_status"]; ok && strings.ToLower(v) == "disable" {
			n.Status = "disable"
		}

		// --- Hostname (Label) -------------------------------------------------
		label := record.GetLabel()
		if label == "" {
			label = "@"
		}
		n.Hostname = label

		// --- Type-specific fields ---------------------------------------------
		switch n.Type {
		case "A":
			ip := record.GetTargetIP()
			if ip == nil || ip.To4() == nil {
				errors = append(errors, fmt.Errorf("a record is missing a valid IPv4 address: %s", record.GetLabelFQDN()))
				continue
			}
			n.IP = ip.String()

		case "AAAA":
			ip := record.GetTargetIP()
			if ip == nil || ip.To16() == nil || ip.To4() != nil {
				errors = append(errors, fmt.Errorf("AAAA record is missing a valid IPv6 address: %s", record.GetLabelFQDN()))
				continue
			}
			n.IPv6 = ip.String()

		case "CNAME":
			target := strings.TrimSuffix(record.GetTargetField(), ".")
			if ascii, err := idna.ToASCII(target); err == nil {
				target = ascii
			}
			n.CanonicalName = target

		case "PTR":
			ip := net.ParseIP(record.GetTargetField())
			if ip == nil {
				errors = append(errors, fmt.Errorf("PTR record %s has invalid target %q: FortiGate only supports IP addresses", record.GetLabelFQDN(), record.GetTargetField()))
				continue
			}
			if ip.To4() != nil {
				n.Type = "PTR"
				n.IP = ip.String()
			} else {
				n.Type = "PTR_V6"
				n.IPv6 = ip.String()
			}

		default:
			errors = append(errors, fmt.Errorf("record type %q is not supported by FortiGate provider: %s", n.Type, record.GetLabelFQDN()))
			continue
		}

		fmt.Printf("recordToNative(): %s => %s ID=%d\n", record.GetLabelFQDN(), n.Hostname, n.ID)
		fmt.Printf("RECORD TO NATIVE: fqdn=%s target=%s ttl=%d\n",
			record.GetLabelFQDN(),
			record.GetTargetField(),
			record.TTL,
		)

		n.ID = id
		id++

		resourceRecords = append(resourceRecords, n)
	}

	return resourceRecords, errors
}

/*
	--------------------------------------------------------------------------
	  ensureDot – make sure an FQDN ends with a trailing dot

---------------------------------------------------------------------------
*/
func ensureDot(fqdn string) string {
	if fqdn == "" || strings.HasSuffix(fqdn, ".") {
		return fqdn
	}
	return fqdn + "."
}
