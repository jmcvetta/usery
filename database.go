// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// Package usery is a user and group managment library.
package usery

type Database interface {
	User(id string) (*User, error)
	Group(id string) (*Group, error)
	Permission(id string) (*Permission, error)

	UserGroups(*User) ([]*Group, error)
	GroupUsers(*Group) ([]*User, error)

	AllPermissions(*User) ([]*Permission, error)
	UserPermissions(*User) ([]*Permission, error)
	GroupPermissions(*Group) ([]*Permission, error)

	CreatePermission(name string) (*Permission, error)
	DeletePermission(*Permission) error

	AddUserPermission(*User, *Permission) error
	AddGroupPermission(*Group, *Permission) error
	DeleteUserPermission(*User, *Permission) error
	DeleteGroupPermission(*Group, *Permission) error
	UpdateUserPermissions(*User, []*Permission) error
	UpdateGroupPermissions(*Group, []*Permission) error
}
