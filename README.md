# Mainservice

It's a repo of dummy-service using corelib, plugin, helperlib and wraplib to show
how go resolves the dependencies.
Corelib is a dependency of all components: plugin, helperlib, wraplib and mainservice.
Helperlib and plugin depend only on corelib.
Wraplib depends on helperlib and corelib.

# Example

Let's see go mod graph. It shows that smaller version of corelib is 1.0.0 from helperlib,
but go chooses 1.2.0 as most senior version (from plugin):
```
$ go mod graph
github.com/godepsresolve/mainservice github.com/godepsresolve/corelib@v1.2.0
github.com/godepsresolve/mainservice github.com/godepsresolve/helperlib@v1.0.0
github.com/godepsresolve/mainservice github.com/godepsresolve/plugin@v1.0.0
github.com/godepsresolve/mainservice github.com/godepsresolve/wraplib@v1.0.0
github.com/godepsresolve/helperlib@v1.0.0 github.com/godepsresolve/corelib@v1.0.0
github.com/godepsresolve/plugin@v1.0.0 github.com/godepsresolve/corelib@v1.2.0
github.com/godepsresolve/wraplib@v1.0.0 github.com/godepsresolve/corelib@v1.1.0
github.com/godepsresolve/wraplib@v1.0.0 github.com/godepsresolve/helperlib@v1.0.0
```

Of course this infomation is not a surprise, because of https://go.dev/ref/mod#minimal-version-selection rules. I don't like naming of this paragraph, because it leads to confusion when remembered, people DO think, that if there exist libX and libY dependant on libZ versions 1.1.0 and 1.2.0 correspondingly, libZ@v1.1.0 will be chosen. But "minimal version" just means that if libZ with version 1.3.0 exists but libX and libY are not dependent on it, it will not be chosen.

So I want to proof it in runtime and demonstrate it clearly:
```
mainservice$ make run
wrap:github.com/godepsresolve/mainservice@v1.0.0 -> github.com/godepsresolve/wraplib@v1.0.0 -> github.com/godepsresolve/helperlib@v1.0.0 -> github.com/godepsresolve/corelib: 1.2.0, HelloWorld
github.com/godepsresolve/wraplib@v1.0.0 -> github.com/godepsresolve/corelib: 1.2.0, HelloWorld
helper:github.com/godepsresolve/mainservice@v1.0.0 -> github.com/godepsresolve/helperlib@v1.0.0 -> github.com/godepsresolve/corelib: 1.2.0, HelloWorld
plugin:github.com/godepsresolve/mainservice@v1.0.0 -> github.com/godepsresolve/plugin@v1.0.0 -> github.com/godepsresolve/corelib: 1.2.0, HelloWorld
core:github.com/godepsresolve/mainservice@v1.0.0 -> github.com/godepsresolve/corelib: 1.2.0, HelloWorld
```
As we can see corelib 1.2.0 was chosen.