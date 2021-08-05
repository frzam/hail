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
hail is a cross-platfrom script management CLI written in go. 

## Features
* Cross-Platform 
* Auto Completion scripts are available (Bash, fish, zsh and powershell)
* Inbuilt fuzzy searching
* Store commands as well as scripts
* Single and small binary to install. (Thanks to go!)
* Run commands or scripts directly from hail
* Portable toml file *.hailconfig*
* Use it as command line bookmark
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
> hail init <title>

# Generate powershell completion script
> hail completion powershell
```


# License
hail is provided under [Apache 2.0](!https://github.com/frzam/hail/blob/master/LICENSE) license.