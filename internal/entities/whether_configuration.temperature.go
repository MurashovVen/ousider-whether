package entities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MurashovVen/outsider-sdk/entities"
)

type WhetherTemperatureConfiguration struct {
	Temperature int
}

func NewWhetherTemperatureConfiguration(temperature int) *WhetherTemperatureConfiguration {
	return &WhetherTemperatureConfiguration{
		Temperature: temperature,
	}
}

func WhetherTemperatureConfigurationParseString(wtc string) (*WhetherTemperatureConfiguration, error) {
	sl := strings.Split(wtc, splitter)
	if len(sl) != 2 {
		return nil, ErrInvalidWhetherTemperatureConfiguration
	}

	temperature, err := strconv.Atoi(sl[1])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidWhetherTemperatureConfiguration, err)
	}

	return NewWhetherTemperatureConfiguration(temperature), nil
}

const splitter = ":"

func (c *WhetherTemperatureConfiguration) String() string {
	return entities.ActionWhetherTemperatureConfigure + splitter + strconv.Itoa(c.Temperature)
}
