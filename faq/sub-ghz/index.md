
### How do I hack my neighbors garage or unlock some random persons car?!?
> Short answer: You don't. That's illegal, and NOT what Flipper was designed for.

### What does "This frequency can only be used for RX in your region" mean?
> Due to legal regulations, Flipper is not allowed to transmit on certain frequencies depending on your device's provisioned location.
> Provisioning occurs whenever you update your firmware via qFlipper or the mobile app and is based on your rough location.

### How do I find the frequency of a device/transponder?
> If it's a commonly used frequency, bring the device *really close* to the Flipper and use the Frequency analyzer.
> If that didn't work, check for the device's FCC ID. It's legally required to be somewhere on the device if it's sold in the US.
> Then, look up that ID on [FCC ID.io](https://fccid.io). 

### I can't tune Flipper to capture a specific frequency.
> You'll need to edit the `setting_user` and `setting_frequency_analyzer_user` to change the frequencies available for selection in the app. The files are located in `subghz/assets` on the SD card.
> Note that this won't magically unlock those frequencies, you're still bound by the device's limitations.

### I captured a garage/car/etc. signal, but it doesn't work when I replay it.
> Unless the item of interest is extremely old, it probably uses rolling codes. Read more below.

### What is a rolling code?
> Think of it like this: Imagine your garage door was programmed to open whenever it received the code "1234" from a transponder.
> This would be a static code, where a replay attack (Read RAW) would be able to open the garage.
> Since replay attacks are so easy, most devices will shuffle the code after each use.
> So the first time you open your garage, the transponder sends "1234" and the second time it sends "5678."
> Rolling codes aren't that simple, but you get the gist.

### I replayed a rolling code and now my original keyfob/transponder doesn't work.
> You'll have to re-sync your old device manually, since it's now lagging behind on the rolling code.

### What is a Debruin/Brute force code?
> A brute force code tries every possible code for a specific bit length, however this is inefficient.
> Example: 0001, 0002, 0003, 0004 ... 9998, 9999.
> Debruin sequences are more efficient by merging multiple codes together.
> Example: 365, 136, and 650 can all be found in 13650 by looking at groups of 3 digits individually.

### Can I attach a more powerful antenna?
> Yes and no. You can't just attach any antenna directly via the GPIO pins, however you could use a separate processor on a protoboard and control it from Flipper, assuming you write your own code to do that.
> For example, you could write your own code on a NRF24 and accompanying Flipper app to control it over GPIO. 

