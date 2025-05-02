package main

import (
	"net/url"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var password = url.QueryEscape("yourStrong(!)Password")
var dsn = "sqlserver://sa:" + password + "@localhost:1433?database=library&encrypt=true&trustservercertificate=true"

var db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
