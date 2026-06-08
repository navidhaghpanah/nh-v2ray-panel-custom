// Package service provides business logic services for the web panel.
package service

import (
	"github.com/navidhaghpanah/nh-v2ray-panel/v3/database/model"
)

// SettingService handles panel settings operations.
type SettingService struct{}

// Get methods for various settings
func (s *SettingService) GetSubDomain() (string, error) { return "", nil }
func (s *SettingService) GetSubPath() (string, error) { return "/", nil }
func (s *SettingService) GetSubJsonPath() (string, error) { return "", nil }
func (s *SettingService) GetSubClashPath() (string, error) { return "", nil }
func (s *SettingService) GetSubJsonEnable() (bool, error) { return false, nil }
func (s *SettingService) GetSubClashEnable() (bool, error) { return false, nil }
func (s *SettingService) GetSubEncrypt() (bool, error) { return false, nil }
func (s *SettingService) GetSubShowInfo() (bool, error) { return false, nil }
func (s *SettingService) GetRemarkModel() (string, error) { return "-io", nil }
func (s *SettingService) GetSubUpdates() (string, error) { return "10", nil }
func (s *SettingService) GetSubJsonMux() (string, error) { return "", nil }
func (s *SettingService) GetSubJsonRules() (string, error) { return "", nil }
func (s *SettingService) GetSubJsonFinalMask() (string, error) { return "", nil }
func (s *SettingService) GetSubClashEnableRouting() (bool, error) { return false, nil }
func (s *SettingService) GetSubClashRules() (string, error) { return "", nil }
func (s *SettingService) GetSubTitle() (string, error) { return "", nil }
func (s *SettingService) GetSubSupportUrl() (string, error) { return "", nil }
func (s *SettingService) GetSubProfileUrl() (string, error) { return "", nil }
func (s *SettingService) GetSubAnnounce() (string, error) { return "", nil }
func (s *SettingService) GetSubEnableRouting() (bool, error) { return false, nil }
func (s *SettingService) GetSubRoutingRules() (string, error) { return "", nil }
func (s *SettingService) GetSubEnable() (bool, error) { return false, nil }
func (s *SettingService) GetSubCertFile() (string, error) { return "", nil }
func (s *SettingService) GetSubKeyFile() (string, error) { return "", nil }
func (s *SettingService) GetSubListen() (string, error) { return "127.0.0.1", nil }
func (s *SettingService) GetSubPort() (int, error) { return 8080, nil }
func (s *SettingService) GetPort() (int, error) { return 8080, nil }
func (s *SettingService) GetBasePath() (string, error) { return "/", nil }
func (s *SettingService) GetCertFile() (string, error) { return "", nil }
func (s *SettingService) GetKeyFile() (string, error) { return "", nil }
func (s *SettingService) GetTgbotEnabled() (bool, error) { return false, nil }
func (s *SettingService) SetTgbotEnabled(enabled bool) error { return nil }
func (s *SettingService) SetTgBotToken(token string) error { return nil }
func (s *SettingService) SetTgbotRuntime(runtime string) error { return nil }
func (s *SettingService) SetTgBotChatId(chatId string) error { return nil }
func (s *SettingService) SetPort(port int) error { return nil }
func (s *SettingService) SetBasePath(path string) error { return nil }
func (s *SettingService) SetCertFile(cert string) error { return nil }
func (s *SettingService) SetKeyFile(key string) error { return nil }
func (s *SettingService) SetSubCertFile(cert string) error { return nil }
func (s *SettingService) SetSubKeyFile(key string) error { return nil }
func (s *SettingService) SetListen(ip string) error { return nil }
func (s *SettingService) SetTwoFactorEnable(enabled bool) error { return nil }
func (s *SettingService) SetTwoFactorToken(token string) error { return nil }
func (s *SettingService) ResetSettings() error { return nil }

// UserService handles user operations.
type UserService struct{}

func (u *UserService) GetFirstUser() (*model.User, error) { return &model.User{}, nil }
func (u *UserService) UpdateFirstUser(username, password string) error { return nil }

// ApiTokenService handles API token operations.
type ApiTokenService struct{}

func (a *ApiTokenService) List() ([]*model.ApiToken, error) { return nil, nil }
func (a *ApiTokenService) Create(name string) (*model.ApiToken, error) { return &model.ApiToken{}, nil }

// InboundService handles inbound operations.
type InboundService struct{}

func (i *InboundService) MigrateDB() error { return nil }

// RegisterSubLinkProvider registers a sub link provider.
func RegisterSubLinkProvider(provider interface{}) {}

// StopBot stops the telegram bot.
func StopBot() {}
