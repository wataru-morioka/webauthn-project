// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Articles", testArticles)
	t.Run("Comments", testComments)
	t.Run("Contacts", testContacts)
	t.Run("Photographs", testPhotographs)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Articles", testArticlesDelete)
	t.Run("Comments", testCommentsDelete)
	t.Run("Contacts", testContactsDelete)
	t.Run("Photographs", testPhotographsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Articles", testArticlesQueryDeleteAll)
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("Contacts", testContactsQueryDeleteAll)
	t.Run("Photographs", testPhotographsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Articles", testArticlesSliceDeleteAll)
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("Contacts", testContactsSliceDeleteAll)
	t.Run("Photographs", testPhotographsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Articles", testArticlesExists)
	t.Run("Comments", testCommentsExists)
	t.Run("Contacts", testContactsExists)
	t.Run("Photographs", testPhotographsExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Articles", testArticlesFind)
	t.Run("Comments", testCommentsFind)
	t.Run("Contacts", testContactsFind)
	t.Run("Photographs", testPhotographsFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Articles", testArticlesBind)
	t.Run("Comments", testCommentsBind)
	t.Run("Contacts", testContactsBind)
	t.Run("Photographs", testPhotographsBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Articles", testArticlesOne)
	t.Run("Comments", testCommentsOne)
	t.Run("Contacts", testContactsOne)
	t.Run("Photographs", testPhotographsOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Articles", testArticlesAll)
	t.Run("Comments", testCommentsAll)
	t.Run("Contacts", testContactsAll)
	t.Run("Photographs", testPhotographsAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Articles", testArticlesCount)
	t.Run("Comments", testCommentsCount)
	t.Run("Contacts", testContactsCount)
	t.Run("Photographs", testPhotographsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accounts", testAccountsHooks)
	t.Run("Articles", testArticlesHooks)
	t.Run("Comments", testCommentsHooks)
	t.Run("Contacts", testContactsHooks)
	t.Run("Photographs", testPhotographsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Articles", testArticlesInsert)
	t.Run("Articles", testArticlesInsertWhitelist)
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("Contacts", testContactsInsert)
	t.Run("Contacts", testContactsInsertWhitelist)
	t.Run("Photographs", testPhotographsInsert)
	t.Run("Photographs", testPhotographsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Articles", testArticlesReload)
	t.Run("Comments", testCommentsReload)
	t.Run("Contacts", testContactsReload)
	t.Run("Photographs", testPhotographsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Articles", testArticlesReloadAll)
	t.Run("Comments", testCommentsReloadAll)
	t.Run("Contacts", testContactsReloadAll)
	t.Run("Photographs", testPhotographsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Articles", testArticlesSelect)
	t.Run("Comments", testCommentsSelect)
	t.Run("Contacts", testContactsSelect)
	t.Run("Photographs", testPhotographsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Articles", testArticlesUpdate)
	t.Run("Comments", testCommentsUpdate)
	t.Run("Contacts", testContactsUpdate)
	t.Run("Photographs", testPhotographsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Articles", testArticlesSliceUpdateAll)
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("Contacts", testContactsSliceUpdateAll)
	t.Run("Photographs", testPhotographsSliceUpdateAll)
}
