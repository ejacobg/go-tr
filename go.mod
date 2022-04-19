module tr

go 1.18

replace control => ./internal/control

replace input => ./internal/input

replace translator => ./internal/translator

require (
	control v0.0.0-00010101000000-000000000000
	input v0.0.0-00010101000000-000000000000
	translator v0.0.0-00010101000000-000000000000
)
