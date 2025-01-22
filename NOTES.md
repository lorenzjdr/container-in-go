# 3 main cores
- Namespaces - provide isolation; appear like they own the environment
- CGroups
- layered filesystems

## Namespaces - 6 
- PID: Creates like a mapping table for programs to see their ID's.
- MNT: Creates it's own mount table. The process doesn't affect other namespaces.
- NET: Own (virtual) network stack. 
- UTS: Allows own view of hostname and domain name.
- IPC: Isolates communication.
- USER: Trick the processes in a container to have root access. Mapped to different underpriveleged host providing isolation.

## CGroups
- Enforce fair or unfair resource sharing.
- Manage processes and give them limits.