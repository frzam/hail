package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"io/ioutil"
	"testing"
)

var (
	pv = `apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv00001
spec:
  capacity:
	storage: 10Gi
  accessModes:
	- ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  nfs:
	path: /mnt/nfs_shares/k8s/root/pv00001
	server: 192.168.10.223`

	findSh = `find -iname *.SH ( find -iname *.Sh /  find -iname *.sH)`

	wasBin = "cd /opt/IBM/BPM/v8.6/profiles/managerProfile/bin"

	diskUsageByType = `find . -type f -empty -prune -o -type f -printf "%s\t" -exec file --brief --mime-type '{}' \; | awk 'BEGIN {printf("%12s\t%12s\n","bytes","type")} {type=$2; a[type]+=$1} END {for (i in a) printf("%12u\t%12s\n", a[i], i)|"sort -nr"}'`

	listGitRepo = `find ~ -name ".git" 2> /dev/null | sed 's/\/.git/\//g' | awk '{print "-------------------------\n\033[1;32mGit Repo:\033[0m " $1; system("git --git-dir="$1".git --work-tree="$1" status")}'`

	ipAddressOnPort80 = "netstat -tn 2>/dev/null | grep :80 | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -nr | head"

	createPasswd = `tr -dc 'a-zA-Z0-9~!@#$%^&*_()+}{?></";.,[]=-' < /dev/urandom | fold -w 32 | head -n 1`

	logsWithException = `find . -name '*.log' -mtime -2 -exec grep -Hc Exception {} \; | grep -v :0$`

	scanOpenPorts = `for i in {1..65535}; do (echo < /dev/tcp/127.0.0.1/$i) &>/dev/null && printf "\n[+] Open Port at\n: \t%d\n" "$i" || printf "."; done`

	debugPod = `apiVersion: v1
	kind: Pod
	metadata:
	  name: dnsutils
	  namespace: default
	spec:
	  containers:
	  - name: dnsutils
		image: gcr.io/kubernetes-e2e-test-images/dnsutils:1.3
		command:
		  - sleep
		  - "3600"
		imagePullPolicy: IfNotPresent
	  restartPolicy: Always`
	serverSh = `#!/bin/bash
	echo|read|{(read t;g=$(echo $t|cut -d' ' -f2)
	while read|grep :;do :;done;[[ -e .$g &&! $g = *..* ]]||exit
	printf "HTTP/1.1 200 OK\nContent-Length: $(stat -c%s .$g)\n\n"
	cat .$g)|nc -l -p $1;}>/dev/fd/0;$0 $1`
)

// testScripts map is contains alias and command pair.
// It is used as already present scripts in hailconfig.
var testScripts = map[string]string{
	"oc-login":              "oc login -u farzam -p",
	"kube-logs":             "kubectl logs -f --tail=00 ",
	"pv":                    pv,
	"find-sh":               findSh,
	"was-bin":               wasBin,
	"disk-usage-by-type":    diskUsageByType,
	"list-git-repo":         listGitRepo,
	"ip-address-on-port-80": ipAddressOnPort80,
	"create-password":       createPasswd,
	"logs-with-exception":   logsWithException,
	"scan-open-ports":       scanOpenPorts,
	"debug-pod":             debugPod,
	"server-sh":             serverSh,
}

func Test_Run(t *testing.T) {
	o := NewGetOptions()
	hc := NewHailConfigDummy()
	b := bytes.NewBufferString("")

	// Test that alias is getting returned properly.
	for alias, wantCmd := range testScripts {
		o.Alias = alias
		o.Run(hc, b)

		out, _ := ioutil.ReadAll(b)

		got := string(out)
		if got != wantCmd+"\n" {
			t.Errorf("want: '%s' while got: '%s'", wantCmd, got)
		}
	}

	// Test validations
	// No alias is found i.e. Empty alias
	o.Alias = ""
	gotErr := o.Run(hc, b).Error()
	wantErr := "error in validation: no alias is found"
	if gotErr != wantErr {
		t.Errorf("want error: '%q', got error: '%q'", wantErr, gotErr)
	}

	// No alias is present
	o.Alias = "my-alias"
	gotErr = o.Run(hc, b).Error()
	wantErr = "alias is not present: no command is found with 'my-alias' alias"
	if gotErr != wantErr {
		t.Errorf("want error: '%q', got error: '%q'", wantErr, gotErr)
	}
}

// NewHailConfigDummy creates a mock hailconfig pointer and then sets testScipts values.
func NewHailConfigDummy() *hailconfig.Hailconfig {
	hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.WithMockHailconfigLoader(""))
	hc.Parse()
	// Add data into our testScripts
	for k, v := range testScripts {
		hc.Add(k, v, "")
	}
	return hc
}
