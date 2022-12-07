
### Feature/Compatability table

| Card name/type    | Read | Write | Save | Emulate | Notes                             |
| :---------------- | :--: | :---: | :--: | :-----: | :-------------------------------- |
| Mifare Classic    | ✅   | ✅    | ✅   | ✅      | Emulation can be a hit or miss    |
| Mifare DESFire    | ✅   |       |      |         | Can read public files             |
| Mifare Ultralight | ✅   |       | ✅   | ✅      | Unlock tags with various methods  |
| NTAG-21X          | ✅   |       | ✅   | ✅      | Very similar to Mifare Ultralight |
| EMV Cards         | ✅   |       | ❌   | ❌      | Can read unencrypted bank cards   |
| NFC-B             |      |       |      | ❌      | No hardware support for emulation |
| iClass/PicoPass   | ✅   |       | ✅   | ❌      | No STM SDK support for emulation  |
| EM4100/EM4102     | ✅   | ❌    | ✅   | ✅      |                                   |
| H10301            | ✅   | ❌    | ✅   | ✅      |                                   |
| Indala            | ✅   | ❌    | ✅   | ✅      |                                   |
| T5577             | ✅   | ✅    | ✅   | ✅      |                                   |
| EM4305            | ✅   |       | ✅   | ✅      |                                   |
| Paxton Net2       | ❌   | ❌    | ❌   | ❌      | No support for Hitag2             |
| Legic Prime       | ❌   | ❌    | ❌   | ❌      | Proprietary protocol              |

***Key:*** *Check = Already implimented as of latest official firmware. No mark = Could be implemented in the future. Cross mark = Unlikely to ever be implemented or impossible.*

### How do I identify which type of card/tag I have?
> To determine the protocol (NFC, RFID, or iClass/PicoPass) you'll need to attempt reading in each corresponding app. If nothing works, check the tag/card for any markings or indications. As a last resort, take a picture of the card/fob and the reader and ask in the [Flipper Discord server](https://flipperzero.one/discord).

### How do I identify which type of NFC tag I have?
> Run the "Read card" action in the NFC app. Only NFC-A type tags are supported (Mifare/NTAG/Some EMV).
> Once successfully read, the tag's type is displayed in bold at the top of the screen.

### Which NFC tags can I write?
> Currently, Mifare Classic's are the only NFC card that can be written to.
> More will be added in the future with firmware updates.

### I was told a Mifare Ultralight/NTAG tag has password-protected sectors. What does that mean?
> Either the read was interrupted, or the tag is actually password protected.
> First, try reading the tag again but make sure it stays on the back of the device until the info screen pops up.
> If you're still seeing the warning, Flipper can unlock *legally distinct NFC-enabled figurines that are pronounced like "Ameebo"* and Xaomi air filter tags, but be warned that there's a risk of **bricking** your tag if you use the wrong password too many times.

### Why does it take so long to read a Mifare Classic?
> Mifare classics are split up into sectors, these sectors are protected by two keys. To read a Mifare Classic, Flipper uses a dictionary attack, which takes a big list currently comprised of 1241 common keys, and checks them individually against each sector on the card. If you know the keys, they can be manually added to the User Dictionary under the "Extra Actions" menu.

### What does it mean when no sectors could be read on a Mifare Classic?
> The data on Mifare Classic cards is split up into sectors, and each sector is protected by two keys.
> If no sectors were read, then Flipper's dictionary attack has failed to find any valid keys.
> If you know the keys, you can manually input them under the "Extra Actions" menu of the NFC app. Otherwise, try attacking the reader with mfkey32v2 as described a few questions down.

### What does it mean when some but not all sectors could be read on a Mifare Classic?
> The data on Mifare Classic cards is split up into sectors, and each sector is protected by two keys.
> The read wasn't successful, but it didn't fail either. Some of the card's data was read and saved, but not all.
> Even if not all sectors were read, you should inspect the dump with the mobile app to see if the missing data is necessary or not. In a few rare cases, semi-read cards can be emulated in place of the original without issue.
> If you still need the rest of the keys, read the next question.

### How to I get Mifare Classic keys from a reader with [mfkey32v2](https://github.com/equipter/mfkey32v2)?
> (WIP, Note to self: https://regex101.com/r/iXmE2N/2)

### Why isn't Mifare Classic emulation working?
> Flipper emulates Mifare Classics according to official specification docs (at 13.56 mhz), however certain card readers operate at slightly different frequencies (such as 13.50 mhz). Since Flipper is unable to detect the frequency (like a real card does), it also can't correct for these minor errors.
> As a result, data transmission doesn't always occur when the reader expects it, and thus emulation is imperfect.
> There are a few theoretical ways to fix this with software, but the best option would require hardware modification.

### Why can't I save/emulate Mifare DESFire?
> DESFire is a very complicated and much more secure chipset. There are no known attacks against it yet.

### What are the .shd files in the NFC directory?
> These are shadow files, and they're created whenever an emulated tag is written to. 
> They store a copy of the original file with whatever was written. This way, the original file remains untouched.

### How do I edit the data in a saved tag?
> You'll need to use a NFC-enabled smartphone with an app that can write tags. One of the easiest to use apps is called NFC Tools, available for both [Android](https://play.google.com/store/apps/details?id=com.wakdev.wdnfc) and [iOS](https://apps.apple.com/us/app/nfc-tools/id1252962749). Due to Mifare Classic emulation quirks, you can only edit the data of saved NTAG and Mifare Ultralight tags. Create an empty NTAG216 with the "Add Manually" action in the NFC app if you don't have one already. Save that tag, then open it from the list. Once you start emulating the tag, you can use the NFC Tools smartphone app to write information on to the emulated tag. This is saved to a .shd file with the same name as the emulated tag. If you need a quick way to generate a tag containing a URL, you can use [Flipper Maker's NFC Creator tool](https://flippermaker.github.io/) online.

### Why doesn't my bank card work when I emulate it?
> EMV Credit/Debit cards are mostly encrypted. The information Flipper reads is the unencrypted portion of the card. This alone is not enough to emulate and complete a transaction. It is impossible to read the encrypted parts.

### Is there any way to save then emulate a bank card to authorize transactions?
> No, as explained in the previous question.

### Why does the NFC feature table say bank cards can be read?
> Most NFC-enabled bank cards expose their card number unencrypted. The expiration date, CVV, and ZIP code are not revealed.
> The card number alone is not enough to create a transaction, thus there's no reason to add a save option.

### Can Flipper emulate a payment terminal and authorize transactions?
> No. Are you starting to see a pattern here?

### Where is the "USB/LibNFC NFC Reader" feature mentioned in the [September blog post](https://blog.flipperzero.one/september-progress/)?
> This was scrapped due to timing issues, more details in [this GitHub issue](https://github.com/flipperdevices/flipperzero-firmware/issues/1173#issuecomment-1127728562).

### Where can I learn more about NFC and RFID technology?
> - Introduction to both Low Frequency and High Frequency: https://blog.flipperzero.one/rfid/
> - Types of NFC https://www.rfwireless-world.com/Tutorials/NFC-Type1-Tag-vs-NFC-Type2-Tag-vs-NFC-Type3-Tag-NFC-Type4-Tag-Types.html
> - Mifare Classic: https://learn.adafruit.com/adafruit-pn532-rfid-nfc/mifare
> - The Mifare Family: https://en.wikipedia.org/wiki/MIFARE
> - Datasheets: http://www.proxmark.org/files/Documents/