package user

// EXIT VALUES
// The useradd command exits with the following values:
// 0   success
// 1   can't update password file
// 2   invalid command syntax
// 3   invalid argument to option
// 4   UID already in use (and no -o)
// 6   specified group doesn't exist
// 9   username already in use
// 10  can't update group file
// 12  can't create home directory
// 14  can't update SELinux user mapping
const (
	USER_ADD_ERROR_INTERNAL_FAILURE                   = -1
	USER_ADD_SUCCESS                                  = 0
	USER_ADD_ERROR_CANNOT_UPDATE_PASSWORD_FILE        = 1
	USER_ADD_ERROR_INVALID_COMMAND_SYNTAX             = 2
	USER_ADD_ERROR_INVALID_ARGUMENT_TO_OPTION         = 3
	USER_ADD_ERROR_UID_ALREADY_IN_USE_WITH_NO_O       = 4
	USER_ADD_ERROR_SPECIFIED_GROUP_DOES_NOT_EXIST     = 6
	USER_ADD_ERROR_USERNAME_ALREADY_IN_USE            = 9
	USER_ADD_ERROR_CANNOT_UPDATE_GROUP_FILE           = 10
	USER_ADD_ERROR_CANNOT_CREATE_HOME_DIRECTORY       = 12
	USER_ADD_ERROR_CANNOT_UPDATE_SELINUX_USRE_MAPPING = 14
)

// EXIT VALUES
// The userdel command exits with the following values:
// 0 success
// 1 can't update password file
// 2 invalid command syntax
// 6 specified user doesn't exist
// 8 user currently logged in
// 10 can't update group file
// 12 can't remove home directory
const (
	USER_DEL_ERROR_INTERNAL_FAILURE        = -1
	USER_DEL_SUCCESS                       = 0
	USER_DEL_CANNOT_UPDATE_PASSWORD_FILE   = 1
	USER_DEL_INVALID_COMMAND_SYNTAX        = 2
	USER_DEL_SPECIFIED_USER_DOES_NOT_EXIST = 6
	USER_DEL_USER_CURRENTLY_LOGGED_IN      = 8
	USER_DEL_CANNOT_UPDATE_GROUP_FILE      = 10
	USER_DEL_CANNOT_REMOVE_HOME_DIRECTORY  = 12
)

type UserReturnCode struct {
	Code     int
	Message  string
	RawError error
}

// type UserConfiguration struct {
//  Username      string
//  BaseDirectory string   // -b, --base-dir BASE_DIR
//  Comment       string   // -c, --comment COMMENT
//  HomeDirectory string   // -d, --home-dir HOME_DIR
//  Defaults      bool     // -D, --defaults
//  ExpireDate    string   // -e, --expiredate EXPIRE_DATE
//  Inactive      bool     // -f, --inactive INACTIVE
//  GID           uint     // -g, --gid GROUP
//  Groups        []string // -G, --groups GROUP1[,GROUP2,...[,GROUPN]]]
//  NoLogInit     bool     // -l, --no-log-init
//  CreateHome    bool     // -m, --create-home
//  NoCreateHome  bool     // -M, --no-create-home
//  NoUserGroup   bool     // -N, --no-user-group
//  NonUnique     bool     // -o, --non-unique
//  Password      string
//  System        bool
//  Root          string
//  Shell         string
//  UID           uint
//  UserGroup     bool

//  // Stuff I don't want to support right now:
//  // SkeletonDirectory string
//  // Keys []string
//  // SELinuxUser string
// }
