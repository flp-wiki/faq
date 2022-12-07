
### What MicroSD Card should I use?
> - It should be a reputable brand (Like SanDisk, Sony, etc.) because often cheaper cards don't fully support the communication protocol Flipper uses. 
> - The card should have a capacity between 4 and 64 GB, but an 8 GB card is MORE than enough.
> - After inserting the card, use the Flipper's setting menu to format (clear) and test the card.
> - Before ejecting the card, unmount it via the Settings menu to ensure data isn't corrupted.
> - Note: You might need a paperclip or similar object to push the SD Card in and out of the device.
> - Read the [official documentation](https://docs.flipperzero.one/basics/sd-card) for more information!

### How do I install databases and dumps?
> Make sure there's a working MicroSD Card in the device first by following the steps above.
> Once you download the dump, you can use qFlipper or the Flipper mobile app to transfer them. If you're transfering a large file or many at once, you can also eject the SD Card from Flipper and insert it in your computer for faster transfers.
> - In qFlipper: Plug your device in, go to the file browser tab, navigate into the SD Card, and drop files in their corresponding folders (The folder names are similar to the file extensions).
> - For mobile apps: Make sure you're connected via Bluetooth, save the file to the app's archive, and synchronize it back to the device.
> - For plugging the SD Card into your PC, drop files in their corresponding folders (The folder names are similar to the file extensions).

### How do I install applications and plugins?
> Assuming the application has been packaged as a `.fap` file, installing it is as easy as placing it inside the `apps` directory on your Flipper's SD card. You can launch the app from the `Applications` app. If the app hasn't been compiled into a `.fap`, either *kindly* ask the author or compile it yourself.

### How do I install custom firmwares?
> First, ask yourself if you really need to. Sure, it might be fun to break out of Sub-GHz transmission restrictions, but how often are you actually going to do that? Is it really worth breaking the law?
> After you've ignored the previous sentences, make sure there's a working MicroSD Card in your Flipper and head over to the repository of your perferred firmware. Look for releases and find the `.dfu` file or update package (typically a `.tar`, `.tar.gz`, or `.zip` file, it always contains a file named `update.fuf`).
> - If you only have a `.dfu`, it must be installed using the "Install from file" option in qFlipper. Select the file and begin the installation.
> - If you have an update package, you can either install it with qFlipper, or install it manually through the Micro SD card by following the steps below.
> - To manually install an update package, extract and transfer the folder (not the original archive file) to the `update` folder on the SD Card (create the folder if it doesn't already exist). Once transferred, go to the desktop/idle screen of the Flipper, press down to access the file browser, then left to view all folders. From there, open the `update` folder (typically at the bottom of the list) and find the folder you just transferred. Lastly, select the file named `update` and choose "Run in app" to install the firmware.
> 
> If there was no pre-compiled update file/package, you'll have to build the firmware yourself. See the next question for details.
> For more information, read the [official documentation](https://docs.flipperzero.one/basics/firmware-update).

### Where and when are developer Q&A sessions held?
> Question and Answer session are held every week on Saturday, at 01:00 and 13:00 (GMT)
> 
> | Time zone      | Side A  |  Side B |
> | :------------: | :-----: | :-----: |
> | GMT/UTC        | 01:00   | 13:00   |
> | Pacific (PDT)  | 6:00 PM | 6:00 AM |
> | Mountain (MDT) | 7:00 PM | 7:00 AM |
> | Central (CDT)  | 8:00 PM | 8:00 AM |
> | Eastern (EDT)  | 9:00 PM | 9:00 AM |
> | China Standard | 09:00   | 21:00   |
> | India Standard | 06:30   | 18:30   |

### Are there archives of past Q&A sessions?
> Archival is a community effort, so only some are available.
> [https://github.com/flipperdevices/flipper-questions-and-answers](https://github.com/flipperdevices/flipper-questions-and-answers)

### How do I write/compile my own applications/plugins/firmware/assets?
<blockquote>
  <em>(The following is a summary of the <a href="https://github.com/flipperdevices/flipperzero-firmware/blob/dev/documentation/fbt.md">official FBT docs</a>.)</em><br>
  Since the introduction of <a href="https://github.com/flipperdevices/flipperzero-firmware/blob/dev/documentation/fbt.md">Flipper Build Tool (FBT)</a>, this has become very easy! You should have a basic understanding of working on a command line before proceeding. The only prerequisite install is <a href="https://git-scm.com/downloads">Git</a>. You should also have an IDE installed, <a href="https://code.visualstudio.com/">VSCode</a> is recommended since the <a href="https://github.com/flipperdevices/flipperzero-firmware">firmware repo</a> has <a href="https://github.com/flipperdevices/flipperzero-firmware/blob/dev/documentation/fbt.md#vscode-integration">config files</a> for it.
  <details>
    <summary>Expand me for the rest of the steps.</summary>
    <em>(WIP, sorry to curb your enthusiasm.)</em>
  </details>
</blockquote>
  
### Can I make my own Flipper instead of buying one?
> Probably not. While the firmware and schematics are mostly public, actually sourcing the components is extremely difficult. Multiple core pieces, such as the screen, were specifically produced to be used in Flipper manufacturing.
  
### How do I get a black-shell Flipper?
> This is no longer possible*, they were Kickstarter-backer exclusives.
> (*No longer possible unless you're willing to shill out hundreds of dollars for one on eBay.)
  
### How do I invert the screen/change backlight color/change case cover, etc.
> These are all hardware mods, generally inaccessible to the average user. Look up/ask around on how to do them if you're really interested, [r/flipperzero](https://old.reddit.com/r/flipperzero/) is a good place to start.

### Will there be future hardware revisions?
> Technically speaking, there's going to be a **very minor** hardware revision in the near future. Functionally speaking, it will be identical to every other Flipper Zero already sold. The revision only replaces a few internal components and doesn't offer any new features compared to existing devices, so don't bother waiting to buy it.
> Besides that, there are concepts for a [Flipper One](https://flipperzero.one/one), but without a timeline for release.

### What is Dummy Mode?
> Currently, it only allows the Snake game to be opened when active. In the future, it will hide every app except games, in case your device is ever inspected or seized.

### My device is frozen, how do I reboot/fix it?
> - To reboot the device: hold the BACK and LEFT buttons, then release simultaneously. If that didn't work, *disconnect the USB cable* and hold BACK for 30 seconds. This will preform a normal reboot.
> - To enter DFU/Recovery mode: Hold BACK and LEFT, then release BACK while still holding LEFT after a few seconds. When the screen lights up, you can release LEFT.
> - To exit DFU/Recovery mode: Follow steps for a normal reboot under the first bullet point.
> 
> If nothing works or the device is completely bricked, first make sure it's charged by plugging it in for 15-30 minutes. As a final resort, if you can't get it to turn on after charging, *unplug the USB cable* and hold OK plus BACK for 30 seconds. **There will be no indication**, but the device is now in recovery mode. Plug it in to a PC and use qFlipper to recover the firmware.
> Read the official docs for [Control](https://docs.flipperzero.one/basics/control), [Reboot](https://docs.flipperzero.one/basics/reboot), and [Firmware recovery](https://docs.flipperzero.one/basics/firmware-update/firmware-recovery).

### How do I access the CLI/Logs?
<blockquote>
  To access the Serial CLI, click one of the following based on your platform.
  <details>
    <summary>Desktop web browser*</summary>
    <em>*Chromium browsers only, such as: Google Chrome, Microsoft Edge, Opera/Opera GX, Brave, and Vivaldi.</em>
    <ul>
      <li>Connect your Flipper via USB.</li>
      <li>Ensure qFlipper and any other serial terminals are closed.</li>
      <li>Open <a href="https://my.flipp.dev/">my.flipp.dev</a> in one of the aforementioned browsers.</li>
      <li>Click <kbd>CONNECT</kbd> and select "USB Serial Device" from the list.</li>
      <li>Wait until you can see your device details on screen.</li>
      <li>Select the 💻 CLI item from the left sidebar.</li>
      <li><strong>Done!</strong></li>
    </ul>
  </details>
  <details>
    <summary>Windows</summary>
    <ul>
      <li>Install <a href="https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html">PuTTY</a> if it isn't already.</li>
      <li>Connect your Flipper via USB.</li>
      <li>Open qFlipper and look for the COM port next to the Flipper's name. <em>(Should say COM followed by a number, like COM1)</em></li>
      <li>Take note of the COM port number.</li>
      <li><strong>CLOSE qFlipper</strong>, otherwise the next steps won't work.</li>
      <li>Open PuTTY and ensure you're on the Session screen.</li>
      <li>Select "Serial" under connection type.</li>
      <li>Set serial line to the COM port. <em>(Just COM followed by the number, like COM1)</em></li>
      <li>Set speed to <code>115200</code></li>
      <li><em>Optional: Save the session settings for easy connection later.</em></li>
      <li>Finally, click <kbd>Open</kbd> to enter the CLI.</li>
      <li><strong>Done!</strong></li>
      <li>If you get an "Access Denied" error, make sure qFlipper isn't running!</li>
    </ul>
  </details>
  <details>
    <summary>MacOS/Linux</summary>
    <em>Note: I'm a filthy Windows user without any way to verify this procedure. Let me know if it's wrong!</em>
    <ul>
      <li>Install <a href="https://www.gnu.org/software/screen/">GNU Screen</a> if it isn't already.</li>
      <li>Connect your Flipper via USB.</li>
      <li>Open qFlipper and look for the device path next to the Flipper's name. <em>(Starts with /dev/tty)</em></li>
      <li><em>Alternatively: Run <code>ls /dev/tty.*</code> in a terminal.</em></li>
      <li>Take note of the full device path.</li>
      <li><strong>CLOSE qFlipper</strong>, otherwise the next steps won't work.</li>
      <li>Open a terminal.</li>
      <li>Run <code>screen PATH 115200</code>, replacing PATH with the device path from earlier.</li>
      <li><strong>Done!</strong></li>
    </ul>
  </details>
  <details>
    <summary>Android</summary>
    <ul>
      <li>Install <a href="https://play.google.com/store/apps/details?id=de.kai_morich.serial_usb_terminal">Serial USB Terminal</a> if it isn't already.</li>
      <li>Open the app and go to the Connections screen in the hamburger menu <em>(3 bars icon)</em></li>
      <li>Connect your Flipper via USB.</li>
      <li>Click the refresh icon if it doesn't automatically show up.</li>
      <li>Allow Serial USB Terminal to access Flipper if prompted.</li>
      <li>If it doesn't automatically connect, click the connect icon in the upper right. <em>(2 plugs icon)</em></li>
      <li><strong>Done!</strong></li>
      <li><em>Note: To exit log mode, you'll have to disconnect and reconnect using the icon.</em></li>
    </ul>
  </details>
  <details>
    <summary>iPhone</summary>
    Unfortunately, iOS is incapable of accessing a serial terminal over USB; try one of the other methods.
  </details>
  On the Flipper, open the settings, go to System, and set Log Level to Debug. <em>(You can keep Debug set to off unless someone asks you to turn it on)</em>
  Once you have the CLI open, type <code>log</code> and press enter to start watching logs. Press <code>Ctrl-C</code> or <code>Cmd-C</code> to exit log mode.
</blockquote>

### How can I tell if I'm running the Iceman edition firmware?
<blockquote>
  From the idle screen, press right to open your Flipper's passport.
  Check for the Iceman logo on the left, like in this screenshot:
  <details>
    <summary>(Click to reveal screenshot)</summary>
    <img src="https://user-images.githubusercontent.com/8518150/203851157-e0ce2065-dd55-4e37-a5aa-5b07ed62e872.png" alt="Iceman firmware screenshot">
  </details>
</blockquote>
