# Build and Configure

## Build

`$ go build .`

## Configuration File

Create JSON configuration file  (see example) with the Provisioner API Server base URL ('base_url') parameter and save as: `$HOME/.provcli`. 


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

`$ ./prov provision --worker <agentID> --url <>` 
`$ ./prov provision -w <agentID> -u <>` 
`$ ./prov provision -w ABC123 -u "http://github.com/my-recipe"` 