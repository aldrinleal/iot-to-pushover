package types

import (
	"encoding/json"
	"time"
)

type DeviceInfo struct {
	Attributes    map[string]string `json:"attributes"`
	DeviceID      string            `json:"deviceId"`
	RemainingLife json.Number       `json:"remainingLife"`
	Type          string            `json:"type"`
}

type DevicePayload struct {
	CertificateID string      `json:"certificateId"`
	ClickType     string      `json:"clickType"`
	RemainingLife json.Number `json:"remainingLife"`
	ReportedTime  int64       `json:"reportedTime"`
	SerialNumber  string      `json:"serialNumber"`
	Topic         string      `json:"topic"`
	Version       string      `json:"version"`
}

type ButtonClicked struct {
	AdditionalInfo struct {
		Version string `json:"version"`
	} `json:"additionalInfo"`
	ClickType    string    `json:"clickType"`
	ReportedTime time.Time `json:"reportedTime"`
}

type DeviceHealthMonitor struct {
	Condition struct {
		RemainingLifeLowerThan json.Number `json:"remainingLifeLowerThan"`
	} `json:"condition"`
}

type PlacementInfo struct {
	Attributes    map[string]string `json:"attributes"`
	Devices       map[string]string `json:"devices"`
	PlacementName string            `json:"placementName"`
	ProjectName   string            `json:"projectName"`
}

type OneClickEvent struct {
	DeviceEvent struct {
		ButtonClicked       *ButtonClicked       `json:"buttonClicked"`
		DeviceHealthMonitor *DeviceHealthMonitor `json:"deviceHealthMonitor"`
	} `json:"deviceEvent"`
	DeviceInfo    DeviceInfo     `json:"deviceInfo"`
	DevicePayload *DevicePayload `json:"devicePayload"`
	PlacementInfo *PlacementInfo `json:"placementInfo"`
}
