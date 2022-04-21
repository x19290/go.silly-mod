# Naming/Versioning confusion (my failure)

## Must/Cannot

To write a public (globally unique) go module, you **must** understand
the [Semantic Version](https://semver.org/).

And you must know that you **cannot** remove/rename your published module
once someone fetches it.

The module namespace of go looks like
one that a git service (GitHub or something) provides.
But the two are very different in some sense.
The former is provided by A **growing-only** proxy service (GOPROXY.)

## The confusion

This **repo** provides **mod**s as follows:
- github/x19290/go.simplest-mod@v0.0.1
- github/x19290/go.silly-mod@v0.0.2
- github/x19290/go.silly-mod@v0.0.3
- ...

github.com/x19290/go.simplest-mod:
- @v0.0.1 - works
- @v0.0.2 - is broken
- @v0.0.3 - not exists
- ...

github.com/x19290/go.silly-mod:
- @v0.0.1 - is broken
- @v0.0.2 - works
- @v0.0.3 - works (has only minor fixes)
- ...

## The road to the confusion

1st time:
1. I created a **repo** github.com/x19290/go.simplest-mod.
1. I declared "**module** github.com/x19290/go.simplest-mod".
1. I tagged, and pushed it as v0.0.1.

2nd time:
1. I changed the **module** declaration from
   ".../go.simplest-mod" to ".../go.silly-mod".
1. I tagged it as v0.0.2.
1. I pushed it to the **repo** github.com/x19290/go.simplest-mod.

At this point, github/x19290/go.simplest-mod@v0.0.2 broke.

3rd time:
1. I renamed the **repo** to github.com/x19290/go.silly-mod.

At this point, github/x19290/go.silly-mod@v0.0.2 worked.

In v0.0.*, only v0.0.2... are semantically correct.
I mean, only v0.0.2... are strictly API compatible.

I should have given versions like v1.0.0... rather than v0.0.2...
