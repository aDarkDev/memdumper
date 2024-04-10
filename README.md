# memdumper
simple memory dumper tool for linux

![test](test.png)

# Usage:

```bash
./memdump <pid>
```

This script uses `/proc/pid/maps` to detect memory addresses and will dump them using the `GDB` debugger.
