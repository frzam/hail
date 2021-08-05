package cmdutil

const (
	// AddExample for add cmd in add.go
	AddExample = `  # Add new Command with alias 'delete-pod' alais
  hail add delete-pods "kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')"
  
  # Add Script from a file.
  hail add -a pv -f pv.yaml
 
  # Add Command with description and alias flag.
  hail add -a delete-pod -d "delete all completed pods"  "kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')"
	
  `
	// CopyExample for copy/cp command in copy.go
	CopyExample = `  # Copy from old alias to new alias.
  hail cp delete-pod dp 

  # Copy from old alias to new alias using flag
  hail copy -o delete-pod -n dp
  `
	// DeleteExample to remove a command and alias.
	DeleteExample = `  # Delete a command with alias 'delete-pod'
  hail delete delete-pod

  # Delete a command basis alias.
  hail rm -a delete-pod
	`
	// EditExample
	EditExample = `  # Edit command with 'delete-alias' in default editor
  hail edit delete-pod	
  
  # Edit using alias
  hail edit -a delete-pod

  # Edit command using fuzzy search
  hail edit
	`
	// GetExample
	GetExample = `  # Get a command basis alias 'pv'
  hail get pv
  
  # Get a command using fuzzy search.
  hail get 
	`
	// InitExample
	InitExample = `  # Init a .hailconfig file at default location $HOME or set env variable HAILCONFIG to path.
  hail init -t my-scripts -i bash

  # Init with a title
  hail init linux-scripts
	`
	// ListExample
	ListExample = `  # List all alias, commands and description 
  hail list

  # List using ls
  hail ls
	`
	// MoveExample
	MoveExample = `  # Move command from old alias to new alias
  hail move delete-pod dl

  # Move using flags
  hail move -o delete-pod -n dl
	`
	// UpdateExample for update command in update.go
	UpdateExample = `  # Update Command with 'delete-pods' alias
  hail update delete-pods "kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')"`
)
