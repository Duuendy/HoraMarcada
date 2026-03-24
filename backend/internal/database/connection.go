package database

import (
	"fmt"
)

const Host = "localhost"
const Port = "5432"
const User = "user"
const Password = "123456"
const DbName = "hora_marcada"

var DataSourceName = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	Host,
	Port,
	User,
	Password,
	DbName,
)