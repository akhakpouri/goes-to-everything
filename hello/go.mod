module example.com/hello

go 1.22.6

replace example/greetings => ../example

require (
	example/greetings v0.0.0-00010101000000-000000000000
	helpers v0.0.0-00010101000000-000000000000
)

replace helpers => ../helpers
