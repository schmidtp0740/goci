package main

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/schmidtp0740/oci-dev/cmd"
)

func main() {
	config := common.DefaultConfigProvider()

	cmd.Execute(config)
	// vcn := createVCN(config, cidrBlock, compartmentID, vcnDisplayName, dnsLabel)
	// deleteVCN(config, vcn)
}
