package serialization

import (
	"github.com/hungyu99/freed/domain/consensus/model/externalapi"
	"github.com/hungyu99/freed/domain/consensus/utils/subnetworks"
)

// DbSubnetworkIDToDomainSubnetworkID converts DbSubnetworkId to DomainSubnetworkID
func DbSubnetworkIDToDomainSubnetworkID(dbSubnetworkID *DbSubnetworkId) (*externalapi.DomainSubnetworkID, error) {
	return subnetworks.FromBytes(dbSubnetworkID.SubnetworkId)
}

// DomainSubnetworkIDToDbSubnetworkID converts DomainSubnetworkID to DbSubnetworkId
func DomainSubnetworkIDToDbSubnetworkID(domainSubnetworkID *externalapi.DomainSubnetworkID) *DbSubnetworkId {
	return &DbSubnetworkId{SubnetworkId: domainSubnetworkID[:]}
}
