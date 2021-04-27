package ldap

type LDAPConfig struct {
	Server             string `mapstructure:"server" json:"server,omitempty" gorm:"column:server" bson:"server,omitempty" dynamodbav:"server,omitempty" firestore:"server,omitempty"`
	BaseDN             string `mapstructure:"base_dn" json:"baseDN,omitempty" gorm:"column:basedn" bson:"baseDN,omitempty" dynamodbav:"baseDN,omitempty" firestore:"baseDN,omitempty"`
	Timeout            int64  `mapstructure:"timeout" json:"timeout,omitempty" gorm:"column:timeout" bson:"timeout,omitempty" dynamodbav:"timeout,omitempty" firestore:"timeout,omitempty"`
	Domain             string `mapstructure:"domain" json:"domain,omitempty" gorm:"column:domain" bson:"domain,omitempty" dynamodbav:"domain,omitempty" firestore:"domain,omitempty"`
	Filter             string `mapstructure:"filter" json:"filter,omitempty" gorm:"column:filter" bson:"filter,omitempty" dynamodbav:"filter,omitempty" firestore:"filter,omitempty"`
	TLS                *bool  `mapstructure:"tls" json:"tls,omitempty" gorm:"column:tls" bson:"tls,omitempty" dynamodbav:"tls,omitempty" firestore:"tls,omitempty"`
	StartTLS           *bool  `mapstructure:"start_tls" json:"startTLS,omitempty" gorm:"column:starttls" bson:"startTLS,omitempty" dynamodbav:"startTLS,omitempty" firestore:"startTLS,omitempty"`
	InsecureSkipVerify *bool  `mapstructure:"insecure_skip_verify" json:"insecureSkipVerify,omitempty" gorm:"column:insecureskipverify" bson:"insecureSkipVerify,omitempty" dynamodbav:"insecureSkipVerify,omitempty" firestore:"insecureSkipVerify,omitempty"`
	Id                 string `mapstructure:"id" json:"id,omitempty" gorm:"column:id" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	DisplayName        string `mapstructure:"display_name" json:"displayName,omitempty" gorm:"column:displayname" bson:"displayName,omitempty" dynamodbav:"displayName,omitempty" firestore:"displayName,omitempty"`
	Contact            string `mapstructure:"contact" json:"contact,omitempty" gorm:"column:contact" bson:"contact,omitempty" dynamodbav:"contact,omitempty" firestore:"contact,omitempty"`
	Email              string `mapstructure:"email" json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty"`
	Phone              string `mapstructure:"phone" json:"phone,omitempty" gorm:"column:phone" bson:"phone,omitempty" dynamodbav:"phone,omitempty" firestore:"phone,omitempty"`
	AccountExpires     string `mapstructure:"account_expires" json:"accountExpires,omitempty" gorm:"column:accountexpires" bson:"accountExpires,omitempty" dynamodbav:"accountExpires,omitempty" firestore:"accountExpires,omitempty"`
	PwdLastSet         string `mapstructure:"pwd_last_set" json:"pwdLastSet,omitempty" gorm:"column:pwdlastset" bson:"pwdLastSet,omitempty" dynamodbav:"pwdLastSet,omitempty" firestore:"pwdLastSet,omitempty"`
	BadPwdCount        string `mapstructure:"bad_pwd_count" json:"badPwdCount,omitempty" gorm:"column:badpwdcount" bson:"badPwdCount,omitempty" dynamodbav:"badPwdCount,omitempty" firestore:"badPwdCount,omitempty"`
	BadPasswordTime    string `mapstructure:"bad_pwd_time" json:"badPasswordTime,omitempty" gorm:"column:badpasswordtime" bson:"badPasswordTime,omitempty" dynamodbav:"badPasswordTime,omitempty" firestore:"badPasswordTime,omitempty"`
	LastLogon          string `mapstructure:"last_logon" json:"lastLogon,omitempty" gorm:"column:lastlogon" bson:"lastLogon,omitempty" dynamodbav:"lastLogon,omitempty" firestore:"lastLogon,omitempty"`
	LockoutTime        string `mapstructure:"lockout_time" json:"lockoutTime,omitempty" gorm:"column:lockouttime" bson:"lockoutTime,omitempty" dynamodbav:"lockoutTime,omitempty" firestore:"lockoutTime,omitempty"`
	WhenCreated        string `mapstructure:"when_created" json:"whenCreated,omitempty" gorm:"column:whencreated" bson:"whenCreated,omitempty" dynamodbav:"whenCreated,omitempty" firestore:"whenCreated,omitempty"`
	Users              string `mapstructure:"users" json:"users,omitempty" gorm:"column:users" bson:"users,omitempty" dynamodbav:"users,omitempty" firestore:"users,omitempty"`
}
