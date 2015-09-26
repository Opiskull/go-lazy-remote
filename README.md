# go-lazy-remote
A simple go-service for executing commands

## Setup
* change config.json and add commands
* start service and all commands in the config will be loaded

## Commands
* /api/commands
> Lists all commands
* /api/commands/info
> Lists all commands and info to all commands
* /api/command-name
> Executes command-name
* /api/command-name/info
> Shows the info for command-name
