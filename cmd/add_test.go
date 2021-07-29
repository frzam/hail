package cmd

import (
	"bytes"
	"hail/internal/hailconfig"
	"os"
	"testing"
)

var addTests = map[string]string{
	"dlogin":     "docker login localhost:5000 -u admin -p",
	"was-bin":    "cd /opt/IBM/BPM/v8.5/profiles/managerprofile/bin",
	"ifconfig":   "ifconfig eth0 192.168.1.12 netmask 255.255.255.0 broadcast 192.168.1.255",
	"dns":        "nslookup -query=mx google.com",
	"wall":       `wall "we will be going down for maintenance for one hour sharply at 03:30 pm"`,
	"rsync":      `rsync -zvr IMG_5267\ copy\=33\ copy\=ok.jpg ~/Desktop/`,
	"mysql-dump": "mysqldump -u root -p --all-databases > /home/server/Desktop/backupfile.sql",
}

func Test_RunAdd(t *testing.T) {
	o := NewAddOptions()
	hc, _ := hailconfig.NewHailconfig(hailconfig.WithMockHailconfigLoader(""))
	defer hc.Close()

	for a, c := range addTests {
		o.Alias = a
		o.Command = c
		o.Run(hc, os.Stdout)

		want := o.Command
		got := hc.Scripts[o.Alias].Command
		assertGotWant(want, got, t)
	}
}
func assertGotWant(want, got string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("want: '%s', got: '%s'", want, got)
	}
}

func Test_getAlias(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAdd(hailconfig.WithMockHailconfigLoader(""), b)
	// Test alias passed in argument
	want := "ping"
	got, _ := getAlias(cmd, []string{want})
	assertGotWant(want, got, t)

	// Test no alias is found
	_, gotErr := getAlias(cmd, []string{})
	wantErr := "no alias is found"
	assertErr(wantErr, gotErr.Error(), t)
}

func Test_getCommand(t *testing.T) {
	// Test With alias and command
	cmd := NewCmdAdd(hailconfig.WithMockHailconfigLoader(""), os.Stdout)
	want := "docker login localhost:5000 -u admin -p"
	got, _ := getCommand(cmd, []string{"dlogin", want})
	assertGotWant(want, got, t)

	// Test with only command and alias as flag
	cmd.Flags().Set("alias", "logs")
	want = "kubectl logs -f --tail=100"
	got, _ = getCommand(cmd, []string{want})
	assertGotWant(want, got, t)
}

func Test_CmdAdd(t *testing.T) {
	// Test add [alias] [cmd]
	// if os.Getenv("TEST") == "1" {
	// 	cmd := NewCmdAdd(hailconfig.WithMockHailconfigLoader(""), os.Stdout)
	// 	cmd.SetArgs([]string{"dlogs", "kubectl logs -f --tail=100"})
	// 	cmd.Execute()
	// }
	// c := exec.Command(os.Args[0], "-test.run=Test_CmdAdd")
	// c.Env = append(c.Env, "TEST=1")
	// stdout, _ := c.StdoutPipe()
	// if err := c.Start(); err != nil {
	// 	t.Fatal(err)
	// }

	// b, _ := ioutil.ReadAll(stdout)
	// got := string(b)
	//want := "Success: command with alias 'dlogs' has been added"

	//fmt.Println("b: ", got)
	// if os.Getenv("ERROR") == "1" {
	// 	cmd := NewCmdAdd(hailconfig.WithMockHailconfigLoader(""), os.Stdout)
	//}
}
