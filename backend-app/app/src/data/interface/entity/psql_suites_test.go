// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("Accounts", testAccountsUpsert)

	t.Run("Articles", testArticlesUpsert)

	t.Run("Comments", testCommentsUpsert)

	t.Run("Contacts", testContactsUpsert)

	t.Run("Photographs", testPhotographsUpsert)
}
