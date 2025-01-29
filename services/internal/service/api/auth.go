package api

import (
	"context"
	"fmt"
	"net"
	"time"

	"app/api/api"
	"app/api/job"
	"app/api/model"
	"app/component/counter"
	"app/pkg/alert"
	"app/pkg/errx"
	"app/pkg/ginx"
	"app/pkg/jwtutil"
	"app/pkg/logger"
	"app/pkg/password"
	"app/pkg/randx"
	"app/services/internal/cache"
	"app/services/internal/component/email"
	"app/services/internal/component/lang"
	"app/services/internal/component/twofa"
	"app/services/internal/config"
	"app/services/internal/pkg/aws_ses"
	"app/services/internal/repo"

	"github.com/go-redis/redis/v8"

	"github.com/mojocn/base64Captcha"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

var _ api.AuthServer = (*Auth)(nil)

type Auth struct {
	Captcha                  *base64Captcha.Captcha
	UserRepo                 *repo.UserRepo
	UserTokenRepo            *repo.UserTokenRepo
	AwsSes                   *aws_ses.Client
	UserRegisterPublisher    job.UserRegisterPublisher
	DB                       *gorm.DB
	Config                   *config.Config
	Email                    *email.Email
	TwoFA                    *twofa.TwoFA
	UserResetPasswordLimiter *cache.UserResetPasswordLimiter
	Redis                    *redis.Client
	Counter                  *counter.Counter
}

func (a *Auth) CheckInviteCode(ctx context.Context, req *api.AuthCheckInviteCodeRequest) (*api.AuthCheckInviteCodeResponse, error) {
	has, err := a.UserRepo.Exists(ctx, "invite_code = ?", req.InviteCode)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthCheckInviteCodeResponse{
		Exists: lo.Ternary[string](has, model.Yes, model.No),
	}, nil
}

func (a *Auth) SafeQuestionVerify(ctx context.Context, req *api.AuthSafeQuestionVerifyRequest) (*api.AuthSafeQuestionVerifyResponse, error) {

	return &api.AuthSafeQuestionVerifyResponse{
		Token: "token",
	}, nil
}

// CheckUsernameV2 产品经理强烈要求，不要限制次数
// 出问题了，请找产品经理
func (a *Auth) CheckUsernameV2(ctx context.Context, req *api.AuthCheckUsernameRequest) (*api.AuthCheckUsernameResponseV2, error) {
	has, err := a.UserRepo.Exists(ctx, "username = ?", req.Username)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthCheckUsernameResponseV2{
		Exists: lo.Ternary[string](has, model.Yes, model.No),
	}, nil
}

func (a *Auth) CheckUsername(ctx context.Context, req *api.AuthCheckUsernameRequest) (*api.AuthCheckUsernameResponse, error) {
	c := ginx.FromContext(ctx)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		remoteAddr := c.Request.RemoteAddr
		// remove port
		var err error
		if ip, _, err = net.SplitHostPort(remoteAddr); err != nil {
			return nil, errx.New(api.ErrBusiness, err)
		}
	}
	// ip limit 3 times
	times, err := a.Counter.Incr(ctx, fmt.Sprintf("check_username:%s:%s", req.Username, ip), time.Hour*24)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	if times > 3 {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "TRY_TOO_MANY_TIMES"))
	}
	remainAttempts := 3 - int(times)

	has, err := a.UserRepo.Exists(ctx, "username = ?", req.Username)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthCheckUsernameResponse{
		Exists:         lo.Ternary[string](has, model.Yes, model.No),
		RemainAttempts: remainAttempts,
	}, nil
}

func (a *Auth) RegisterSafeQuestion(ctx context.Context, req *api.AuthRegisterSafeQuestionRequest) (*api.AuthRegisterSafeQuestionResponse, error) {

	return &api.AuthRegisterSafeQuestionResponse{}, nil
}

func (a *Auth) EmailSend(ctx context.Context, req *api.AuthEmailSendRequest) (*api.AuthEmailSendResponse, error) {
	if err := a.Email.Send(ctx, req.Email, email.Scene(req.Scene)); err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthEmailSendResponse{}, nil
}

func (a *Auth) CaptchaSend(ctx context.Context, req *api.AuthCaptchaSendRequest) (*api.AuthCaptchaSendResponse, error) {
	id, b64s, err := a.Captcha.Generate()
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthCaptchaSendResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, nil
}

