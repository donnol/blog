ls /proc/58607

[manual](https://man7.org/linux/man-pages/man5/proc.5.html)

|文件名| 内容 | 是否需要root权限 | 说明 |
|--|--|--|--|
| arch_status | 空 | | |
| clear_refs | | 需 | |
| cpuset | / | | |
| fd | 文件描述符目录，里面有名称为0~18的链接文件 | | |
| limits | 四列数据：Limit, Soft Limit, Hard Limit, Units，关于cpu, 文件, 栈等配额信息 | | |
| mem | cat: /proc/58607/mem: Input/output error | | |
| net | 目录，有各种协议的文件: tcp, udp, icmp... | | |
| oom_score_adj | 0 | | |
| root | 目录，指向系统根目录的链接 | | |
| setgroups | allow | | |
| stat | 58607 (jdmgr) S 58382 58382 9872 34820 58382 1077936128 6387 0 26 0 6552 6932 0 0 20 0 15 0 95648476 1838948352 11496 18446744073709551615 4300800 25256385 140731858714304 0 0 0 0 0 2143420159 0 0 0 17 8 0 0 0 0 0 49557040 50499008 83853312 140731858719703 140731858719777 140731858719777 140731858722800 0 | | |
| task | 目录，里面的子目录与本进程的目录结构类似 | | |
| uid_map |          0          0 4294967295 | | |
| attr | 目录，里面有：current, execr, fscreate, keycreate, prev, sockcreate | | |
| cmdline | ./jdmgr--config=/home/jd/Project/jdmgr/data/conf/jdmgr-local.tomlserver | | |
| cwd | 目录，里面有启动进程所在目录的文件 | | |
| fdinfo | 目录，里面有名称为0~18的文件，每个文件里有以下信息：pos, flags, mnt_id, ino | | |
| loginuid | 4294967295 | | |
| mountinfo | 磁盘挂载情况 | | |
| ns | 目录，里面有：cgroup, net, time, uts, ipc, pid, time_for_children, mnt, pid_for_children, user | | |
| pagemap | 一堆问号 | | |
| sched | jdmgr (58607, #threads: 15) 和一系列键值 | | |
| smaps | 很多键值数据 | | |
| statm | 448962 11467 6999 5117 0 46089 0 | | |
| timens_offsets | monotonic 0 0; boottime 0 0 | | |
| wchan | futex_wait_queue_me | | |
| auxv | 一堆问号 | | |
| comm | jdmgr | | |
| environ | 环境变量 | | |
| gid_map |          0          0 4294967295 | | |
| map_files | 目录 | 需 | |
| mounts | 挂载情况：none /mnt/wsl tmpfs rw,relatime 0 0 | | |
| oom_adj | 0 | | |
| personality | 00000000 | | |
| schedstat | 14831015600 63888900 226226 | | |
| smaps_rollup | 一系列键值：Rss: 47844 kB;Pss: 47781 kB;... | | |
| status | Name:  jdmgr; Umask: 0022; State: S (sleeping); Tgid:  58607 | | |
| timers | 空 | | |
| cgroup | 15:name=systemd:/; 14:misc:/; 13:rdma:/; 12:pids:/; 11:hugetlb:/; 10:net_prio:/; 9:perf_event:/; 8:net_cls:/; 7:freezer:/; 6:devices:/; 5:memory:/; 4:blkio:/; 3:cpuacct:/; 2:cpu:/; 1:cpuset:/; 0::/ | | |
| coredump_filter | 00000033 | | |
| exe | 一堆问号 | | |
| io | rchar: 65191135; wchar: 10478455; syscr: 668811; syscw: 233308; read_bytes: 352256; write_bytes: 1781760; cancelled_write_bytes: 0 | | |
| maps | 00400000-0041a000 r--p 00000000 08:20 2198269 /home/jd/Project/jdmgr/jdmgr; ... | | |
| mountstats | device none mounted on /mnt/wsl with fstype tmpfs; ... | | |
| oom_score | 668 | | |
| projid_map |          0          0 4294967295 | | |
| sessionid | 4294967295 | | |
| stack | [<0>] do_epoll_wait+0x5ce/0x710; [<0>] do_compat_epoll_pwait.part.0+0xe/0x80; [<0>] __x64_sys_epoll_pwait+0x7f/0x130; [<0>] do_syscall_64+0x38/0xc0; [<0>] entry_SYSCALL_64_after_hwframe+0x62/0xcc | 需 | |
| syscall | 202 0x303c728 0x80 0x0 0x0 0x0 0x0 0x7ffeb0727b10 0x48d183 | | |
| timerslack_ns | 50000 | 需 | |
