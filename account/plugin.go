package account

import (
	"github.com/backstage/backstage/db"
	"github.com/backstage/backstage/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PluginConfig struct {
	Name    string                 `json:"name"`
	Service string                 `json:"service"`
	Config  map[string]interface{} `json:"config,omitempty"`
}

// Save associates a middleware with a service.
//
// It requires to inform the fields: Subdomain and Middleware name.
// It is not allowed to associates two middlewares with the same subdomain.
func (m *PluginConfig) Save(user *User) error {
	err := m.validate(user)
	if err != nil {
		return err
	}

	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.PluginsConfig().Upsert(bson.M{"service": m.Service, "name": m.Name}, m)
	return err
}

// Delete removes an existing middleware config from the server.
func (m *PluginConfig) Delete(user *User) error {
	err := m.validate(user)
	if err != nil {
		return err
	}

	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.PluginsConfig().Remove(m)
	if err == mgo.ErrNotFound {
		return &errors.ValidationError{Payload: "Configuration not found."}
	}
	if err != nil {
		return &errors.ValidationError{Payload: err.Error()}
	}
	return err
}

// DeletePluginsByService removes an existing plugin config from the server based on given subdomain.
func DeletePluginsByService(subdomain string) error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.PluginsConfig().RemoveAll(bson.M{"service": subdomain})
	if err != nil {
		return err
	}

	return nil
}

func (m *PluginConfig) validate(user *User) error {
	if m.Name == "" {
		return &errors.ValidationError{Payload: "Name cannot be empty."}
	}
	if m.Service == "" {
		return &errors.ValidationError{Payload: "Service cannot be empty."}
	}

	service, err := FindServiceBySubdomain(m.Service)
	if err != nil {
		return err
	}

	_, err = FindTeamByAlias(service.Team, user)
	if err != nil {
		return err
	}
	return nil
}
