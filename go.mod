module github.com/godepsresolve/mainservice_replace_fork

go 1.20

require (
	github.com/godepsresolve/corelib v1.2.0
	github.com/godepsresolve/helperlib v1.0.0
	github.com/godepsresolve/plugin v1.0.0
	github.com/godepsresolve/wraplib v1.0.0
)

replace github.com/godepsresolve/corelib => github.com/godepsresolve/corelib_fork v1.4.0
