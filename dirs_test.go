package appconf

import (
	"testing"
)

func setup() *AppConf {
	return NewConf("apptest", WithAuthor("daemotron"), WithVersion("1.0"))
}

func TestUserDataDir(t *testing.T) {
	app := setup()
	_, err := app.UserDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestSiteDataDir(t *testing.T) {
	app := setup()
	_, err := app.SiteDataDir(false)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestGlobalDataDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestUserConfigDir(t *testing.T) {
	app := setup()
	_, err := app.UserConfigDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestSiteConfigDir(t *testing.T) {
	app := setup()
	_, err := app.SiteConfigDir(false)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestGlobalConfigDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalConfigDir(true)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestUserStateDir(t *testing.T) {
	app := setup()
	_, err := app.UserStateDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestUserCacheDir(t *testing.T) {
	app := setup()
	_, err := app.UserCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestGlobalCacheDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestUserLogDir(t *testing.T) {
	app := setup()
	_, err := app.UserLogDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestConfigDirs(t *testing.T) {
	app := setup()
	_, err := app.ConfigDirs(true)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}
