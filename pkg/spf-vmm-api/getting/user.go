package getting

// UserAndRole vmm spf object
type UserAndRole struct {
	UserName string
	RoleName string
	RoleID   string
}

// OwnerOrGuardian object
type OwnerOrGuardian struct {
	Name                  string
	EncryptionCertificate string
	SigningCertificate    string
}
