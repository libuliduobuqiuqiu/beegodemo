package types

type DeviceNodeInfo struct {
	Type          string `json:"provider_name" db:"provider_name"`
	Platform      string `json:"platform" db:"platform"` //Platform : 平台
	Version       string `json:"version" db:"version"`
	LicenseStatus string `json:"license_status" db:"license_status"`
	SerialNumber  string `json:"serial_number" db:"sn"` //序列号
	HostName      string `json:"host_name" db:"host_name"`
	Name          string `json:"name" db:"name"`
	Status        string `json:"status" db:"status"`
	DeviceStatus  string `json:"dev_status" db:"dev_status"` //设备返回的原始数据
}

type Device struct {
	ID        string `json:"id" db:"id"`
	Address   string `json:"address" db:"ip"`            // 设备访问地址
	ApiPort   int    `json:"api_port" db:"api_port"`     // 设备访问API端口
	SSHPort   int    `json:"ssh_port" db:"ssh_port"`     // 设备SSH端口
	ProjectID string `json:"project_id" db:"project_id"` // 设备项目ID
	WebUrl    string `json:"web_url" db:"web_url"`       // Web管理页面地址
	DeviceNodeInfo
}
