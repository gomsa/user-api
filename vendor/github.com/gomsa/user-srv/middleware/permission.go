package middleware

// Permission 权限
type Permission struct {
	Service     string `json:"service"`
	Method      string `json:"method"`
	Auth        bool   `json:"auth"`
	Policy      bool   `json:"policy"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
