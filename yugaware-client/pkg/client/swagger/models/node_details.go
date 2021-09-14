// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeDetails Details of a cloud node
//
// swagger:model NodeDetails
type NodeDetails struct {

	// The availability zone's UUID
	// Format: uuid
	AzUUID strfmt.UUID `json:"azUuid,omitempty"`

	// Node information, as reported by the cloud provider
	CloudInfo *CloudSpecificInfo `json:"cloudInfo,omitempty"`

	// True if cron jobs were properly configured for this node
	CronsActive bool `json:"cronsActive,omitempty"`

	// True if this node is a master
	IsMaster bool `json:"isMaster,omitempty"`

	// True if this node is a REDIS server
	IsRedisServer bool `json:"isRedisServer,omitempty"`

	// True if this node is a Tablet server
	IsTserver bool `json:"isTserver,omitempty"`

	// True if this node is a YCQL server
	IsYqlServer bool `json:"isYqlServer,omitempty"`

	// True if this node is a YSQL server
	IsYsqlServer bool `json:"isYsqlServer,omitempty"`

	// Machine image name
	MachineImage string `json:"machineImage,omitempty"`

	// Master HTTP port
	MasterHTTPPort int32 `json:"masterHttpPort,omitempty"`

	// Master RCP port
	MasterRPCPort int32 `json:"masterRpcPort,omitempty"`

	// Node exporter port
	NodeExporterPort int32 `json:"nodeExporterPort,omitempty"`

	// Node ID
	NodeIdx int32 `json:"nodeIdx,omitempty"`

	// Node name
	NodeName string `json:"nodeName,omitempty"`

	// Node UUID
	// Format: uuid
	NodeUUID strfmt.UUID `json:"nodeUuid,omitempty"`

	// UUID of the cluster to which this node belongs
	// Format: uuid
	PlacementUUID strfmt.UUID `json:"placementUuid,omitempty"`

	// REDIS HTTP port
	RedisServerHTTPPort int32 `json:"redisServerHttpPort,omitempty"`

	// REDIS RPC port
	RedisServerRPCPort int32 `json:"redisServerRpcPort,omitempty"`

	// Node state
	// Example: Provisioned
	// Enum: [ToBeAdded ToJoinCluster Provisioned SoftwareInstalled UpgradeSoftware UpdateGFlags Live Stopping Starting Stopped Unreachable ToBeRemoved Removing Removed Adding BeingDecommissioned Decommissioned UpdateCert ToggleTls Resizing]
	State string `json:"state,omitempty"`

	// Tablet server HTTP port
	TserverHTTPPort int32 `json:"tserverHttpPort,omitempty"`

	// Tablet server RPC port
	TserverRPCPort int32 `json:"tserverRpcPort,omitempty"`

	// YCQL HTTP port
	YqlServerHTTPPort int32 `json:"yqlServerHttpPort,omitempty"`

	// YCQL RPC port
	YqlServerRPCPort int32 `json:"yqlServerRpcPort,omitempty"`

	// YSQL HTTP port
	YsqlServerHTTPPort int32 `json:"ysqlServerHttpPort,omitempty"`

	// YSQL RPC port
	YsqlServerRPCPort int32 `json:"ysqlServerRpcPort,omitempty"`
}

// Validate validates this node details
func (m *NodeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacementUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeDetails) validateAzUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.AzUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("azUuid", "body", "uuid", m.AzUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetails) validateCloudInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudInfo) { // not required
		return nil
	}

	if m.CloudInfo != nil {
		if err := m.CloudInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudInfo")
			}
			return err
		}
	}

	return nil
}

func (m *NodeDetails) validateNodeUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("nodeUuid", "body", "uuid", m.NodeUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetails) validatePlacementUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.PlacementUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("placementUuid", "body", "uuid", m.PlacementUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

var nodeDetailsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ToBeAdded","ToJoinCluster","Provisioned","SoftwareInstalled","UpgradeSoftware","UpdateGFlags","Live","Stopping","Starting","Stopped","Unreachable","ToBeRemoved","Removing","Removed","Adding","BeingDecommissioned","Decommissioned","UpdateCert","ToggleTls","Resizing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeDetailsTypeStatePropEnum = append(nodeDetailsTypeStatePropEnum, v)
	}
}

const (

	// NodeDetailsStateToBeAdded captures enum value "ToBeAdded"
	NodeDetailsStateToBeAdded string = "ToBeAdded"

	// NodeDetailsStateToJoinCluster captures enum value "ToJoinCluster"
	NodeDetailsStateToJoinCluster string = "ToJoinCluster"

	// NodeDetailsStateProvisioned captures enum value "Provisioned"
	NodeDetailsStateProvisioned string = "Provisioned"

	// NodeDetailsStateSoftwareInstalled captures enum value "SoftwareInstalled"
	NodeDetailsStateSoftwareInstalled string = "SoftwareInstalled"

	// NodeDetailsStateUpgradeSoftware captures enum value "UpgradeSoftware"
	NodeDetailsStateUpgradeSoftware string = "UpgradeSoftware"

	// NodeDetailsStateUpdateGFlags captures enum value "UpdateGFlags"
	NodeDetailsStateUpdateGFlags string = "UpdateGFlags"

	// NodeDetailsStateLive captures enum value "Live"
	NodeDetailsStateLive string = "Live"

	// NodeDetailsStateStopping captures enum value "Stopping"
	NodeDetailsStateStopping string = "Stopping"

	// NodeDetailsStateStarting captures enum value "Starting"
	NodeDetailsStateStarting string = "Starting"

	// NodeDetailsStateStopped captures enum value "Stopped"
	NodeDetailsStateStopped string = "Stopped"

	// NodeDetailsStateUnreachable captures enum value "Unreachable"
	NodeDetailsStateUnreachable string = "Unreachable"

	// NodeDetailsStateToBeRemoved captures enum value "ToBeRemoved"
	NodeDetailsStateToBeRemoved string = "ToBeRemoved"

	// NodeDetailsStateRemoving captures enum value "Removing"
	NodeDetailsStateRemoving string = "Removing"

	// NodeDetailsStateRemoved captures enum value "Removed"
	NodeDetailsStateRemoved string = "Removed"

	// NodeDetailsStateAdding captures enum value "Adding"
	NodeDetailsStateAdding string = "Adding"

	// NodeDetailsStateBeingDecommissioned captures enum value "BeingDecommissioned"
	NodeDetailsStateBeingDecommissioned string = "BeingDecommissioned"

	// NodeDetailsStateDecommissioned captures enum value "Decommissioned"
	NodeDetailsStateDecommissioned string = "Decommissioned"

	// NodeDetailsStateUpdateCert captures enum value "UpdateCert"
	NodeDetailsStateUpdateCert string = "UpdateCert"

	// NodeDetailsStateToggleTLS captures enum value "ToggleTls"
	NodeDetailsStateToggleTLS string = "ToggleTls"

	// NodeDetailsStateResizing captures enum value "Resizing"
	NodeDetailsStateResizing string = "Resizing"
)

// prop value enum
func (m *NodeDetails) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeDetailsTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeDetails) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node details based on the context it is used
func (m *NodeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeDetails) contextValidateCloudInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudInfo != nil {
		if err := m.CloudInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeDetails) UnmarshalBinary(b []byte) error {
	var res NodeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}