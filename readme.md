# Cookbook Server

Recipe management database engine.

## Building

- Install [go](https://golang.org/doc/install).
- Create a file named "users.json" and populate it with a map of usernames to passwords you want to use.
- Build the project `go build .`

## Running

- `./cookbook`

## Terminology

- **Recipe** - a general blueprint how to create something. It consists of the *product* (result),
*ingredients* (components that are required to make the recipe), *rates* (abstract relative amounts)
and *process* (that is required for the recipe).

## Features

The following list is a defacto requirements document for this project.

- [x] Storing recipes in a performant manner.
- [x] Retrieving recipes by name in a performant manner.
- [ ] Retrieving recipes by ingredients in a performant manner.
- [ ] Querying recipes fitting a certain criteria.
- [x] Displaying a single recipe on a web-page.
- [x] Authenticating existing users.
- [x] Allowing authenticated users to add new recipes.
- [ ] Allowing authenticated users to modify existing recipes.
- [ ] Presenting recipes in a visually-appealing manner.
- [ ] Integrated about/documentation page.
