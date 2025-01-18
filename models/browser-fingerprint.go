package models

type MexcBrowserFingerprint struct {
	Mtoken            string  `json:"mtoken"`
	Mhash             string  `json:"mhash"` // 为mtoken的md5
	Sys               string  `json:"sys"`
	SysVer            string  `json:"sys_ver"`
	BrowserName       string  `json:"browser_name"`
	BrowserVer        string  `json:"browser_ver"`
	KernelName        string  `json:"kernel_name"`
	KernelVer         string  `json:"kernel_ver"`
	GpuType           string  `json:"gpu_type"`
	Language          string  `json:"language"`
	DisplayResolution string  `json:"display_resolution"`
	ColorDepth        string  `json:"color_depth"`
	TotalMemory       string  `json:"total_memory"`
	PixelRatio        float64 `json:"pixel_ratio"`
	TimeZone          string  `json:"time_zone"`
	SessionEnable     string  `json:"session_enable"`
	StorageEnable     string  `json:"storage_enable"`
	IndexeddbEnable   string  `json:"indexeddb_enable"`
	WebsqlEnable      string  `json:"websql_enable"`
	DoNotTrack        bool    `json:"do_not_track"`
	IsAlphago         bool    `json:"is_alphago"`
	CanvasCrc         string  `json:"canvas_crc"`
	Fonts             string  `json:"fonts"`
	EDevices          string  `json:"e_devices"`
	AudioHash         string  `json:"audio_hash"`
	WebglHash         string  `json:"webgl_hash"`
	MemberID          string  `json:"member_id"`
	EnvInfo           string  `json:"env_info"`
	Hostname          string  `json:"hostname"`
	SdkV              string  `json:"sdk_v"`
	ProductType       int     `json:"product_type"`
	PlatformType      int     `json:"platform_type"`
}

func (mfp *MexcBrowserFingerprint) ToMap() map[string]interface{} {
	// 将所有字段都放在一个map中
	dataMap := make(map[string]interface{})
	dataMap["mtoken"] = mfp.Mtoken
	dataMap["mhash"] = mfp.Mhash
	dataMap["sys"] = mfp.Sys
	dataMap["sys_ver"] = mfp.SysVer
	dataMap["browser_name"] = mfp.BrowserName
	dataMap["browser_ver"] = mfp.BrowserVer
	dataMap["kernel_name"] = mfp.KernelName
	dataMap["kernel_ver"] = mfp.KernelVer
	dataMap["gpu_type"] = mfp.GpuType
	dataMap["language"] = mfp.Language
	dataMap["display_resolution"] = mfp.DisplayResolution
	dataMap["color_depth"] = mfp.ColorDepth
	dataMap["total_memory"] = mfp.TotalMemory
	dataMap["pixel_ratio"] = mfp.PixelRatio
	dataMap["time_zone"] = mfp.TimeZone
	dataMap["session_enable"] = mfp.SessionEnable
	dataMap["storage_enable"] = mfp.StorageEnable
	dataMap["indexeddb_enable"] = mfp.IndexeddbEnable
	dataMap["websql_enable"] = mfp.WebsqlEnable
	dataMap["do_not_track"] = mfp.DoNotTrack
	dataMap["is_alphago"] = mfp.IsAlphago
	dataMap["canvas_crc"] = mfp.CanvasCrc
	dataMap["fonts"] = mfp.Fonts
	dataMap["e_devices"] = mfp.EDevices
	dataMap["audio_hash"] = mfp.AudioHash
	dataMap["webgl_hash"] = mfp.WebglHash
	dataMap["member_id"] = mfp.MemberID
	dataMap["env_info"] = mfp.EnvInfo
	dataMap["hostname"] = mfp.Hostname
	dataMap["sdk_v"] = mfp.SdkV
	dataMap["product_type"] = mfp.ProductType
	dataMap["platform_type"] = mfp.PlatformType

	return dataMap
}
