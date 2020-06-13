package auth

import "time"

type UserInfo struct {
	UserId              string     `json:"userId,omitempty" gorm:"column:userid" bson:"_id,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Username            string     `json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty"`
	Contact             string     `json:"contact,omitempty" gorm:"column:contact" bson:"contact,omitempty" dynamodbav:"contact,omitempty" firestore:"contact,omitempty"`
	DisplayName         string     `json:"displayName,omitempty" gorm:"column:displayname" bson:"displayName,omitempty" dynamodbav:"displayName,omitempty" firestore:"displayName,omitempty"`
	Password            string     `json:"password,omitempty" gorm:"column:password" bson:"password,omitempty" dynamodbav:"password,omitempty" firestore:"password,omitempty"`
	Disable             bool       `json:"disable,omitempty" gorm:"column:disable" bson:"disable,omitempty" dynamodbav:"disable,omitempty" firestore:"disable,omitempty"`
	Deactivated         bool       `json:"deactivated,omitempty" gorm:"column:deactivated" bson:"deactivated,omitempty" dynamodbav:"deactivated,omitempty" firestore:"deactivated,omitempty"`
	Suspended           bool       `json:"suspended,omitempty" gorm:"column:suspended" bson:"suspended,omitempty" dynamodbav:"suspended,omitempty" firestore:"suspended,omitempty"`
	LockedUntilTime     *time.Time `json:"lockedUntilTime,omitempty" gorm:"column:lockeduntiltime" bson:"lockedUntilTime,omitempty" dynamodbav:"lockedUntilTime,omitempty" firestore:"lockedUntilTime,omitempty"`
	SuccessTime         *time.Time `json:"successTime,omitempty" gorm:"column:successtime" bson:"successTime,omitempty" dynamodbav:"successTime,omitempty" firestore:"successTime,omitempty"`
	FailTime            *time.Time `json:"failTime,omitempty" gorm:"column:failtime" bson:"failTime,omitempty" dynamodbav:"failTime,omitempty" firestore:"failTime,omitempty"`
	FailCount           int        `json:"failCount,omitempty" gorm:"column:failcount" bson:"failCount,omitempty" dynamodbav:"failCount,omitempty" firestore:"failCount,omitempty"`
	PasswordChangedTime *time.Time `json:"passwordChangedTime,omitempty" gorm:"column:passwordchangedtime" bson:"passwordChangedTime,omitempty" dynamodbav:"passwordChangedTime,omitempty" firestore:"passwordChangedTime,omitempty"`
	MaxPasswordAge      int        `json:"maxPasswordAge,omitempty" gorm:"column:maxpasswordage" bson:"maxPasswordAge,omitempty" dynamodbav:"maxPasswordAge,omitempty" firestore:"maxPasswordAge,omitempty"`
	UserType            string     `json:"userType,omitempty" gorm:"column:usertype" bson:"userType,omitempty" dynamodbav:"userType,omitempty" firestore:"userType,omitempty"`
	Roles               *[]string  `json:"roles,omitempty" gorm:"column:roles" bson:"roles,omitempty" dynamodbav:"roles,omitempty" firestore:"roles,omitempty"`
	Privileges          *[]string  `json:"privileges,omitempty" gorm:"column:privileges" bson:"privileges,omitempty" dynamodbav:"privileges,omitempty" firestore:"privileges,omitempty"`
	AccessDateFrom      *time.Time `json:"accessDateFrom,omitempty" gorm:"column:accessdatefrom" bson:"accessDateFrom,omitempty" dynamodbav:"accessDateFrom,omitempty" firestore:"accessDateFrom,omitempty"`
	AccessDateTo        *time.Time `json:"accessDateTo,omitempty" gorm:"column:accessDateTo" bson:"accessDateTo,omitempty" dynamodbav:"accessDateTo,omitempty" firestore:"accessDateTo,omitempty"`
	AccessTimeFrom      *time.Time `json:"accessTimeFrom,omitempty" gorm:"column:accesstimefrom" bson:"accessTimeFrom,omitempty" dynamodbav:"accessTimeFrom,omitempty" firestore:"accessTimeFrom,omitempty"`
	AccessTimeTo        *time.Time `json:"accessTimeTo,omitempty" gorm:"column:accesstimeto" bson:"accessTimeTo,omitempty" dynamodbav:"accessTimeTo,omitempty" firestore:"accessTimeTo,omitempty"`
	IsTwoFactor         bool       `json:"isTwoFactor,omitempty" gorm:"column:istwofactor" bson:"isTwoFactor,omitempty" dynamodbav:"isTwoFactor,omitempty" firestore:"isTwoFactor,omitempty"`
}