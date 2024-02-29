module river_supavisor_issue

go 1.21.0

toolchain go1.21.2

require (
	github.com/golang-cz/devslog v0.0.4
	github.com/jackc/pgx/v5 v5.5.3
	github.com/riverqueue/river v0.0.22
	github.com/riverqueue/river/riverdriver/riverpgxv5 v0.0.22
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/riverqueue/river/riverdriver v0.0.22 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

// replace github.com/henomis/pinecone-go => github.com/JanRuettinger/pinecone-go v0.0.0-20231107040808-e1f1e8f98c90

replace github.com/danielgtaylor/huma/v2 v2.2.1-0.20240119234608-85278c95c7c6 => github.com/JanRuettinger/huma/v2 v2.0.0-20240122190219-1bb68018164f

// replace github.com/danielgtaylor/huma/v2 v2.2.1-0.20240119234608-85278c95c7c6 => /Users/jan/Documents/dataleap/yosemite/third_party_libs/huma
