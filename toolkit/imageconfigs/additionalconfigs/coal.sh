#!/bin/bash
# Setup the data partition
mkdir /data/overlays
mkdir -p /data/overlays/etc/upper
mkdir -p /data/overlays/etc/work
mkdir -p /data/overlays/home/upper
mkdir -p /data/overlays/home/work
mkdir -p /data/overlays/var/upper
mkdir -p /data/overlays/var/work
mkdir -p /data/overlays/root/upper
mkdir -p /data/overlays/root/work

# Make rootfs mount as read-only
sed -i "s/\/ ext4 defaults/\/ ext4 defaults,ro/" /etc/fstab

# Ensure data partition is mounted in initrd along with overlay
sed -i "s/data ext4 defaults/data ext4 defaults,x-initrd.mount/" /etc/fstab
echo "overlay /etc overlay x-initrd.mount,x-systemd.requires-mounts-for=/sysroot/data,lowerdir=/sysroot/etc,upperdir=/sysroot/data/overlays/etc/upper,workdir=/sysroot/data/overlays/etc/work 0 0" >> /etc/fstab
echo "overlay /home overlay x-initrd.mount,x-systemd.requires-mounts-for=/sysroot/data,lowerdir=/sysroot/home,upperdir=/sysroot/data/overlays/home/upper,workdir=/sysroot/data/overlays/home/work 0 0" >> /etc/fstab
echo "overlay /var overlay x-initrd.mount,x-systemd.requires-mounts-for=/sysroot/data,lowerdir=/sysroot/var,upperdir=/sysroot/data/overlays/var/upper,workdir=/sysroot/data/overlays/var/work 0 0" >> /etc/fstab
echo "overlay /root overlay x-initrd.mount,x-systemd.requires-mounts-for=/sysroot/data,lowerdir=/sysroot/root,upperdir=/sysroot/data/overlays/root/upper,workdir=/sysroot/data/overlays/root/work 0 0" >> /etc/fstab

# Enable initrd to break into a shell
#sed -i "s/rd.shell=0 rd.emergency=reboot/rd.shell=1 rd.break=pre-pivot/" /boot/grub2/grub.cfg

# Ensure overlay driver is available in initrd
# echo "add_drivers+=\" overlay \"" >> /etc/dracut.conf.d/01-coal.conf
# Enable systemd-repart in the initrd
# echo "add_dracutmodules+=\" systemd-repart \"" >> /etc/dracut.conf.d/01-coal.conf

# Regenerate initrd with locale in it
# dracut --force --regenerate-all --include /usr/lib/locale /usr/lib/locale
