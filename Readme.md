# README

## To reproduce issue
1. Sign up with supabase and create a project
2. Replace database URL with project URL
3. Run river migration
4. Start app: `go run main.go`
5. Go to supabase and run `SELECT * FROM pg_stat_activity where query  like '%river%' order by pid desc` to see the number of increasing connections caused by the river queue.
