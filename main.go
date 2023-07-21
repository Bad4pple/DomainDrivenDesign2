package main

import (
	"errors"
	"log"
	"time"
)

var (
	ErrBatchAvaliableQuantityNotEnoght = errors.New("batch's avaliable quantity not enoght")
	ErrOrderLineDoesNotExits           = errors.New("order line does not exits")
)

type OrderLine struct {
	Reference string
	Sku       string
	Quantity  int
}

func NewOrderLine(rf, sku string, quantity int) OrderLine {
	return OrderLine{
		Reference: rf,
		Sku:       sku,
		Quantity:  quantity,
	}
}

type Batch struct {
	Reference         string
	Sku               string
	purchasedQuantity int
	allocations       []OrderLine
	Eta               time.Time
}

func NewBatch(rf, sku string, avaliableQuantity int, Eta time.Time) Batch {
	return Batch{
		Reference:         rf,
		Sku:               sku,
		purchasedQuantity: avaliableQuantity,
		Eta:               Eta,
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

func (b Batch) AllocatedQuantity() int {
	var allocatedQuantity int

	for _, line := range b.allocations {
		allocatedQuantity += line.Quantity
	}
	return allocatedQuantity
}

func (b Batch) AvaliableQuantity() int {
	return b.purchasedQuantity - b.AllocatedQuantity()
}

// func (b *Batch) OldAllocate(line OrderLine) (bool, error) {
// 	if b.purchasedQuantity < line.Quantity {
// 		return false, ErrBatchAvaliableQuantityNotEnoght
// 	}

// 	b.purchasedQuantity -= line.Quantity
// 	return true, nil
// }

func main() {
	batch := NewBatch("batch-001", "SMALL-TABLE", 20, time.Now())
	log.Println(batch)
	line := NewOrderLine("order-202", "SMALL-TABLE", 2)
	success, err := batch.Allocate(line)
	if err != nil {
		log.Println(err)
	}
	if !success {
		log.Println("Allocation unsuccessful")
	}
	log.Println(batch)

	line = NewOrderLine("order-101", "SMALL-TABLE", 2)
	success, err = batch.Deallocate(line)
	if err != nil {
		log.Println(err)
	}
	if !success {
		log.Println("Deallocation unsuccessful")
	}

	log.Println(batch.AvaliableQuantity())
}
