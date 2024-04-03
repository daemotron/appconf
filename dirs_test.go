package appconf

import (
	"testing"
)

func setup() *AppConf {
	return NewConf("apptest", WithAuthor("daemotron"), WithVersion("1.0"))
}

func TestAppConf_UserDataDir(t *testing.T) {
	app := setup()
	_, err := app.UserDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_SiteDataDir(t *testing.T) {
	app := setup()
	_, err := app.SiteDataDir(false)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_GlobalDataDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_UserConfigDir(t *testing.T) {
	app := setup()
	_, err := app.UserConfigDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_SiteConfigDir(t *testing.T) {
	app := setup()
	_, err := app.SiteConfigDir(false)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_GlobalConfigDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalConfigDir(true)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_UserStateDir(t *testing.T) {
	app := setup()
	_, err := app.UserStateDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_UserCacheDir(t *testing.T) {
	app := setup()
	_, err := app.UserCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_GlobalCacheDir(t *testing.T) {
	app := setup()
	_, err := app.GlobalCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_UserLogDir(t *testing.T) {
	app := setup()
	_, err := app.UserLogDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}

func TestAppConf_ConfigDirs(t *testing.T) {
	app := setup()
	_, err := app.ConfigDirs(true)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
}
