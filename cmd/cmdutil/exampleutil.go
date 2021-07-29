package cmdutil

const UpdateExample = `  # Update Command with 'delete-pods' alias
  hail update delete-pods 'kubectl delete pod $(kubectl get pods | grep Completed | awk '{print $1}')'`
