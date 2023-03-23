package main

const (
	BlockAir = iota
	BlockVoid
	BlockGrass
)

type BlockSide uint32

const (
	SideNorth BlockSide = iota
	SideSouth
	SideEast
	SideWest
	SideTop
	SideBottom
)

func BlockGenFace(side BlockSide, x, y, z float32) ([]float32) {
	xP := float32( 0.5 + x);
	xN := float32(-0.5 + x);
	yP := float32( 0.5 + y);
	yN := float32(-0.5 + y);
	zP := float32( 0.5 + z);
	zN := float32(-0.5 + z);

	switch side {
	case SideNorth:
		return []float32 {
			xN, yP, zN, 1.0, 1.0,
			xP, yN, zN, 0.0, 0.0,
			xN, yN, zN, 1.0, 0.0,
			xN, yP, zN, 1.0, 1.0,
			xP, yP, zN, 0.0, 1.0,
			xP, yN, zN, 0.0, 0.0,
		}
	case SideSouth:
		return []float32 {
			xN, yP, zP, 0.0, 1.0,
			xN, yN, zP, 0.0, 0.0,
			xP, yN, zP, 1.0, 0.0,
			xN, yP, zP, 0.0, 1.0,
			xP, yN, zP, 1.0, 0.0,
			xP, yP, zP, 1.0, 1.0,
		}
	case SideEast:
		return []float32 {
			xP, yP, zN, 1.0, 1.0,
			xP, yN, zP, 0.0, 0.0,
			xP, yN, zN, 1.0, 0.0,
			xP, yP, zP, 0.0, 1.0,
			xP, yN, zP, 0.0, 0.0,
			xP, yP, zN, 1.0, 1.0,
		}
	case SideWest:
		return []float32 {
			xN, yP, zN, 0.0, 1.0,
			xN, yN, zN, 0.0, 0.0,
			xN, yN, zP, 1.0, 0.0,
			xN, yP, zP, 1.0, 1.0,
			xN, yP, zN, 0.0, 1.0,
			xN, yN, zP, 1.0, 0.0,
		}
	case SideTop:
		return []float32 {
			xN, yP, zP, 0.0, 0.0,
			xP, yP, zP, 1.0, 0.0,
			xN, yP, zN, 0.0, 1.0,
			xN, yP, zN, 0.0, 1.0,
			xP, yP, zP, 1.0, 0.0,
			xP, yP, zN, 1.0, 1.0,
		}
	case SideBottom:
		return []float32 {
			xN, yN, zP, 0.0, 1.0,
			xN, yN, zN, 0.0, 0.0,
			xP, yN, zP, 1.0, 1.0,
			xN, yN, zN, 0.0, 0.0,
			xP, yN, zN, 1.0, 0.0,
			xP, yN, zP, 1.0, 1.0,
		}
	}

	return nil
}