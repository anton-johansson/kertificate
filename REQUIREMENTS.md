# Requirements

* Register existing CA's
  - Also allow creating a default set of CA's
* Generate new CA's
* Generate certificates using a self-signed CA
  - Is it possible to generate certificates using a public CA too? If so how?
  - Should contain fields for profile below and:
    * Common name
    * Password (and confirm)
* Add profiles to generate CA's for, containing:
  - Country code
  - State or province
  - Locality (eg. city)
  - Organization name (eg. company)
  - Organizational unit (eg. section)
  - E-mail address
  - Valid for (days)
  - Keylength
  - Digest
* Login with different providers
  - Static users
  - LDAP
* Connect certificates to clients/consumer
* Autorenew certificates
  - send these to consumers, etc
