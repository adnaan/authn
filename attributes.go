package authn

import (
	"fmt"
	"net/http"

	"github.com/adnaan/authn/models"
)

type M map[string]interface{}

func (m M) String(key string) (string, bool) {
	v, ok := m[key]
	if !ok {
		return "", false
	}
	vstr, ok := v.(string)
	if !ok {
		return "", false
	}
	return vstr, true
}

func (m M) Int(key string) (int, bool) {
	v, ok := m[key]
	if !ok {
		return 0, false
	}
	vstr, ok := v.(int)
	if !ok {
		return 0, false
	}
	return vstr, true
}

func (m M) Bool(key string) (bool, bool) {
	v, ok := m[key]
	if !ok {
		return false, false
	}
	vstr, ok := v.(bool)
	if !ok {
		return false, false
	}
	return vstr, true
}

func (a *Account) Attributes() *Attributes {
	return &Attributes{
		req:          a.req,
		acc:          a.acc,
		sessionStore: a.sessionStore,
		sessionAttributes: &SessionAttributes{
			req:          a.req,
			acc:          a.acc,
			sessionStore: a.sessionStore,
		},
	}
}

type SessionAttributes struct {
	req          *http.Request
	acc          *models.Account
	sessionStore SessionsStore
}

// Get gets an attribute from the temporary session store for the account
func (a *SessionAttributes) Map() (M, error) {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return nil, fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}
	m := make(M)
	for k, v := range session.Values {
		ks, ok := k.(string)
		if !ok {
			continue
		}
		m[ks] = v
	}

	return m, nil
}

// Get gets an attribute from the temporary session store for the account
func (a *SessionAttributes) Get(key string) (interface{}, error) {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return nil, fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	val, ok := session.Values[key]
	if !ok {
		return nil, fmt.Errorf("err: %v, %w", err, ErrAttributeNotFound)
	}

	return val, nil
}

// Set sets an attribute in the temporary session store for the account
func (a *SessionAttributes) Set(w http.ResponseWriter, key string, val interface{}) error {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values[key] = val
	err = session.Save(a.req, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// Del deletes an attribute from the temporary session store for the account
func (a *SessionAttributes) Del(w http.ResponseWriter, key string) error {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}
	delete(session.Values, key)
	err = session.Save(a.req, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	return nil
}

type Attributes struct {
	req               *http.Request
	acc               *models.Account
	sessionStore      SessionsStore
	sessionAttributes *SessionAttributes
}

// Get gets an attribute from the database for the account
func (a *Attributes) Get(key string) (interface{}, error) {
	v, ok := a.acc.Attributes[key]
	if !ok {
		return nil, fmt.Errorf("%w", ErrAttributeNotFound)
	}
	return v, nil
}

// All gets all attributes from the database for the account
func (a *Attributes) Map() M {
	return a.acc.Attributes
}

// Set stores an attribute in the database for the account
func (a *Attributes) Set(key string, val interface{}) error {
	currentAttributes := a.acc.Attributes
	currentAttributes[key] = val

	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// Del deletes an attribute from the database for the account
func (a *Attributes) Del(key string) error {
	currentAttributes := a.acc.Attributes
	delete(currentAttributes, key)

	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// SetMany stores  many attributes in the database for the account
func (a *Attributes) SetMany(attrs map[string]interface{}) error {
	currentAttributes := a.acc.Attributes
	for k, v := range attrs {
		currentAttributes[k] = v
	}
	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// DelMany deletes many attributes from the database for the account
func (a *Attributes) DelMany(keys []string) error {
	currentAttributes := a.acc.Attributes
	for _, k := range keys {
		delete(currentAttributes, k)
	}
	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// SetBytes stores byte data in the database for the account
func (a *Attributes) SetBytes(attr []byte) error {
	err := a.acc.Update().SetAttributeBytes(attr).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// GetBytes gets byte data from the database for the account
func (a *Attributes) GetBytes() []byte {
	return a.acc.AttributeBytes
}

// DelBytes delete byte data from the database for the account
func (a *Attributes) DelBytes() error {
	err := a.acc.Update().ClearAttributeBytes().Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

func (a *Attributes) Session() *SessionAttributes {
	return a.sessionAttributes
}
