# Mainservice fork

It's a repo of forked dummy-service using forked corelib (corelib_fork), plugin, helperlib and wraplib to show
how go resolves the replaced forked dependencies.
Corelib is a dependency of all components: plugin, helperlib, wraplib and mainservice.
Helperlib and plugin depend only on corelib.
Wraplib depends on helperlib and corelib.

# Example of replace

Let's see go mod graph. It shows that smaller version of corelib is 1.0.0 from helperlib,
but go chooses 1.2.0 as most senior version (from plugin):
```
$ go mod graph
github.com/godepsresolve/mainservice_replace_fork github.com/godepsresolve/corelib@v1.2.0
github.com/godepsresolve/mainservice_replace_fork github.com/godepsresolve/helperlib@v1.0.0
github.com/godepsresolve/mainservice_replace_fork github.com/godepsresolve/plugin@v1.0.0
github.com/godepsresolve/mainservice_replace_fork github.com/godepsresolve/wraplib@v1.0.0
github.com/godepsresolve/helperlib@v1.0.0 github.com/godepsresolve/corelib@v1.0.0
github.com/godepsresolve/plugin@v1.0.0 github.com/godepsresolve/corelib@v1.2.0
github.com/godepsresolve/wraplib@v1.0.0 github.com/godepsresolve/corelib@v1.1.0
github.com/godepsresolve/wraplib@v1.0.0 github.com/godepsresolve/helperlib@v1.0.0

```

Then let's see what is in go.mod:
```
module github.com/godepsresolve/mainservice_replace_fork

go 1.20

require (
	github.com/godepsresolve/corelib v1.2.0
	github.com/godepsresolve/helperlib v1.0.0
	github.com/godepsresolve/plugin v1.0.0
	github.com/godepsresolve/wraplib v1.0.0
)

replace github.com/godepsresolve/corelib => github.com/godepsresolve/corelib_fork v1.4.0
```

As we can see we replaced vulnerable original library corelib with it's fork corelib_fork. As you can see, `go mod graph` does not reflects the replace.

But would fork of corelib used in all other libraries dependent on it since we don't change in them nothing?

So I want to proof it in runtime and demonstrate it clearly that it will:
```
mainservice_replace_fork$ make run
wrap:github.com/godepsresolve/mainservice_replace_fork@v1.0.0 -> github.com/godepsresolve/wraplib@v1.0.0 -> github.com/godepsresolve/helperlib@v1.0.0 -> github.com/godepsresolve/corelib_fork@v1.4.0: HelloWorld
github.com/godepsresolve/wraplib@v1.0.0 -> github.com/godepsresolve/corelib_fork@v1.4.0: HelloWorld
helper:github.com/godepsresolve/mainservice_replace_fork@v1.0.0 -> github.com/godepsresolve/helperlib@v1.0.0 -> github.com/godepsresolve/corelib_fork@v1.4.0: HelloWorld
plugin:github.com/godepsresolve/mainservice_replace_fork@v1.0.0 -> github.com/godepsresolve/plugin@v1.0.0 -> github.com/godepsresolve/corelib_fork@v1.4.0: HelloWorld
core:github.com/godepsresolve/mainservice_replace_fork@v1.0.0 -> github.com/godepsresolve/corelib_fork@v1.4.0: HelloWorld

```
As we can see corelib_fork 1.4.0 was chosen. It works as described in https://go.dev/ref/mod#mvs-replace. So we don't need to fork all dependency tree to fix something in dependency, it's enough to make those changes in our fork.
