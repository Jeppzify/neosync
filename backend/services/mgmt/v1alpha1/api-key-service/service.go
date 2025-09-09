package v1alpha1_apikeyservice

import (
	"github.com/Jeppzify/neosync/backend/internal/userdata"
	"github.com/Jeppzify/neosync/internal/neosyncdb"
)

type Service struct {
	cfg            *Config
	db             *neosyncdb.NeosyncDb
	userdataclient userdata.Interface
}

type Config struct {
	IsAuthEnabled bool
}

func New(
	cfg *Config,
	db *neosyncdb.NeosyncDb,
	userdataclient userdata.Interface,
) *Service {
	return &Service{cfg: cfg, db: db, userdataclient: userdataclient}
}
