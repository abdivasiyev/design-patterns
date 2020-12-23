package builder

// PersonalComputer product to build
type PersonalComputer struct {
	ramCap int
	hddCap int
	cpu    string
	os     string
	gpu    string
}

// PCBuilder is used by each builders
type PCBuilder interface {
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetGPU() PCBuilder
	GetPC() PersonalComputer
}

// HomeEditionPCBuilder ...
type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

// SetRAM ...
func (b *HomeEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 4
	return b
}

// SetHDD ...
func (b *HomeEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

// SetCPU ...
func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i3"
	return b
}

// SetOS ...
func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Home Edition"
	return b
}

// SetGPU ...
func (b *HomeEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "Intel Graphic Card"
	return b
}

// GetPC ...
func (b *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}

// GameEditionPCBuilder ...
type GameEditionPCBuilder struct {
	pc PersonalComputer
}

// SetRAM ...
func (b *GameEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 16
	return b
}

// SetHDD ...
func (b *GameEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

// SetCPU ...
func (b *GameEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i7"
	return b
}

// SetOS ...
func (b *GameEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Ultimate"
	return b
}

// SetGPU ...
func (b *GameEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "AMD Radeon X80"
	return b
}

// GetPC ...
func (b *GameEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}

// Manufacturer ...
type Manufacturer struct {
	builder PCBuilder
}

// SetBuilder ...
func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.builder = builder
}

// ConstructPC ...
func (m *Manufacturer) ConstructPC() {
	m.builder.SetCPU().SetGPU().SetHDD().SetRAM().SetOS()
}
