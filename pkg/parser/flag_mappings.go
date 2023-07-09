package parser

import "gopkg.in/yaml.v3"

// FlagType represents the type of flag.
type FlagType int

// Flag types
const (
	ArrayType FlagType = 1 << iota
	ArrayOrMap
	BoolType
	Float64Type
	IntType
	StringType
	MapType
	MountType
	DurationType
	FileType
	UlimitType
)

func (f FlagType) YamlKind() yaml.Kind {
	switch f {
	case ArrayType, FileType:
		return yaml.SequenceNode
	case BoolType, Float64Type, IntType, StringType, DurationType:
		return yaml.ScalarNode
	case MapType, MountType, UlimitType:
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

// specialComposeTypes are top-level or parent keys that have special handling in the compose file.
// "^" indicates top level
// "$" indicates that it can be replaced
var specialComposeTypes = map[string]FlagType{
	"annotations":  MapType,
	"^services":    MapType,
	"environment":  MapType,
	"$service":     StringType,
	"^volumes":     MapType,
	"volumes":      ArrayType,
	"logging":      MapType,
	"options":      MapType, // logging.options
	"healthcheck":  MapType,
	"^networks":    MapType,
	"networks":     ArrayType,
	"default":      StringType, // networks.default
	"ports":        ArrayType,
	"blkio_config": MapType,
	"storage_opt":  MapType,
	"deploy":       MapType,
	"resources":    MapType, // deploy.resources
	"limits":       MapType, // deploy.limits or deploy.resources.limits
	"reservations": MapType, // deploy.reservations
	"sysctls":      MapType,
}

// map docker run flags to docker compose file flags
// Defined according to the specification here: https://github.com/compose-spec/compose-spec/blob/master/spec.md
var dockerRunFlags = map[string]DockerFlag{
	"a": {
		Reference: "attach",
	},
	"add-host": {
		Type:        ArrayType,
		ComposeName: "^services.$service.extra_hosts.$var",
	},
	"annotation": {
		Type:        MapType,
		ComposeName: "^services.$service.annotations.$var",
	},
	"attach": {
		Type:        ArrayType,
		ComposeName: "^services.$service.attach.$var",
	},
	"blkio-weight": {
		Type:        IntType,
		ComposeName: "^services.$service.blkio_config.weight",
	},
	"blkio-weight-device": {
		// TODO: check the spec format here https://github.com/compose-spec/compose-spec/blob/master/spec.md#blkio_config
		Type:        ArrayType,
		ComposeName: "^services.$service.blkio_config.weight_device.$var",
	},
	"c": {
		Reference: "cpu-shares",
	},
	"cap-add": {
		Type:        ArrayType,
		ComposeName: "^services.$service.cap_add.$var",
	},
	"cap-drop": {
		Type:        ArrayType,
		ComposeName: "^services.$service.cap_drop.$var",
	},
	"cgroup-parent": {
		Type:        StringType,
		ComposeName: "^services.$service.cgroup_parent",
	},
	"cgroupns": {
		Type:        StringType,
		ComposeName: "^services.$service.cgroupns_mode",
	},
	"cidfile": {
		Type:        StringType,
		ComposeName: "^services.$service.container_id_file",
	},
	"cpu-period": {
		Type:        IntType,
		ComposeName: "^services.$service.cpu_period",
	},
	"cpu-quota": {
		Type:        IntType,
		ComposeName: "^services.$service.cpu_quota",
	},
	"cpu-rt-period": {
		Type:        IntType,
		ComposeName: "^services.$service.cpu_rt_period",
	},
	"cpu-rt-runtime": {
		Type:        IntType,
		ComposeName: "^services.$service.cpu_rt_runtime",
	},
	"cpu-shares": {
		Type:        IntType,
		ComposeName: "^services.$service.cpu_shares",
		Alias:       "c",
	},
	"cpus": {
		Type:        Float64Type,
		ComposeName: "^services.$service.deploy.limits.cpus",
	},
	"cpuset-cpus": {
		Type:        StringType,
		ComposeName: "^services.$service.cpuset",
	},
	"cpuset-mems": { // TODO: not sure if this is correct
		Type:        StringType,
		ComposeName: "^services.$service.cpuset_mems",
	},
	"d": {
		Reference: "detach",
	},
	"detach": {
		Type:        BoolType,
		ComposeName: "^services.$service.detach",
		Alias:       "d",
	},
	"device": {
		Type:        ArrayType,
		ComposeName: "^services.$service.devices.$var",
	},
	"device-cgroup-rule": {
		Type:        ArrayType,
		ComposeName: "^services.$service.device_cgroup_rules.$var",
	},
	"device-read-bps": {
		Type:        ArrayType,
		ComposeName: "^services.$service.blkio_config.device_read_bps.$var",
	},
	"device-read-iops": {
		Type:        ArrayType,
		ComposeName: "^services.$service.blkio_config.device_read_iops.$var",
	},
	"device-write-bps": {
		Type:        ArrayType,
		ComposeName: "^services.$service.blkio_config.device_write_bps.$var",
	},
	"device-write-iops": {
		Type:        ArrayType,
		ComposeName: "^services.$service.blkio_config.device_write_iops.$var",
	},
	"disable-content-trust": { // TODO: not supported in compose?
		Type: BoolType,
		// ComposeName: "^services.$service.disable_content_trust",
	},
	"dns": {
		Type:        ArrayType,
		ComposeName: "^services.$service.dns.$var",
	},
	"dns-option": {
		Type:        ArrayType,
		ComposeName: "^services.$service.dns_opt.$var",
	},
	"dns-search": {
		Type:        ArrayType,
		ComposeName: "^services.$service.dns_search.$var",
	},
	"domainname": {
		Type:        StringType,
		ComposeName: "^services.$service.domainname",
	},
	"e": {
		Reference: "env",
	},
	"entrypoint": {
		Type:        StringType,
		ComposeName: "^services.$service.entrypoint",
	},
	"env": {
		Type:        MapType,
		ComposeName: "^services.$service.environment.$var",
		Alias:       "e",
	},
	"env-file": {
		Type:        ArrayType,
		ComposeName: "^services.$service.env_file.$var",
	},
	"expose": {
		Type:        ArrayType,
		ComposeName: "^services.$service.expose",
	},
	"gpus": { // TODO: to be supported
		Type: StringType,
		//ComposeName: "^services.$service.deploy.resources.reservations.gpus",
	},
	"group-add": { // TODO: not supported in docker compose v3
		Type:        StringType,
		ComposeName: "^services.$service.group_add",
	},
	"h": {
		Reference: "health-cmd",
	},
	"health-cmd": {
		Type:        StringType,
		ComposeName: "^services.$service.healthcheck.test",
	},
	"health-interval": {
		Type:        DurationType,
		ComposeName: "^services.$service.healthcheck.interval",
	},
	"health-retries": {
		Type:        StringType,
		ComposeName: "^services.$service.healthcheck.retries",
	},
	"health-start-period": {
		Type:        DurationType,
		ComposeName: "^services.$service.healthcheck.start_period",
	},
	"health-timeout": {
		Type:        DurationType,
		ComposeName: "^services.$service.healthcheck.timeout",
	},
	"hostname": {
		Type:        StringType,
		ComposeName: "^services.$service.hostname",
		Alias:       "h",
	},
	"i": {
		Reference: "interactive",
	},
	"init": {
		Type:        BoolType,
		ComposeName: "^services.$service.init",
	},
	"interactive": {
		Type:        BoolType,
		ComposeName: "^services.$service.stdin_open",
		Alias:       "i",
	},
	"ip": {
		Type:        StringType,
		ComposeName: "^services.$service.networks.default.ipv4_address",
	},
	"ip6": {
		Type:        StringType,
		ComposeName: "^services.$service.networks.default.ipv6_address",
	},
	"ipc": {
		Type:        StringType,
		ComposeName: "^services.$service.ipc",
	},
	"isolation": {
		Type:        StringType,
		ComposeName: "^services.$service.isolation",
	},
	"kernel-memory": { // TODO: to be supported
		Type: StringType,
		//ComposeName: "^services.$service.deploy.resources.reservations.memory",
	},
	"l": {
		Reference: "label",
	},
	"label": {
		Type:        ArrayType,
		ComposeName: "^services.$service.labels.$var",
		Alias:       "l",
	},
	"labels-file": {
		Type:        FileType,
		ComposeName: "^services.$service.labels",
	},
	"link": {
		Type:        ArrayType,
		ComposeName: "^services.$service.links.$var",
	},
	"link-local-ip": {
		Type:        ArrayType,
		ComposeName: "^services.$service.networks.default.link_local_ips.$var",
	},
	"log-driver": {
		Type:        StringType,
		ComposeName: "^services.$service.logging.driver",
	},
	"log-opt": {
		Type:        MapType,
		ComposeName: "^services.$service.logging.options.$var",
	},
	"mac-address": {
		Type:        StringType,
		ComposeName: "^services.$service.mac_address",
	},
	"memory": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "^services.$service.deploy.resources.limits.memory",
	},
	"memory-reservation": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "^services.$service.deploy.resources.reservations.memory",
	},
	"memory-swap": { // TODO: is this correct? For which version?
		Type:        StringType,
		ComposeName: "^services.$service.memswap_limit",
	},
	"memory-swappiness": { // TODO: is this correct? For which version?
		Type:        IntType,
		ComposeName: "^services.$service.mem_swappiness",
	},
	"mount": {
		Type:        MountType,
		ComposeName: "^services.$service.volumes.$var",
	},
	"name": {
		Type:        StringType,
		ComposeName: "^services.$service.container_name",
	},
	"network": {
		Type:        StringType,
		ComposeName: "^services.$service.network_mode",
	},
	"network-alias": {
		Type:        ArrayType,
		ComposeName: "^services.$service.networks.default.aliases.$var",
	},
	"no-healthcheck": {
		Type:        BoolType,
		ComposeName: "^services.$service.healthcheck.disable",
	},
	"oom-kill-disable": {
		Type:        BoolType,
		ComposeName: "^services.$service.oom_kill_disable",
	},
	"oom-score-adj": {
		Type:        IntType,
		ComposeName: "^services.$service.oom_score_adj",
	},
	"p": {
		Reference: "publish",
	},
	"pid": {
		Type:        StringType,
		ComposeName: "^services.$service.pid",
	},
	"pids-limit": {
		Type:        IntType,
		ComposeName: "^services.$service.pids_limit",
	},
	"platform": {
		Type:        StringType,
		ComposeName: "^services.$service.platform",
	},
	"privileged": {
		Type:        BoolType,
		ComposeName: "^services.$service.privileged",
	},
	"publish": {
		Type:        ArrayType,
		ComposeName: "^services.$service.ports.$var",
		Alias:       "p",
	},
	"publish-all": { // TODO: not sure how to handle this
		Type: BoolType,
		//ComposeName: "^services.$service.ports",
	},
	"read-only": {
		Type:        BoolType,
		ComposeName: "^services.$service.read_only",
	},
	"restart": {
		Type:        StringType,
		ComposeName: "^services.$service.restart",
	},
	"rm": { // TODO: not sure how to handle this
		Type: BoolType,
		//ComposeName: "^services.$service.remove",
	},
	"runtime": {
		Type:        StringType,
		ComposeName: "^services.$service.runtime",
	},
	"security-opt": {
		Type:        ArrayType,
		ComposeName: "^services.$service.security_opt.$var",
	},
	"shm-size": {
		Type:        StringType,
		ComposeName: "^services.$service.shm_size",
	},
	"sig-proxy": { // TODO: not supported in compose?
		Type: BoolType,
	},
	"stop-signal": {
		Type:        StringType,
		ComposeName: "^services.$service.stop_signal",
	},
	"stop-timeout": {
		Type:        DurationType,
		ComposeName: "^services.$service.stop_grace_period",
	},
	"storage-opt": {
		Type:        MapType,
		ComposeName: "^services.$service.storage_opt.$var",
	},
	"sysctl": {
		Type:        MapType,
		ComposeName: "^services.$service.sysctls.$var",
	},
	"t": {
		Reference: "tty",
	},
	"tmpfs": {
		Type:        ArrayType,
		ComposeName: "^services.$service.tmpfs.$var",
	},
	"tty": {
		Type:        BoolType,
		ComposeName: "^services.$service.tty",
		Alias:       "t",
	},
	"u": {
		Reference: "user",
	},
	"ulimit": {
		Type:        UlimitType,
		ComposeName: "^services.$service.ulimits.$var",
	},
	"user": {
		Type:        StringType,
		ComposeName: "^services.$service.user",
		Alias:       "u",
	},
	"userns": {
		Type:        StringType,
		ComposeName: "^services.$service.userns_mode",
	},
	"uts": {
		Type:        StringType,
		ComposeName: "^services.$service.uts",
	},
	"v": {
		Reference: "volume",
	},
	"volume": {
		Type:        ArrayType,
		ComposeName: "^services.$service.volumes.$var",
		Alias:       "v",
	},
	"volume-driver": { // TODO: figure out how to support this
		Type:        StringType,
		ComposeName: "",
	},
	"volumes-from": {
		Type:        ArrayType,
		ComposeName: "^services.$service.volumes_from.$var",
	},
	"w": {
		Reference: "workdir",
	},
	"workdir": {
		Type:        StringType,
		ComposeName: "^services.$service.working_dir",
		Alias:       "w",
	},
}
