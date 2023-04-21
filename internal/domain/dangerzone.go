package domain

type DangerZoneCreateReq struct {
	DeviceID  string  `json:"device_id"`
	CompanyID string  `json:"company_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
	TTL       int64   `json:"ttl"`
}

type DangerZone struct {
	DeviceID  string  `json:"device_id" bson:"device_id"`
	CompanyID string  `json:"company_id" bson:"company_id"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Radius    float64 `json:"radius" bson:"radius"`
	EndTs     int64   `json:"end_timestamp" bson:"end_ts"`
}

func (r *DangerZoneCreateReq) ToDangerZone() *DangerZone {
	return &DangerZone{
		DeviceID:  r.DeviceID,
		CompanyID: r.CompanyID,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
		Radius:    r.Radius,
	}
}
func (r *DangerZoneCreateReq) IsValid() bool {
	return r.DeviceID != "" &&
		r.CompanyID != "" &&
		r.Radius > 0 &&
		r.TTL > 0
}
