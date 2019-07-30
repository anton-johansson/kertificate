package pki

type Certificate struct {
	CountryCode        string
	State              string
	Locality           string
	StreetAddress      string
	PostalCode         string
	Organization       string
	OrganizationalUnit string
	CommonName         string
	ValidFor           int
	KeySize            int
}
