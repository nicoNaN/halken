// Contains instruction opcodes & definitions
// Reference: http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
package cpu

// Holds all relevant information for CPU instructions
// Don't think we need to store opcode in the struct since
// it will be equal to the byte key of the instruction in the map
type Instruction struct {
	Mnemonic    string
	// Number of T cycles instruction takes to execute
	// Divide by 4 to get number of M cycles
	tCycles     int
	// Currently this isn't actually used - can't think of a way to reference
	// this value by the time we are inside of an executor function
	NumOperands int
	Executor    func() // Executes appropriate function
}

// Non-CB prefixed instructions
// Parentheses indicate an address
// i8 is 8-bit immediate, i16 is 16-bit immediate
// a16 is a 16-bit address, a8 is an 8-bit address added to $FF00
// s8 is 8-bit signed data, added to PC to move it
func (gbcpu *GBCPU) loadInstructions() {
	gbcpu.Instrs = map[byte]Instruction{
		0x00: Instruction{"NOP", 4, 0, func() {}},
		0x01: Instruction{"LD BC,i16", 12, 2, func() { gbcpu.LDrr_nn(&gbcpu.Regs.b, &gbcpu.Regs.c) }},
		0x02: Instruction{"LD (BC),A", 8, 0, func() { gbcpu.LDrr_r(&gbcpu.Regs.b, &gbcpu.Regs.c, &gbcpu.Regs.a) }},
		0x03: Instruction{"INC BC", 8, 0, func() { gbcpu.INCrr(&gbcpu.Regs.b, &gbcpu.Regs.c) }},
		0x04: Instruction{"INC B", 4, 0, func() { gbcpu.INCr(&gbcpu.Regs.b) }},
		0x05: Instruction{"DEC B", 4, 0, func() { gbcpu.DECr(&gbcpu.Regs.b) }},
		0x06: Instruction{"LD B,i8", 8, 1, func() { gbcpu.LDr_n(&gbcpu.Regs.b) }},
		0x07: Instruction{"RLCA", 4, 0, func() { gbcpu.RLCA(&gbcpu.Regs.a) }},
		0x08: Instruction{"LD (a16),SP", 20, 2, func() { gbcpu.LDnn_SP() }},
		0x09: Instruction{"ADD HL,BC", 8, 0, func() { gbcpu.ADDrr_rr(&gbcpu.Regs.h, &gbcpu.Regs.l, &gbcpu.Regs.b, &gbcpu.Regs.c) }},
		0x0A: Instruction{"LD A,(BC)", 8, 0, func() { gbcpu.LDr_rr(&gbcpu.Regs.a, &gbcpu.Regs.b, &gbcpu.Regs.c) }},
		0x0B: Instruction{"DEC BC", 8, 0, func() { gbcpu.DECrr(&gbcpu.Regs.b, &gbcpu.Regs.c) }},
		0x0C: Instruction{"INC C", 4, 0, func() { gbcpu.INCr(&gbcpu.Regs.c) }},
		0x0D: Instruction{"DEC C", 4, 0, func() {}},
		0x0E: Instruction{"LD C,i8", 8, 1, func() {}},
		0x0F: Instruction{"RRCA", 4, 0, func() {}},
		0x10: Instruction{"STOP", 4, 1, func() {}},
		0x11: Instruction{"LD DE,i16", 12, 2, func() {}},
		0x12: Instruction{"LD (DE),A", 8, 0, func() {}},
		0x13: Instruction{"INC DE", 4, 0, func() {}},
		0x14: Instruction{"INC D", 4, 0, func() {}},
		0x15: Instruction{"DEC D", 4, 0, func() {}},
		0x16: Instruction{"LD D,i8", 8, 1, func() {}},
		0x17: Instruction{"RLA", 4, 0, func() {}},
		0x18: Instruction{"JR s8", 12, 1, func() {}},
		0x19: Instruction{"ADD HL,DE", 8, 0, func() {}},
		0x1A: Instruction{"LD A,(DE)", 8, 0, func() {}},
		0x1B: Instruction{"DEC DE", 8, 0, func() {}},
		0x1C: Instruction{"INC E", 4, 0, func() {}},
		0x1D: Instruction{"DEC E", 4, 0, func() {}},
		0x1E: Instruction{"LD E,i8", 8, 1, func() {}},
		0x1F: Instruction{"RRA", 4, 0, func() {}},
		0x20: Instruction{"JR NZ,s8", 8, 0, func() {}},
		0x21: Instruction{"LD HL,i16", 12, 2, func() {}},
		0x22: Instruction{"LD (HL+),A", 8, 0, func() {}},
		0x23: Instruction{"INC HL", 8, 0, func() {}},
		0x24: Instruction{"INC H", 4, 0, func() {}},
		0x25: Instruction{"DEC H", 4, 0, func() {}},
		0x26: Instruction{"LD H,i8", 8, 1, func() {}},
		0x27: Instruction{"DAA", 4, 0, func() {}},
		0x28: Instruction{"JR Z,s8", 8, 1, func() {}},
		0x29: Instruction{"ADD HL,HL", 8, 0, func() {}},
		0x2A: Instruction{"LD A,(HL+)", 8, 0, func() {}},
		0x2B: Instruction{"DEC HL", 8, 0, func() {}},
		0x2C: Instruction{"INC L", 4, 0, func() {}},
		0x2D: Instruction{"DEC L", 4, 0, func() {}},
		0x2E: Instruction{"LD L,i8", 8, 1, func() {}},
		0x2F: Instruction{"CPL", 4, 0, func() {}},
		0x30: Instruction{"JR NC,s8", 8, 1, func() {}},
		0x31: Instruction{"LD SP,i16", 12, 2, func() {}},
		0x32: Instruction{"LD (HL-),A", 8, 0, func() {}},
		0x33: Instruction{"INC SP", 8, 0, func() {}},
		0x34: Instruction{"INC (HL)", 12, 0, func() {}},
		0x35: Instruction{"DEC (HL)", 12, 0, func() {}},
		0x36: Instruction{"LD (HL),i8", 12, 1, func() {}},
		0x37: Instruction{"SCF", 4, 0, func() {}},
		0x38: Instruction{"JR C,s8", 8, 1, func() {}},
		0x39: Instruction{"ADD HL,SP", 8, 0, func() {}},
		0x3A: Instruction{"LD A,(HL-)", 8, 0, func() {}},
		0x3B: Instruction{"DEC SP", 8, 0, func() {}},
		0x3C: Instruction{"INC A", 4, 0, func() {}},
		0x3D: Instruction{"DEC A", 4, 0, func() {}},
		0x3E: Instruction{"LD A,i8", 8, 1, func() {}},
		0x3F: Instruction{"CCF", 4, 0, func() {}},
		0x40: Instruction{"LD B,B", 4, 0, func() {}},
		0x41: Instruction{"LD B,C", 4, 0, func() {}},
		0x42: Instruction{"LD B,D", 4, 0, func() {}},
		0x43: Instruction{"LD B,E", 4, 0, func() {}},
		0x44: Instruction{"LD B,H", 4, 0, func() {}},
		0x45: Instruction{"LD B,L", 4, 0, func() {}},
		0x46: Instruction{"LD B,(HL)", 8, 0, func() {}},
		0x47: Instruction{"LD B,A", 4, 0, func() {}},
		0x48: Instruction{"LD C,B", 4, 0, func() {}},
		0x49: Instruction{"LD C,C", 4, 0, func() {}},
		0x4A: Instruction{"LD C,D", 4, 0, func() {}},
		0x4B: Instruction{"LD C,E", 4, 0, func() {}},
		0x4C: Instruction{"LD C,H", 4, 0, func() {}},
		0x4D: Instruction{"LD C,L", 4, 0, func() {}},
		0x4E: Instruction{"LD C,(HL)", 8, 0, func() {}},
		0x4F: Instruction{"LD C,A", 4, 0, func() {}},
		0x50: Instruction{"LD D,B", 4, 0, func() {}},
		0x51: Instruction{"LD D,C", 4, 0, func() {}},
		0x52: Instruction{"LD D,D", 4, 0, func() {}},
		0x53: Instruction{"LD D,E", 4, 0, func() {}},
		0x54: Instruction{"LD D,H", 4, 0, func() {}},
		0x55: Instruction{"LD D,L", 4, 0, func() {}},
		0x56: Instruction{"LD D,(HL)", 8, 0, func() {}},
		0x57: Instruction{"LD D,A", 4, 0, func() {}},
		0x58: Instruction{"LD E,B", 4, 0, func() {}},
		0x59: Instruction{"LD E,C", 4, 0, func() {}},
		0x5A: Instruction{"LD E,D", 4, 0, func() {}},
		0x5B: Instruction{"LD E,E", 4, 0, func() {}},
		0x5C: Instruction{"LD E,H", 4, 0, func() {}},
		0x5D: Instruction{"LD E,L", 4, 0, func() {}},
		0x5E: Instruction{"LD E,(HL)", 8, 0, func() {}},
		0x5F: Instruction{"LD E,A", 4, 0, func() {}},
		0x60: Instruction{"LD H,B", 4, 0, func() {}},
		0x61: Instruction{"LD H,C", 4, 0, func() {}},
		0x62: Instruction{"LD H,D", 4, 0, func() {}},
		0x63: Instruction{"LD H,E", 4, 0, func() {}},
		0x64: Instruction{"LD H,H", 4, 0, func() {}},
		0x65: Instruction{"LD H,L", 4, 0, func() {}},
		0x66: Instruction{"LD H,(HL)", 8, 0, func() {}},
		0x67: Instruction{"LD H,A", 4, 0, func() {}},
		0x68: Instruction{"LD L,B", 4, 0, func() {}},
		0x69: Instruction{"LD L,C", 4, 0, func() {}},
		0x6A: Instruction{"LD L,D", 4, 0, func() {}},
		0x6B: Instruction{"LD L,E", 4, 0, func() {}},
		0x6C: Instruction{"LD L,H", 4, 0, func() {}},
		0x6D: Instruction{"LD L,L", 4, 0, func() {}},
		0x6E: Instruction{"LD L,(HL)", 8, 0, func() {}},
		0x6F: Instruction{"LD L,A", 4, 0, func() {}},
		0x70: Instruction{"LD (HL),B", 8, 0, func() {}},
		0x71: Instruction{"LD (HL),C", 8, 0, func() {}},
		0x72: Instruction{"LD (HL),D", 8, 0, func() {}},
		0x73: Instruction{"LD (HL),E", 8, 0, func() {}},
		0x74: Instruction{"LD (HL),H", 8, 0, func() {}},
		0x75: Instruction{"LD (HL),L", 8, 0, func() {}},
		0x76: Instruction{"HALT", 4, 0, func() {}},
		0x77: Instruction{"LD (HL),A", 8, 0, func() {}},
		0x78: Instruction{"LD A,B", 4, 0, func() {}},
		0x79: Instruction{"LD A,C", 4, 0, func() {}},
		0x7A: Instruction{"LD A,D", 4, 0, func() {}},
		0x7B: Instruction{"LD A,E", 4, 0, func() {}},
		0x7C: Instruction{"LD A,H", 4, 0, func() {}},
		0x7D: Instruction{"LD A,L", 4, 0, func() {}},
		0x7E: Instruction{"LD A,(HL)", 8, 0, func() {}},
		0x7F: Instruction{"LD A,A", 4, 0, func() {}},
		0x80: Instruction{"ADD A,B", 4, 0, func() {}},
		0x81: Instruction{"ADD A,C", 4, 0, func() {}},
		0x82: Instruction{"ADD A,D", 4, 0, func() {}},
		0x83: Instruction{"ADD A,E", 4, 0, func() {}},
		0x84: Instruction{"ADD A,H", 4, 0, func() {}},
		0x85: Instruction{"ADD A,L", 4, 0, func() {}},
		0x86: Instruction{"ADD A,(HL)", 8, 0, func() {}},
		0x87: Instruction{"ADD A,A", 4, 0, func() {}},
		0x88: Instruction{"ADC A,B", 4, 0, func() {}},
		0x89: Instruction{"ADC A,C", 4, 0, func() {}},
		0x8A: Instruction{"ADC A,D", 4, 0, func() {}},
		0x8B: Instruction{"ADC A,E", 4, 0, func() {}},
		0x8C: Instruction{"ADC A,H", 4, 0, func() {}},
		0x8D: Instruction{"ADC A,L", 4, 0, func() {}},
		0x8E: Instruction{"ADC A,(HL)", 8, 0, func() {}},
		0x8F: Instruction{"ADC A,A", 4, 0, func() {}},
		0x90: Instruction{"SUB B", 4, 0, func() {}},
		0x91: Instruction{"SUB C", 4, 0, func() {}},
		0x92: Instruction{"SUB D", 4, 0, func() {}},
		0x93: Instruction{"SUB E", 4, 0, func() {}},
		0x94: Instruction{"SUB H", 4, 0, func() {}},
		0x95: Instruction{"SUB L", 4, 0, func() {}},
		0x96: Instruction{"SUB (HL)", 8, 0, func() {}},
		0x97: Instruction{"SUB A", 4, 0, func() {}},
		0x98: Instruction{"SBC A,B", 4, 0, func() {}},
		0x99: Instruction{"SBC A,C", 4, 0, func() {}},
		0x9A: Instruction{"SBC A,D", 4, 0, func() {}},
		0x9B: Instruction{"SBC A,E", 4, 0, func() {}},
		0x9C: Instruction{"SBC A,H", 4, 0, func() {}},
		0x9D: Instruction{"SBC A,L", 4, 0, func() {}},
		0x9E: Instruction{"SBC A,(HL)", 8, 0, func() {}},
		0x9F: Instruction{"SBC A,A", 4, 0, func() {}},
		0xA0: Instruction{"AND B", 4, 0, func() {}},
		0xA1: Instruction{"AND C", 4, 0, func() {}},
		0xA2: Instruction{"AND D", 4, 0, func() {}},
		0xA3: Instruction{"AND E", 4, 0, func() {}},
		0xA4: Instruction{"AND H", 4, 0, func() {}},
		0xA5: Instruction{"AND L", 4, 0, func() {}},
		0xA6: Instruction{"AND (HL)", 8, 0, func() {}},
		0xA7: Instruction{"AND A", 4, 0, func() {}},
		0xA8: Instruction{"XOR B", 4, 0, func() {}},
		0xA9: Instruction{"XOR C", 4, 0, func() {}},
		0xAA: Instruction{"XOR D", 4, 0, func() {}},
		0xAB: Instruction{"XOR E", 4, 0, func() {}},
		0xAC: Instruction{"XOR H", 4, 0, func() {}},
		0xAD: Instruction{"XOR L", 4, 0, func() {}},
		0xAE: Instruction{"XOR (HL)", 8, 0, func() {}},
		0xAF: Instruction{"XOR A", 4, 0, func() {}},
		0xB0: Instruction{"OR B", 4, 0, func() {}},
		0xB1: Instruction{"OR C", 4, 0, func() {}},
		0xB2: Instruction{"OR D", 4, 0, func() {}},
		0xB3: Instruction{"OR E", 4, 0, func() {}},
		0xB4: Instruction{"OR H", 4, 0, func() {}},
		0xB5: Instruction{"OR L", 4, 0, func() {}},
		0xB6: Instruction{"OR (HL)", 8, 0, func() {}},
		0xB7: Instruction{"OR A", 4, 0, func() {}},
		0xB8: Instruction{"CP B", 4, 0, func() {}},
		0xB9: Instruction{"CP C", 4, 0, func() {}},
		0xBA: Instruction{"CP D", 4, 0, func() {}},
		0xBB: Instruction{"CP E", 4, 0, func() {}},
		0xBC: Instruction{"CP H", 4, 0, func() {}},
		0xBD: Instruction{"CP L", 4, 0, func() {}},
		0xBE: Instruction{"CP (HL)", 8, 0, func() {}},
		0xBF: Instruction{"CP A", 4, 0, func() {}},
		0xC0: Instruction{"RET NZ", 8, 0, func() {}},
		0xC1: Instruction{"POP BC", 12, 0, func() {}},
		0xC2: Instruction{"JP NZ,a16", 16, 2, func() {}},
		0xC3: Instruction{"JP a16", 16, 2, func() { gbcpu.JPaa() }},
		0xC4: Instruction{"CALL NZ,a16", 12, 2, func() {}},
		0xC5: Instruction{"PUSH BC", 16, 0, func() {}},
		0xC6: Instruction{"ADD A,i8", 8, 1, func() {}},
		0xC7: Instruction{"RST 00H", 16, 0, func() {}},
		0xC8: Instruction{"RET Z", 8, 0, func() {}},
		0xC9: Instruction{"RET", 16, 0, func() {}},
		0xCA: Instruction{"JP Z,a16", 12, 2, func() {}},
		0xCB: Instruction{"PREFIX CB", 4, 0, func() {}},
		0xCC: Instruction{"CALL Z,a16", 12, 2, func() {}},
		0xCD: Instruction{"CALL a16", 24, 2, func() {}},
		0xCE: Instruction{"ADC A,i8", 8, 1, func() {}},
		0xCF: Instruction{"RST 08H", 16, 0, func() {}},
		0xD0: Instruction{"RET NC", 8, 0, func() {}},
		0xD1: Instruction{"POP DE", 12, 0, func() {}},
		0xD2: Instruction{"JP NC,a16", 12, 2, func() {}},
		// 0xD3: no corresponding instruction
		0xD4: Instruction{"CALL NC,a16", 12, 2, func() {}},
		0xD5: Instruction{"PUSH DE", 16, 0, func() {}},
		0xD6: Instruction{"SUB i8", 8, 1, func() {}},
		0xD7: Instruction{"RST 10H", 16, 0, func() {}},
		0xD8: Instruction{"RET C", 8, 0, func() {}},
		0xD9: Instruction{"RETI", 16, 0, func() {}},
		0xDA: Instruction{"JP C,a16", 12, 2, func() {}},
		// 0xDB: no corresponding instruction
		0xDC: Instruction{"CALL C,a16", 12, 2, func() {}},
		// 0xDD: no corresponding instruction
		0xDE: Instruction{"SBC A,i8", 8, 1, func() {}},
		0xDF: Instruction{"RST 18H", 16, 0, func() {}},
		0xE0: Instruction{"LDH ($FF00+a8),A", 12, 1, func() {}},
		0xE1: Instruction{"POP HL", 12, 0, func() {}},
		0xE2: Instruction{"LD ($FF00+C),A", 8, 0, func() {}},
		// 0xE3: no corresponding instruction
		// 0xE4: no corresponding instruction
		0xE5: Instruction{"PUSH HL", 16, 0, func() {}},
		0xE6: Instruction{"AND i8", 8, 1, func() {}},
		0xE7: Instruction{"RST 20H", 16, 0, func() {}},
		0xE8: Instruction{"ADD SP,s8", 16, 1, func() {}},
		0xE9: Instruction{"JP (HL)", 4, 0, func() {}},
		0xEA: Instruction{"LD (a16),A", 16, 2, func() {}},
		// 0xEB: no corresponding instruction
		// 0xEC: no corresponding instruction
		// 0xED: no corresponding instruction
		0xEE: Instruction{"XOR i8", 8, 1, func() {}},
		0xEF: Instruction{"RST 28H", 16, 0, func() {}},
		0xF0: Instruction{"LDH A,($FF00+a8)", 12, 1, func() {}},
		0xF1: Instruction{"POP AF", 12, 0, func() {}},
		0xF2: Instruction{"LD A,($FF00+C)", 8, 0, func() {}},
		0xF3: Instruction{"DI", 4, 0, func() {}},
		// 0xF4: no corresponding instruction
		0xF5: Instruction{"PUSH AF", 16, 0, func() {}},
		0xF6: Instruction{"OR i8", 8, 1, func() {}},
		0xF7: Instruction{"RST 30H", 16, 0, func() {}},
		0xF8: Instruction{"LD HL,SP+s8", 12, 1, func() {}},
		0xF9: Instruction{"LD SP,HL", 8, 0, func() {}},
		0xFA: Instruction{"LD A,(a16)", 16, 2, func() {}},
		0xFB: Instruction{"EI", 4, 0, func() {}},
		// 0xFC: no corresponding instruction
		// 0xFD: no corresponding instruction
		0xFE: Instruction{"CP i8", 8, 1, func() {}},
		0xFF: Instruction{"RST 38H", 16, 0, func() {}}}
}

// CB prefixed instructions
// CB is the prefix byte. Like the Z80, the Sharp LR35902 will
// look up a CB prefixed instruction in a different instruction bank
// More info: http://www.z80.info/decoding.htm
var instructionsCB map[byte]Instruction = map[byte]Instruction{
	0x00: Instruction{"RLC B", 8, 0, func() {}},
}
