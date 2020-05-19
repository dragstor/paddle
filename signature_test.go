package paddle

import (
	"crypto/rsa"
	"net/url"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// verifierMock is an autogenerated mock type for the verifier type
type verifierMock struct {
	mock.Mock
}

// Verify provides a mock function with given fields: publicKey, signature, form
func (_m *verifierMock) Verify(publicKey *rsa.PublicKey, signature string, form url.Values) error {
	ret := _m.Called(publicKey, signature, form)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rsa.PublicKey, string, url.Values) error); ok {
		r0 = rf(publicKey, signature, form)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

const publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4
GNADCBiQKBgQCWCVF4mm8LI2vg
4yF9T8VkcZSdlaTwKglId0iWRI
8+bdxLF0jtrj3CVJpiTC81mvX0
fxk6b/Si2aC6BQjf6GSN90MOD+
BazuMVHMpeIYCgzQ1u8fQtxh4m
aS9Df5xdpOfh4HRjbSNQhceAXL
QT2meCTRhW4JJrKyzAMxXU3dZU
AQIDAQAB
-----END PUBLIC KEY-----`

func TestValidVerify(t *testing.T) {
	// TODO: enter valid signature
	t.Skip("cannot test it without valid signature")
	pubKey, err := parsePublicKey(publicKey)
	require.NoError(t, err)
	v := new(realVerifier)
	err = v.Verify(pubKey, "c29tZQ==", url.Values{
		"123": {"12345"},
	})
	require.NoError(t, err)
}

func TestInvalidVerify(t *testing.T) {
	pubKey, err := parsePublicKey(publicKey)
	require.NoError(t, err)
	v := new(realVerifier)
	err = v.Verify(pubKey, "c29tZQ==", url.Values{
		"123": {"12345"},
	})
	require.Error(t, err)
	assert.Equal(t, ErrInvalidSignature, errors.Cause(err))
}