func (a *Auth) PasswordReset(ctx context.Context, req *api.AuthPasswordResetRequest) (*api.AuthPasswordResetResponse, error) {
	// find user
	user, err := a.UserRepo.Find(ctx, "username = ?", req.Username)
	if err == gorm.ErrRecordNotFound {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_EMPTY"))
	}
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}

	// check password
	if req.NewPassword != req.ConfirmPassword {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "PASSWORD_IS_NOT_SAME"))
	}
	// reset password
	pass, err := password.GeneratePassword(req.NewPassword)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	if err = a.UserResetPasswordLimiter.Incr(ctx, user.UserId); err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	// update user
	_, err = a.UserRepo.UpdateById(ctx, &model.User{
		UserId:   user.UserId,
		Password: string(pass),
	})
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthPasswordResetResponse{}, nil
}

func (a *Auth) PasswordResetV2(ctx context.Context, req *api.AuthPasswordResetRequestV2) (*api.AuthPasswordResetResponse, error) {
	return &api.AuthPasswordResetResponse{}, nil
}

func (a *Auth) Login(ctx context.Context, req *api.AuthLoginRequest) (*api.AuthLoginResponse, error) {
	// check req
	if req.Username == "" || req.Password == "" {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_EMPTY"))
	}

	var ip = ginx.IpFromContext(ctx)
	// counter
	lockKey := fmt.Sprintf("login:%s:%s", req.Username, ip)
	times, err := a.Counter.Get(ctx, lockKey)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}

	var incLock = func() {
		_, err = a.Counter.Incr(ctx, lockKey, time.Minute*10)
		if err != nil {
			alert.Alert(ctx, "incr login times error", []string{
				fmt.Sprintf("username: %s", req.Username),
			})
		}
	}

	if times > 5 {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "TRY_TOO_MANY_TIMES"))
	}

	// get user
	user, err := a.UserRepo.Find(ctx, "username = ?", req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			incLock()
			return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_ERROR"))
		}
		logger.Printf(ctx, "get user error: %s", err)
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_ERROR"))
	}
	ok, err := a.isAppStoreCheatDeleteUser(ctx, user.UserId)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	if ok {
		return nil, errx.New(api.ErrUserLogOff, lang.T(ctx, "USER_IS_LOG_OFF"))
	}
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	// check password count

	// check username and password
	if err = password.ValidatePassword(req.Password, user.Password); err != nil {
		incLock()
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_ERROR"))
	}

	// check status
	if user.Status == model.UserStatusOff {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "USER_IS_DISABLE"))
	}

	token, err := jwtutil.Sign(a.Config.Jwt.OpenApiSecret, time.Hour*24*30, &api.AuthUser{
		UserId: user.UserId,
	})
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	// update user token
	err = a.DB.Exec("insert into user_token (user_id, token, ip, created_at) values (?, ?, ?, ?) on duplicate key update token = ?, ip = ?, updated_at = ?",
		user.UserId, token, ip, time.Now(), token, ip, time.Now()).Error
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	// generate token
	return &api.AuthLoginResponse{
		Token:        lo.Ternary[string](user.SecretTwoFA != "", "", token),
		Is2FAEnabled: lo.Ternary[string](user.SecretTwoFA != "", model.Yes, model.No),
		Avatar:       lo.Ternary[string](user.SecretTwoFA != "", "", user.Avatar),
		Nickname:     lo.Ternary[string](user.SecretTwoFA != "", "", user.Nickname),
		CustomerId:   lo.Ternary[string](user.SecretTwoFA != "", "", user.CustomerId),
		Username:     lo.Ternary[string](user.SecretTwoFA != "", "", user.Username),
	}, nil
}

const superInviteCode = "cG9kJ4kO4hN4iQ6gV3sZ8vR0fC5pE1pS"

