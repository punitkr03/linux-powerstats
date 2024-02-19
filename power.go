package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readBatteryPower() (float64, error) {
	power_now_filePath := "/sys/class/power_supply/BAT0/power_now"
	pwr, err := os.ReadFile(power_now_filePath)

	if err != nil {
		return 0, err
	}

	powerStr := strings.TrimSpace(string(pwr))
	powerInt, err := strconv.Atoi(powerStr)
	if err != nil {
		return 0, err
	}

	return float64(powerInt) / 1000000, nil
}

func readBatteryCapacity() (int, error) {
	capacity_filepath := "/sys/class/power_supply/BAT0/capacity"
	capacity, err := os.ReadFile(capacity_filepath)
	if err != nil {
		return 0, err
	}

	capacityStr := strings.TrimSpace(string(capacity))
	capacityInt, err := strconv.Atoi(capacityStr)

	if err != nil {
		return 0, err
	}

	return capacityInt, nil
}

func readBatteryStatus() (string, error) {
	status_filepath := "/sys/class/power_supply/BAT0/status"
	status, err := os.ReadFile(status_filepath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(status)), nil
}

func displayBatteryPower(power float64) {
	fmt.Printf("Battery Power: %.3f W\n", power)
}

func displayBatteryCapacity(capacity int) {
	fmt.Printf("Battery Capacity: %d%%\n", capacity)
}

func displayBatteryStatus(status string) {
	fmt.Printf("Battery Status: %s\n", status)
}

func main() {
	for {
		clearTerminal()
		batteryPower, _ := readBatteryPower()
		batteryStatus, _ := readBatteryStatus()
		capacity, err := readBatteryCapacity()

		if err != nil {
			fmt.Printf("Error reading battery stats: %v\n", err)
			os.Exit(1)
		}
		displayBatteryStatus(batteryStatus)
		displayBatteryPower(batteryPower)
		displayBatteryCapacity(capacity)
		time.Sleep(500 * time.Millisecond)
	}
}
