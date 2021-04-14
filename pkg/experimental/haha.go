package experimental

var ENTER_PART1 = []byte{
	0x58, 0x06, 0x23, 0x27, 0xD2, 0xE4, 0xEB, 0x0B,  0x73, 0x75, 0x73, 0x68, 0x31, 0x00, 0x18, 0x00, //X.#'....  sush1...
	0x8D, 0x56, 0x47, 0x00, 0x8C, 0xF8, 0x18, 0x00,  0x30, 0xB9, 0xE2, 0x40, 0x94, 0xB6, 0xE2, 0x40, //.VG.....  0..@...@
	0x98, 0xF8, 0x18, 0x00, 0x91, 0x01, 0x00, 0x00,  0x01, 0x01, 0x01, 0x01, 0x35, 0x00, 0x01, 0x00, //........  ....5...

	0x48, 0x14, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00,  0x00, 0x00, 0x28, 0x43, 0x00, 0x00, 0x00, 0x00, //H...`...  ..(C....
	0x00, 0x00, 0xC8, 0xC2, 0x46, 0x05, 0x00, 0x00,  0xBE, 0x00, 0x00, 0x00, 0x46, 0x05, 0x00, 0x00, //....F...  ....F...
	0xBE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0xC8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........

	0x00, 0x00, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x00,  0x34, 0x08, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, //....<...  4.......
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //........  ........
	0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x13, 0x4F, 0xA7, 0x08, 0x42, 0x00, 0x00, 0x00, //........  .O..B...
	0x54, 0x0B, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //T.......  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
													// WEAPON
	0x14, 0x4F, 0xA7, 0x08, 0x42, 0x00, 0x00, 0x00,  0x78, 0x05, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, //.O..B...  x.......
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //........  ........
	0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x15, 0x4F, 0xA7, 0x08, 0x42, 0x00, 0x00, 0x00, //........  .O..B...
	// COSTUME
	0x68, 0x1F, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //h.......  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //........  ........
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var ENTER_PART2_UNK = []byte{
													 0xA4, 0xEF, 0xFA, 0x0D, 0xF0, 0x25, 0x10, 0x63, //........  .....%.c
	0xBB, 0xAB, 0x46, 0x00, 0x30, 0xB9, 0xE2, 0x40,  0x40, 0xD7, 0x09, 0x13, 0xF0, 0x3C, 0x7E, 0x1A, //..F.0..@  @....<~.
	0x20, 0xD7, 0x09, 0x13, 0x91, 0x01, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x18, 0x66, 0xF4, 0x0F, 0x5C, 0x14, 0x56, 0x00,  0x00, 0x00, 0x80, 0x3F, 0x00, 0x00, 0x00, 0x00, // .f..\.V.  ...?....
	0x91, 0x01, 0x00, 0x00, 0x65, 0x14, 0x56, 0x00,  0x03, 0x00, 0x00, 0x00, 0xF8, 0x3A, 0x45, 0x63, // ....e.V.  .....:Ec

	0x94, 0xB6, 0xE2, 0x40, 0x00, 0x00, 0x00, 0x00,  0x18, 0x66, 0xF4, 0x0F, 0x4C, 0xD7, 0x09, 0x13, // ...@....  .f..L...
	0x40, 0xD7, 0x09, 0x13, 0xF0, 0x37, 0x7E, 0x1A,  0x5D, 0x74, 0x39, 0x00, 0xD7, 0x5B, 0x47, 0x00, // @....7~.  ]t9..[G.
	0x08, 0x00, 0x00, 0x00, 0x6E, 0xF9, 0x18, 0x00,  0xB2, 0xF9, 0x18, 0x00, 0xF4, 0x84, 0x5B, 0x00, // ....n...  ......[.
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0xE8, 0x03, 0x66, 0x00, // ........  ......f.
	0xC0, 0x05, 0x66, 0x00, 0x8C, 0xF9, 0x18, 0x00,  0x05, 0x00, 0x00, 0x00, 0x46, 0xF9, 0x18, 0x00, // ..f.....  ....F...
	0xC6, 0x05, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..f.....  ........
	0x00, 0x00, 0x00, 0x00, 0x90, 0xF9, 0x18, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0xCC, 0xCC, 0xCC, 0xCC,  0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xCC, 0xFB, 0x3F, // ........  .......?
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x6E, 0xF9, 0x18, 0x00, 0xF0, 0xF9, 0x18, 0x00,  0xE0, 0xF9, 0x18, 0x00, 0xF2, 0x71, 0x5B, 0x00, // n.......  .....q[.
	0xF0, 0xF9, 0x18, 0x00, 0x16, 0x00, 0x00, 0x00,  0x90, 0xF9, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x88, 0xAE, 0xFE, 0x3F, 0x18, 0x00,  0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // .....?..  ........
	0xC6, 0xFA, 0x18, 0x00, 0x01, 0x00, 0x00, 0x00,  0xBC, 0xF9, 0x18, 0x00, 0x82, 0x46, 0x5B, 0x00, // ........  .....F[.
	0xC6, 0xFA, 0x18, 0x00, 0xC5, 0xFA, 0x18, 0x00,  0x07, 0x00, 0x00, 0x00, 0xC5, 0xFA, 0x18, 0x00, // ........  ........
	0xC5, 0xFA, 0x18, 0x00, 0x2F, 0x4D, 0x5B, 0x00,  0xC8, 0xF9, 0x18, 0x00, 0xCA, 0xB4, 0x47, 0x00, // ..../M[.  ......G.
	0x48, 0xFC, 0x18, 0x00, 0xA0, 0x96, 0xFA, 0x17,  0xA0, 0x96, 0xFA, 0x17, 0x29, 0xA5, 0x48, 0x00, // H.......  ....).H.
	0xEC, 0x47, 0x76, 0x20, 0x54, 0xCB, 0xFA, 0x17,  0x78, 0xA5, 0x48, 0x00, 0x00, 0x02, 0x00, 0x00, // .Gv.T...  x.H.....
	0x68, 0x1F, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // h.......  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x35, 0x00, 0x29, 0x77,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ....5.)w  ........
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x93, 0x4F, 0x5B, 0x00, 0x58, 0xFA, 0x18, 0x00, // ........  .O[.X...
	0xC4, 0xFA, 0x18, 0x00, 0x64, 0xFA, 0x18, 0x00,  0xDE, 0xBA, 0x59, 0x00, 0x54, 0xFA, 0x18, 0x00, // ....d...  ..Y.T...
	0x64, 0xFA, 0x18, 0x00, 0x8A, 0xBC, 0x59, 0x00,  0xB7, 0x00, 0x00, 0x00, 0x7C, 0xFA, 0x18, 0x00, // d.....Y.  ....|...
	0xDE, 0xBA, 0x59, 0x00, 0x6C, 0xFA, 0x18, 0x00,  0x7C, 0xFA, 0x18, 0x00, 0x8A, 0xBC, 0x59, 0x00, // ..Y.l...  |.....Y.
	0xB7, 0x00, 0x00, 0x00, 0xBF, 0x1F, 0x63, 0x00,  0x6C, 0xFA, 0x18, 0x00, 0x95, 0xBC, 0x59, 0x00, // ......c.  l.....Y.
	0x6C, 0xFA, 0x18, 0x00, 0x48, 0x17, 0x59, 0x00,  0x24, 0xFD, 0x18, 0x00, 0x3A, 0x36, 0x5A, 0x00, // l...H.Y.  $...:6Z.
	0x00, 0x00, 0x00, 0x00, 0xB0, 0xEE, 0x51, 0x61,  0x08, 0x14, 0x3F, 0x00, 0xA8, 0x05, 0x3F, 0x00, // ......Qa  ..?...?.
	0x00, 0x00, 0x00, 0x00, 0x10, 0xFB, 0x18, 0x00,  0x02, 0x00, 0x00, 0x00, 0x20, 0xFD, 0x18, 0x00, // ........  ........
	0x25, 0xDF, 0x59, 0x00, 0xF2, 0xE9, 0x59, 0x00,  0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // %.Y...Y.  ........
	0xF0, 0xBE, 0xC7, 0x1A, 0x9E, 0xEB, 0x59, 0x00,  0x14, 0xFD, 0x18, 0x00, 0xB0, 0xEE, 0x51, 0x61, // ......Y.  ......Qa
	0x08, 0x14, 0x3F, 0x00, 0xA8, 0x05, 0x3F, 0x00,  0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..?...?.  ........
	0x18, 0xFD, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x3F, 0x00, 0x5E, 0x01, 0x00, 0x00, // ........  ..?.^...
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0xC1, 0x1F, 0x63, 0x00, 0x06, 0x00, 0x00, 0x00, // ........  ..c.....
	0x20, 0xFD, 0x18, 0x00, 0x38, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ....8...  ........
	0xC1, 0x1F, 0x63, 0x00, 0x18, 0xFB, 0x18, 0x00,  0x64, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, // ..c.....  d.......
	0x00, 0x00, 0x00, 0x06, 0xB8, 0xFC, 0x18, 0x00,  0x68, 0xCD, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  h.......
	0x0C, 0x0B, 0x2E, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0x20, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, // ........  ........
	0x08, 0x0B, 0x2E, 0x00, 0x00, 0x00, 0x00, 0x00,  0xAC, 0xFB, 0x18, 0x00, 0xA9, 0x0F, 0x29, 0x77, // ........  ......)w
	0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 0.......  0.......
	0xE0, 0x0E, 0x29, 0x77, 0x00, 0x00, 0x00, 0x00,  0x48, 0xF7, 0xF8, 0x62, 0x02, 0x00, 0x00, 0x00, // ..)w....  H..b....
	0x07, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, // ........  ....8...
	0x14, 0xFD, 0x18, 0x00, 0x00, 0x00, 0x00, 0x01,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, // ........  ........
	0x23, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,  0x04, 0x00, 0x00, 0x00, 0xF0, 0x38, 0x7F, 0x1A, // #.......  .....8..
	0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x83, 0x8A, 0x64, 0x60, // ....0...  ......d`
	0x00, 0x00, 0x00, 0x00, 0x20, 0xD7, 0x09, 0x13,  0x18, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x3F, 0x00, 0x24, 0x14, 0x56, 0x00,  0x58, 0x02, 0x00, 0x00, 0x48, 0x24, 0x45, 0x63, // ..?.$.V.  X...H$Ec
	0xEC, 0xF3, 0x9A, 0x1F, 0x30, 0x00, 0x00, 0x00,  0xD0, 0x07, 0x2E, 0x00, 0x18, 0x00, 0x00, 0x00, // ....0...  ........
	0x00, 0x00, 0x00, 0x00, 0xA0, 0x9B, 0xEC, 0x5E,  0xD4, 0x20, 0xD1, 0x60, 0x2F, 0x00, 0x00, 0x00, // .......^  ...`/...
	0x00, 0x00, 0x2E, 0x00, 0x00, 0x00, 0x2E, 0x00,  0x10, 0x20, 0xD1, 0x60, 0x01, 0x00, 0x00, 0x00, // ........  ...`....
	0xC8, 0x20, 0xD1, 0x60, 0x4D, 0x05, 0x00, 0x00,  0x3C, 0x0A, 0x2E, 0x00, 0xC4, 0x01, 0x2E, 0x00, // ...`M...  <.......
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x38, 0x0A, 0x2E, 0x00, 0x40, 0xD3, 0x59, 0x03, // ........  8...@.Y.
	0x84, 0xFC, 0x18, 0x00, 0xA9, 0x0F, 0x29, 0x77,  0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ......)w  ........
	0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0xE0, 0x0E, 0x29, 0x77, 0x00, 0x00, 0x00, 0x00, // ........  ..)w....
	0xF8, 0x95, 0xD1, 0x60, 0x02, 0x00, 0x00, 0x00,  0x03, 0x00, 0x00, 0x00, 0x56, 0x06, 0x00, 0x00, // ...`....  ....V...
	0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x06, 0xEC, 0x63, 0x60, // ........  ......c`
	0x00, 0x00, 0x00, 0x00, 0x9C, 0xFC, 0x18, 0x00,  0x06, 0xEC, 0x63, 0x60, 0x00, 0x00, 0x00, 0x00, // ........  ..c`....
	0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x80, 0xFC, 0x18, 0x00, 0xC4, 0x0A, 0x59, 0x00, // x.......  ......Y.
	0x79, 0x00, 0x00, 0x00, 0x6C, 0xFC, 0x18, 0x00,  0x00, 0x00, 0x00, 0x00, 0x20, 0xFD, 0x18, 0x00, // y...l...  ........
	0x00, 0x00, 0x00, 0x00, 0x06, 0xEC, 0x63, 0x60,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ......c`  ........
	0x00, 0x00, 0x00, 0x00, 0x50, 0x46, 0x00, 0x00,  0xE4, 0xFC, 0x18, 0x00, 0x10, 0x00, 0x00, 0x00, // ....PF..  ........
	0x64, 0x6C, 0x63, 0x5E, 0xF8, 0x95, 0xD1, 0x60,  0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, // dlc^...`  ........
	0x50, 0x59, 0x43, 0x46, 0x3C, 0xFD, 0x18, 0x00,  0x60, 0xD8, 0x08, 0x01, 0xBC, 0xFC, 0x18, 0x00, // PYCF<...  `.......
	0x37, 0x00, 0x59, 0x00, 0x10, 0x00, 0x00, 0x00,  0x79, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, // 7.Y.....  y.......
	0x59, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,  0xE4, 0xBE, 0x4A, 0x00, 0x10, 0x00, 0x00, 0x00, // Y.......  ..J.....
	0x56, 0x06, 0x00, 0x00, 0xE9, 0x44, 0x59, 0x00,  0x56, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // V....DY.  V.......
	0x50, 0xFE, 0x18, 0x00, 0x08, 0xD8, 0x5B, 0x00,  0xFF, 0xFF, 0xFF, 0xFF, 0x06, 0xB4, 0x4A, 0x00, // P.....[.  ......J.
	0x60, 0xD8, 0x08, 0x01, 0x64, 0xD8, 0x08, 0x01,  0x54, 0xFD, 0x18, 0x00, 0xD8, 0x2E, 0x3F, 0x00, // `...d...  T.....?.
	0x3D, 0x5A, 0x46, 0xD5, 0x1D, 0xF3, 0x01, 0x00,  0x64, 0xD8, 0x08, 0x01, 0x54, 0xFD, 0x18, 0x00, // =ZF.....  d...T...
	0x64, 0xD8, 0x08, 0x01, 0x50, 0x59, 0x43, 0x46,  0x24, 0x99, 0x75, 0x20, 0x08, 0x74, 0x3F, 0x00, // d...PYCF  $.u..t?.
	0x64, 0xD8, 0x08, 0x01, 0x50, 0x59, 0x43, 0x46,  0x42, 0xAD, 0x4A, 0x00, 0x60, 0xD8, 0x08, 0x01, // d...PYCF  B.J.`...
	0x44, 0xFD, 0x18, 0x00, 0x3C, 0xFD, 0x18, 0x00,  0x24, 0x99, 0x75, 0x20, 0x24, 0x99, 0x75, 0x20, // D...<...  $.u.$.u.
	0x08, 0x07, 0x00, 0x00, 0x08, 0x07, 0x00, 0x00,  0x04, 0x01, 0x59, 0x28, 0x30, 0x2C, 0x31, 0x2C, // ........  ..Y(0,1,
	0x32, 0x2C, 0x34, 0x2C, 0x33, 0x38, 0x2C, 0x34,  0x34, 0x2C, 0x32, 0x30, 0x37, 0x00, 0x59, 0x00, // 2,4,38,4  4,207.Y.
	0x58, 0x65, 0x63, 0x5E, 0x04, 0xF6, 0x18, 0x00,  0x88, 0x9F, 0xD1, 0x60, 0x28, 0xF6, 0x18, 0x00, // Xec^....  ...`(...
	0x00, 0xC2, 0x59, 0x00, 0x47, 0x05, 0x00, 0x00,  0x90, 0x9B, 0xEC, 0x5E, 0x64, 0x00, 0x00, 0x00, // ..Y.G...  ...^d...
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x10, 0x20, 0xD1, 0x60, 0x83, 0x00, 0x43, 0x05, // ........  ...`..C.
	0xE7, 0x28, 0x53, 0x00, 0x28, 0x9F, 0xD1, 0x60,  0x28, 0x65, 0x63, 0x5E, 0x04, 0xF6, 0x18, 0x00, // .(S.(..`  (ec^....
	0x74, 0x29, 0x98, 0x5B, 0x08, 0xF6, 0x18, 0x00,  0x84, 0x3F, 0x59, 0x00, 0x00, 0x00, 0x3F, 0x00, // t).[....  .?Y...?.
	0x00, 0x00, 0x00, 0x00, 0xA3, 0x3F, 0x59, 0x00,  0xE0, 0x66, 0x63, 0x5E, 0x04, 0xF6, 0x18, 0x00, // .....?Y.  .fc^....
	0x74, 0x29, 0x98, 0x5B, 0x28, 0xF6, 0x18, 0x00,  0xB4, 0xFD, 0x18, 0x00, 0xDC, 0xF5, 0x18, 0x00, // t).[(...  ........
	0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x2C, 0xF6, 0x18, 0x00, 0x32, 0xB5, 0x5A, 0x00, // ........  ,...2.Z.
	0x80, 0x82, 0x71, 0x19, 0x40, 0xA3, 0x0A, 0x01,  0x00, 0x06, 0x00, 0x00, 0x88, 0x9F, 0xD1, 0x60, // ..q.@...  .......`
	0x58, 0x00, 0x00, 0x00, 0x40, 0xA9, 0x0A, 0x01,  0x80, 0x88, 0x71, 0x19, 0x9C, 0x62, 0xFE, 0x17, // X...@...  ..q..b..
	0x84, 0x62, 0xFE, 0x17, 0x9C, 0x62, 0xFE, 0x17,  0x84, 0x62, 0xFE, 0x17, 0x58, 0x06, 0x00, 0x00, // .b...b..  .b..X...
	0x04, 0x62, 0xFE, 0x17, 0xD8, 0x8C, 0x58, 0x00,  0x9C, 0x62, 0xFE, 0x17, 0xE8, 0x61, 0xFE, 0x17, // .b....X.  .b...a..
	0x04, 0x62, 0xFE, 0x17, 0x58, 0x06, 0x00, 0x00,  0x9C, 0x62, 0xFE, 0x17, 0xE8, 0x61, 0xFE, 0x17, // .b..X...  .b...a..
	0x84, 0x62, 0xFE, 0x17, 0xB4, 0x66, 0x63, 0x5E,  0x24, 0x99, 0x75, 0x20, 0x50, 0xFE, 0x18, 0x00, // .b...fc^  $.u.P...
	0xBB, 0x48, 0x5C, 0x00, 0x84, 0x66, 0x63, 0x5E,  0x00, 0x00, 0x00, 0x00, 0x78, 0x03, 0x31, 0x27, // .H\..fc^  ....x.1'
	0x60, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x58, 0x00, // `.......  ......X.
	0x00, 0x00, 0x00, 0x00,
}

