VMIMG := macosx-tiger.qcow2
VMMEM := 2G

VMSYSTEM := qemu-system-ppc
VMFLAGS := -M mac99,via=pmu -cpu g4 -m $(VMMEM) -netdev user,id=net0 -device sungem,netdev=net0

all:

run:
	# The image can be found here https://archive.org/details/mac-osx-tiger-10.4-ppc-installed-qcow2-image
	# For more info read https://wiki.qemu.org/Documentation/Platforms/PowerPC#Mac_OS_9.x_to_Mac_OS_X_10.4
	# To install new software read https://www.macintoshrepository.org/

	$(VMSYSTEM) $(VMFLAGS) \
		-hda $(VMIMG) \
		-boot c
