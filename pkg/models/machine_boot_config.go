// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MachineBootConfig Machine boot config that will be passed to the instance that can be used to perform common automated configuration tasks and even run scripts after the instance starts.
//
// swagger:model MachineBootConfig
type MachineBootConfig struct {

	// A valid cloud config data in json-escaped yaml syntax
	// Example: #cloud-config\nrepo_update: true\nrepo_upgrade: all\n\npackages:\n - mysql-server\n\nruncmd:\n - sed -e '/bind-address/ s/^#*/#/' -i /etc/mysql/mysql.conf.d/mysqld.cnf\n - service mysql restart\n - mysql -e \"GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'mysqlpassword';\"\n - mysql -e \"FLUSH PRIVILEGES;\"\n
	Content string `json:"content,omitempty"`
}

// Validate validates this machine boot config
func (m *MachineBootConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this machine boot config based on context it is used
func (m *MachineBootConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MachineBootConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineBootConfig) UnmarshalBinary(b []byte) error {
	var res MachineBootConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
