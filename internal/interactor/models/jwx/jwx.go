package jwx

// JWX struct is used to create token
type JWX struct {
	// 使用者ID
	UserID *string `json:"user_id,omitempty"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 角色
	Role *string `json:"role,omitempty"`
}

// Token return structure file
type Token struct {
	// 授權令牌
	AccessToken string `json:"access_token,omitempty"`
	// 刷新令牌
	RefreshToken string `json:"refresh_token,omitempty"`
	// 角色
	Role string `json:"role,omitempty"`
}

// Refresh struct is used to refresh token
type Refresh struct {
	// 刷新令牌
	RefreshToken string `json:"refresh_token,omitempty" binding:"required" validate:"required"`
}
