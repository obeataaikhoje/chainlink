package pg

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestTxDBDriver(t *testing.T) {
	dbURL, ok := os.LookupEnv("CL_DATABASE_URL")
	require.True(t, ok, "CL_DATABASE_URL must be set")
	db := NewSqlxDB(t, dbURL)
	dropTable := func() error {
		_, err := db.Exec(`DROP TABLE IF EXISTS txdb_test`)
		return err
	}
	// clean up, if previous tests failed
	err := dropTable()
	assert.NoError(t, err)
	_, err = db.Exec(`CREATE TABLE txdb_test (id TEXT NOT NULL)`)
	assert.NoError(t, err)
	t.Cleanup(func() {
		_ = dropTable()
	})
	_, err = db.Exec(`INSERT INTO txdb_test VALUES ($1)`, uuid.New().String())
	assert.NoError(t, err)
	ensureValuesPresent := func(t *testing.T, db *sqlx.DB) {
		var ids []string
		err = db.Select(&ids, `SELECT id from txdb_test`)
		assert.NoError(t, err)
		assert.Len(t, ids, 1)
	}

	ensureValuesPresent(t, db)
	t.Run("Cancel of tx's context does not trigger rollback of driver's tx", func(t *testing.T) {
		ctx, cancel := context.WithCancel(tests.Context(t))
		_, err := db.BeginTx(ctx, nil)
		assert.NoError(t, err)
		cancel()
		// BeginTx spawns separate goroutine that rollbacks the tx and tries to close underlying connection, unless
		// db driver says that connection is still active.
		// This approach is not ideal, but there is no better way to wait for independent goroutine to complete
		time.Sleep(time.Second * 10)
		ensureValuesPresent(t, db)
	})

	t.Run("Make sure calling sql.Register() can be called twice", func(t *testing.T) {
		require.NoError(t, RegisterTxDb(tests.Context(t), "foo"))
		require.NoError(t, RegisterTxDb(tests.Context(t), "bar"))
		drivers := sql.Drivers()
		assert.Contains(t, drivers, "txdb")
	})
}
