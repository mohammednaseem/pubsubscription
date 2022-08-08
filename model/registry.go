package model

import (
	cloudiot "google.golang.org/api/cloudiot/v1"
)

type Registry struct {
	ProjectID  string `json:"projectid" validate:"required"`
	Region     string `json:"region" validate:"required"`
	RegistryID string `json:"registryid" validate:"required"`
	TopicName  string `json:"topicname"  validate:"required"`
}

type RegistryCreate struct {
	Parent                   string                              `json:"parent" validate:"required"`
	Project                  string                              `json:"project" validate:"required"`
	Region                   string                              `json:"region" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:"required"`
	EventNotificationConfigs []*cloudiot.EventNotificationConfig `json:"eventNotificationConfigs" validate:"required"`
	StateNotificationConfig  *cloudiot.StateNotificationConfig   `json:"stateNotificationConfig"  validate:""`
	MqttConfig               cloudiot.MqttConfig                 `json:"mqttConfig"  validate:""`
	HttpConfig               cloudiot.HttpConfig                 `json:"httpConfig"  validate:""`
	LogLevel                 string                              `json:"logLevel"  validate:""`
	CreatedOn                string                              `json:"createdOn"  validate:""`
	Credentials              []*cloudiot.RegistryCredential      `json:"credentials"  validate:""`
}
type RegistryUpdate struct {
	UpdateMask               string                              `json:"updatemask" validate:"required"`
	Project                  string                              `json:"project" validate:"required"`
	Region                   string                              `json:"region" validate:"required"`
	Parent                   string                              `json:"parent" validate:"required"`
	Id                       string                              `json:"id" validate:"required"`
	Name                     string                              `json:"name" validate:"required"`
	EventNotificationConfigs []*cloudiot.EventNotificationConfig `json:"eventNotificationConfigs" validate:""`
	StateNotificationConfig  *cloudiot.StateNotificationConfig   `json:"stateNotificationConfig"  validate:""`
	MqttConfig               cloudiot.MqttConfig                 `json:"mqttConfig"  validate:""`
	HttpConfig               cloudiot.HttpConfig                 `json:"httpConfig"  validate:""`
	LogLevel                 string                              `json:"logLevel"  validate:""`
	CreatedOn                string                              `json:"createdOn"  validate:""`
	Credentials              []*cloudiot.RegistryCredential      `json:"credentials"  validate:""`
}
type RegistryDelete struct {
	Id      string `json:"id" validate:"required"`
	Parent  string `json:"parent" validate:"required"`
	Project string `json:"project" validate:"required"`
	Region  string `json:"region" validate:"required"`
}
type DeviceCreate struct {
	Project     string                       `json:"project" validate:"required"`
	Parent      string                       `json:"parent" validate:"required"`
	NumId       string                       `json:"numId" validate:""`
	Region      string                       `json:"region" validate:"required"`
	Registry    string                       `json:"registry" validate:"required"`
	Id          string                       `json:"id" validate:"required"`
	Name        string                       `json:"name" validate:"required"`
	Credentials []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel    string                       `json:"logLevel"  validate:""`
	Blocked     bool                         `json:"blocked"  validate:""`
	Metadata    map[string]string            `json:"metadata"  validate:""`
	CreatedOn   string                       `json:"createdOn"  validate:""`
}
type DeviceUpdate struct {
	UpdateMask  string                       `json:"updatemask" validate:"required"`
	Project     string                       `json:"project" validate:"required"`
	Region      string                       `json:"region" validate:"required"`
	Registry    string                       `json:"registry" validate:"required"`
	Parent      string                       `json:"parent" validate:"required"`
	Id          string                       `json:"id" validate:"required"`
	Name        string                       `json:"name" validate:"required"`
	Credentials []*cloudiot.DeviceCredential `json:"credentials" validate:"required"`
	LogLevel    string                       `json:"logLevel"  validate:""`
	Blocked     bool                         `json:"blocked"  validate:""`
	Metadata    map[string]string            `json:"metadata"  validate:""`
	CreatedOn   string                       `json:"createdOn"  validate:""`
}
type DeviceDelete struct {
	Id       string `json:"id" validate:"required"`
	Parent   string `json:"parent" validate:"required"`
	Project  string `json:"project" validate:"required"`
	Region   string `json:"region" validate:"required"`
	Registry string `json:"registry" validate:"required"`
}
