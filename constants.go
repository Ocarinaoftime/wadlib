package wadlib

// SignatureType allows specification of the type of signature to be parsed in a ticket.
type SignatureType uint32

const (
	// There's far more types than these internally!
	// See https://git.io/JfJzH or, from acer_cloud_wifi_copy,
	// /sw_x/es_core/esc/core/base/include/esitypes.h#L74
	// However, only RSA 2048 is used in the Wii's title system.
	SignatureRSA2048 SignatureType = 0x00010001
)

// ESLicenseType describes the current title's license type.
type ESLicenseType uint8

const (
	LicensePermanent ESLicenseType = iota
	LicenseDemo
	LicenseTrial
	LicenseRental
	LicenseSubscription
	LicenseService
)

type ContentType uint16

const (
	TitleTypeNormal ContentType = 0x0001
	TitleTypeShared             = 0x8001
)
