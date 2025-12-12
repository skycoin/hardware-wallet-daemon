# Linux Setup Guide for Skycoin Hardware Wallet

This guide explains how to set up your Linux system to use the Skycoin Hardware Wallet (SkyWallet) without requiring root/sudo privileges.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Understanding the Issue](#understanding-the-issue)
- [Installation Steps](#installation-steps)
  - [1. Install udev Rules](#1-install-udev-rules)
  - [2. Verify Installation](#2-verify-installation)
- [Troubleshooting](#troubleshooting)
  - [Error: "libusb: bad access [code -3]"](#error-libusb-bad-access-code--3)
  - [Device Not Found](#device-not-found)
  - [Permission Denied](#permission-denied)
- [Advanced: Manual Kernel Driver Unbinding](#advanced-manual-kernel-driver-unbinding)
- [Technical Details](#technical-details)

## Prerequisites

- Linux system with udev (most modern distributions)
- libusb-1.0 or later
- Go 1.10+ (for building from source)
- SkyWallet device

## Quick Start

For most users, simply install the udev rules:

```bash
cd $GOPATH/src/github.com/skycoin/hardware-wallet-daemon
sudo cp udev/51-skywallet.rules /etc/udev/rules.d/
sudo udevadm control --reload-rules
sudo udevadm trigger
```

Then unplug and replug your SkyWallet device.

## Understanding the Issue

When you connect a SkyWallet to a Linux system, two things need to happen:

1. **USB Permissions**: Your user account needs permission to access the USB device
2. **Kernel Driver**: The Linux kernel automatically binds a HID (Human Interface Device) driver to the SkyWallet, which prevents userspace programs from accessing it

The udev rules in this repository solve both problems automatically.

## Installation Steps

### 1. Install udev Rules

Copy the provided udev rules file to your system's udev configuration directory:

```bash
sudo cp udev/51-skywallet.rules /etc/udev/rules.d/
```

Reload udev rules to activate the changes:

```bash
sudo udevadm control --reload-rules
sudo udevadm trigger
```

**Note**: The second command triggers the rules for already-connected devices. If your SkyWallet is already plugged in, you may need to unplug and replug it.

### 2. Verify Installation

After installing the udev rules and reconnecting your device, verify the setup:

**Check USB permissions:**

```bash
# Find your device
lsusb | grep 313a

# Check device file permissions (adjust bus/device numbers as shown by lsusb)
ls -l /dev/bus/usb/001/012
```

You should see permissions like `crw-rw-rw-` or `crw-rw-r--+` (the `+` indicates ACLs).

**Check kernel driver status:**

```bash
# Find the device in sysfs
find /sys/bus/usb/devices -name "*313a*" -type d

# Check if a driver is bound (should be empty or not exist)
ls -l /sys/bus/usb/devices/1-1.5:1.0/driver 2>/dev/null
```

If the last command shows a symlink to `usbhid`, the kernel driver unbind didn't work. See [Troubleshooting](#troubleshooting) below.

**Test the daemon:**

```bash
cd $GOPATH/src/github.com/skycoin/hardware-wallet-daemon
make run-usb
```

In another terminal:

```bash
curl http://127.0.0.1:9510/api/v1/features
```

You should see JSON output with your device's features, not an error.

## Troubleshooting

### Error: "libusb: bad access [code -3]"

This usually means the kernel HID driver is still bound to the device.

**Solution 1: Manual Unbind (Temporary)**

Create a helper script:

```bash
#!/bin/bash
# unbind-skywallet.sh

INTERFACE=$(find /sys/bus/usb/devices -type l -name "driver" 2>/dev/null | \
    while read -r link; do \
        iface=$(dirname "$link"); \
        if grep -q "313a" "$iface/../idVendor" 2>/dev/null && \
           grep -q "0001" "$iface/../idProduct" 2>/dev/null && \
           grep -q "03" "$iface/bInterfaceClass" 2>/dev/null; then \
            echo "$iface"; \
            break; \
        fi; \
    done)

if [ -z "$INTERFACE" ]; then
    echo "SkyWallet not found or already unbound"
    exit 0
fi

IFACE_NAME=$(basename "$INTERFACE")
echo "Unbinding $IFACE_NAME from usbhid driver..."
echo "$IFACE_NAME" | sudo tee /sys/bus/usb/drivers/usbhid/unbind > /dev/null

echo "Done. You can now use the hardware wallet without sudo."
```

Make it executable and run it:

```bash
chmod +x unbind-skywallet.sh
./unbind-skywallet.sh
```

**Note**: This solution is temporary - you'll need to run the script each time you plug in the device.

**Solution 2: Check udev Rules**

Verify the udev rules are installed correctly:

```bash
ls -l /etc/udev/rules.d/51-skywallet.rules
```

If the file doesn't exist, reinstall it as described in [Installation Steps](#installation-steps).

Check udev logs for errors:

```bash
sudo udevadm test /sys/bus/usb/devices/1-1.5:1.0 2>&1 | grep -i error
```

(Adjust the device path to match your system - find it with `lsusb -t`)

**Solution 3: Reload udev and Reconnect**

Sometimes udev needs a more forceful reload:

```bash
sudo udevadm control --reload-rules
sudo systemctl restart systemd-udevd
```

Then unplug and replug your device.

### Device Not Found

If the daemon reports "device not found":

**Check USB connection:**

```bash
lsusb | grep 313a
```

You should see:
```
Bus 001 Device 012: ID 313a:0001 SkycoinFoundation SKYWALLET
```

If not, try a different USB cable or port.

**Check dmesg for errors:**

```bash
sudo dmesg | tail -20
```

Look for USB-related errors or warnings.

### Permission Denied

If you get "permission denied" errors:

**Check group membership:**

Some systems require users to be in specific groups (e.g., `plugdev`, `dialout`):

```bash
groups
```

If you're not in these groups, add yourself:

```bash
sudo usermod -a -G plugdev $USER
sudo usermod -a -G dialout $USER
```

Then log out and log back in for changes to take effect.

**Verify udev rule syntax:**

```bash
sudo udevadm test $(udevadm info -q path -n /dev/bus/usb/001/012) 2>&1 | grep skywallet
```

(Adjust device path as needed)

## Advanced: Manual Kernel Driver Unbinding

If automatic unbinding via udev doesn't work, you can unbind the kernel driver manually after each device connection.

### Option 1: Simple Unbind Script

See [Solution 1](#error-libusb-bad-access-code--3) above.

### Option 2: Kernel Module Quirks

You can tell the kernel to ignore the device entirely using module parameters.

**Create a modprobe config file:**

```bash
sudo tee /etc/modprobe.d/skywallet-quirks.conf << 'EOF'
# Prevent usbhid from binding to SkyWallet
options usbhid quirks=0x313a:0x0001:0x0004
EOF
```

**Reload the usbhid module:**

```bash
sudo rmmod usbhid
sudo modprobe usbhid
```

**Note**: This is system-wide and permanent. The quirks flag `0x0004` tells usbhid to ignore the device.

### Option 3: udev RUN Action

This is what the provided udev rules do automatically. If you need to customize it, edit `/etc/udev/rules.d/51-skywallet.rules`:

```
SUBSYSTEM=="usb", ATTR{idVendor}=="313a", ATTR{idProduct}=="0001", RUN+="/path/to/your/unbind/script.sh"
```

## Technical Details

### Why Does This Happen?

The SkyWallet uses USB HID (Human Interface Device) class for communication. When Linux detects a HID device, it automatically loads the `usbhid` kernel driver to handle it. This driver is designed for keyboards, mice, and game controllers.

However, for hardware wallets, we need **direct userspace access** to the USB device via libusb. The kernel driver and userspace libusb cannot both access the device simultaneously.

### What Do the udev Rules Do?

The provided udev rules (`51-skywallet.rules`) do two things:

1. **Set Permissions**: `MODE="0666", TAG+="uaccess"`
   - Makes the device readable and writable by all users
   - The `uaccess` tag provides better security on systemd-based systems

2. **Unbind Kernel Driver**: `RUN+="/bin/sh -c '...'"`
   - Automatically detaches the `usbhid` driver when the device is plugged in
   - Allows libusb to claim the device immediately

### Security Considerations

**MODE="0666" vs TAG+="uaccess":**

- `MODE="0666"` makes the device world-readable/writable (less secure)
- `TAG+="uaccess"` limits access to the logged-in user (more secure)
- Both are included for compatibility; modern systems prefer `uaccess`

**Automatic Driver Unbinding:**

- The udev rule runs a shell command with root privileges
- The command is carefully crafted to only unbind the SkyWallet interface
- Review the script in `51-skywallet.rules` before installing if you have security concerns

### Alternative: Using sudo

If you prefer not to install udev rules, you can run the daemon with sudo:

```bash
sudo make run-usb
```

However, this is **not recommended** for security reasons:
- The daemon has network access (HTTP API)
- Running network services as root increases attack surface
- udev rules provide a more secure, user-friendly solution

### Supported Devices

These rules are specifically for:
- **Vendor ID**: `0x313a` (Skycoin Foundation)
- **Product ID**: `0x0001` (SKYWALLET)

If you have a different hardware wallet, you'll need different udev rules.

## Getting Help

If you encounter issues not covered in this guide:

1. **Check logs**: Look for errors in `dmesg`, `journalctl -xe`, and daemon output
2. **Verify device**: Use `lsusb -v -d 313a:0001` to inspect device details
3. **Test permissions**: Try running the daemon with `sudo` to rule out permission issues
4. **Report issues**: Open an issue on GitHub with:
   - Your Linux distribution and version
   - Output of `lsusb -v -d 313a:0001`
   - Output of `udevadm info --query=all --name=/dev/bus/usb/XXX/YYY`
   - Daemon error messages

## Contributing

If you've found a solution that works for your Linux distribution, please contribute:

1. Test the solution thoroughly
2. Document the steps clearly
3. Submit a pull request with updates to this guide

## References

- [libusb documentation](https://libusb.info/)
- [Writing udev rules](http://reactivated.net/writing_udev_rules.html)
- [systemd uaccess](https://www.freedesktop.org/software/systemd/man/systemd-udevd.service.html)
- [USB HID specification](https://www.usb.org/hid)
