module main

go 1.20

require github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker v0.0.0

require github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/db v0.0.0

replace github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker v0.0.0 => ./poker

replace github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/db v0.0.0 => ./db
