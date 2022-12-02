package temp

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Block struct{}

func New() *Block {
	return &Block{}
}

func (b *Block) Name() string {
	return "temp"
}

func (b *Block) Render() (string, error) {
	out, err := exec.Command("sensors", "-j").Output()
	if err != nil {
		return "", fmt.Errorf("exec sensors failed: %v", err)
	}

	var sensors = struct {
		Thinkpad struct {
			CPU struct {
				Temp float64 `json:"temp1_input"`
			} `json:"CPU"`
		} `json:"thinkpad-isa-0000"`
		NVME struct {
			Composite struct {
				Temp float64 `json:"temp1_input"`
			} `json:"Composite"`
		} `json:"nvme-pci-0400"`
	}{}
	if err := json.Unmarshal(out, &sensors); err != nil {
		return "", fmt.Errorf("failed to json unmarshal: %v", err)
	}

	return fmt.Sprintf(
		" %.0f℃  %.0f℃",
		sensors.Thinkpad.CPU.Temp, sensors.NVME.Composite.Temp,
	), nil
}
