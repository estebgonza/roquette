package roqclient

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PqClient struct {
	conn *sql.DB
}

func (p *PqClient) Execute(query string) (int32, error) {
	_, err := p.conn.Exec(query)
	return 0, err
}

func (p *PqClient) Connect(host string, port int, user string, pass string, db string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	conn, err := sql.Open("postgres", psqlInfo)
	p.conn = conn
	if err != nil {
		log.Println(err)
	}
	return err
}