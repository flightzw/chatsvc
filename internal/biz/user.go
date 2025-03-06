package biz

import (
	"context"
	"crypto/rsa"
	"fmt"
	"os"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/flightzw/chatsvc/internal/biz/query"
	"github.com/flightzw/chatsvc/internal/conf"
	"github.com/flightzw/chatsvc/internal/entity"
	"github.com/flightzw/chatsvc/internal/enum"
	"github.com/flightzw/chatsvc/internal/utils/hash"
	"github.com/flightzw/chatsvc/internal/utils/jwt"
	"github.com/flightzw/chatsvc/internal/vo"
	"github.com/flightzw/chatsvc/internal/ws/client"
)

// User
type User struct {
	ID          int32              `json:"id"`
	Username    string             `json:"username"`      // 用户名
	Password    string             `json:"password"`      // 密码
	AvatarURL   string             `json:"avatar_url"`    // 头像url
	Nickname    string             `json:"nickname"`      // 昵称
	Gender      int32              `json:"gender"`        // 性别：0未知，1男，2女
	Signature   string             `json:"signature"`     // 个性签名
	Status      enum.AccountStatus `json:"status"`        // 状态: 1正常，2封禁
	LastLoginAt *gtime.Time        `json:"last_login_at"` // 最后上线时间
	LastLoginIP string             `json:"last_login_ip"` // 最后上线ip
	CreatedAt   *gtime.Time        `json:"created_at"`    // 创建时间
	UpdatedAt   *gtime.Time        `json:"updated_at"`    // 更新时间
	DeletedAt   *gtime.Time        `json:"deleted_at"`    // 删除时间
}

// UserRepo
type UserRepo interface {
	TransactionInterface

	CreateUser(ctx context.Context, user *User) (id int32, err error)
	GetUser(ctx context.Context, id int32) (data *User, err error)
	GetUserByUsername(ctx context.Context, username string) (data *User, err error)
	ListUser(ctx context.Context, queryFunc query.QueryFunc, page, pageSize int) (data []*User, total int64, err error)
	UpdateUser(ctx context.Context, id int32, data entity.AnyMap) (err error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo       UserRepo
	friendRepo FriendRepo

	log        *log.Helper
	conf       *conf.Server
	chatClient *client.ChatClient

	privateKey        *rsa.PrivateKey
	refreshPrivateKey *rsa.PrivateKey
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, friendRepo FriendRepo, conf *conf.Server, chatClient *client.ChatClient, logger log.Logger) *UserUsecase {
	helper := log.NewHelper(log.With(logger, "module", "chatsvc/biz/UserUsecase"))

	prifileContent, _ := os.ReadFile(conf.Jwt.AccessToken.Prifile)
	privateKey, err := jwtv5.ParseRSAPrivateKeyFromPEM(prifileContent)
	if err != nil {
		helper.Fatalf("Failed to parse private_key: %v", err)
	}
	refreshPrifileContent, _ := os.ReadFile(conf.Jwt.RefreshToken.Prifile)
	refreshPrivateKey, err := jwtv5.ParseRSAPrivateKeyFromPEM(refreshPrifileContent)
	if err != nil {
		helper.Fatalf("Failed to parse refresh_private_key: %v", err)
	}
	return &UserUsecase{
		repo:              repo,
		friendRepo:        friendRepo,
		log:               helper,
		conf:              conf,
		chatClient:        chatClient,
		privateKey:        privateKey,
		refreshPrivateKey: refreshPrivateKey,
	}
}

func (uc *UserUsecase) Register(ctx context.Context, user *User) error {
	if data, _ := uc.repo.GetUserByUsername(ctx, user.Username); data != nil {
		return errno.ErrorUserUnameBeUsed("用户名已被使用")
	}

	user.Password = generatePasswordStr(user.Password)
	user.Status = enum.AccountStatusNormal
	if _, err := uc.repo.CreateUser(ctx, user); err != nil {
		return errno.ErrorUserRegisterFailed("用户注册时出错").WithCause(err)
	}
	return nil
}

func (uc *UserUsecase) Login(ctx context.Context, user *User, rememberMe bool) (*vo.LoginVO, error) {
	data, err := uc.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, errno.ErrorUserNotFound("用户名或密码错误").WithCause(err)
	}
	hashPassword, salt := getHashPasswordAndSalt(data)
	if !hash.PasswordCheck(hashPassword, user.Password, salt) {
		return nil, errno.ErrorUserNotFound("用户名或密码错误")
	}
	if data.Status != enum.AccountStatusNormal {
		return nil, errno.ErrorParamInvalid(fmt.Sprintf("账号异常（状态：%s）", data.Status.Map()))
	}

	token, refreshToken, err := uc.generateUserToken(ctx, data, rememberMe)
	if err != nil {
		return nil, errno.ErrorUserTokenSignFailed("签发认证凭据时出错").WithCause(err)
	}
	return &vo.LoginVO{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserUsecase) RefreshToken(ctx context.Context) (*vo.LoginVO, error) {
	id, _ := jwt.GetUserInfo(ctx)
	if id == 0 {
		return nil, errno.ErrorUserNotFound("用户凭证数据有误")
	}

	data, err := uc.repo.GetUser(ctx, int32(id))
	if err != nil {
		return nil, errno.ErrorUserNotFound("未找到有效的用户信息").WithCause(err)
	}
	token, refreshToken, err := uc.generateUserToken(ctx, data, true)
	if err != nil {
		return nil, errno.ErrorUserTokenSignFailed("签发认证凭据时出错").WithCause(err)
	}
	return &vo.LoginVO{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserUsecase) GetUserSelf(ctx context.Context) (data *User, err error) {
	id, _ := jwt.GetUserInfo(ctx)
	if id == 0 {
		return nil, errno.ErrorUserNotFound("用户凭证数据有误")
	}
	if data, err = uc.repo.GetUser(ctx, id); err != nil {
		return nil, errno.ErrorUserNotFound("未找到有效的用户信息").WithCause(err)
	}
	return data, nil
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int32) (data *vo.UserVO, err error) {
	user, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		return nil, errno.ErrorUserNotFound("未找到有效的用户信息").WithCause(err)
	}
	onlineMap := uc.chatClient.IsOnline(ctx, user.ID)
	return newUserVO(user, onlineMap[id]), nil
}

func (uc *UserUsecase) ListUserInfo(ctx context.Context, name string) (data []*vo.UserVO, err error) {
	qf := func(do query.QueryChain) query.QueryChain {
		u := query.NewUserQuery()
		return do.Where(u.Username.Like("%" + name + "%")).Or(u.Nickname.Like("%" + name + "%"))
	}
	users, _, err := uc.repo.ListUser(ctx, qf, 1, 100)
	if err != nil {
		return nil, errno.ErrorDataQueryFailed("获取用户列表时出错").WithCause(err)
	}
	userIds := make([]int32, 0, len(users))
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}
	onlineMap := uc.chatClient.IsOnline(ctx, userIds...)
	for _, user := range users {
		data = append(data, newUserVO(user, onlineMap[user.ID]))
	}
	return data, nil
}

