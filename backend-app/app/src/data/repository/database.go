package repository

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
	"context"
	"log"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/lib/pq"

	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/config"
	entity "github.com/wataru-morioka/webauthn-project/backend-app/app/src/data/interface/entity"
)

var connection *sql.DB

type DbRepository struct {
}

func NewDbRepository() *DbRepository {
	var once sync.Once
	once.Do(func() {
		createConnection()
    })
	return &DbRepository{}
}

func createConnection() {
	log.Print("コネクション生成")
	env := NewEnv()
	connection, _ = sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			env.PostgresEndpoint,
			5432,
			env.PostgresDatabase,
			env.PostrgresUser,
			env.PostgresPassword,
		),
	)
}

func (a DbRepository) CreateAccount(uid *string) error {
	log.Print("データ登録")
	now := time.Now()
	article := &entity.Account{
		IsDeleted: false,
		IsAdmin: true,
		Account: "test",
		Name: "test",
		// State: null.StringFrom("test"),
		LoginCount: 0,
		UID: null.StringFrom(*uid),
		Thumbnail: null.BytesFrom([]byte("test")),
		LatestLoginAt: now,
		CreatedAt: now,
		ModifiedAt: now,
	}

	if err := article.Insert(context.Background(), connection, boil.Infer()); err != nil {
		log.Print("データ登録エラー: %+v", err)
		return err
	}

	return nil
}