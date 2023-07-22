package old

import (
	"errors"
	"log"
	"time"
)

var (
	ErrBatchAvaliableQuantityNotEnoght = errors.New("batch's avaliable quantity not enoght")
	ErrOrderLineDoesNotExits           = errors.New("order line does not exits")
)

type Quantity int
type Sku string
type Reference string

type OrderLine struct {
	Reference Reference
	Sku       Sku
	Quantity  Quantity
}

func NewOrderLine(rf Reference, sku Sku, quantity Quantity) OrderLine {
	return OrderLine{
		Reference: rf,
		Sku:       sku,
		Quantity:  quantity,
	}
}

type Batch struct {
	Reference         Reference
	Sku               Sku
	purchasedQuantity Quantity
	allocations       []OrderLine
	Eta               time.Time
}

func NewBatch(rf Reference, sku Sku, qty Quantity, eta time.Time) Batch {
	return Batch{
		Reference:         rf,
		Sku:               sku,
		purchasedQuantity: qty,
		Eta:               eta,
	}
}

func (batch Batch) CanAllocate(line OrderLine) bool {
	for _, allocatedLine := range batch.allocations {
		if allocatedLine.Reference == line.Reference {
			return false
		}
	}
	if batch.Sku == line.Sku && batch.purchasedQuantity >= line.Quantity {
		return true
	}
	return false
}

func (b *Batch) Allocate(line OrderLine) (bool, error) {
	if !b.CanAllocate(line) {
		return false, ErrBatchAvaliableQuantityNotEnoght
	}
	b.allocations = append(b.allocations, line)
	return true, nil
}

func (b *Batch) Deallocate(line OrderLine) (bool, error) {
	for index, item := range b.allocations {
		if item.Reference == line.Reference {
			b.allocations = append(b.allocations[:index], b.allocations[index+1:]...)
			return true, nil
		}
	}
	return false, ErrOrderLineDoesNotExits
}

func (b Batch) AllocatedQuantity() Quantity {
	var allocatedQuantity Quantity

	for _, line := range b.allocations {
		allocatedQuantity += line.Quantity
	}
	return allocatedQuantity
}

func (b Batch) AvaliableQuantity() Quantity {
	return b.purchasedQuantity - b.AllocatedQuantity()
}

func main() {
	smallest_batch := NewBatch("batch-001", "SMALL-TABLE", 200, time.Now())
	smallest_line := NewOrderLine("order-202", "SMALL-TABLE", 2)

	fatest_batch := NewBatch("batch-002", "COFEE-TABLE", 20, time.Now())
	fatest_line := NewOrderLine("order-203", "COFEE-TABLE", 2)

	success, err := smallest_batch.Allocate(smallest_line)

	if err != nil {
		log.Println(err)
	}
	if !success {
		log.Println("Allocation unsuccessful")
	}

	success, err = fatest_batch.Allocate(fatest_line)

	if err != nil {
		log.Println(err)
	}
	if !success {
		log.Println("Allocation unsuccessful")
	}

	log.Println(smallest_batch.AvaliableQuantity())
	log.Println(fatest_batch.AvaliableQuantity())

}

// could not find or load main class org.gradle.wrapper
