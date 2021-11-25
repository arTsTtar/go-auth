package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-auth/entity"
	"go-auth/models/dto/request"
	"go-auth/models/dto/response"
	enumRole "go-auth/models/struct/role"
	"go-auth/repository"
	"go-auth/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(data map[string]string) (*fiber.Cookie, error, int)
	BackupCodeLogin(data map[string]string) (*fiber.Cookie, error, int)
	Register(c request.UserRequest) (response.UserCreationResponse, error)
	Logout() *fiber.Cookie
	GetUserDetailsFromToken(token jwt.Token) (response.SimpleUserResponse, error)
	ChangePassword(id string, data request.ChangePassword) (*entity.User, error)
	Disable2FA(user *entity.User) error
}

type authService struct {
	backupCodeRepository repository.BackupCodeRepository
	userRepository       repository.UserRepository
	roleRepository       repository.RoleRepository
}

func NewAuthService(ur repository.UserRepository, rr repository.RoleRepository,
	bcr repository.BackupCodeRepository) AuthService {
	return authService{
		userRepository:       ur,
		roleRepository:       rr,
		backupCodeRepository: bcr,
	}
}

func (a authService) Login(data map[string]string) (*fiber.Cookie, error, int) {
	user, err := a.getUserByEmail(data["email"])
	if err != nil {
		return nil, err, 400
	}

	err = utils.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil || user.Id == 0 {
		return nil, err, 400
	}

	return utils.CreateAuthCookieAndHandleError(user, 30)
}

func (a authService) BackupCodeLogin(data map[string]string) (*fiber.Cookie, error, int) {

	user, err := a.getUserByEmail(data["email"])
	if err != nil {
		return nil, err, 400
	}

	var backupCodes entity.BackupCodes
	backupCodes, err = a.backupCodeRepository.FindByUser(*user)

	for i := 0; i < len(backupCodes); i++ {
		err = utils.CompareHashAndPassword(backupCodes[i].BackupCode, []byte(data["backupCode"]))
		if err != nil {
			continue
		}

		err := a.backupCodeRepository.DeleteById(backupCodes[i].Id)

		if err != nil {
			return nil, errors.New("could not delete backup code"), 500
		}

		return utils.CreateAuthCookieAndHandleError(user, 5)
	}
	return nil, errors.New("unauthorized"), 400
}

func (a authService) Register(data request.UserRequest) (response.UserCreationResponse, error) {
	var existingUser entity.User
	existingUser, _ = a.userRepository.FindUserByEmail(*data.Email)

	if existingUser.Id != 0 {
		return response.UserCreationResponse{}, fiber.NewError(400, "Client already exists")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(*data.Password), 12)

	var backupPasswords = utils.CreateBackupCodes()

	var qrData = utils.GenerateB64Qr(data)

	clientRole, err := a.roleRepository.FindByName(enumRole.Client.String())

	if err != nil {
		return response.UserCreationResponse{}, fiber.NewError(500, "Role not found - "+enumRole.Client.String())
	}
	user := entity.User{
		Name:           *data.Name,
		Email:          *data.Email,
		Password:       password,
		Roles:          []entity.Role{clientRole},
		TwoFactEnabled: qrData.TwoFactEnabled,
		TwoFactSecret:  qrData.Secret,
	}
	user, _ = a.userRepository.Save(user)

	var userBackupCodes = entity.BackupCodes{}

	for i := 0; i < len(backupPasswords); i++ {
		backupPasswd, _ := bcrypt.GenerateFromPassword([]byte(backupPasswords[i]), 12)
		backupCode := entity.BackupCode{
			UserId:     user.Id,
			BackupCode: backupPasswd,
		}
		userBackupCodes = append(userBackupCodes, &backupCode)
	}
	_, _ = a.backupCodeRepository.SaveAll(userBackupCodes)

	userResponse := response.UserCreationResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: qrData.TwoFactEnabled,
		Secret:         qrData.Secret,
		QrCode:         qrData.QrCode,
		BackupCodes:    backupPasswords,
	}
	return userResponse, nil
}

func (a authService) Logout() *fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute * 5),
		HTTPOnly: true,
	}
	return &cookie
}

func (a authService) GetUserDetailsFromToken(token jwt.Token) (response.SimpleUserResponse, error) {
	claims := token.Claims.(jwt.MapClaims)

	var user entity.User

	user, _ = a.userRepository.FindUserById(claims["Issuer"].(string))

	userResponse := response.SimpleUserResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: user.TwoFactEnabled,
	}
	return userResponse, nil
}

func (a authService) ChangePassword(id string, data request.ChangePassword) (*entity.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(*data.Password), 12)
	if err != nil {
		return nil, errors.New("something went wrong while encrypting password")
	}
	err = a.userRepository.UpdatePassword(id, password)

	if err != nil {
		return nil, err
	}

	user, _ := a.userRepository.FindUserById(id)
	return &user, nil
}

func (a authService) Disable2FA(user *entity.User) error {
	err := a.userRepository.DisableUser2FA(user)
	if err != nil {
		return err
	}
	return nil
}

func (a authService) getUserByEmail(email string) (*entity.User, error) {
	user, err := a.userRepository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
