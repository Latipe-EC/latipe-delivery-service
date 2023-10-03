package entities

type ProvinceDetail struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Type         string `json:"type"`
	NameWithType string `json:"name_with_type"`
	Code         string `json:"code"`
}

type DistrictDetail struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Slug         string `json:"slug"`
	NameWithType string `json:"name_with_type"`
	Path         string `json:"path"`
	PathWithType string `json:"path_with_type"`
	Code         string `json:"code"`
	ParentCode   string `json:"parent_code"`
}

type WardDetail struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Slug         string `json:"slug"`
	NameWithType string `json:"name_with_type"`
	Path         string `json:"path"`
	PathWithType string `json:"path_with_type"`
	Code         string `json:"code"`
	ParentCode   string `json:"parent_code"`
}
