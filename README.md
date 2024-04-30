# 🤖 CLAI - AI Assistant for CLI Commands

CLAI is an AI assistant for CLI commands. It is designed to help users to work with CLI commands more efficiently. 
How often do you struggle to remember how to execute a specific command? Or what the right flags and arguments are?
Just write in natural language what you want to achieve, and CLAI will try to find the right command.

It takes into account your current shell and operating system to provide the most relevant commands. 
However, you can also specify a different shell and OS to target a different environment.

## Examples

```shell
./clai set up an ssh tunnel to be able to connect to MySQL on my remote server
Set up an SSH tunnel to remotely connect to a MySQL server on port 3306, forwarded to localhost:3307

$ ssh -L 3307:localhost:3306 username@remote_server -N
```

```shell
./clai -os=Windows -shell=powershell rename all files with the extension .yml to .yaml in current directory
Rename all files with .yml extension to .yaml in the current directory.

$ Get-ChildItem *.yml | Rename-Item -NewName {$_.BaseName+'.yaml'}
```

## Installation

Just download a pre-compiled binary for your system (macOS, Linux, Windows) from the [releases page](https://github.com/pteich/clai/releases).
You only need to unpack the archive and run the binary in a terminal of your choice.

## OpenAI API

CLAI uses the OpenAI API to generate the commands. Per default, it uses Groq with the "llama3-70b-8192" model.
You need to provide a valid token to use the API (either per config file, see below or environment variable CLAI_TOKEN).

To use OpenAI, provide an empty endpoint or use https://api.openai.com/v1 and a valid OpenAI model.

You can also provide a custom endpoint, to use any OpenAI compatible API.

## Config

CLAI can be configured using a YAML config file, environment variables, or command line arguments.
Values provided by config file are overwritten by environment variables and command line arguments.

The config file needs to be named `.clai.yaml` and can be placed in the user's home directory or in the current working directory.

The following values can be configured:

```yaml
endpoint: https://api.openai.com/v1
model: gpt-4-turbo
token: <YOUR_API_KEY>
shell: bash
os: Linux
```

All values can also be set using environment variables. The following environment variables are available:

- `CLAI_ENDPOINT`
- `CLAI_MODEL`
- `CLAI_TOKEN`
- `CLAI_SHELL`
- `CLAI_OS`

At least a token is required.

## Credits

Original idea found at https://gist.github.com/Sh4yy/3941bf5014bc8c980fad797d85149b65