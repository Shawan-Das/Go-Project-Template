package util

import (
	_ "context"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4/log/logrusadapter"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// main database connection
func CreateDBConnection() *gorm.DB {
	fmt.Println("Connecting....")

	// Fetch database configuration details
	dbHost := ViperReturnStringConfigVariableFromLocalConfigJSON("db_host")
	dbPort := ViperReturnIntegerConfigVariableFromLocalConfigJSON("db_port")
	dbName := ViperReturnStringConfigVariableFromLocalConfigJSON("db_name")
	dbUsername := ViperReturnStringConfigVariableFromLocalConfigJSON("db_username")
	dbPassword := ViperReturnStringConfigVariableFromLocalConfigJSON("db_password")

	// Build data source string
	dataSourceName := "host=" + dbHost + " user=" + dbUsername +
		" password=" + dbPassword + " dbname=" + dbName +
		" port=" + strconv.Itoa(dbPort) + " sslmode=disable"

	// Set up logging for GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	return db
}

// CreateConnectionUsingGormToProcurementSchema creates database connection using gorm to procurement schema
func CreateConnectionUsingGormToCommonSchema() *gorm.DB {
	fmt.Println("Connecting....")
	dbHost := ViperReturnStringConfigVariableFromLocalConfigJSON("db_host")
	dbPort := ViperReturnIntegerConfigVariableFromLocalConfigJSON("db_port")
	dbName := ViperReturnStringConfigVariableFromLocalConfigJSON("db_name")
	dbUsername := ViperReturnStringConfigVariableFromLocalConfigJSON("db_username")
	dbPassword := ViperReturnStringConfigVariableFromLocalConfigJSON("db_password")

	dataSourceName := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + strconv.Itoa(dbPort) + " sslmode=disable"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,
		},
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ViperReturnStringConfigVariableFromLocalConfigJSON("common_schema_name") + ".",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// fmt.Println("failed to connect database")
		panic(err)
	} else {
		return db
	}
}

// CreateConnectionUsingGormToProcurementSchema creates database connection using gorm to procurement schema
func CreateConnectionUsingGormToAccountsSchema() *gorm.DB {
	fmt.Println("Connecting....")
	dbHost := ViperReturnStringConfigVariableFromLocalConfigJSON("db_host")
	dbPort := ViperReturnIntegerConfigVariableFromLocalConfigJSON("db_port")
	dbName := ViperReturnStringConfigVariableFromLocalConfigJSON("db_name")
	dbUsername := ViperReturnStringConfigVariableFromLocalConfigJSON("db_username")
	dbPassword := ViperReturnStringConfigVariableFromLocalConfigJSON("db_password")

	dataSourceName := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + strconv.Itoa(dbPort) + " sslmode=disable"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,
		},
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ViperReturnStringConfigVariableFromLocalConfigJSON("accounts_schema_name") + ".",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// fmt.Println("failed to connect database")
		panic(err)
	} else {
		return db
	}
}

// CreateConnectionUsingGormToProcurementSchema creates database connection using gorm to procurement schema
func CreateConnectionUsingGormToSitlPosSchema() *gorm.DB {
	fmt.Println("Connecting....")
	dbHost := ViperReturnStringConfigVariableFromLocalConfigJSON("db_host")
	dbPort := ViperReturnIntegerConfigVariableFromLocalConfigJSON("db_port")
	dbName := ViperReturnStringConfigVariableFromLocalConfigJSON("db_name")
	dbUsername := ViperReturnStringConfigVariableFromLocalConfigJSON("db_username")
	dbPassword := ViperReturnStringConfigVariableFromLocalConfigJSON("db_password")

	dataSourceName := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + strconv.Itoa(dbPort) + " sslmode=disable"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,
		},
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ViperReturnStringConfigVariableFromLocalConfigJSON("sitlpos_schema_name") + ".",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// fmt.Println("failed to connect database")
		panic(err)
	} else {
		return db
	}
}

// CreateConnectionUsingGormToProcurementSchema creates database connection using gorm to procurement schema
func CreateConnectionUsingGormToiServiceSchema() *gorm.DB {
	fmt.Println("Connecting....")
	dbHost := ViperReturnStringConfigVariableFromLocalConfigJSON("db_host")
	dbPort := ViperReturnIntegerConfigVariableFromLocalConfigJSON("db_port")
	dbName := ViperReturnStringConfigVariableFromLocalConfigJSON("db_name")
	dbUsername := ViperReturnStringConfigVariableFromLocalConfigJSON("db_username")
	dbPassword := ViperReturnStringConfigVariableFromLocalConfigJSON("db_password")

	dataSourceName := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + strconv.Itoa(dbPort) + " sslmode=disable"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,
		},
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ViperReturnStringConfigVariableFromLocalConfigJSON("iservice_schema_name") + ".",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// fmt.Println("failed to connect database")
		panic(err)
	} else {
		return db
	}
}

// ViperReturnStringConfigVariableFromLocalConfigJSON returns values of string variable from local-config.json
func ViperReturnStringConfigVariableFromLocalConfigJSON(key string) string {
	// viper.SetConfigFile("local-config.json")
	var fileDetails string = ConfigDetails
	// fmt.Println("File Name1 :", fileDetails)
	var (
		fileName string
		fileType string
		location string
	)
	if fileDetails != "" {
		fileName, fileType, location = CreateFileDetails(fileDetails)
	}
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType(fileType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(location) // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	// viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println(key)
		fmt.Println(value)
		log.Fatalf("Invalid type assertion")
		return ""
	}
	return value
}

// ViperReturnIntegerConfigVariableFromLocalConfigJSON returns values of int variable from local-config.json
func ViperReturnIntegerConfigVariableFromLocalConfigJSON(key string) int {
	// viper.SetConfigFile("local-config.json")
	var fileDetails string = ConfigDetails
	// fmt.Println("File Name1 :", fileDetails)
	var (
		fileName string
		fileType string
		location string
	)
	if fileDetails != "" {
		fileName, fileType, location = CreateFileDetails(fileDetails)
	}

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType(fileType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(location) // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value := viper.GetInt(key)
	return value
}
