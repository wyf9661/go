// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo.exe -godefs -- -DSYLIXOS -ID:/WS/SylixOS_WS/base1/libsylixos/SylixOS -ID:/WS/SylixOS_WS/base1/libsylixos/SylixOS/include -ID:/WS/SylixOS_WS/base1/libsylixos/SylixOS/include/network types_sylixos.go

//go:build sylixos

package syscall

const (
	sizeofPtr      = 0x8
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x8
	sizeofLongLong = 0x8
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur int64
	Max int64
}

type _Gid_t uint32

type Stat_t struct {
	Dev     uint64
	Ino     uint64
	Mode    int32
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    uint64
	Size    int64
	Atime   int64
	Mtime   int64
	Ctime   int64
	Blksize int64
	Blocks  int64
	Resv1   *byte
	Resv2   *byte
	Resv3   *byte
}

type Statfs_t struct {
	Type    int64
	Bsize   int64
	Blocks  int64
	Bfree   int64
	Bavail  int64
	Files   int64
	Ffree   int64
	Fsid    Fsid
	Flag    int64
	Namelen int64
	Spare   [7]int64
}

type Flock_t struct {
	Type   int16
	Whence int16
	Start  int64
	Len    int64
	Pid    int32
	Xxx    [4]int64
}

type Dirent struct {
	Name      [513]uint8
	Type      uint8
	Shortname [13]uint8
	Resv      **byte
}

type Fsid struct {
	Val [2]int32
}

const (
	pathMax = 0x200
)

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [104]uint8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]uint8
	Rcf    uint16
	Route  [16]uint16
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [26]uint8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [80]uint8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex int32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

const (
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x36
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
)

type FdSet struct {
	Bits [32]uint64
}

const (
	SizeofIfMsghdr         = 0xa8
	SizeofIfData           = 0x98
	SizeofIfaMsghdr        = 0x14
	SizeofIfAnnounceMsghdr = 0x18
	SizeofRtMsghdr         = 0x98
	SizeofRtMetrics        = 0x70
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Data    IfData
}

type IfData struct {
	Type       uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Recvquota  uint8
	Xmitquota  uint8
	Mtu        uint64
	Metric     uint64
	Baudrate   uint64
	Ipackets   uint64
	Ierrors    uint64
	Opackets   uint64
	Oerrors    uint64
	Collisions uint64
	Ibytes     uint64
	Obytes     uint64
	Imcasts    uint64
	Omcasts    uint64
	Iqdrops    uint64
	Noproto    uint64
	Recvtiming uint64
	Xmittiming uint64
	Lastchange Timeval
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Metric  int32
}

type IfAnnounceMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Name    [16]uint8
	What    uint16
}

type RtMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Flags   int32
	Addrs   int32
	Pid     int32
	Seq     int32
	Errno   int32
	Use     int32
	Inits   uint64
	Rmx     RtMetrics
}

type RtMetrics struct {
	Locks    uint64
	Mtu      uint64
	Hopcount uint64
	Expire   uint64
	Recvpipe uint64
	Sendpipe uint64
	Ssthresh uint64
	Rtt      uint64
	Rttvar   uint64
	Pksent   uint64
	Filler   [4]uint64
}

type Termios struct {
	Iflag uint32
	Oflag uint32
	Cflag uint32
	Lflag uint32
	Line  uint8
	Cc    [19]uint8
}

type Utsname struct {
	Sysname  [16]uint8
	Nodename [513]uint8
	Release  [64]uint8
	Version  [128]uint8
	Machine  [64]uint8
}

type posix_spawnopt_t struct {
	ISigNo       int32
	UlId         uint64
	UlMainOption uint64
	StStackSize  uint64
}

type sched_param struct {
	Priority        int32
	Ss_low_priority int32
	Ss_repl_period  Timespec
	Ss_init_budget  Timespec
	Ss_max_repl     int32
	Pad             [12]uint64
}

type posix_spawnattr_t struct {
	SFlags        int16
	PidGroup      int32
	SigsetDefault uint64
	SigsetMask    uint64
	Schedparam    sched_param
	IPolicy       int32
	PcWd          *uint8
	Presraw       *uint8
	Opt           posix_spawnopt_t
	UlExts        uint64
	UlPad         [5]uint64
}

type posix_spawn_file_actions_t struct {
	PplineActions **uint8
	Presraw       *uint8
	IInited       int32
	UlPad         [16]uint64
}