var ENTER_PART2 = []byte{
	// INVETORY START???
	//0x96, 0x11,
	  				  0x03, 0x0A, 0x00, 0x00,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // ........  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x01, 0x01, 0x03, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x16, 0x4F, 0xA7, 0x08,  0x42, 0x00, 0x00, 0x00, 0xAC, 0x11, 0x00, 0x00, // .....O..  B.......
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // ........  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x01, 0x01, 0x03,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x17, 0x4F, 0xA7, 0x08, // ........  .....O..
	0x42, 0x00, 0x00, 0x00, 0xB2, 0x11, 0x00, 0x00,  0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // B.......  ........
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,  0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x01, 0x01, 0x03, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x18, 0x4F, 0xA7, 0x08,  0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // .....O..  B.......
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x78, 0x03, 0x31, 0x27,  0x60, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, // ....x.1'  `.......
	0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x58, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ......X.  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........

	// PART 3
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ....
}


var ENTER_PART4 = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ........  ........
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27, // ........  ....(.H'
	0xCA, 0xB4, 0x47, 0x00, 0x00, 0x00, 0x00, 0x00,  0xA0, 0x96, 0xFA, 0x17, 0xA0, 0x96, 0xFA, 0x17, // ..G.....  ........
	0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,  0x54, 0xCB, 0xFA, 0x17, 0x13, 0x4F, 0xA7, 0x08, // ........  T....O..
	0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27,  0xCA, 0xB4, 0x47, 0x00, 0x00, 0x00, 0x00, 0x00, // B...(.H'  ..G.....
	0xA0, 0x96, 0xFA, 0x17, 0xA0, 0x96, 0xFA, 0x17,  0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, // ........  ........
	0x54, 0xCB, 0xFA, 0x17, 0x14, 0x4F, 0xA7, 0x08,  0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27, // T....O..  B...(.H'
	0xCA, 0xB4, 0x47, 0x00, 0x00, 0x00, 0x00, 0x00,  0xA0, 0x96, 0xFA, 0x17, 0xA0, 0x96, 0xFA, 0x17, // ..G.....  ........
	0x00, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00,  0x54, 0xCB, 0xFA, 0x17, 0x15, 0x4F, 0xA7, 0x08, // ........  T....O..
	0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27,  0xCA, 0xB4, 0x47, 0x00, 0x00, 0x00, 0x00, 0x00, // B...(.H'  ..G.....
	0xA0, 0x96, 0xFA, 0x17, 0xA0, 0x96, 0xFA, 0x17,  0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, // ........  ........
	0x54, 0xCB, 0xFA, 0x17, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27, // T.......  ....(.H'
	0xCA, 0xB4, 0x47, 0x00, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..G.....  ........
	0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00,  0x54, 0xCB, 0xFA, 0x17, 0x16, 0x4F, 0xA7, 0x08, // ........  T....O..
	0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27,  0xCA, 0xB4, 0x47, 0x00, 0x01, 0x00, 0x00, 0x00, // B...(.H'  ..G.....
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, // ........  ........
	0x54, 0xCB, 0xFA, 0x17, 0x17, 0x4F, 0xA7, 0x08,  0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27, // T....O..  B...(.H'
	0xCA, 0xB4, 0x47, 0x00, 0x01, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, // ..G.....  ........
	0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00,  0x54, 0xCB, 0xFA, 0x17, 0x18, 0x4F, 0xA7, 0x08, // ........  T....O..
	0x42, 0x00, 0x00, 0x00, 0x28, 0x00, 0x48, 0x27,  0xCA, 0xB4, 0x47, 0x00, 0x58, 0x00, 0x00, 0x00, // B...(.H'  ..G.X...
	0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,  0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, // ........  ........
	0x54, 0xCB, 0xFA, 0x17, 0x18, 0x4F, 0xA7, 0x08,  0x42, 0x00, 0x00, 0x00, 0x34, 0x00, 0x46, 0x27, // T....O..  B...4.F'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  0x0C, 0x00, 0x00, 0x00, 0x2B, 0x00, 0x00, 0x00, // ........  ....+...
	0x19, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,  0x07, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, // ........  ........
	0x79, 0x00, 0x00, 0x00, 0xD0, 0xC0, 0xBA, 0x5C,  0xC6, 0x06, 0x00, 0x00, 0x04, 0x09, 0x2E, 0x00, // y......\  ........
	0x08, 0x00, 0xFD, 0x27, 0x02, 0x00, 0x00, 0x00,  0x0C, 0x00, 0xD4, 0x27, 0x00, 0x00, 0x00, 0x00, // ...'....  ...'....
	0x00, 0x00, 0x00, 0x00, 0x0C, 0x00, 0xD4, 0x27,  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // .......'  ........
	0x04, 0x00, 0x15, 0x27,             //                           ...'
}