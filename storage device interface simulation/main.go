// storagedevice interface [interface working]
package main

import "fmt"

// Signature needed inside: Write(payload string) bool
type StorageDevice interface {
	Write(payload string) bool
}

// 2. Concrete Type 1: RAMBuffer
type RAMBuffer struct {
	Capacity int
	Data     []string
}

// Write(payload string) bool for *RAMBuffer (Pointer Receiver)
func (buf *RAMBuffer) Write(payload string) bool {
	if len(buf.Data) >= buf.Capacity {
		fmt.Println("[RAMBuffer] buffer is full")
		return false
	}
	buf.Data = append(buf.Data, payload)
	return true
}

// Concrete Type 2: DiskDrive
type DiskDrive struct {
	DriveLetter string
	IsOnline    bool
}

// Write(payload string) bool for DiskDrive
func (buf DiskDrive) Write(payload string) bool {
	if !buf.IsOnline {
		return false
	}
	fmt.Printf("--> Saved [%s] to disk [%s]\n", payload, buf.DriveLetter)
	return true
}

// 4. The Polymorphic Pipeline
// BroadcastBackup(devices []StorageDevice, payload string)
func BroadcastBackup(devices []StorageDevice, payload string) {
	for _, dev := range devices {
		success := dev.Write(payload)
		if !success {
			fmt.Println("-->ALERT! A device Failed to write ")
		}

	}
}

// It should loop through all devices, call Write(payload), and alert if any return false.

func main() {
	// Initialize our concrete structs
	ram := &RAMBuffer{Capacity: 2, Data: make([]string, 0)}
	disk1 := &DiskDrive{DriveLetter: "C", IsOnline: true}
	disk2 := &DiskDrive{DriveLetter: "D", IsOnline: false} // This one is offline!

	// Pack completely different structs into a single slice of interfaces!
	// Why does this work? Because all 3 satisfy the StorageDevice interface!
	myDevices := []StorageDevice{ram, disk1, disk2}

	fmt.Println("--- Broadcast 1 ---")
	BroadcastBackup(myDevices, "CRITICAL_SYS_LOG_01")

	fmt.Println("\n--- Broadcast 2 ---")
	BroadcastBackup(myDevices, "CRITICAL_SYS_LOG_02")

	fmt.Println("\n--- Broadcast 3 (Should trigger RAM overflow!) ---")
	BroadcastBackup(myDevices, "CRITICAL_SYS_LOG_03")

	// Inspect the RAM buffer directly at the end
	fmt.Println("\nFinal RAM Buffer Contents:", ram.Data)
}
