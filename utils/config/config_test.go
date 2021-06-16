package config

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configuration, err := LoadConfig("../../environment/")
	require.NoError(t, err)

	// test postgres
	require.NotZero(t, configuration.PostgresPort)
	require.NotZero(t, configuration.PostgresHost)
	require.NotZero(t, configuration.PostgresPassword)
	require.NotZero(t, configuration.PostgresUser)

	// test redis
	require.NotZero(t, configuration.RedisHost)
	require.NotZero(t, configuration.RedisPort)

	// test backend
	require.NotZero(t, configuration.BackendHost)
	require.NotZero(t, configuration.BackendPort)

	fmt.Printf("%#v \n", configuration)
}
