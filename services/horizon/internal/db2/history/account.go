package history

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/BonexIO/go/services/horizon/internal/db2"
	"github.com/BonexIO/go/xdr"
	"database/sql"
)

// Accounts provides a helper to filter rows from the `history_accounts` table
// with pre-defined filters.  See `AccountsQ` methods for the available filters.

// Account is a row of data from the `history_accounts` table
type Account struct {
	ID      	int64	`db:"id"`
	Address 	string 	`db:"address"`
	AccountType uint32  `db:"accounttype"`
}

func NewAccount(id int64, address string, accountType uint32) *Account {
	return &Account{
		ID: 		 id,
		Address:     address,
		AccountType: accountType,
	}
}

func (q *Q) Accounts() *AccountsQ {
	return &AccountsQ{
		parent: q,
		sql:    selectAccount,
	}
}

// AccountByAddress loads a row from `history_accounts`, by address
func (q *Q) AccountByAddress(dest interface{}, addy string) error {
	sql := selectAccount.Limit(1).Where("ha.address = ?", addy)
	return q.Get(dest, sql)
}

// AccountByID loads a row from `history_accounts`, by id
func (q *Q) AccountByID(dest interface{}, id int64) error {
	sql := selectAccount.Limit(1).Where("ha.id = ?", id)
	return q.Get(dest, sql)
}

// Page specifies the paging constraints for the query being built by `q`.
func (q *AccountsQ) Page(page db2.PageQuery) *AccountsQ {
	if q.Err != nil {
		return q
	}

	q.sql, q.Err = page.ApplyTo(q.sql, "ha.id")
	return q
}

// Select loads the results of the query specified by `q` into `dest`.
func (q *AccountsQ) Select(dest interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	q.Err = q.parent.Select(dest, q.sql)
	return q.Err
}

// AccountsByAddresses loads a rows from `history_accounts`, by addresses
func (q *Q) AccountsByAddresses(dest interface{}, addresses []string) error {
	sql := selectAccount.Where(map[string]interface{}{
		"ha.address": addresses, // ha.address IN (...)
	})
	return q.Select(dest, sql)
}

// CreateAccountWithType inserts account id and account type to history_accounts table
func (q *Q) CreateAccountWithTypes(address string, accountType xdr.AccountType) (sql.Result, error) {
	sql := sq.Insert("history_accounts").Columns("address", "accounttype")
	sql = sql.Values(address, accountType)
	return q.Exec(sql)
}

// CreateAccounts creates rows for addresses in history_accounts table and
// put
func (q *Q) CreateAccounts(dest interface{}, addresses map[string]uint32) error {
	sql := sq.Insert("history_accounts").Columns("address", "accounttype")

	added := 0

	for address, accountType := range addresses {

		var existing Account

		err := q.AccountByAddress(&existing, address)

		//account already exists, return id
		if err == nil {
			continue
		}

		added += 1

		sql = sql.Values(address, accountType)
	}

	if added > 0 {
		sql = sql.Suffix("RETURNING *")

		return q.Select(dest, sql)
	}

	return nil
}
// Return id for account. If account doesn't exist, it will be created and the new id returned.
func (q *Q) GetCreateAccountID(
	aid xdr.AccountId,
) (result int64, err error) {

	var existing Account

	err = q.AccountByAddress(&existing, aid.Address())

	//account already exists, return id
	if err == nil {
		result = existing.ID
		return
	}

	// unexpected error
	if !q.NoRows(err) {
		return
	}

	//insert account and return id
	err = q.GetRaw(
		&result,
		`INSERT INTO history_accounts (address) VALUES (?) RETURNING id`,
		aid.Address(),
	)

	return
}

var selectAccount = sq.Select("ha.*").From("history_accounts ha")
