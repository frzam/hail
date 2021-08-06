<div align="center">
    <h1 style= "font-weight:700">hail</h1>
    <p><b>Cross-Platform Scripts Management CLI written in go</b></p>
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/frzam/hail)](https://goreportcard.com/report/github.com/frzam/hail)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=frzam_hail&metric=alert_status)](https://sonarcloud.io/dashboard?id=frzam_hail)
![example workflow](https://github.com/frzam/hail/actions/workflows/release.yml/badge.svg)
 [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GO Version]( https://img.shields.io/github/go-mod/go-version/frzam/hail
)]( https://img.shields.io/github/go-mod/go-version/frzam/hail
)
 [![CII](https://bestpractices.coreinfrastructure.org/badge_static/72)](https://bestpractices.coreinfrastructure.org/en/projects/5082)
> Cross-Platform script management CLI written in go.

* [About](#about)
* [Features](#features)
* [Installation](#installation)
    * [Linux Or Unix](#linux-or-unix)
    * [Windows](#windows)
*  [Usage](#usage)
*  [License](#license)

## About
**hail** is a cross-platfrom script management CLI written in go. We spend lot of time in terminal, 

```
> hail -h
hail is a cross-platform script management tool

Usage:
  hail [flags]
  hail [command]

Available Commands:
  add         add is used to add a new command in collection
  completion  generate the autocompletion script for the specified shell
  config      it is used to list or update configurations
  copy        copy/cp is  used to copy one command/script to a new alias
  delete      delete/rm removes command from hail basis alias
  edit        edit previously added command or script in text editor
  get         get retrieves command basis the alias
  help        Help about any command
  init        init initializes an empty .hailconfig file with title as provided
  list        list/ls prints all the alias and commands
  move        move/mv used to move command with old alias to new alias
  run         it is used to directly run a command from alias
  update      updates already present command
  version     version prints the current version of hail

Flags:
  -h, --help   help for hail

Use "hail [command] --help" for more information about a command.
```

```
> hail ls
+--------------------+----------------------------------------------------+----------------------+
| ALIAS              | COMMAND                                            | DESCRIPTION          |
+--------------------+----------------------------------------------------+----------------------+
| list-git-repo      | find ~ -name ".git" 2> /dev/null | sed 's/\/.git/\ |                      |
|                    | //g' | awk '{print "-------------------------\n\03 |                      |
|                    | 3[1;32mGit Repo:\033[0m " $1; system("git --git-di |                      |
|                    | r="$1".git --work-tree="$1" status")}'             |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| log-with-exception | find . -name '*.log' -mtime -2 -exec grep -Hc Exce |                      |
|                    | ption {} \; | grep -v :0$                          |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| pv                 | apiVersion: v1                                     | Persistence volume i |
|                    | kind: PersistentVolume                             | n ocp                |
|                    | metadata:                                          |                      |
|                    |   name: pv00001                                    |                      |
|                    | spec:                                              |                      |
|                    |   capacity:                                        |                      |
|                    |     storage: 10Gi                                  |                      |
|                    |   accessModes:                                     |                      |
|                    |     - ReadWriteOnce                                |                      |
|                    |   persistentVolumeReclaimPolicy: Retain            |                      |
|                    |   nfs:                                             |                      |
|                    |     path: /mnt/nfs_shares/k8s/root/pv00001         |                      |
|                    |     server: 192.168.10.223                         |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| scan-ports         | for i in {1..65535}; do (echo < /dev/tcp/127.0.0.1 |                      |
|                    | /$i) &>/dev/null && printf "\n[+] Open Port at\n:  |                      |
|                    | \t%d\n" "$i" || printf "."; done                   |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| server-sh          | `#!/bin/bash                                       |                      |
|                    |     echo|read|{(read t;g=$(echo $t|cut -d' ' -f2)  |                      |
|                    |     while read|grep :;do :;done;[[ -e .$g &&! $g = |                      |
|                    |  *..* ]]||exit                                     |                      |
|                    |     printf "HTTP/1.1 200 OK\nContent-Length: $(sta |                      |
|                    | t -c%s .$g)\n\n"                                   |                      |
|                    |     cat .$g)|nc -l -p $1;}>/dev/fd/0;$0 $1         |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| was-bin            | cd /opt/IBM/BPM/v8.6/profiles/managerProfile/bin   |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| create-password    | tr -dc 'a-zA-Z0-9~!@#$%^&*_()+}{?></";.,[]=-' < /d | generate a password  |
|                    | ev/urandom | fold -w 32 | head -n 1                |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
| disk-usage-by-type | find . -type f -empty -prune -o -type f -printf "% |                      |
|                    | s\t" -exec file --brief --mime-type '{}' \; | awk  |                      |
|                    | 'BEGIN {printf("%12s\t%12s\n","bytes","type")} {ty |                      |
|                    | pe=$2; a[type]+=$1} END {for (i in a) printf("%12u |                      |
|                    | \t%12s\n", a[i], i)|"sort -nr"}'                   |                      |
|                    |                                                    |                      |
+--------------------+----------------------------------------------------+----------------------+
```

## Features
* Cross-Platform 
* Auto Completion scripts are available (Bash, fish, zsh and powershell)
* Inbuilt fuzzy searching
* Store commands as well as scripts
* Single and small binary to install. (Thanks to go!)
* Run commands or scripts directly from hail
* Portable toml file *.hailconfig*
* Format of scripts or command is also preserved eg *yaml* 

## Installation
### Linux or Unix
```python
# Download the latest version of hail from releases https://github.com/frzam/hail/releases/ and unzip the file
> wget -c https://github.com/frzam/hail/releases/download/v0.1.11/hail_0.1.11_Linux_x86_64.tar.gz -O - | tar -xz


# Give execute permission to hail and move the binary into bin folder so it is accessible everywhere.
> chmod +x hail | mv hail /usr/local/bin/

# Test if hail is working properly.
> hail version

# Initialize hailconfig, it will create .hailconfig file under $HOME. If you want to 
# create .hailconfig anywhere else then set env HAILCONFIG to that path.
> hail init <title>
```
[OPTIONAL] Set up tab auto completion for bash. Auto completion scripts are also available for fish and zsh.
```python
    
# Generate bash script in a file
> hail completion bash > ~/.hail

# Open ~/.bashrc in editor
> vi ~/.bashrc

# Add below line in .bashrc
> source ~/.hail

# Refresh .bashrc by doing
> source ~/.bashrc
```
### Windows
```python
# Download the latest binary from releases https://github.com/frzam/hail/releases/ 

# Unzip the hail__Windows_x86_64.tar.gz file into *hail.exe*

# Place hail.exe into the PATH

# Test if hail is working properly. Open Command Prompt, Powershell or Git bash and Run.
> hail version

# Initialize hailconfig, it will create .hailconfig file under $USERPROFILE.
# If you want to create .hailconfig anywhere else then set env HAILCONFIG to that path
> hail init -t <title> -i bash

# Generate powershell completion script
> hail completion powershell
```
## Usage
1. Init

    ```> hail init -t my-config -i bash```

2. Add a command with description

    ```> hail add -a create-password -d "create a password" "tr -dc 'a-zA-Z0-9~!@#$%^&*_()+}{?></";.,[]=-' < /dev/urandom | fold -w 32 | head -n 1"```
3. Get a command with alias 'create-password'

    ```> hail get create-password```

4. Edit a command with alias 'create-password'

    ```> hail edit create-password```

5. Run a command with alias 'create-password'

    ```> hail run create-password```

6. Delete an entry with alias 'create-password'

    ```> hail rm create-password```

7. List all aliases with commands and descriptions.

    ```> hail ls```

# License
hail is provided under [Apache 2.0](https://github.com/frzam/hail/blob/master/LICENSE) license.