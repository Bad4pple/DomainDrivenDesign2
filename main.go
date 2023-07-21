package main

import (
	"errors"
	"log"
	"time"
)

var (
	ErrBatchAvaliableQuantityNotEnoght = errors.New("batch's avaliable quantity not enoght")
)

type Batch struct {
	Reference         string
	Sku               string
	AvaliableQuantity int
	Eta               time.Time
}

func NewBatch(rf, sku string, avaliableQuantity int, Eta time.Time) Batch {
	return Batch{
		Reference:         rf,
		Sku:               sku,
		AvaliableQuantity: avaliableQuantity,
		Eta:               Eta,
	}
}

func (b *Batch) Allocate(line OrderLine) (bool, error) {
	if b.AvaliableQuantity < line.Quantity {
		return false, ErrBatchAvaliableQuantityNotEnoght
	}
	b.AvaliableQuantity = b.AvaliableQuantity - line.Quantity
	return true, nil
}

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

func main() {
	batch := NewBatch("batch-001", "SMALL-TABLE", 20, time.Now())
	log.Println(batch)
	line := NewOrderLine("order-202", "SMALL-TABLE", 21)
	success, err := batch.Allocate(line)
	if err != nil {
		log.Println(err)
	}
	if !success {
		log.Println("Allocation unsuccessful")
	}
	log.Println(batch)
}
