// BlueZ D-Bus LE Advertising API Description [advertising-api.txt]
// Advertising packets are structured data which is broadcast on the LE Advertising
// channels and available for all devices in range.  Because of the limited space
// available in LE Advertising packets (31 bytes), each packet's contents must be
// carefully controlled.
// BlueZ acts as a store for the Advertisement Data which is meant to be sent.
// It constructs the correct Advertisement Data from the structured
// data and configured the kernel to send the correct advertisement.
// Advertisement Data objects are registered freely and then referenced by BlueZ
// when constructing the data sent to the kernel.
package advertising
