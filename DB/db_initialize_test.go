package DB

import (
	"testing"

	"github.com/spf13/viper"
)

func catchExit(f func()) (exited bool) {
	oldExit := osExit
	defer func() { osExit = oldExit }()

	exited = false
	osExit = func(code int) {
		exited = true
		panic("os.Exit called")
	}

	defer func() {
		if r := recover(); r != nil {
		}
	}()

	f()
	return exited
}

func TestLoadYMLConfig(t *testing.T) {
	viper.Reset()

	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}

	username := viper.GetString("prod.username")
	if username == "" {
		t.Error("Expected prod.username in config")
	}

	dbPort := viper.GetInt("prod.db_port")
	if dbPort == 0 {
		t.Error("Expected prod.db_port in config")
	}
}

func TestInitialize_Success(t *testing.T) {
	viper.Reset()

	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	Initialize()

	if DB == nil {
		t.Fatal("Expected DB connection to be initialized")
	}
}

func TestLoadYMLConfig_Failure(t *testing.T) {
	viper.Reset()

	viper.AddConfigPath("./invalid_path")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	exited := catchExit(func() {
		loadYMLConfig()
	})

	if !exited {
		t.Error("Expected os.Exit on config load failure")
	}
}
