# Cookbook Server

Recipe management database engine.

## Building

- Install [go](https://golang.org/doc/install).
- Build the project `go build .`

## Running

- `./cookbook`

## Terminology

- **Recipe** - an instruction how to create something. For the purposes of this project a recipe
consists of the product (the result), the ingredients (requirements), rates (how much is produced/
consumed) and the station (or process) by which the product is made.

## Features

The following list is a defacto requirements document for this project.

- [ ] Storing recipes in a performant manner.
- [ ] Retrieving recipes by name in a performant manner.
- [ ] Retrieving recipes by ingredients in a performant manner.
- [ ] Querying recipes fitting a certain criteria.
- [ ] Authenticating existing users.
- [ ] Allowing authenticated users to add new recipes.
- [ ] Allowing authenticated users to modify existing recipes.
- [ ] Presenting recipes in a visual manner.
