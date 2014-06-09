DockMon agent
=============

Docker monitoring made non-intrusive - agent part

# What is DockMon ?

DockMon is a solution to monitor any docker container without having to embed some monitoring stuff.
It only uses standard and available data sources (cgroup, namespaces, Docker API) to gather matrics on running and stopped containers on the host the agent is running on.

# What is this "agent part" thing ?

DockMon needs its agent to be running on every host running containers you want to monitor. Every time the agent is launched, it lists every container existing on the current host (running or exited) and gather metrics on each. It sends the collected data to the specified server.

# How can I configure it ?

The whole agent configuration is done in a single json file. By default, it looks in /etc/dockmon_agent.conf but it can be overriden by using -c <CONF_FILE_PATH> argument on agent invocation.
The following can be configured :
 - Docker API endpoint (either tcp or socket)
 - cgroup sysfs prefix (i.e. : if the memory.stat file path of the containers is /sys/fs/cgroup/memory/*docker*/$CID/memory.stat, the prefix should be *docker*)

# Is it safe to run it on my production server ?

Obviously not, but is it safe to run any closed source agent on your production server ? ;)

# I want to contribute !

Glad to hear that !
Everything you need to start hacking this agent is a go development environment and a text editor.
To run unittests, you'll need https://github.com/stretchr/testify : 
```
go get github.com/stretchr/testify
```

