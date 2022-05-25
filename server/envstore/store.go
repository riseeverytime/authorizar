package envstore

import (
	"sync"

	"github.com/authorizerdev/authorizer/server/constants"
)

var (
	// ARG_DB_URL is the cli arg variable for the database url
	ARG_DB_URL *string
	// ARG_DB_TYPE is the cli arg variable for the database type
	ARG_DB_TYPE *string
	// ARG_ENV_FILE is the cli arg variable for the env file
	ARG_ENV_FILE *string
	// ARG_LOG_LEVEL is the cli arg variable for the log level
	ARG_LOG_LEVEL *string
)

// Store data structure
type Store struct {
	StringEnv map[string]string   `json:"string_env"`
	BoolEnv   map[string]bool     `json:"bool_env"`
	SliceEnv  map[string][]string `json:"slice_env"`
}

// EnvStore struct
type EnvStore struct {
	mutex sync.Mutex
	store *Store
}

var defaultStore = &EnvStore{
	store: &Store{
		StringEnv: map[string]string{
			constants.EnvKeyAdminCookieName:  "authorizer-admin",
			constants.EnvKeyJwtRoleClaim:     "role",
			constants.EnvKeyOrganizationName: "Authorizer",
			constants.EnvKeyOrganizationLogo: "https://www.authorizer.dev/images/logo.png",
		},
		BoolEnv: map[string]bool{
			constants.EnvKeyDisableBasicAuthentication: false,
			constants.EnvKeyDisableMagicLinkLogin:      false,
			constants.EnvKeyDisableEmailVerification:   false,
			constants.EnvKeyDisableLoginPage:           false,
			constants.EnvKeyDisableSignUp:              false,
		},
		SliceEnv: map[string][]string{},
	},
}

// EnvStoreObj.GetBoolStoreEnvVariable global variable for EnvStore
var EnvStoreObj = defaultStore

// UpdateEnvStore to update the whole env store object
func (e *EnvStore) UpdateEnvStore(store Store) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	// just override the keys + new keys

	for key, value := range store.StringEnv {
		e.store.StringEnv[key] = value
	}

	for key, value := range store.BoolEnv {
		e.store.BoolEnv[key] = value
	}

	for key, value := range store.SliceEnv {
		e.store.SliceEnv[key] = value
	}
}

// UpdateEnvVariable to update the particular env variable
func (e *EnvStore) UpdateEnvVariable(storeIdentifier, key string, value interface{}) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	switch storeIdentifier {
	case constants.StringStoreIdentifier:
		e.store.StringEnv[key] = value.(string)
	case constants.BoolStoreIdentifier:
		e.store.BoolEnv[key] = value.(bool)
	case constants.SliceStoreIdentifier:
		e.store.SliceEnv[key] = value.([]string)
	}
}

// GetStringStoreEnvVariable to get the env variable from string store object
func (e *EnvStore) GetStringStoreEnvVariable(key string) string {
	// e.mutex.Lock()
	// defer e.mutex.Unlock()
	return e.store.StringEnv[key]
}

// GetBoolStoreEnvVariable to get the env variable from bool store object
func (e *EnvStore) GetBoolStoreEnvVariable(key string) bool {
	// e.mutex.Lock()
	// defer e.mutex.Unlock()
	return e.store.BoolEnv[key]
}

// GetSliceStoreEnvVariable to get the env variable from slice store object
func (e *EnvStore) GetSliceStoreEnvVariable(key string) []string {
	// e.mutex.Lock()
	// defer e.mutex.Unlock()
	return e.store.SliceEnv[key]
}

// GetEnvStoreClone to get clone of current env store object
func (e *EnvStore) GetEnvStoreClone() Store {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	result := *e.store
	return result
}

func (e *EnvStore) ResetStore() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.store = defaultStore.store
}
