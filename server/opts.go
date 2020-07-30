package server

import (
	"errors"
	"flag"
	"fmt"

	"github.com/gobench-io/gobench/logger"
)

type serverType string

var (
	wkType serverType = "worker"
	mtType serverType = "master"
)

// Options block for gobench server
type Options struct {
	ServerType  serverType
	Addr        string
	Port        int
	ClusterPort int    // master address to solicit and connect
	Route       string // master address for worker to connect to
	Logger      logger.Logger
}

func setBaselineOptions(opts *Options) {
	opts.Logger = logger.NewStdLogger()

	if opts.Addr == "" {
		opts.Addr = DEFAULT_HOST
	}

	if opts.ServerType == wkType {
		if opts.Route == "" {
			opts.Route = fmt.Sprintf("0.0.0.0:%d", DEFAULT_CLUSTER_PORT)
		}
		return
	}

	// by default, it is a master
	opts.ServerType = mtType
	if opts.Port == 0 {
		opts.Port = DEFAULT_PORT
	}
	if opts.ClusterPort == 0 {
		opts.ClusterPort = DEFAULT_CLUSTER_PORT
	}
}

func DefaultMasterOptions() *Options {
	opts := &Options{
		ServerType: mtType,
	}
	setBaselineOptions(opts)
	return opts
}

func DefaultWorkerOptions() *Options {
	opts := &Options{
		ServerType: wkType,
	}
	setBaselineOptions(opts)
	return opts
}

// ConfigureOptions accepts a flag set and augments it with gobench specific
// flags. On success, an options structure is returned configured based on the
// selected flags
func ConfigureOptions(fs *flag.FlagSet, args []string, printVersion, printHelp func()) (*Options, error) {
	opts := &Options{}
	setBaselineOptions(opts)

	var (
		showVersion bool
		showHelp    bool
		port        int
		isMaster    bool
		isWorker    bool
	)

	fs.BoolVar(&showVersion, "v", false, "Print version information.")
	fs.BoolVar(&showVersion, "version", false, "Print version information.")
	fs.BoolVar(&showHelp, "h", false, "Show this message.")
	fs.BoolVar(&showHelp, "help", false, "Show this message.")
	fs.IntVar(&port, "p", DEFAULT_PORT, "Port of the master server.")
	fs.IntVar(&port, "port", DEFAULT_PORT, "Port of the master server.")

	fs.BoolVar(&isMaster, "m", false, "Run this server as a master.")
	fs.BoolVar(&isMaster, "master", false, "Run this server as a master.")
	fs.BoolVar(&isWorker, "w", false, "Run this server as a worker.")
	fs.BoolVar(&isWorker, "worker", false, "Run this server as a worker.")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if showVersion {
		printVersion()
		return nil, nil
	}

	if showHelp {
		printHelp()
		return nil, nil
	}

	opts.Port = port

	if isMaster && isWorker {
		return nil, errors.New("a worker cannot be master and worker at the same time")
	}

	if isMaster || !isWorker {
		opts.ServerType = mtType
	}
	if isWorker {
		opts.ServerType = wkType
	}

	return opts, nil
}
