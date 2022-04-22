# Just running (not importing)

## No one can import a main package

Like tests, you can include and publish main package(s) in your module.

In your module dir, you can easily **run** sucn ***main**s.
See [21main.sh](../2demo/21main.sh).

But you should know that no one can **import** such main.
You can only **run** it.

## Running from outside of your module dir

A main can be run from outside of your module dir.

```sh
git clone --recursive https://github.com/x19290/go.silly-mod-demo
cd go.silly-mod-demo/2more

rm -f go.*
go mod init user
go mod edit -replace github.com/x19290/go.silly-mod=../go.silly-mod # option
go get github.com/x19290/go.silly-mod

go run github.com/x19290/go.silly-mod/2demo
```

## Testing from outside of your module dir

Module tests can also be run from outside of your module dir.

```sh
git clone --recursive https://github.com/x19290/go.silly-mod-demo
cd go.silly-mod-demo/2more

rm -f go.*
go mod init user
go mod edit -replace github.com/x19290/go.silly-mod=../go.silly-mod # option
go get -t github.com/x19290/go.silly-mod  # -t required

go test github.com/x19290/go.silly-mod github.com/x19290/go.silly-mod/1testdata
```

