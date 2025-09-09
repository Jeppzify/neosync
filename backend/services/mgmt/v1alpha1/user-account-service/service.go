package v1alpha1_useraccountservice

import (
	auth_client "github.com/Jeppzify/neosync/backend/internal/auth/client"
	"github.com/Jeppzify/neosync/backend/internal/userdata"
	"github.com/Jeppzify/neosync/internal/authmgmt"
	"github.com/Jeppzify/neosync/internal/billing"
	"github.com/Jeppzify/neosync/internal/ee/license"
	"github.com/Jeppzify/neosync/internal/ee/rbac"
	"github.com/Jeppzify/neosync/internal/neosyncdb"
	"github.com/Jeppzify/neosync/internal/temporal/clientmanager"
)

type Service struct {
	cfg                    *Config
	db                     *neosyncdb.NeosyncDb
	temporalConfigProvider clientmanager.ConfigProvider
	authclient             auth_client.Interface
	authadminclient        authmgmt.Interface
	billingclient          billing.Interface
	rbacClient             rbac.Interface
	licenseclient          license.EEInterface
}

type Config struct {
	IsAuthEnabled            bool
	IsNeosyncCloud           bool
	DefaultMaxAllowedRecords *int64
}

func New(
	cfg *Config,
	db *neosyncdb.NeosyncDb,
	temporalConfigProvider clientmanager.ConfigProvider,
	authclient auth_client.Interface,
	authadminclient authmgmt.Interface,
	billingclient billing.Interface,
	rbacClient rbac.Interface,
	licenseclient license.EEInterface,
) *Service {
	return &Service{
		cfg:                    cfg,
		db:                     db,
		temporalConfigProvider: temporalConfigProvider,
		authclient:             authclient,
		authadminclient:        authadminclient,
		billingclient:          billingclient,
		rbacClient:             rbacClient,
		licenseclient:          licenseclient,
	}
}

func (s *Service) UserDataClient() userdata.Interface {
	return userdata.NewClient(s, s.rbacClient, s.licenseclient)
}
