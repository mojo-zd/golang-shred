Linux FS

- Bootfs

1. Bootloader 引导加载kernel
2. kernel 当kernel被加载到内存中后umount bootfs

- rootfs

1. /dev、/proc、/bin、/etc等标准目录和文件
2. 对于不同的linux发行版, bootfs基本是一致的, 但rootfs会有差别

Docker启动

存储驱动

- AUFS
- OverlayFS
- Device Mapper
- Btrfs
- ZFS

练习:
```
mkdir overlayfs
mkdir upper lower merged work
echo "from lower" > lower/in_lower.txt
echo "from upper" > upper/in_upper.txt
echo "from lower" > lower/in_both.txt
echo "from upper" > upper/in_both.txt
sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
```
