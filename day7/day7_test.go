package main

import (
	"reflect"
	"testing"
)

func TestCountChildBags(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	got := CountChildBags(input, "shiny gold")
	want := 32
	assertIntEqual(t, got, want)
}

func TestCountContainingBags(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	got := CountContainingBags(input, "shiny gold")
	want := 4
	assertIntEqual(t, got, want)
}

func TestParseBags(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	got := ParseBags(input)
	want := 9
	assertIntEqual(t, len(got), want)
}

func TestParseBag(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.`
	cache := Cache{}
	ParseBag(input, &cache)

	got := cache["light red"]
	want := Bag{color: "light red", children: []*Bag{
		cache["bright white"],
		cache["muted yellow"],
		cache["muted yellow"],
	}}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Wanted %#v, got %#v", want, got)
	}
}

func TestParseChild(t *testing.T) {
	input := `1 bright white bag`
	parent := Bag{color: "Parent"}
	cache := Cache{}
	ParseChild(input, &parent, &cache)

	got := cache["bright white"]
	want := Bag{color: "bright white", parents: []*Bag{&parent}}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Wanted %#v, got %#v", want, got)
	}
}

func TestCache(t *testing.T) {
	t.Run("GetOrAdd should create new bag when cache miss", func(t *testing.T) {
		cache := Cache{}
		got := cache.GetOrAdd("bright pink")

		if got == nil {
			t.Error("Wanted a new bag")
		}
	})

	t.Run("GetOrAdd should return cached bag when cache hit", func(t *testing.T) {
		cache := Cache{}
		want := cache.GetOrAdd("bright pink")
		got := cache.GetOrAdd("bright pink")

		if want != got {
			t.Errorf("Wanted existing bag %p to be added, but got %p", &want, &got)
		}
	})
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
