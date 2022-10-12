package account

type UserAccountInfo struct {
	Id        string `json:"id"`
	LoginName string `json:"login_name"`
	CnName    string `json:"name"`
	Phone     string `josn:"phone"`
	RoleId    string `json:"role_id"`
	RoleName  string `json:"role_name"`
}
