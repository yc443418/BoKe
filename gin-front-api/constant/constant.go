// 系统常量定义
// @author chen

package constant

const (
	LoginCode         = "loginCode:"      // 登录验证码key
	EmailCode         = "emailCode:"      // 邮箱验证码key
	EmailSubjectKey   = "emailSubject"    // 邮箱主体key
	RegisterMessage   = "registerMessage" // 用户注册语key
	Reply             = 1                 // 回复我的
	StarArticle       = 2                 // 赞了文章
	StarComment       = 3                 // 赞了评论
	SystemMessage     = 4                 // 系统消息
	NotReadStatus     = 1                 // 未读
	ReadStatus        = 2                 // 已读
	ContextKeyUserObj = "authedUserObj"   // 用户认证标识
	Status            = 2                 // 用户状态（1-正常，2-停用）
	NotTopType        = 1                 // 未置顶
	TopType           = 2                 // 置顶
	Delete            = 1                 // 已删除
	NotDelete         = 2                 // 未删除
	HOT               = 1                 // 热榜
	PUBLISH           = 2                 // 发布时间
	NEW               = 3                 // 最新
	LikeSuccess       = "点赞成功"            // 点赞
	NotLikeSuccess    = "取消点赞成功"          // 取消点赞
	ZERO              = 0                 // 父级
	LikeComment       = 1                 // 未点赞
	CancelLikeComment = 2                 // 已点赞
	ArticleList       = 1                 // 文章列表
	CommentList       = 2                 // 评论列表
	StarList          = 3                 // 点赞列表
)
