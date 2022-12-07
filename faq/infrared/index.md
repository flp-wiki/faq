
### How do I add more devices to the "Universal Remotes" menu?
> While it isn't possible to add new items under the universal menu, there exist plenty of repositories containing many dumps of IR remotes. The most popular is [Flipper-IRDB](https://github.com/logickworkshop/Flipper-IRDB).
> (Note: When downloading, it's *highly recommended* to unmount the SD Card from your Flipper and directly plug it in to your computer.) If you only need a remote for one device, you can use [Flipper Maker's IR Device tool](https://flippermaker.github.io/) to create and transfer it on the go.

### The universal TV remote doesn't work besides the power button.
> The stock universal tv remote database mostly contains power codes, and very few of everything else. This file (Located at `infrared/assets/tv.ir` on the SD Card) be manually replaced with one containing extra codes for all buttons. To do so, download [this file](https://raw.githubusercontent.com/UberGuidoZ/Flipper/main/Infrared/tv.ir) and use qFlipper to transfer it into the path from the previous sentence.

### What are CSV/Pronto/IR Plus codes?
> All three are different formats of infrared databases. They are not natively compatible with Flipper, but repositories exist that hold converted and compatible versions, such as [Flipper-IRDB](https://github.com/logickworkshop/Flipper-IRDB).


