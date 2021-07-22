package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type OperatorSearchHint struct {
	Operator string
	Phone string
}

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64

	attacker := func(stopAt time.Time) {
		for {
			_, err := search(ctx, dbpool, "alex", 5)
			if err != nil { log.Fatal(err) }

			atomic.AddUint64(&queries, 1)

			if time.Now().After(stopAt) {
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}

	wg.Wait()

	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func search(ctx context.Context, dbpool *pgxpool.Pool, prefix string, limit int) ([]OperatorSearchHint, error) {
	const sql = `
select
	operator,
	phone
from operators
where operator like $1
order by operator asc
limit $2;
`

	pattern := prefix + "%"
	rows, err := dbpool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()

	// В слайс hints будут собраны все строки, полученные из базы
	var hints []OperatorSearchHint

	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint OperatorSearchHint

		// Scan записывает значения столбцов в свойства структуры hint
		err = rows.Scan(&hint.Operator, &hint.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		hints = append(hints, hint)
	}

	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}

	return hints, nil
}


func main()  {
	ctx := context.Background()
	cdn := "postgres://app:qwerty@localhost:5433/app"

	cfg, err := pgxpool.ParseConfig(cdn)
	if err != nil { log.Fatal(err) }

	// Pool соединений обязательно ограничивать сверху,
	// так как иначе есть потенциальная опасность превысить лимит соединений с базой.
	cfg.MaxConns = 80
	cfg.MinConns = 1


	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil { log.Fatal(err) }
	defer dbpool.Close()

	duration := time.Duration(10 * time.Second)
	threads := 1000
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)

	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)

}