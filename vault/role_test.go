//go:build integration
package vault

import "testing"

const roleName = "test-role"

// TestRole tests Role
func TestRole(t *testing.T) {
	role, err := dsv.Role(roleName)

	if err != nil {
		t.Errorf("calling roles.Role(\"%s\"): %s", roleName, err)
		return
	}

	if role.Name != roleName {
		t.Errorf("expecting %s but roles.Role was %s", roleName, role.Name)
		return
	}
}

// TestNonexistentRole asks for the "nonexistent" role and fails if it gets it
func TestNonexistentRole(t *testing.T) {
	roleName := "nonexistent"
	_, err := dsv.Role(roleName)

	if err == nil {
		t.Errorf("role '%s' exists but but it should not", roleName)
		return
	}
}
