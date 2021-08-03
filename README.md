# hail
[![Go Report Card](https://goreportcard.com/badge/github.com/frzam/hail)](https://goreportcard.com/report/github.com/frzam/hail)
> Cross-Platform script management CLI written in go.

* [About](#about)
* [Features](#features)
* [Installation](#installation)
    * [Linux or Unix](#linuxorunix)
    * [Windows](#windows)
*  [Usage](#usage)
    * [Add](#add)
    * [Get](#get)
    * [List](#list)
    * [Delete](#delete)
    * [Edit](#edit)
*  [License](#license)


## Features
* Cross-Platform
* Auto Completion scripts are available
* Inbuilt fuzzy searching
* Store commands as well as scripts
* Single and small binary to install
* Run commands or scripts directly from hail
* Portable toml file *.hailconfig*
* Use it as command line bookmark
* Format of scripts or command is preserved eg *yaml* 

## Installation
### Linux or Unix
* Download the latest version of hail from releases https://github.com/frzam/hail/releases/ 

    ```wget https://github.com/frzam/hail/releases/download/v0.1.11/hail_0.1.11_Linux_x86_64.tar.gz```
* Unzip the tar.gz file

    ```tar -xf hail_0.1.11_Linux_x86_64.tar.gz```    
* Give execute permission to *hail*

    ```chmod +x hail```    
* Move the binary into *bin* folder so the it is accessible everywhere.

    ```mv hail /usr/local/bin/```
* Test if hail is working properly.

    ```hail version```
* Initialize hailconfig, it will create *.hailconfig* file under **$HOME**. If you want to create *.hailconfig* anywhere else then set env **HAILCONFIG** to that path

    ```hail init <title>```
* [OPTIONAL] Set up tab auto completion for bash. 
    * Generate bash script in a file

      ```hail completion bash > ~/.hail```

    * Open ~/.bashrc in editor

        ```vi ~/.bashrc```
    * Add in .bashrc

       ```source ~/.hail```
     * Refresh .bashrc by doing

        ```source ~/.bashrc```  

### Windows
* Download the latest binary from releases https://github.com/frzam/hail/releases/ 
* Unzip the *hail_`<version>`_Windows_x86_64.tar.gz* file into *hail.exe*
* Place *hail.exe* into the *PATH*
* Test if hail is working properly. Open Command Prompt, Powershell or Git bash and run.

    ```hail version```
* Initialize hailconfig, it will create .hailconfig file under **$USERPROFILE**. If you want to create .hailconfig anywhere else then set env **HAILCONFIG** to that path

    ```hail init <title>```
