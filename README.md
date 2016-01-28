# Memory Errors
Measure the random bit flips that appear in memory from the time to time.
This flips are believed to be caused by [background radiation](https://en.wikipedia.org/wiki/ECC_memory#Problem_background). 

As the number of physical bits increases so does the chance of catching such
an error.

For our program to be successful in a short period of time, it needs to have a
lot of physical bits in use. This is accomplished by allocating a large region of
memory and then doing some OS specific things to make sure the physical memory
belongs to our program the entire time.

## ECC
Make sure you aren't using ECC memory for this experiment. Othewise, it will correct the error
before this program has a chance to see it.

## Disable Swap File
We must prevent our physical memory from being swapped out to disk, which would reduce the
number of physical bits we're able to use to detect errors. On
Linux and OS X, this can be achieved with the [mlock](http://linux.die.net/man/2/mlock)
system call.

## OS X Compressed Memory
If you're running on a recent version of OS X, your OS probably uses memory compression. Since
this memory catcher program initializes all memory to a constant (0) this makes it very
compressable. However, we don't want to compress memory in this case, we can to use as
many physical bits as possible to cast a wider net for high energy particles to hit.

To see if memory compression is being used, run:

    sysctl -a vm.compressor_mode

If that reports something besides `vm.compressor_mode: 1` then you are using compressed memory.

To disable compressed memory, run:

    sudo nvram boot-args="vm_compressor=1"

and reboot. Use the `sysctl` command above to verify it's now set to 1 (disabled).

If you want to restore memory compression when you're done with this experiment, run:

    sudo nvram boot-args="vm_compressor=4"

and reboot.

## Verifying Memory Use
There are several variables here, and we want to make sure our program is actually getting and
keeping the physical memory it's requesting. On OS X, one way to do this is to use
Activity Monitor and look at Memory, Compressed Memory, Real Mem, and Private Mem. Compressed
Memory should be 0 and the others should be equal to each other.
