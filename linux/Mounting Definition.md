Mounting is the attaching of an additional filesystem to the currently accessible filesystem of a computer.

A filesystem is a hierarchy of directories (also referred to as a directory tree) that is used to organize files on a computer or storage media (e.g., a CDROM or floppy disk). On computers running Linux or other Unix-like operating systems, the directories start with the root directory, which is the directory that contains all other directories and files on the system and which is designated by a forward slash ( / ). The currently accessible filesystem is the filesystem that can be accessed on a computer at a given time.

In order to gain access to files on a storage device, **the user must first inform the operating system where in the directory tree to mount the device.** A device in a mounting context can be a partition (i.e., a logically independent section) on a hard disk drive (HDD), a CDROM, a floppy disk, a USB (universal serial bus) key drive, a tape drive, or any other external media. For example, to access the files on a CDROM, the user must inform the system to make the filesystem on the CDROM appear in some directory, typically /mnt/cdrom (which exists for this very purpose).

The mount point is the directory (usually an empty one) in the currently accessible filesystem to which a additional filesystem is mounted. It becomes the root directory of the added directory tree, and that tree becomes accessible from the directory to which it is mounted (i.e., its mount point,so a mount point is basicly an entry point). Any original contents of a directory that is used as a mount point become invisible and inaccessible while the filesystem is still mounted.

The /mnt directory exists by default on all Unix-like systems. It, or usually its subdirectories (such as /mnt/floppy and /mnt/usb), are intended specifically for use as mount points for removable media such as CDROMs, USB key drives and floppy disks.

On some operating systems, everything is mounted automatically by default so that users are never even aware that there is any such thing as mounting. Linux and other Unix-like systems can likewise be configured so that everything is mounted by default, as a major feature of such systems is that they are highly configurable. However, they are not usually set up this way, for both safety and security reasons. Moreover, only the root user (i.e., administrative user) is generally permitted by default to mount devices and filesystems on such systems, likewise as safety and security measures.

In the simplest case, such as on some personal computers, the entire filesystem on a computer running a Unix-like operating system resides on just a single partition, as is typical for Microsoft Windows systems. More commonly, it is spread across several partitions, possibly on different physical disks or even across a network. Thus, for example, the system may have one partition for the root directory, a second for the /usr directory, a third for the /home directory and a fourth for use as swap space. (Swap space is a part of HDD that is used for virtual memory, which is the simulation of additional main memory).

The only partition that can be accessed immediately after a computer boots (i.e., starts up) is the root partition, which contains the root directory, and usually at least a few other directories as well. The other partitions must be attached to this root filesystem in order for an entire, multiple-partition filesystem to be accessible. Thus, about midway through the boot process, the operating system makes these non-root partitions accessible by mounting them on to specified directories in the root partition.

Systems can be set up so that external storage devices can be mounted automatically upon insertion. This is convenient and is usually satisfactory for home computers. However, it can cause security problems, and thus it is usually not (or, at least, should not be) permitted for networked computers in businesses and other organizations. Rather, such devices must be mounted manually after insertion, and such manual mounting can only be performed by the root account.

Mounting can often be performed manually by the root user by merely using the mount command followed by the name of the device to be mounted and its mounting destination (but in some cases it is also necessary to specify the type of filesystem). For example, to mount the eighth partition on the first HDD, which is designated by /dev/hda8, using a directory named /dir8 as the mount point, the following could be used:
```
mount /dev/hda8 /dir8
```
Removing the connection between the mounted device and the rest of the filesystem is referred to as unmounting. It is performed by running the umount (with no letter n after the first u) command, likewise followed by the name of the device to be unmounted and its mount point. For example, to unmount the eighth partition from the root filesystem, the following would be used:
```
umount /dev/hda8 /dir8
```
A list of the devices that are currently mounted can be seen by viewing the /etc/fstab file. This plain text configuration file also shows the mount points and other information about the devices, and it is employed during the boot process to tell the system which partitions to automatically mount. It can be safely viewed by using the cat command, i.e.,
```
cat /etc/fstab
```