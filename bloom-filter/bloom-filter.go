package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const BitSize = 100

type Product struct {
	ID string
	Name string
	Slug string
}

func createProduct (name string) *Product {
	slug := strings.Join(strings.Split(name, " "), "-")
	timeNow := time.Now()
	IDProduct := strconv.Itoa(timeNow.Year()) +
				 strconv.Itoa(int(timeNow.Month())) +
				 strconv.Itoa(timeNow.Day())
	milis := strconv.Itoa(timeNow.Nanosecond())
	
	return &Product{
		ID : IDProduct + milis,
		Name:  name,
		Slug: slug,
	}
}

type Db struct {
	Products []*Product
	Bloom *Bloom
}

func (dB *Db) MaybeExist (p *Product) {
	maybeExist := dB.Bloom.probabilistic(p);
	if maybeExist {
		fmt.Println("You can continue to search in database the item probably exists")
	} else {
		fmt.Println("You can't continue to search because item doesn't not exist")
	}
}

func (dB *Db) AddProduct (p *Product) {
	hashSlug := dB.Bloom.hashSlug(p.Slug) // it can be hash with murmurHash to avoid collision
	hashSlugReverse := dB.Bloom.hashSlugReverse(p.Slug)
	hashId := dB.Bloom.hashId(p.ID)
	hashIdReverse := dB.Bloom.hashIdReverse(p.ID)
	hashName := dB.Bloom.hashMidLetter(p.Name)
	dB.Bloom.onBit(hashSlug, hashSlugReverse, hashId, hashIdReverse, hashName)
	dB.Products = append(dB.Products, p)
}

func createDb () *Db {
	bloom := createBloom();
	return &Db{
		Bloom: bloom,
	}
}

type Bloom struct {
	bit [BitSize]uint8
}

func createBloom () *Bloom {
	return &Bloom{}
}

func (b *Bloom) probabilistic (p *Product) bool {
	hashSlug := b.hashSlug(p.Slug) // it can be hash with murmurHash to avoid collision
	hashSlugReverse := b.hashSlugReverse(p.Slug)
	hashId := b.hashId(p.ID)
	hashIdReverse := b.hashIdReverse(p.ID)
	hashName := b.hashMidLetter(p.Name)

	if b.bit[hashSlug] == 0 ||
	   b.bit[hashSlugReverse] == 0 ||
	   b.bit[hashId] == 0 ||
	   b.bit[hashIdReverse] == 0 ||
	   b.bit[hashName] == 0  {
		return false
	   }

	   return true
}

func (b *Bloom) hashSlug(slug string) int {
	sum := 0
	for _, v := range slug[:3] {
		sum += int(v)
	}
	return sum % BitSize
}

func (b *Bloom) hashSlugReverse(slug string) int {
	sum := 0
	for _, v := range slug[len(slug) - 3:] {
		sum += int(v)
	}
	return sum % BitSize
}

func (b *Bloom) hashId(ID string) int {
	sum := 0
	for _, v := range ID[:3] {
		sum += int(v)
	}
	return sum % BitSize
}

func (b *Bloom) hashIdReverse(ID string) int {
	sum := 0
	for _, v := range ID[len(ID) - 3:] {
		sum += int(v)
	}
	return sum % BitSize
}

func (b *Bloom) hashMidLetter(name string) int {
	sum := int(name[len(name) / 2])
	return sum % BitSize
}

func (b *Bloom) onBit (indexes ...int) {
	for _, v := range indexes {
		b.bit[v] = 1
		fmt.Println("v = ", v,b.bit[v])
	}
}


func main () {
	db1 := createDb();
	p1 := createProduct("Yakult Super");
	p2 := createProduct("Yakult");
	db1.AddProduct(p1)
	db1.MaybeExist(p2)
	db1.MaybeExist(p1)
	// fmt.Println(p1)
	fmt.Println(db1.Bloom)
}