package api

import "context"

// AuthServer 认证服务(release)
type AuthServer interface {
	//x:api post /api/auth/login 登录
	Login(ctx context.Context, req *AuthLoginRequest) (*AuthLoginResponse, error)
	//x:api post /api/auth/register 注册
	Register(ctx context.Context, req *AuthRegisterRequest) (*AuthRegisterResponse, error)
	//x:api post /api/auth/register/safe_question 安全问题
	RegisterSafeQuestion(ctx context.Context, req *AuthRegisterSafeQuestionRequest) (*AuthRegisterSafeQuestionResponse, error)
	//x:api post /api/auth/safe_question/verify 安全问题验证
	SafeQuestionVerify(ctx context.Context, req *AuthSafeQuestionVerifyRequest) (*AuthSafeQuestionVerifyResponse, error)
	//x:api post /api/auth/email/send 发送邮箱验证码
	EmailSend(ctx context.Context, req *AuthEmailSendRequest) (*AuthEmailSendResponse, error)
	//x:api post /api/auth/captcha/send 发送图形验证码
	CaptchaSend(ctx context.Context, req *AuthCaptchaSendRequest) (*AuthCaptchaSendResponse, error)
	//x:api post /api/auth/password/reset 重置密码
	PasswordReset(ctx context.Context, req *AuthPasswordResetRequest) (*AuthPasswordResetResponse, error)
	//x:api post /api/auth/password/reset/v2 重置密码
	PasswordResetV2(ctx context.Context, req *AuthPasswordResetRequestV2) (*AuthPasswordResetResponse, error)

	//x:api post /api/auth/two_fa_login 两步验证登录
	TwoFALogin(ctx context.Context, req *AuthTwoFALoginRequest) (*AuthTwoFALoginResponse, error)
	//x:api post /api/auth/check_username 检查用户名是否存在
	CheckUsername(ctx context.Context, req *AuthCheckUsernameRequest) (*AuthCheckUsernameResponse, error)
	//x:api post /api/auth/check_username/v2 检查用户名是否存在
	CheckUsernameV2(ctx context.Context, req *AuthCheckUsernameRequest) (*AuthCheckUsernameResponseV2, error)
	//x:api post /api/auth/check_invite_code 检查邀请码是否存在
	CheckInviteCode(ctx context.Context, req *AuthCheckInviteCodeRequest) (*AuthCheckInviteCodeResponse, error)
}

type AuthCheckInviteCodeRequest struct {
	// invite code 邀请码
	InviteCode string `json:"invite_code" binding:"required"`
}

type AuthCheckInviteCodeResponse struct {
	// exists: Y / N
	Exists string `json:"exists"`
}

type AuthSafeQuestionVerifyRequest struct {
	// username
	Username string `json:"username" binding:"required"`
	// answers
	Answers []*AnswerItem `json:"answers" binding:"required"`
}

type AuthSafeQuestionVerifyResponse struct {
	// token
	Token string `json:"token"`
}

type AuthCheckUsernameRequest struct {
	// username
	Username string `json:"username" binding:"required"`
}

type AuthCheckUsernameResponse struct {
	// exists: Y / N
	Exists string `json:"exists"`
	// remain attempts
	RemainAttempts int `json:"remain_attempts"`
}

type AuthCheckUsernameResponseV2 struct {
	// exists: Y / N
	Exists string `json:"exists"`
}

type AuthRegisterSafeQuestionRequest struct{}

type QuestionItem struct {
	// question_id
	QuestionId int `json:"question_id"`
	// title
	Title string `json:"title"`
}

type AuthRegisterSafeQuestionResponse struct {
	// question1
	Question1 []*QuestionItem `json:"question1"`
	// question2
	Question2 []*QuestionItem `json:"question2"`
	// question3
	Question3 []*QuestionItem `json:"question3"`
	// rule
	Rule string `json:"rule"`
	// Tnc
	Tnc string `json:"tnc"`
	// privacy_policy
	PrivacyPolicy string `json:"privacy_policy"`
}

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AuthLoginResponse struct {
	Token        string `json:"token"`
	Is2FAEnabled string `json:"is_2fa_enabled"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	CustomerId   string `json:"customer_id"`
	Nickname     string `json:"nickname"`
}

type AnswerItem struct {
	// question_id
	QuestionId int `json:"question_id"`
	// answer
	Answer string `json:"answer"`
}

type AuthRegisterRequest struct {
	// username
	Username string `json:"username" binding:"required"`
	// password
	Password string `json:"password" binding:"required"`
	// confirm password
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	// invite code 邀请码
	InviteCode string `json:"invite_code" binding:"required"`
	// answers
	Answers []*AnswerItem `json:"answers" binding:"required"`
}

type AuthRegisterResponse struct {
	Token string `json:"token"`
}

type AuthEmailSendRequest struct {
	// email
	Email string `json:"email" validate:"required,email"`
	// scene 验证码场景 LOGIN,REGISTER,FORGOT_PASSWORD, FORGOT_PAY_PASSWORD, SET_PAY_PASSWORD
	Scene string `json:"scene" validate:"required,oneof=LOGIN REGISTER FORGOT_PASSWORD FORGOT_PAY_PASSWORD SET_PAY_PASSWORD"`
}

type AuthEmailSendResponse struct{}

type AuthCaptchaSendRequest struct{}

type AuthCaptchaSendResponse struct {
	// captcha id
	CaptchaId string `json:"captcha_id"`
	// captcha image
	CaptchaImage string `json:"captcha_image"`
}

type AuthPasswordResetRequestV2 struct {
	// new password
	NewPassword string `json:"new_password" binding:"required"`
	// confirm password
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	// token
	Token string `json:"token" binding:"required"`
}

type AuthPasswordResetRequest struct {
	// username
	Username string `json:"username" binding:"required"`
	// new password
	NewPassword string `json:"new_password" binding:"required"`
	// confirm password
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	// answers
	Answers []*AnswerItem `json:"answers" binding:"required"`
}

type AuthPasswordResetResponse struct{}

type AuthTwoFALoginRequest struct {
	// Email
	UserName string `json:"user_name"`
	// password
	Password string `json:"password"`
	// passcode
	Passcode string `json:"passcode"`
}

type AuthTwoFALoginResponse struct {
	Token      string `json:"token"`
	Avatar     string `json:"avatar"`
	Username   string `json:"username"`
	CustomerId string `json:"customer_id"`
	Nickname   string `json:"nickname"`
}
