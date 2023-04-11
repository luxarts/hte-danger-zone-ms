package domain

type DangerZoneCreateReq struct {
	DeviceID  string  `json:"device_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
	TTL       int64   `json:"ttl"`
}

type DangerZone struct {
	DeviceID  string  `json:"did"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    float64 `json:"r"`
	EndTs     int64   `json:"e_ts"`
}

func (r *DangerZoneCreateReq) ToDangerZone() *DangerZone {
	return &DangerZone{
		DeviceID:  r.DeviceID,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Radius:    r.Radius,
	}
}
