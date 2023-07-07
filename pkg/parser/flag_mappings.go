package parser

import "gopkg.in/yaml.v3"

// FlagType represents the type of flag.
type FlagType int

// Flag types
const (
	ArrayType FlagType = iota
	BoolType
	Float64Type
	IntType
	StringType
	MapType
	MountType
	DurationType
	FileType
)

func (f FlagType) YamlKind() yaml.Kind {
	switch f {
	case ArrayType, FileType:
		return yaml.SequenceNode
	case BoolType, Float64Type, IntType, StringType, DurationType:
		return yaml.ScalarNode
	case MapType, MountType:
		return yaml.MappingNode
	}
	return yaml.ScalarNode
}

// DockerFlag represents a docker run flag.
type DockerFlag struct {
	Type FlagType
	// ComposeName is the name of the flag in the docker compose file.
	ComposeName string
	// Reference points to another flag which is an alias or uses the same compose name.
	Reference string
	// Alias is the alias of the flag in the docker run command.
	Alias string
}

// map docker run flags to docker compose file flags
var dockerRunFlags = map[string]DockerFlag{
	"a": {
		Reference: "attach",
	},
	"add-host": {
		Type:        ArrayType,
		ComposeName: "services.$service.extra_hosts",
	},
	"annotation": {
		Type:        StringType,
		ComposeName: "services.$service.annotations",
	},
	"attach": {
		Type:        ArrayType,
		ComposeName: "services.$service.attach",
	},
	"blkio-weight": {
		Type:        IntType,
		ComposeName: "services.$service.blkio_weight",
	},
	"blkio-weight-device": {
		Type:        ArrayType,
		ComposeName: "services.$service.blkio_config.weight_device",
	},
	"c": {
		Reference: "cpu-shares",
	},
	"cap-add": {
		Type:        ArrayType,
		ComposeName: "services.$service.cap_add",
	},
	"cap-drop": {
		Type:        ArrayType,
		ComposeName: "services.$service.cap_drop",
	},
	"cgroup-parent": {
		Type:        StringType,
		ComposeName: "services.$service.cgroup_parent",
	},
	"cgroupns": {
		Type:        StringType,
		ComposeName: "services.$service.cgroupns_mode",
	},
	"cidfile": {
		Type:        StringType,
		ComposeName: "services.$service.container_id_file",
	},
	"cpu-period": {
		Type:        IntType,
		ComposeName: "services.$service.cpu_period",
	},
	"cpu-quota": {
		Type:        IntType,
		ComposeName: "services.$service.cpu_quota",
	},
	"cpu-rt-period": {
		Type:        IntType,
		ComposeName: "services.$service.cpu_rt_period",
	},
	"cpu-rt-runtime": {
		Type:        IntType,
		ComposeName: "services.$service.cpu_rt_runtime",
	},
	"cpu-shares": {
		Type:        IntType,
		ComposeName: "services.$service.cpu_shares",
		Alias:       "c",
	},
	"cpus": {
		Type:        Float64Type,
		ComposeName: "services.$service.cpu_count",
	},
	"cpuset-cpus": {
		Type:        StringType,
		ComposeName: "services.$service.cpuset",
	},
	"cpuset-mems": { // TODO: not sure if this is correct
		Type:        StringType,
		ComposeName: "services.$service.cpuset_mems",
	},
	"d": {
		Reference: "detach",
	},
	"detach": {
		Type:        BoolType,
		ComposeName: "services.$service.detach",
		Alias:       "d",
	},
	"device": {
		Type:        ArrayType,
		ComposeName: "services.$service.devices",
	},
	"device-cgroup-rule": {
		Type:        ArrayType,
		ComposeName: "services.$service.device_cgroup_rules",
	},
	"device-read-bps": {
		Type:        ArrayType,
		ComposeName: "services.$service.blkio_config.device_read_bps",
	},
	"device-read-iops": {
		Type:        ArrayType,
		ComposeName: "services.$service.blkio_config.device_read_iops",
	},
	"device-write-bps": {
		Type:        ArrayType,
		ComposeName: "services.$service.blkio_config.device_write_bps",
	},
	"device-write-iops": {
		Type:        ArrayType,
		ComposeName: "services.$service.blkio_config.device_write_iops",
	},
	"disable-content-trust": { // TODO: not supported in compose?
		Type: BoolType,
		// ComposeName: "services.$service.disable_content_trust",
	},
	"dns": {
		Type:        ArrayType,
		ComposeName: "services.$service.dns",
	},
	"dns-option": {
		Type:        ArrayType,
		ComposeName: "services.$service.dns_opt",
	},
	"dns-search": {
		Type:        ArrayType,
		ComposeName: "services.$service.dns_search",
	},
	"domainname": {
		Type:        StringType,
		ComposeName: "services.$service.domainname",
	},
	"e": {
		Reference: "env",
	},
	"entrypoint": {
		Type:        StringType,
		ComposeName: "services.$service.entrypoint",
	},
	"env": {
		Type:        ArrayType,
		ComposeName: "services.$service.environment",
		Alias:       "e",
	},
	"env-file": {
		Type:        ArrayType,
		ComposeName: "services.$service.env_file",
	},
	"expose": {
		Type:        ArrayType,
		ComposeName: "services.$service.expose",
	},
	"gpus": { // TODO: to be supported
		Type: StringType,
		//ComposeName: "services.$service.deploy.resources.reservations.gpus",
	},
	"group-add": { // TODO: not supported in docker compose v3
		Type:        StringType,
		ComposeName: "services.$service.group_add",
	},
	"h": {
		Reference: "health-cmd",
	},
	"health-cmd": {
		Type:        StringType,
		ComposeName: "services.$service.healthcheck.test",
	},
	"health-interval": {
		Type:        DurationType,
		ComposeName: "services.$service.healthcheck.interval",
	},
	"health-retries": {
		Type:        StringType,
		ComposeName: "services.$service.healthcheck.retries",
	},
	"health-start-period": {
		Type:        DurationType,
		ComposeName: "services.$service.healthcheck.start_period",
	},
	"health-timeout": {
		Type:        DurationType,
		ComposeName: "services.$service.healthcheck.timeout",
	},
	"hostname": {
		Type:        StringType,
		ComposeName: "services.$service.hostname",
		Alias:       "h",
	},
	"i": {
		Reference: "interactive",
	},
	"init": {
		Type:        BoolType,
		ComposeName: "services.$service.init",
	},
	"interactive": {
		Type:        BoolType,
		ComposeName: "services.$service.stdin_open",
		Alias:       "i",
	},
	"ip": {
		Type:        StringType,
		ComposeName: "services.$service.networks.default.ipv4_address",
	},
	"ip6": {
		Type:        StringType,
		ComposeName: "services.$service.networks.default.ipv6_address",
	},
	"ipc": {
		Type:        StringType,
		ComposeName: "services.$service.ipc",
	},
	"isolation": {
		Type:        StringType,
		ComposeName: "services.$service.isolation",
	},
	"kernel-memory": { // TODO: to be supported
		Type: StringType,
		//ComposeName: "services.$service.deploy.resources.reservations.memory",
	},
	"l": {
		Reference: "label",
	},
	"label": {
		Type:        ArrayType,
		ComposeName: "services.$service.labels",
		Alias:       "l",
	},
	"labels-file": {
		Type:        FileType,
		ComposeName: "services.$service.labels",
	},
	"link": {
		Type:        ArrayType,
		ComposeName: "services.$service.links",
	},
	"link-local-ip": {
		Type:        ArrayType,
		ComposeName: "services.$service.networks.default.link_local_ips",
	},
	"log-driver": {
		Type:        StringType,
		ComposeName: "services.$service.logging.driver",
	},
	"log-opt": {
		Type:        ArrayType,
		ComposeName: "services.$service.logging.options",
	},
	"mac-address": {
		Type:        StringType,
		ComposeName: "services.$service.mac_address",
	},
	"memory": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "services.$service.deploy.resources.limits.memory",
	},
	"memory-reservation": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "services.$service.deploy.resources.reservations.memory",
	},
	"memory-swap": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "services.$service.memswap_limit",
	},
	"memory-swappiness": { // TODO: is this correct? For which version?
		Type:        IntType,
		ComposeName: "services.$service.mem_swappiness",
	},
	"mount": {
		Type:        MountType,
		ComposeName: "services.$service.volumes.$mount",
	},
	"name": {
		Type:        StringType,
		ComposeName: "services.$service.container_name",
	},
	"network": {
		Type:        StringType,
		ComposeName: "services.$service.network_mode",
	},
	"network-alias": {
		Type:        ArrayType,
		ComposeName: "services.$service.networks.default.aliases",
	},
	"no-healthcheck": {
		Type:        BoolType,
		ComposeName: "services.$service.healthcheck.disable",
	},
	"oom-kill-disable": {
		Type:        BoolType,
		ComposeName: "services.$service.oom_kill_disable",
	},
	"oom-score-adj": {
		Type:        IntType,
		ComposeName: "services.$service.oom_score_adj",
	},
	"p": {
		Reference: "publish",
	},
	"pid": {
		Type:        StringType,
		ComposeName: "services.$service.pid",
	},
	"pids-limit": {
		Type:        IntType,
		ComposeName: "services.$service.pids_limit",
	},
	"platform": {
		Type:        StringType,
		ComposeName: "services.$service.platform",
	},
	"privileged": {
		Type:        BoolType,
		ComposeName: "services.$service.privileged",
	},
	"publish": {
		Type:        ArrayType,
		ComposeName: "services.$service.ports",
		Alias:       "p",
	},
	"publish-all": { // TODO: not sure how to handle this
		Type: BoolType,
		//ComposeName: "services.$service.ports",
	},
	"read-only": {
		Type:        BoolType,
		ComposeName: "services.$service.read_only",
	},
	"restart": {
		Type:        StringType,
		ComposeName: "services.$service.restart",
	},
	"rm": { // TODO: not sure how to handle this
		Type: BoolType,
		//ComposeName: "services.$service.remove",
	},
	"runtime": {
		Type:        StringType,
		ComposeName: "services.$service.runtime",
	},
	"security-opt": {
		Type:        ArrayType,
		ComposeName: "services.$service.security_opt",
	},
	"shm-size": {
		Type:        StringType,
		ComposeName: "services.$service.shm_size",
	},
	"sig-proxy": { // TODO: not supported in compose?
		Type: BoolType,
	},
	"stop-signal": {
		Type:        StringType,
		ComposeName: "services.$service.stop_signal",
	},
	"stop-timeout": {
		Type:        DurationType,
		ComposeName: "services.$service.stop_grace_period",
	},
	"storage-opt": {
		Type:        MapType,
		ComposeName: "services.$service.storage_opt",
	},
	"sysctl": {
		Type:        MapType,
		ComposeName: "services.$service.sysctls",
	},
	"t": {
		Reference: "tty",
	},
	"tmpfs": {
		Type:        ArrayType,
		ComposeName: "services.$service.tmpfs",
	},
	"tty": {
		Type:        BoolType,
		ComposeName: "services.$service.tty",
		Alias:       "t",
	},
	"u": {
		Reference: "user",
	},
	"ulimit": {
		Type:        MapType,
		ComposeName: "services.$service.ulimits",
	},
	"user": {
		Type:        StringType,
		ComposeName: "services.$service.user",
		Alias:       "u",
	},
	"userns": {
		Type:        StringType,
		ComposeName: "services.$service.userns_mode",
	},
	"uts": {
		Type:        StringType,
		ComposeName: "services.$service.uts",
	},
	"v": {
		Reference: "volume",
	},
	"volume": {
		Type:        ArrayType,
		ComposeName: "services.$service.volumes",
		Alias:       "v",
	},
	"volume-driver": { // TODO: figure out how to support this
		Type:        StringType,
		ComposeName: "",
	},
	"volumes-from": {
		Type:        ArrayType,
		ComposeName: "services.$service.volumes_from",
	},
	"w": {
		Reference: "workdir",
	},
	"workdir": {
		Type:        StringType,
		ComposeName: "services.$service.working_dir",
		Alias:       "w",
	},
}
