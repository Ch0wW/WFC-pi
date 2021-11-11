# WFC-pi

WFC-pi is a small project based on Raspberry Pi OS Lite that provides a very simple headless AP that can be used to replay online with the Nintendo DS / Nintendo Wii/vWii.

You can also use that access point to play WiiU / N3DS / Switch online without compatibility or connection issues.

**Please note that the Wii still requires patching your games in order to play online!**

# Downloads
The releases can be downloaded @ http://dl.baseq.fr/wfcpi/. 7Zip or WinRAR is required to extract the archive.

# Features
* Creates a hidden guest access point (`wfcpi`) that is isolated with the rest of your network.
* Connects you automatically to Wiimmfi & Riiconnect24 services (as long as you installed/patched them).
    * If you don't know how to hack it, [Wii.guide](https://wii.guide/) gives all you need to know in order to hack it in less than 10 minutes.

# Prequirements 
In order for WFC-pi to properly work, you will need:
* An Ethernet adapter
* A MicroSD with 4GB of capacity or more

# Compatibility
WFC-pi works with the following devices:

### Out-of-the-box
* Raspberry Pi 3B+ (all versions)
* Raspberry Pi 4B+ (all versions)

### Requires a compatible USB-Ethernet adapter
* Raspberry Pi Zero W
* Raspberry Pi Zero W 2

### Requires a compatible USB-Wifi adapter
* Rasbperry Pi 1B (all versions)
* Raspberry Pi 2B (all versions)

# USB Compatibility-list
Please check the Wiki to check the compatibility list, both for the WiFi and Ethernet drivers.
Additionally, if your device isn't listed but you tested it, please create an issue so that we could update it!

# Installation
WFC-pi is installed the exact same way as any image available for Raspberry Pis. We heavily recommand you using [USBImager](https://gitlab.com/bztsrc/usbimager) rather than that bloaty BalenaEtcher to install that image to your SD Card.

If you want to manually add WFC-pi to your local image, simply execute the scripts in order located in `/scripts`. Please note you'll need to modify some values. (more info in its `README.md` file).


# Tools and manual installation

WFC-pi includes some tools and their sources used in the image provided:

* **wfcpi-dns-updater** : Automatic DNS updater written in golang

# ToDo
- [ ] Find a way to heaviliy compress or reduce the total size of the distribution
- [ ] Find a way to include sudomemo support while also being able to connect to Wiimmfi
- [ ] A Web interface?
- [ ] (when it's done) Streetpass2 compatiblity for Nintendo 3DS?
- [ ] (when it's done) Support for Activision's Wii's Masterserver replacement?