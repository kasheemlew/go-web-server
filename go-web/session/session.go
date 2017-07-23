package session

import (
	"sync"
	"fmt"
	"io"
	"encoding/base64"
)

type Manager struct {
	cookieName string // private cookieName
	lock sync.Mutex // protects session
	provider Provider
	maxlifetime int64
}

type Provider interface {
	// initialization of a session, and returns a new session if succeeds.
	SessionInit(sid string) (Session, error)
	// returns a session represented by the corresponding sid. Creates a new session and returns it if it does not already exist
	SessionRead(sid string) (Session, error)
	// given an sid, deletes the corresponding session.
	SeesionDestroy(sid string) (Seesion, error)
	// deletes expired session variables according to maxLifeTime.
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// unique session id
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

var globalSessions *session.Manager
var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider" + name)
	}
}

func init() {
	// initialize session manager
	globalSessions = NewManager("memory", "gosessionid", 3600)
}