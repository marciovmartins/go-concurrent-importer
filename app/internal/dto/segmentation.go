package dto

type UserSegmentationsResponse struct {
	UserID        int64              `json:"user_id"`
	Segmentations []UserSegmentation `json:"segmentations"`
}

type UserSegmentation struct {
	Patients    []Patient   `json:"patients"`
	Specialties []Specialty `json:"specialties"`
	Drugs       []Drug      `json:"drugs"`
}

type Patient struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Specialty struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type Drug struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}
