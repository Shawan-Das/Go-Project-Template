package service

// import (
// 	"crypto/sha256"
// 	"encoding/json"
// 	"fmt"
// 	// "net/http"
// 	"strings"
// 	"time"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// 	"github.com/tools/system/model"
// 	"github.com/tools/system/util"
// )

// var _asLogger = logrus.New()

// // AuthenticationRESTService provides reference data related rest services
// type AuthenticationRESTService struct {
// 	dbUtil                 *util.PGSqlDBUtil
// 	jwtSigningKey          []byte
// 	bypassAuth             map[string]bool
// 	adminEmail             string
// 	adminPassword          string
// 	adminEmpCode           string
// 	cache                  map[string]UserRoleInfo
// 	acl                    map[string]bool
// 	aclEnabled             map[string]bool
// 	empSpecificActionCache map[string]map[string]bool
// 	// EmpCode-->Action->true/false
// }

// // type authDataInput struct {
// // 	Email          string      `json:"email"`
// // 	Password       string      `json:"pwd"`
// // 	NewPassword    string      `json:"newPwd"`
// // 	Role           interface{} `json:"role,omitempty"`
// // 	ExtraInfo      interface{} `json:"extrainfo,omitempty"`
// // 	PasswordStatus *string     `json:"pwdstat,omitempty"`
// // }

// // AuthorizationClaims JWTTokenClaims
// type AuthorizationClaims struct {
// 	Empcode string `json:"empcode"`
// 	Email   string `json:"email"`
// 	jwt.StandardClaims
// }

// type authServiceConfig struct {
// 	JWTKey        *string  `json:"jwtKey"`
// 	BypassAuth    []string `json:"bypassAuth"`
// 	AdminEmail    string   `json:"adminEmailId"`
// 	AdminPassword string   `json:"adminPassword"`
// 	AdminEmpCode  string   `json:"adminEmpCode"`
// }

// // UserRoleInfo contains user role related inforation
// type UserRoleInfo struct {
// 	RoleMap  map[string]bool
// 	EmpCode  string
// 	EmpEmail string
// 	// Other stuff
// }

// // func buildUserRoleInfo(authInfo *model.AuthenticationInfo) UserRoleInfo {
// // 	roles := make(map[string]bool)
// // 	jb, _ := json.Marshal(authInfo.Role)
// // 	json.Unmarshal(jb, &roles)
// // 	if len(roles) == 0 {
// // 		roles["USER"] = true
// // 	}
// // 	return UserRoleInfo{EmpCode: authInfo.Empcode, EmpEmail: authInfo.EmaiID, RoleMap: roles}
// // }

// // NewAuthenticationRESTService retuens a new initialized version of the service
// func NewAuthenticationRESTService(config []byte, dbUtil *util.PGSqlDBUtil, verbose bool) *AuthenticationRESTService {
// 	service := new(AuthenticationRESTService)
// 	if err := service.Init(config, dbUtil, verbose); err != nil {
// 		_asLogger.Errorf("Unable to intialize service instance %v", err)
// 		return nil
// 	}
// 	return service
// }

// // Init intializes the service instance
// func (s *AuthenticationRESTService) Init(config []byte, dbUtil *util.PGSqlDBUtil, verbose bool) error {
// 	if verbose {
// 		_asLogger.SetLevel(logrus.DebugLevel)
// 	}
// 	if dbUtil == nil {
// 		return fmt.Errorf("null DB Util reference passed ")
// 	}
// 	s.dbUtil = dbUtil
// 	var conf authServiceConfig
// 	err := json.Unmarshal(config, &conf)
// 	if err != nil {
// 		_asLogger.Errorf("Unable to parse config json file ", err)
// 		return err
// 	}
// 	if conf.JWTKey != nil && len(*conf.JWTKey) > 0 {
// 		s.jwtSigningKey = []byte(*conf.JWTKey)
// 	}
// 	s.bypassAuth = make(map[string]bool)
// 	s.bypassAuth["/"] = true
// 	if conf.BypassAuth != nil && len(conf.BypassAuth) > 0 {
// 		for _, url := range conf.BypassAuth {
// 			s.bypassAuth[url] = true
// 		}
// 	}
// 	s.adminEmail = conf.AdminEmail
// 	s.adminPassword = conf.AdminPassword
// 	s.adminEmpCode = conf.AdminEmpCode
// 	s.cache = make(map[string]UserRoleInfo)
// 	s.acl = make(map[string]bool)
// 	s.aclEnabled = make(map[string]bool)
// 	s.empSpecificActionCache = make(map[string]map[string]bool)
// 	s.loadACLInfo()
// 	_asLogger.Infof("Successfully initialized AuthenticationRESTService")
// 	return nil
// }

// // AddRouters add api end points specific to this service
// func (s *AuthenticationRESTService) AddRouters(router *gin.Engine) {
// }

// // CreateAuthInfo creates a authentication entry with default password
// func (s *AuthenticationRESTService) CreateAuthInfo(emailID, empCode string, role map[string]interface{}, isMigration bool) (bool, string, string, []interface{}, string) {
// 	if len(emailID) == 0 || len(empCode) == 0 || role == nil || len(role) == 0 || !model.IsValidRole(role) {
// 		return false, "EmployeeID/Email/Role should not be blank or invalid", "", nil, ""
// 	}
// 	// Check for email id
// 	records, err := s.dbUtil.QueryRecords("select email from hrm.authentication_info where email=$1 ", emailID)
// 	if err != nil {
// 		_asLogger.Errorf("Error in query hrm.authentication_info %v", err)
// 		return false, "Error in checking records", "", nil, ""
// 	}
// 	if len(records) > 0 {
// 		return false, "EmailID aready taken", "", nil, ""
// 	}
// 	var authRecord model.AuthenticationInfo
// 	defaultPassword := "passw0rd" // TOBE Randomly generated
// 	authRecord.Empcode = empCode
// 	authRecord.EmaiID = emailID
// 	authRecord.Password = s.getHashOf(defaultPassword)
// 	authRecord.Role = role
// 	authRecord.ExtraInfo = map[string]string{}
// 	authRecord.PasswordExpiry = s.getExpiryDate(90)
// 	authRecord.PasswordStatus = "DEFAULT"
// 	sql, params := authRecord.GetInsertStatement()

