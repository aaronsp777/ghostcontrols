# ghostcontrols
Decoder/Encoder for GhostControls Gate Remotes

[GhostControls](https://www.ghostcontrols.com) makes a variety of automatic gate operators, transmitters and keypads & receivers.

The protocol has a number of noteworthy features such as toggling a gate open/closed.
* Gates can be pinned open for “party” mode.
* Visor remotes can be locked out for “vacation” mode.
* Remotes can clone other remotes.
* All remotes on a given system use the same shared secret key.
* Remotes can generate a new secret key when requested.

# Receiving
* Get an inexpensive RTL2832U dongle
* Install https://github.com/merbanan/rtl_433
* run rtl_433 -c rtl_433.conf


# Transmitting
* Get a raspberry pi
* Install https://github.com/F5OEO/rpitx
* Add an antenna
* Compile & run send.go
