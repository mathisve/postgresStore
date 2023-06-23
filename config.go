package postgresStore

type ConnectionConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	SslMode     string
	StorageMode string
	Unlogged    bool
	ConnStr     string
}

// StorageMode
// see createSchema function
const (
	StorageModeExtended = "EXTENDED"
	StorageModeExternal = "EXTERNAL"
)

var DefaultConnectionConfig = ConnectionConfig{
	Host:        "localhost",
	Port:        5432,
	Username:    "postgres",
	Password:    "password",
	DBName:      "postgres",
	SslMode:     "disable",
	StorageMode: StorageModeExtended,
	Unlogged:    false,
}

// SetHost sets the host parameter in the connectionConfig
func (c ConnectionConfig) SetHost(host string) ConnectionConfig {
	c.Host = host
	return c
}

// SetPort sets the port parameter in the connectionConfig
func (c ConnectionConfig) SetPort(port int) ConnectionConfig {
	c.Port = port
	return c
}

// SetUsername sets the username parameter in the connectionConfig
func (c ConnectionConfig) SetUsername(username string) ConnectionConfig {
	c.Username = username
	return c
}

// SetPassword sets the password parameter in the connectionConfig
func (c ConnectionConfig) SetPassword(password string) ConnectionConfig {
	c.Password = password
	return c
}

// SetDBName sets the DBName parameter in the connectionConfig
func (c ConnectionConfig) SetDBName(dbname string) ConnectionConfig {
	c.DBName = dbname
	return c
}

// SetSsLMode sets the SsLMode parameter in the connectionConfig
func (c ConnectionConfig) SetSslMode(sslMode string) ConnectionConfig {
	c.SslMode = sslMode
	return c
}

// SetStorageMode sets the StorageMode parameter in the connectionConfig
func (c ConnectionConfig) SetStorageMode(storageMode string) ConnectionConfig {
	c.StorageMode = storageMode
	return c
}

// SetUnlogged sets the unlogged parameter in the connectionConfig
func (c ConnectionConfig) SetUnlogged(unlogged bool) ConnectionConfig {
	c.Unlogged = unlogged
	return c
}
