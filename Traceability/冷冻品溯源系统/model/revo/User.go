package revo

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username    string  `json:"username" binding:"required,min=5,max=50"`                                     // 用户名
	Password    string  `json:"password" binding:"required,min=8,max=50"`                                     // 密码
	Confirm     string  `json:"confirm" binding:"required,eqfield=Password"`                                  // 确认密码
	UserType    string  `json:"user_type" binding:"required,oneof=factory dealer supervision consumer admin"` // 用户类型
	ContactInfo Contact `json:"contact_info" binding:"required"`                                              // 联系信息
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// UserResponse 用户响应结构体
type UserResponse struct {
	UserID    string  `json:"user_id"`    // 用户ID
	Username  string  `json:"username"`   // 用户名
	UserType  string  `json:"user_type"`  // 用户类型
	Status    int     `json:"status"`     // 用户状态
	CreatedAt string  `json:"created_at"` // 创建时间
	Contact   Contact `json:"contact"`    // 联系信息
}

// Contact 联系信息结构体
type Contact struct {
	Name  string `json:"name"`  // 联系人姓名
	Email string `json:"email"` // 电子邮箱
	Phone string `json:"phone"` // 联系电话
}
