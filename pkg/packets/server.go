package packets

/*
	First two byte is packet size?
	Third byte is the OPCODE
 */

var SERVER_LIST = []byte{
	0x06, 0x00, 0x03, 0x00, 0x01, 0x00,


	0x30, 0x00, 0x18, 0x27,
	0x54, 0x65, 0x6D, 0x70,
	0x65, 0x73, 0x74, // server name

	//0x00, 0x14, 0x48, 0x3A,
	//0x43, 0x01, 0x00, 0x00,
	//0x00, 0x78, 0x47, 0x3A,
	//0x43, 0x94, 0x47, 0x3A,
	//0x43, 0xC1, 0x5A, 0x41,
	//0x00, 0x94, 0x47, 0x3A,
	//0x43, 0x15, 0x05, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00,
	0x01, // ID
	0x01, // Recommended
	0x26,
	0x00, 0x0F, 0x01, 0x00,
	0x00,

	0x30, 0x00, 0x18, 0x27, 0x4C, 0x65, 0x70, 0x75, 0x73, 0x00,
	0x00, 0x00, 0x65, 0x00, 0x3A, 0x43, 0x01, 0x00, 0x00, 0x00, 0x78, 0x47, 0x3A, 0x43, 0x94, 0x47,
	0x3A, 0x43, 0xC1, 0x5A, 0x41, 0x00, 0x94, 0x47, 0x3A, 0x43, 0x15, 0x05, 0x00, 0x00, 0x26, 0x00,
	0x26, 0x00, 0x0F, 0x01, 0x01, 0x00,
}

// Returns Server IP
//var NEW_GAME_SERVER = []byte{
//	0x48, 0x00, 0x11, 0x27, 0x23,
//
//	// IP ADDRESS
//	0x31, 0x32, 0x37, 0x2E,
//	0x30, 0x2E, 0x30, 0x2E,
//	0x31, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//
//	0x00, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//	0x00, 0x00, 0x00, 0x00,
//	//0x00, 0x00, 0x00, 0x68,
//	//0xDE, 0x74, 0x66, 0xC2,
//	//0xE2, 0x42, 0x00, 0xD0,
//	//0xDE, 0x74, 0x66, 0x58,
//	//0x1B, 0x00, 0x00, 0x00,
//
//	// UNK
//	0x31, 0x39, 0x39, 0x35,
//	0x36, 0x66, 0x35, 0x32,
//	0x33, 0x32, 0x62, 0x66,
//	0x32, 0x64, 0x33, 0x38,
//	0x37, 0x37, 0x37, 0x61,
//	0x31, 0x33, 0x65, 0x31,
//	0x00,
//
//	// END
//	0x74, 0x66,
//}

var SEND_SERVER = []byte{
	0x54, 0x00, // Packet Length?
	0x07, // Packet ID?
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00,
	0x54, 0x65, 0x6D, 0x70, 0x65, 0x73, 0x74, // Server Name
	0x00, 0xC4, 0xD1, 0x40, 0x00,
	0xB0, 0xD9, 0x85, 0x65, 0x80, 0x5B, 0xF4, 0x01, 0x88, 0x5B, 0xF4, 0x01, 0x88, 0xD9, 0x85, 0x65,
	0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xBF, 0x40, 0x00, 0x0B, 0x00, 0x00, 0x00,
	0x02, 0x00, 0x01, 0x01,
}

// GAME SERVER
var NEW_GAME_SERVER = []byte{
	0x48, 0x00, 0x11, 0x27, 0x23,

	// IP ADDRESS
	//0x31, 0x32, 0x37, 0x2E,
	//0x30, 0x2E, 0x30, 0x2E,
	//0x31, 0x00, 0x00, 0x00,
	//0x00, 0x00, 0x00, 0x00,

	//52.221.199.168
	0x35, 0x32, 0x2e, 0x32,
	0x32, 0x31, 0x2e, 0x31,
	0x39, 0x39, 0x2e, 0x31,
	0x36, 0x38, 0x00, 0x00,

	//0x31, 0x33, 0x2E, 0x32,
	//0x33, 0x30, 0x2E, 0x32,
	//0x32, 0x34, 0x2E, 0x31,
	//0x36, 0x36, 0x00, 0x00,

	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0xDE, 0x74, 0x66, //  81......  ....h.tf
	0xC2, 0xE2, 0x42, 0x00, 0xD0, 0xDE, 0x74, 0x66,  0x58, 0x1B, 0x00, 0x00, 0x00,

	0x4B, 0x41, 0x4C, 0x33,
	0x6A, 0x63, 0x49, 0x7A,
	0x71, 0x47, 0x67, 0x4B,
	0x76, 0x4F, 0x66, 0x31,
	0x64, 0x62, 0x59, 0x5A,
	0x4B, 0x43, 0x38, 0x63,
	0x53,

	0x74, 0x66,
}

var NEW_GAME_SERVER_COPY = []byte{
	0x48, 0x00, 0x11, 0x27, 0x23,

	0x31, 0x32, 0x37, 0x2E,
	0x30, 0x2E, 0x30, 0x2E,
	0x31, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0xDE, 0x74, 0x66, //  81......  ....h.tf
	0xC2, 0xE2, 0x42, 0x00, 0xD0, 0xDE, 0x74, 0x66,  0x58, 0x1B, 0x00, 0x00, 0x00,

	0x31, 0x39, 0x39, 0x35,
	0x36, 0x66, 0x35, 0x32,
	0x33, 0x32, 0x62, 0x66,
	0x32, 0x64, 0x33, 0x38,
	0x37, 0x37, 0x37, 0x61,
	0x31, 0x33, 0x65, 0x31,
	0x00,

	0x74, 0x66,
}