// 	return true, "", sql, params, defaultPassword
// }

// func (s *AuthenticationRESTService) getHashOf(password string) string {
// 	shaBytes := sha256.Sum256([]byte(password))
// 	return fmt.Sprintf("%x", shaBytes)
// }

// func (s *AuthenticationRESTService) getExpiryDate(days int) string {
// 	return time.Now().AddDate(0, 0, days).Format("20060102")
// }

// // func (s *AuthenticationRESTService) checkAuth(c *gin.Context) bool {
// // 	_asLogger.Infof("URL %s", c.Request.URL)

// // 	url := c.Request.URL
// // 	uri := url.RequestURI()
// // 	if s.jwtSigningKey == nil || strings.EqualFold(uri, "/") {
// // 		//jwt is not available
// // 		return true
// // 	}
// // 	if _, isFound := s.bypassAuth[uri]; isFound {
// // 		return true
// // 	}
// // 	authHeader := c.Request.Header.Get("Authorization")
// // 	_asLogger.Infof("Authorization header received %s", authHeader)
// // 	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer") {
// // 		_asLogger.Infof("Invalid authotization header %s", authHeader)

// // 		return false
// // 	}
// // 	tokenStrs := strings.Split(authHeader, " ")
// // 	if len(tokenStrs) != 2 {
// // 		_asLogger.Infof("Unable to parse authorization token %s", authHeader)
// // 		return false
// // 	}
// // 	tokenStr := tokenStrs[1]
// // 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

// // 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// // 			return nil, fmt.Errorf("UnexpectetokenStrd signing method: %v", token.Header["alg"])
// // 		}
// // 		return s.jwtSigningKey, nil
// // 	})
// // 	if err != nil {
// // 		_asLogger.Errorf("Token parse error %v", err)
// // 		return false
// // 	}
// // 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// // 		empCode := fmt.Sprintf("%v", claims["empcode"])
// // 		empEmail := fmt.Sprintf("%v", claims["email"])
// // 		//TODO:Verify empCode and emailID in the login cache
// // 		//TODO:Also get the user role
// // 		if role, isFound := s.cache[empCode]; isFound {

// // 			c.Set("__ROLE_INFO__", role)
// // 			_asLogger.Infof("Logged in user %s %s", empCode, empEmail)
// // 			return true
// // 		}
// // 		_asLogger.Errorf("Role entry not found")
// // 	}
// // 	return false
// // }

// // GetLoggedInUserEmpCode returns logged-in users empcode
// func (s *AuthenticationRESTService) GetLoggedInUserEmpCode(c *gin.Context) string {
// 	if usrInfo, isExisting := c.Get("__ROLE_INFO__"); isExisting {
// 		if usrDetails, isOk := usrInfo.(UserRoleInfo); isOk {
// 			return usrDetails.EmpCode
// 		}
// 	}
// 	return "----"
// }

// // GetLoggedInUserRoleInfo returns logged-in users empcode
// func (s *AuthenticationRESTService) GetLoggedInUserRoleInfo(c *gin.Context) (string, map[string]bool) {
// 	if usrInfo, isExisting := c.Get("__ROLE_INFO__"); isExisting {
// 		if usrDetails, isOk := usrInfo.(UserRoleInfo); isOk {
// 			return usrDetails.EmpCode, usrDetails.RoleMap
// 		}
// 	}
// 	return "", map[string]bool{}
// }

// // HasPriviledge checks if the logged in user has right priviledge for the input action
// func (s *AuthenticationRESTService) HasPriviledge(action string, c *gin.Context) bool {
// 	if _, isFound := s.aclEnabled[action]; !isFound {
// 		return true
// 	}
// 	if usrInfo, isExisting := c.Get("__ROLE_INFO__"); isExisting {
// 		if usrDetails, isOk := usrInfo.(UserRoleInfo); isOk {
// 			return s.hasEntry(action, usrDetails.RoleMap)
// 		}
// 	}
// 	return false
// }

// func (s *AuthenticationRESTService) hasEntry(action string, roles map[string]bool) bool {
// 	for role := range roles {
// 		k := fmt.Sprintf("%s:%s", action, role)
// 		if _, isFound := s.acl[k]; isFound {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (s *AuthenticationRESTService) loadACLInfo() {
// 	recs, err := s.dbUtil.QueryRecords("select * from hrm.acl_info")
// 	if err != nil || len(recs) == 0 {
// 		_asLogger.Warnf("Unable to load acl_info records %v", err)
// 		return
// 	}
// 	aclRecords := model.BuildACLInfo(recs)
// 	for _, aclentry := range aclRecords {
// 		s.aclEnabled[aclentry.Action] = true
// 		roles := strings.Split(aclentry.Role, ",")
// 		if aclentry.EmpCode == "*" {
// 			for _, role := range roles {
// 				k := fmt.Sprintf("%s:%s", aclentry.Action, role)
// 				s.acl[k] = true
// 				_asLogger.Info("Loaded acl entry ", k)
// 			}
// 		}
// 		// else {
// 		// 	//This is employee specific entry
// 		// }
// 	}
// }
