package libkb

import (
	"fmt"

	"github.com/keybase/client/go/protocol/keybase1"
)

func (mctx MetaContext) LogoutKillSecrets() (err error) {
	return mctx.LogoutWithOptions(LogoutOptions{KeepSecrets: false})
}

func (mctx MetaContext) LogoutKeepSecrets() (err error) {
	return mctx.LogoutWithOptions(LogoutOptions{KeepSecrets: true})
}

type LogoutOptions struct {
	KeepSecrets bool
	Force       bool
}

func (mctx MetaContext) LogoutWithOptions(options LogoutOptions) (err error) {
	username := mctx.ActiveDevice().Username(mctx)
	return mctx.LogoutUsernameWithOptions(username, options)
}

func (mctx MetaContext) LogoutUsernameWithOptions(username NormalizedUsername, options LogoutOptions) (err error) {
	mctx = mctx.WithLogTag("LOGOUT")
	defer mctx.Trace(fmt.Sprintf("MetaContext#LogoutWithOptions(%#v)", options), func() error { return err })()

	g := mctx.G()
	defer g.switchUserMu.Acquire(mctx, "Logout")()

	mctx.Debug("MetaContext#logoutWithSecretKill: after switchUserMu acquisition (username: %s, options: %#v)",
		username, options)

	if options.KeepSecrets {
		mctx.Debug("Keeping secrets after this login, so not checking HasRandomPW")
	} else {
		mctx.Debug("MetaContext#logoutWithSecretKill: checking if CanLogout")
		canLogoutRes := CanLogout(mctx)
		mctx.Debug("MetaContext#logoutWithSecretKill: CanLogout res: %#v", canLogoutRes)
		if !canLogoutRes.CanLogout {
			return fmt.Errorf("Cannot log out: %s", canLogoutRes.Reason)
		}
	}

	var keychainMode KeychainMode
	keychainMode, err = g.ActiveDevice.ClearGetKeychainMode()
	if err != nil {
		return err
	}

	g.LocalSigchainGuard().Clear(mctx.Ctx(), "Logout")

	mctx.Debug("+ MetaContext#logoutWithSecretKill: calling logout hooks")
	g.CallLogoutHooks(mctx)
	mctx.Debug("- MetaContext#logoutWithSecretKill: called logout hooks")

	g.ClearPerUserKeyring()

	// NB: This will acquire and release the cacheMu lock, so we have to make
	// sure nothing holding a cacheMu ever looks for the switchUserMu lock.
	g.FlushCaches()

	if keychainMode == KeychainModeOS {
		mctx.logoutSecretStore(username, options.KeepSecrets)
	} else {
		mctx.Debug("Not clearing secret store in mode %d", keychainMode)
	}

	// reload config to clear anything in memory
	if err := g.ConfigReload(); err != nil {
		mctx.Debug("Logout ConfigReload error: %s", err)
	}

	// send logout notification
	g.NotifyRouter.HandleLogout(mctx.Ctx())

	g.FeatureFlags.Clear()

	g.IdentifyDispatch.OnLogout()

	g.Identify3State.OnLogout()

	err = g.GetUPAKLoader().OnLogout()
	if err != nil {
		return err
	}

	g.Pegboard.OnLogout(mctx)

	return nil
}

func (m MetaContext) logoutSecretStore(username NormalizedUsername, noKillSecrets bool) {

	g := m.G()
	g.secretStoreMu.Lock()
	defer g.secretStoreMu.Unlock()

	if g.secretStore == nil || username.IsNil() {
		return
	}

	if noKillSecrets {
		g.switchedUsers[username] = true
		return
	}

	if err := g.secretStore.ClearSecret(m, username); err != nil {
		m.Debug("clear stored secret error: %s", err)
		return
	}

	// If this user had previously switched into his account and wound up in the
	// g.switchedUsers map (see just above), then now it's fine to delete them,
	// since they are deleted from the secret store successfully.
	delete(g.switchedUsers, username)
}

// LogoutSelfCheck checks with the API server to see if this uid+device pair should
// logout.
func (mctx MetaContext) LogoutSelfCheck() error {
	g := mctx.G()
	uid := g.ActiveDevice.UID()
	if uid.IsNil() {
		mctx.Debug("LogoutSelfCheck: no uid")
		return nil
	}
	deviceID := g.ActiveDevice.DeviceID()
	if deviceID.IsNil() {
		mctx.Debug("LogoutSelfCheck: no device id")
		return nil
	}

	arg := APIArg{
		Endpoint: "selfcheck",
		Args: HTTPArgs{
			"uid":       S{Val: uid.String()},
			"device_id": S{Val: deviceID.String()},
		},
		SessionType: APISessionTypeREQUIRED,
	}
	res, err := g.API.Post(mctx, arg)
	if err != nil {
		return err
	}

	logout, err := res.Body.AtKey("logout").GetBool()
	if err != nil {
		return err
	}

	mctx.Debug("LogoutSelfCheck: should log out? %v", logout)
	if logout {
		mctx.Debug("LogoutSelfCheck: logging out...")
		return mctx.LogoutKillSecrets()
	}

	return nil
}

func CanLogout(mctx MetaContext) (res keybase1.CanLogoutRes) {
	if !mctx.G().ActiveDevice.Valid() {
		mctx.Debug("CanLogout: looks like user is not logged in")
		res.CanLogout = true
		return res
	}

	if mctx.G().ActiveDevice.KeychainMode() == KeychainModeNone {
		mctx.Debug("CanLogout: ok to logout since the key used doesn't user the keychain")
		res.CanLogout = true
		return res
	}

	if err := CheckCurrentUIDDeviceID(mctx); err != nil {
		switch err.(type) {
		case DeviceNotFoundError, UserNotFoundError,
			KeyRevokedError, NoDeviceError, NoUIDError:
			mctx.Debug("CanLogout: allowing logout because of CheckCurrentUIDDeviceID returning: %s", err.Error())
			return keybase1.CanLogoutRes{CanLogout: true}
		default:
			// Unexpected error like network connectivity issue, fall through.
			// Even if we are offline here, we may be able to get cached value
			// `false` from LoadHasRandomPw and be allowed to log out.
			mctx.Debug("CanLogout: CheckCurrentUIDDeviceID returned: %q, falling through", err.Error())
		}
	}

	hasRandomPW, err := LoadHasRandomPw(mctx, keybase1.LoadHasRandomPwArg{
		ForceRepoll: false,
	})

	if err != nil {
		return keybase1.CanLogoutRes{
			CanLogout: false,
			Reason:    fmt.Sprintf("We couldn't ensure that your account has a passphrase: %s", err.Error()),
		}
	}

	if hasRandomPW {
		return keybase1.CanLogoutRes{
			CanLogout:     false,
			SetPassphrase: true,
			Reason:        "You signed up without a password and need to set a password first",
		}
	}

	res.CanLogout = true
	return res
}
