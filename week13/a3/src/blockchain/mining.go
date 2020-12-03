package blockchain

import (
	"work_queue"
)

// Mine in a very simple way: check sequentically until a valid hash is found.
// This doesn't *need* to be used in any way, but could be used to do some mining
// before your .Mine is complete. Results should be the same as .Mine (but slower).
func (blk *Block) mineSequential() {
	proof := uint64(0)
	for !blk.validHashProof(proof) {
		proof += 1
	}
	blk.SetProof(proof)
}

type miningWorker struct {
	// TODO
	startValue uint64
	endValue   uint64
	block      Block
}

func (task miningWorker) Run() interface{} {
	// TODO
	result := new(MiningResult)
	result.Found = false

	// end (inclusive)
	for i := task.startValue; i <= task.endValue; i++ {
		task.block.SetProof(i)
		if task.block.ValidHash() {
			result.Proof = i
			result.Found = true
			break
		}
	}
	return *result
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // was a valid proof-of-work found?
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	wq := work_queue.Create(uint(workers), uint(chunks))
	chunkSize := (start - end) / chunks

	for i := start; i <= end; i += chunkSize {
		work := new(miningWorker)
		work.startValue = start
		work.endValue = end
		work.block = blk
		wq.Enqueue(*work)
	}

	result := new(MiningResult)
	for v := range wq.Results {
		if v.(MiningResult).Found == true {
			wq.Shutdown()
			*result = v.(MiningResult)
			break
		}
	}
	return *result
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4567)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
