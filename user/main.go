package user

import (
	// "bytes"
	"errors"
	// "os"
	"os/exec"
	// "strings"
	"fmt" // for debugging
)

// func CreateUser(u *UserConfiguration) *UserReturnCode {
func CreateUser(username string, options []string) *UserReturnCode {
	if username == "" {
		return &UserReturnCode{
			Code:     USER_ADD_ERROR_INTERNAL_FAILURE,
			Message:  "No username provided",
			RawError: errors.New("No username provided"),
		}
	}

	_cmd_flags := make([]string, 0)
	_cmd_flags = append(_cmd_flags, username)
	_cmd_flags = append(_cmd_flags, options...)

	fmt.Printf("I will execute: adduser %s", _cmd_flags)

	err := exec.Command("adduser", _cmd_flags...).Run()

	if err != nil {
		_exit_condition := new(UserReturnCode)

		switch err.Error() {
		case "exit status 1":
			_exit_condition.Code = USER_ADD_ERROR_CANNOT_UPDATE_PASSWORD_FILE
			_exit_condition.Message = "can't update password file"
			_exit_condition.RawError = err
			break

		case "exit status 2":
			_exit_condition.Code = USER_ADD_ERROR_INVALID_COMMAND_SYNTAX
			_exit_condition.Message = "invalid command syntax"
			break

		case "exit status 3":
			_exit_condition.Code = USER_ADD_ERROR_INVALID_ARGUMENT_TO_OPTION
			_exit_condition.Message = "can't update password file"
			_exit_condition.RawError = err
			break

		case "exit status 4":
			_exit_condition.Code = USER_ADD_ERROR_UID_ALREADY_IN_USE_WITH_NO_O
			_exit_condition.Message = "UID already in use (and no -o)"
			_exit_condition.RawError = err
			break

		case "exit status 6":
			_exit_condition.Code = USER_ADD_ERROR_SPECIFIED_GROUP_DOES_NOT_EXIST
			_exit_condition.Message = "specified group doesn't exist"
			_exit_condition.RawError = err
			break

		case "exit status 9":
			_exit_condition.Code = USER_ADD_ERROR_USERNAME_ALREADY_IN_USE
			_exit_condition.Message = "username already in use"
			_exit_condition.RawError = err
			break

		case "exit status 10":
			_exit_condition.Code = USER_ADD_ERROR_CANNOT_UPDATE_GROUP_FILE
			_exit_condition.Message = "can't update group file"
			_exit_condition.RawError = err
			break

		case "exit status 12":
			_exit_condition.Code = USER_ADD_ERROR_CANNOT_CREATE_HOME_DIRECTORY
			_exit_condition.Message = "can't create home directory"
			_exit_condition.RawError = err
			break

		case "exit status 14":
			_exit_condition.Code = USER_ADD_ERROR_CANNOT_UPDATE_SELINUX_USRE_MAPPING
			_exit_condition.Message = "can't update SELinux user mapping"
			_exit_condition.RawError = err
			break

		default:
			_exit_condition.Code = USER_ADD_ERROR_INTERNAL_FAILURE
			_exit_condition.Message = err.Error()
			_exit_condition.RawError = err
		}

		return _exit_condition
	}

	return &UserReturnCode{
		Code:     USER_ADD_SUCCESS,
		Message:  "OK",
		RawError: nil,
	}
}
