/*
NanoCMS is a micro-configuration management system, which is embedded
into the clbd. It runs nanostates that calls pre-installed Ansible modules on
the client machines.
*/

package backend

import (
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/isbm/uyuni-ncd/nanostate"
	"github.com/isbm/uyuni-ncd/nanostate/compiler"
	"github.com/isbm/uyuni-ncd/runners"
	"golang.org/x/crypto/ssh"
	"os"
	"os/user"
	"path"
)

type NanoCms struct {
	BaseBackend
	stateindex  *nanostate.NanoStateIndex
	bootstrapId string
}

func NewNanoCmsBackend() *NanoCms {
	cms := new(NanoCms)
	cms.stateindex = nanostate.NewNanoStateIndex()
	return cms
}

// SetBootstrapStateId sets bootstrap id
func (n *NanoCms) SetBootstrapStateId(id string) {
	n.bootstrapId = id
}

// GetStateIndex returns an instance of a state index.
// XXX: This is subject to change, because nanostate API
// are unstable, so is the indexer.
func (n *NanoCms) GetStateIndex() *nanostate.NanoStateIndex {
	return n.stateindex
}

// RunStateSSH is to run nanostate over SSH
func (n *NanoCms) RunStateSSH(state *nanostate.Nanostate, fqdns ...string) *runners.RunnerResponse {
	logger.Debugf("Running state '%s' on %d machines", state.Id, len(fqdns))
	shr := runners.NewSSHRunner().SetRemoteUsername("root").SetSSHHostVerification(false)
	for _, fqdn := range fqdns {
		shr.AddHost(fqdn)
	}
	logger.Debugf("Run success: %t", shr.Run(state))

	return shr.Response()
}

// LoadNstFile loads a nanostate from a file path
func (n *NanoCms) LoadNstFile(statepath string) *nanostate.Nanostate {
	// Compile nanostate
	cpr := nstcompiler.NewNstCompiler()
	err := cpr.LoadFile(statepath)
	if err != nil {
		panic(err)
	}

	// Load compiled tree
	nst := nanostate.NewNanostate()
	err = nst.Load(cpr.Tree())
	if err != nil {
		panic(err)
	}

	return nst
}

// SshCopyId is copying SSH keys to the target machine, using username and password
func (n *NanoCms) SshCopyId(fqdn string, username string, password string) bool {
	tempkey := path.Join("/root", ".ssh", fqdn+".pub")

	conf, _ := auth.PasswordKey(username, password, ssh.InsecureIgnoreHostKey())
	cnt := scp.NewClient(fqdn+":22", &conf)
	err := cnt.Connect()
	if err != nil {
		logger.Errorln("Error connecting to the remote machine:", err.Error())
		return false
	}
	logger.Infof("Connected to the remote '%s' via SSH as user '%s'", fqdn, username)

	usr, _ := user.Current()
	fh, err := os.Open(path.Join(usr.HomeDir, ".ssh", "id_rsa.pub"))
	if err != nil {
		logger.Errorln("Error accessing id_rsa.pub key:", err.Error())
	}

	defer fh.Close()
	defer cnt.Close()

	err = cnt.CopyFile(fh, tempkey, "0600")
	if err != nil {
		logger.Errorln("Error copying pub", err.Error())
	}
	logger.Debugln("Copied", tempkey)

	rcnt, err := ssh.Dial("tcp", cnt.Host, cnt.ClientConfig)
	if err != nil {
		logger.Errorln("Error opening runner connection:", err.Error())
		return false
	}
	defer rcnt.Close()

	rssh, err := rcnt.NewSession()
	if err != nil {
		logger.Errorln("Error opening runner session:", err.Error())
		return false
	}
	defer rssh.Close()

	err = rssh.Run(fmt.Sprintf("cat %s >> /root/.ssh/authorized_keys;rm %s", tempkey, tempkey))
	if err != nil {
		logger.Errorln("Error running command:", err.Error())
		return false
	}
	logger.Debugln("Updated authorized_keys with", tempkey)

	return true
}

// Initialise cluster schema
func (n *NanoCms) StartUp() {}

// StageNode runs node stanging by calling the Node Controller
// to wipe it out and prepare for being added to the cluster.
//
// To stage any Uyuni Server, one needs to do the following:
//   - Run Uyuni Server in cluster node mode
//   - Configure and start Node Controller on it
//
// After stage node message is received on Node Controller,
// it should:
//   - reset the database to the initial state (remove all
//     systems, channels etc)
//   - Confirm this done and ack the response
//
// All this is achieved by installing a ncd binary and let it
// run a nanostate.
func (n *NanoCms) Bootstrap(fqdn string, user string, password string) *ClusterNode {
	//n.SshCopyId(fqdn, user, password)

	/*.
	NST:
		1. curl binary for arch to a location (ncd)
		2. curl service script
		3. configuration
		4. start service
		5. check working
	*/
	// That goes from the config

	logger.Debugln("Accessing state by ID:", n.bootstrapId)
	logger.Debugln("NST file:", n.GetStateIndex().GetStateById(n.bootstrapId).Path)

	n.RunStateSSH(n.LoadNstFile(n.GetStateIndex().GetStateById(n.bootstrapId).Path), fqdn)

	logger.Debugln("Bootstrap node", fqdn)
	return nil
}