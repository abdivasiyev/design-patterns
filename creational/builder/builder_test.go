package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManufacturer(t *testing.T) {
	manufacturer := &Manufacturer{}

	homePCBuilder := &HomeEditionPCBuilder{}
	manufacturer.SetBuilder(homePCBuilder)
	manufacturer.ConstructPC()

	homePC := homePCBuilder.GetPC()

	assert.Equal(t, homePC.cpu, "i3")
	assert.Equal(t, homePC.gpu, "Intel Graphic Card")

	gamePCBuilder := &GameEditionPCBuilder{}
	manufacturer.SetBuilder(gamePCBuilder)
	manufacturer.ConstructPC()

	gamePC := gamePCBuilder.GetPC()

	assert.Equal(t, gamePC.cpu, "i7")
	assert.Equal(t, gamePC.gpu, "AMD Radeon X80")
}