func (uc *UserUsecase) UpdateUserInfo(ctx context.Context, user *User) (err error) {
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Signature = strings.TrimSpace(user.Signature)
	return uc.friendRepo.Transaction(ctx, func(ctx context.Context) error {
		userID, _ := jwt.GetUserInfo(ctx)
		// 更新 Friend
		if err = uc.friendRepo.UpdateFriendByFriendID(ctx, userID, entity.AnyMap{
			"friend_nickname": user.Nickname,
		}); err != nil {
			return errno.ErrorDataUpdateFailed("更新好友信息时出错").WithCause(err)
		}
		// 更新 User
		if err = uc.repo.UpdateUser(ctx, userID, entity.AnyMap{
			"nickname":  user.Nickname,
			"gender":    user.Gender,
			"signature": user.Signature,
		}); err != nil {
			return errno.ErrorDataUpdateFailed("更新用户信息时出错").WithCause(err)
		}
		return nil
	})
}

func (uc *UserUsecase) UpdatePassword(ctx context.Context, oldPassword, newPassword string) (err error) {
	if oldPassword == newPassword {
		return errno.ErrorParamInvalid("新旧密码不得相同")
	}

	userID, _ := jwt.GetUserInfo(ctx)
	data, err := uc.repo.GetUser(ctx, userID)
	if err != nil {
		return errno.ErrorUserNotFound("获取用户信息时出错").WithCause(err)
	}
	hashPassword, salt := getHashPasswordAndSalt(data)
	if !hash.PasswordCheck(hashPassword, oldPassword, salt) {
		return errno.ErrorParamInvalid("密码错误，请重新输入")
	}
	if err = uc.repo.UpdateUser(ctx, userID, entity.AnyMap{
		"password": generatePasswordStr(newPassword),
	}); err != nil {
		return errno.ErrorDataUpdateFailed("修改密码时出错").WithCause(err)
	}
	return nil
}

func (uc *UserUsecase) generateUserToken(_ context.Context, user *User, rememberMe bool) (token, refreshToken string, err error) {
	signMethod := jwtv5.SigningMethodRS256
	token, err = jwt.SignedString(signMethod, user.ID, user.Username,
		uc.conf.Jwt.AccessToken.ExpireIn,
		uc.privateKey,
	)
	if err != nil {
		return "", "", err
	}
	if rememberMe {
		refreshToken, err = jwt.SignedString(signMethod, user.ID, user.Username,
			uc.conf.Jwt.RefreshToken.ExpireIn,
			uc.refreshPrivateKey,
		)
		if err != nil {
			return "", "", err
		}
	}
	return
}

func getHashPasswordAndSalt(user *User) (string, string) {
	items := strings.Split(user.Password, ":")
	return items[0], items[1]
}

func newUserVO(user *User, isOnline bool) *vo.UserVO {
	return &vo.UserVO{
		ID:        user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		Nickname:  user.Nickname,
		Gender:    user.Gender,
		Signature: user.Signature,
		IsOnline:  isOnline,
	}
}