func (a *Auth) Register(ctx context.Context, req *api.AuthRegisterRequest) (*api.AuthRegisterResponse, error) {
	// check invite code
	if req.InviteCode == "" {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "INVITE_CODE_IS_EMPTY"))
	}

	var pid int
	if req.InviteCode != superInviteCode {
		parent, err := a.UserRepo.Find(ctx, "invite_code = ?", req.InviteCode)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "INVITE_CODE_IS_NOT_EXISTS"))
			}
			return nil, errx.New(api.ErrBusiness, err)
		}
		pid = parent.UserId
	}

	// verify email verification code
	// if a.Config.Env != "local" {
	// 	if err := a.Email.Verify(ctx, req.Email, email.Register, req.EmailVerificationCode, true); err != nil {
	// 		return nil, errx.New(api.ErrInvalidParam, err)
	// 	}
	// }
	// find user
	has, err := a.UserRepo.Exists(ctx, "username = ?", req.Username)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	if has {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_IS_EXISTS"))
	}
	// generate password
	pass, err := password.GeneratePassword(req.Password)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	var userId int
	// query question

	err = a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// create user
		u := &model.User{
			Pid:              pid,
			CustomerId:       "customId",
			Status:           model.UserStatusOn,
			Username:         req.Username,
			Nickname:         req.Username,
			Avatar:           "",
			Password:         string(pass),
			IsValid:          model.No,
			EnableLevelGrade: model.Yes,
			IsClubOwner:      model.No,
			KycStatus:        model.UserKycStatusUnauth,
		}
		if err = a.UserRepo.TxCreate(ctx, a.DB, u); err != nil {
			return err
		}

		// update invite code
		if _, err = a.UserRepo.TxUpdateById(ctx, a.DB, &model.User{
			UserId:     u.UserId,
			InviteCode: fmt.Sprintf("%s%04d", randx.Alpha(2), u.UserId),
		}); err != nil {
			return err
		}

		userId = u.UserId
		// pub
		err = a.UserRegisterPublisher.Publish(ctx, &job.UserRegisterRequest{
			UserId: u.UserId,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	token, err := jwtutil.Sign(a.Config.Jwt.OpenApiSecret, time.Hour*24*30, &api.AuthUser{
		UserId: userId,
	})
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthRegisterResponse{
		Token: token,
	}, nil
}

func (a *Auth) TwoFALogin(ctx context.Context, req *api.AuthTwoFALoginRequest) (*api.AuthTwoFALoginResponse, error) {

	// get user
	user, err := a.UserRepo.Find(ctx, "username = ?", req.UserName)
	if err == gorm.ErrRecordNotFound {
		return nil, errx.New(api.ErrInvalidParam, "username or password is error")
	}
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	ok, err := a.isAppStoreCheatDeleteUser(ctx, user.UserId)
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	if ok {
		return nil, errx.New(api.ErrUserLogOff, lang.T(ctx, "USER_IS_LOG_OFF"))
	}
	// check username and password
	if err = password.ValidatePassword(req.Password, user.Password); err != nil {
		return nil, errx.New(api.ErrInvalidParam, lang.T(ctx, "USERNAME_OR_PASSWORD_IS_EMPTY"))
	}

	// check status
	if user.Status == model.UserStatusOff {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "USER_IS_DISABLE"))
	}

	// check 2fa
	if user.SecretTwoFA == "" {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "2FA_IS_NOT_ENABLED"))
	}

	// check 2fa code
	if !a.TwoFA.Verify(ctx, req.Passcode, user.SecretTwoFA) {
		return nil, errx.New(api.ErrBusiness, lang.T(ctx, "2FA_CODE_IS_ERROR"))
	}

	token, err := jwtutil.Sign(a.Config.Jwt.OpenApiSecret, time.Hour*24*30, &api.AuthUser{
		UserId: user.UserId,
	})
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	ip := ginx.IpFromContext(ctx)
	// update user token
	err = a.DB.Exec("insert into user_token (user_id, token, ip, created_at) values (?, ?, ?, ?) on duplicate key update token = ?, ip = ?, created_at = ?",
		user.UserId, token, ip, time.Now(), token, ip, time.Now()).Error
	if err != nil {
		return nil, errx.New(api.ErrBusiness, err)
	}
	return &api.AuthTwoFALoginResponse{
		Token:      token,
		Avatar:     user.Avatar,
		Nickname:   user.Nickname,
		CustomerId: user.CustomerId,
		Username:   user.Username,
	}, nil
}

func (a *Auth) isAppStoreCheatDeleteUser(ctx context.Context, userid int) (bool, error) {
	// exists
	return false, nil
}

func genCustomId() string {
	// YY + 7位随机数
	return fmt.Sprintf("%s%s", time.Now().Format("06"), randx.Digit(7))
}
