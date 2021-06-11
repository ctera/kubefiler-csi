package driver

// constants of keys in secrets
const (
	FilerUsernameKey = "username"
	FilerPasswordKey = "password"
)

// constants of keys in volume parameters
const (
	// Hostname or IP address of the Ctera Filer
	FilerAddressKey = "fileraddress"
	// Path on the Ctera Filer to share
	PathKey = "path"

	// KmsKeyId represents key for KMS encryption key
	KmsKeyIDKey = "kmskeyid"
)

// constants for default command line flag values
const (
	DefaultCSIEndpoint = "unix://tmp/csi.sock"
)
