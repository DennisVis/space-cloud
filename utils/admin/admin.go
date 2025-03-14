package admin

import (
	"errors"
	"net/http"
	"sync"

	"github.com/spaceuptech/space-cloud/config"
)

// Manager manages all admin transactions
type Manager struct {
	lock  sync.RWMutex
	admin *config.Admin
}

// New creates a new admin manager instance
func New() *Manager {
	return &Manager{}
}

// SetConfig sets the admin config
func (m *Manager) SetConfig(admin *config.Admin) {
	m.lock.Lock()
	m.admin = admin
	m.lock.Unlock()
}

// Login handles the admin login operation
func (m *Manager) Login(user, pass string) (int, string, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, u := range m.admin.Users {
		if u.User == user && u.Pass == pass {
			token, err := m.createToken(map[string]interface{}{"id": user, "role": user})
			if err != nil {
				return http.StatusInternalServerError, "", err
			}
			return http.StatusOK, token, nil
		}
	}

	return http.StatusUnauthorized, "", errors.New("invalid credentials provided")
}
