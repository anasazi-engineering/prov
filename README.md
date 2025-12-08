### Configuration File

Create JSOn configuration file with 'base_url' parameter and save as: `$HOME/.provcli`

## Project File Structure




``` code
├── cmd/
│   ├── root.go
│   ├── login.go
│   ├── user.go
│   ├── user_get.go
│   └── user_list.go
├── internal/
│   ├── api/
│   │   ├── client.go
│   │   └── types.go
│   └── config/
│       ├── config.go
│       └── tokens.go
├── go.mod
└── main.go
```

### Adding Commands

Use the following command to add CLI commands to the project using the Cobra CLI tool:

`$ cobra-cli add [command_name]`


## Commands

### Login

Login does a combination of calling login endpoint with username and password and then asking 2FA TOTP token and calling the verify TOTP endpoint. On success, access and refresh tokens are received and stored locally in the `.provcli` configuration file.

`$ ./prov login -u mgmillsa -p P@ssw0rd! -o org1`

`$ ./prov getdevices` 

### Logout

`$ ./prov logout` 

## Authorize BootBox

`$ ./prov authbb ABC123` 

## Apply Recipe

`$ ./prov provision --node <agentID> --url <>` 
`$ ./prov provision -n <agentID> -u <>` 


short: Apply recipe to node

long: 
The provision command is used to select a recipe
that will be used to provision the device specific.
The command requires the node's agentID and the URL
for the recipe to be applied.
