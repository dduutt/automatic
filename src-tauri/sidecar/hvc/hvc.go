package hvc

import (
	"time"
)

type Device struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	TableName  string  `json:"table_name"`
	SkipLength float64 `json:"skip_length"`
}

type DeviceInfo struct {
	ID         int       `json:"id"`
	DeviceName string    `json:"device_name"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type ChangeInfo struct {
	ID         int       `json:"id"`
	DeviceName string    `json:"device_name"`
	Length     float64   `json:"length"`
	DeviceTime time.Time `json:"device_time"`
	Voltage    float64   `json:"voltage"`
	CreatedAt  time.Time `json:"created_at"`
}

type Record struct {
	StartTime  time.Time  `json:"start_time"`
	StopTime   time.Time  `json:"stop_time"`
	DeviceName string     `json:"device_name"`
	Length     string     `json:"length"`
	Details    DetailList `json:"details"`
}

type DetailList struct {
	Detail []DetailItem
}
type DetailItem struct {
	T string    `json:"T"` // 事件类型（如 "Fault"）
	I time.Time `json:"I"` // 时间（解析为 time.Time）
	L float64   `json:"L"` // 数值 L（解析为 float64）
	B float64   `json:"B"` // 数值 B（解析为 float64）
	F float64   `json:"F"` // 数值 F（解析为 float64）
}

func (d *Device) GetInfo() (*DeviceInfo, error) {
	q := `SELECT TOP 1 * FROM device_info WHERE device_name = ?`
	deviceInfo := &DeviceInfo{}
	err := db.Get(deviceInfo, q, d.Name)
	if err != nil {
		return nil, err
	}
	return deviceInfo, nil
}
