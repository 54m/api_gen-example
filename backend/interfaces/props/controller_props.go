// Package props is a scaffold file for props of controllers
package props

import "github.com/54m/api_gen-example/backend/domain/user"

// ControllerProps is passed from Bootstrap() to all controllers
type ControllerProps struct {
	UserService *user.Service
}
