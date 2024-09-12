package blocks

type Blocks interface {
	BlockId() int
	CategoryId() int
	StartTime() int
	EndTime() int
	Duration() int
	Color() string
	BackgroundImageUrl() string
	BlockMemo() string
}

type blocksImpl struct {
	blockId            int
	categoryId         int
	startTime          int
	endTime            int
	duration           int
	color              string
	backgroundImageUrl string
	blockMemo          string
}

func NewBlock(startTime, endTime, duration int, color, backgroundImageUrl, blockMemo string) Blocks {
	i := new(blocksImpl)
	i.startTime = startTime
	i.endTime = endTime
	i.duration = duration
	i.color = color
	i.backgroundImageUrl = backgroundImageUrl
	i.blockMemo = blockMemo
	return i
}

func (b blocksImpl) BlockId() int {
	return b.blockId
}

func (b blocksImpl) CategoryId() int {
	return b.categoryId
}

func (b blocksImpl) StartTime() int {
	return b.startTime
}

func (b blocksImpl) EndTime() int {
	return b.endTime
}

func (b blocksImpl) Duration() int {
	return b.duration
}

func (b blocksImpl) Color() string {
	return b.color
}

func (b blocksImpl) BackgroundImageUrl() string {
	return b.backgroundImageUrl
}

func (b blocksImpl) BlockMemo() string {
	return b.blockMemo
}
