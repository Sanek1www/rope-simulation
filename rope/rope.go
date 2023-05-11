package rope

const (
	BlockWidth = 5
)

type block struct {
	X      float64
	Y      float64
	Vector float64
}

type Rope struct {
	blocks []block
}

func NewRope(size int) *Rope {
	blocks := make([]block, size)

	for i := 0; i < len(blocks); i++ {
		blocks[i] = block{
			float64(i) * BlockWidth,
			0,
			0,
		}
	}

	return &Rope{blocks}
}

func (r *Rope) makeWave() {
	for i := 5; i < 30; i++ {
		r.blocks[i].Y = r.blocks[i-1].Y + 5
	}

	for i := 30; i < 55; i++ {
		r.blocks[i].Y = r.blocks[i-1].Y - 5
	}
}

func (r *Rope) updateVectors() {
	for i := 0; i < len(r.blocks); i++ {
		if i == 0 {
			r.blocks[i].Vector += r.blocks[i+1].Y - r.blocks[i].Y
			continue
		}

		if i == len(r.blocks)-1 {
			r.blocks[i].Vector += r.blocks[i-1].Y - r.blocks[i].Y
			continue
		}

		r.blocks[i].Vector += (r.blocks[i-1].Y+r.blocks[i+1].Y)/2 - r.blocks[i].Y
	}
}

func (r *Rope) updateBlocksPos() {
	for i := 0; i < len(r.blocks); i++ {
		r.blocks[i].Y = r.blocks[i].Vector + r.blocks[i].Y
	}
}
