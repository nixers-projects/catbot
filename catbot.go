package main


import (
    "missing"
    "fmt"
    "strings"
    "os/exec"
    "regexp"
    "io/ioutil"
    "net/http"
    "time"
    "github.com/thoj/go-ircevent"
    "last"
)

var (
    server          string  = "irc.iotek.org:6667"
    admins                  = []string {}
    auth_users              = []string {}
    memolist                = []string {}
    memomsg                 = []string {}
    channel         string  = ""
    nickname        string  = ""
    username        string  = ""
    nickserv_pass   string  = ""
    cmdPrefix       string  = "."
    s_nick          string  = ""
    s_cmd           string  = ""
)

func shellcom (cmd string, ircobj *irc.Connection) {
    out, err := exec.Command(cmd).Output()
    if (err != nil) { fmt.Println(err) }
    s := strings.Split(string(out), "\n")
    for i := 0; i < len(s); i++ { ircobj.Privmsg(channel, s[i]) }
    time.Sleep(4000 * time.Millisecond)
}


func check_ident (nick string, ircobj *irc.Connection) {
    ircobj.Privmsgf("nickserv", "STATUS %s", nick)
    return
}

func getTitle (url string) string {
    r, err := http.Get(url)
    if err != nil { return "err" }
    p, err := ioutil.ReadAll(r.Body)
    if err != nil { return "err" }
    r.Body.Close()
    re := regexp.MustCompile("<title>.*?</title>")
    buf := re.FindString(string(p))
    fmt.Println(buf)
    if (len(buf) > 16) {
        return buf[7:len(buf)-8]
    }
    return "err"
}

func handleHttp (e *irc.Event, ircobj *irc.Connection) {
    t := getTitle(e.Message)
    if (t != "err") { ircobj.Privmsg(channel, t) }

}

func handleCmd (s_nick string, s_cmd []string, ircobj *irc.Connection) {
    target := channel

    fmt.Println(s_nick)
    fmt.Println(s_cmd)

    cmd	:= strings.Replace(s_cmd[0], cmdPrefix, "", 1) //strings.Replace(strings.Split(e.Message, " ")[0], cmdPrefix, "", 1)
    cmdArgs := s_cmd[1:] //strings.Split(e.Message, " ")[1:]

    fmt.Printf("%s\n", cmd)
//	if (!missing.Present(admins, s_nick)) {
//		fmt.Println(s_nick + " is not in admins.")
//		if (!missing.Present(auth_users, s_nick)) {
//			fmt.Println(s_nick + " is not in auth_users")
//			return
//		}
//	}
	// auth_users command tree
    switch cmd {
    case "memo":
        n := len(cmdArgs)
        if n > 0 {
            if len(memolist) == len(memomsg) {
                if cmdArgs[0] == "send" && n > 2 {
                    memolist = append(memolist, cmdArgs[1])
                    cmdArgs = append(cmdArgs, "[" + s_nick + "]")
                    memomsg = append(memomsg, strings.Join(cmdArgs[2:], " "))
                    fmt.Printf("%s", memomsg)
                }
                if cmdArgs[0] == "list" { ircobj.Privmsgf(target, "memos for: %s", memolist) }
                if cmdArgs[0] == "read" {
                    for i := 0; i < len(memolist); i++ {
                        if (memolist[i] == s_nick) {
                            ircobj.Privmsgf(s_nick, "%s: %s", memolist[i], memomsg[i])
                            memolist = missing.Remove(memolist, memolist[i])
                            memomsg = missing.Remove(memomsg, memomsg[i])
                        }
                    }
                }
            }
        } else { ircobj.Privmsg(target, "available options: list | send | read") }

    case "fortune":
	//	ircobj.Privmsg(channel, shellcom("fortune"))
        shellcom("fortune", ircobj)

    case "test":
        ircobj.Privmsgf("NICKSERV", "STATUS %s", cmdArgs[0])
        break
    case "np":
        if len(cmdArgs) < 1 { break }
        r,x := last.Last(cmdArgs[0])
        if x != nil { break }
        ircobj.Privmsg(channel, r)
        break

    }

    if (!missing.Present(admins, s_nick)) { return }
	// admin command tree
        switch cmd {
        case "user":
            if (len(cmdArgs) > 0) {
                if (cmdArgs[0] == "add") { if(!missing.Present(auth_users, string(cmdArgs[1]))) { auth_users = append(auth_users, string(cmdArgs[1]))} }
                if (cmdArgs[0] == "del") { auth_users = missing.Remove(auth_users, string(cmdArgs[1]))}
                if (cmdArgs[0] == "list") { ircobj.Privmsgf(target, "admins: %s || users: %s", admins, auth_users) }
            } else { ircobj.Privmsg(target, "available options: add | del | list") }
        }
}

func main () {
    ircobj := irc.IRC(nickname, username)
    ircobj.Connect(server)
    ircobj.VerboseCallbackHandler = false

    ircobj.AddCallback("001", func(e *irc.Event) {
        if (ircobj.GetNick()!=nickname) {

        }
        if (nickserv_pass!="") {
            ircobj.Privmsgf("nickserv", "identify %s", nickserv_pass)
            ircobj.Privmsg("hostserv", "on")
            ircobj.Privmsg("chanserv", "invite #wizards")
            time.Sleep(750 * time.Millisecond)
        }
        ircobj.Join(channel)
    })
    ircobj.AddCallback("NOTICE",  func(e *irc.Event) {
        if (e.Nick == "NickServ" && strings.Split(e.Message, " ")[2] == "3" && s_cmd != "") {
			// true auth
            fmt.Println(string(s_nick) + " is authed")
            handleCmd(s_nick, strings.Split(s_cmd, " "), ircobj)
            s_nick = ""
            s_cmd = ""
        }
    })
    ircobj.AddCallback("PRIVMSG", func(e *irc.Event) {
		// insert handlers
    if (strings.HasPrefix(e.Message, "http://"))	{ handleHttp(e, ircobj) }
    if (strings.HasPrefix(e.Message, "https://"))	{ handleHttp(e, ircobj) }
    if (strings.HasPrefix(e.Message, cmdPrefix))	{
        s_nick = e.Nick
        s_cmd = e.Message
        check_ident(s_nick, ircobj)
    }
    time.Sleep(350 * time.Millisecond)
    })

    ircobj.Loop()
}
