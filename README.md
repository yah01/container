# Container

## Get Started

### Iterator
```golang
slice := cslice.NewSlice(5, 4, 3, 2, 1)


for iter := slice.Iter(); iter.Valid(); iter.Next() {
    fmt.Print(iter.Value()," ")
}

// output:
// 5 4 3 2 1
```

Some utils are very useful, use Map() to double all elements:
```golang
slice := cslice.NewSlice(5, 4, 3, 2, 1)

iter := Map(slice.Iter(), func(elem int) int {
		return elem * 2
	})

for ; iter.Valid(); iter.Next() {
    fmt.Print(iter.Value()," ")
}

// output:
// 10 8 6 4 2
```

Use Reduce() to get bit-OR of all elements:
```golang
slice := cslice.NewSlice(5, 4, 3, 2, 1)

result := Reduce(slice.Iter(), func(result, elem int) int {
		return result|elem
	})

fmt.Print(result)

// output:
// 7
```

Use Enumerate() to get an iterator with index:
```golang
slice := cslice.NewSlice(5, 4, 3, 2, 1)

for iter := Enumerate(slice.Iter()); iter.Valid(); iter.Next() {
    v := iter.Value()
    fmt.Println(v.Idx, v.Value)
}

// output:
// 0 5
// 1 4
// 2 3
// 3 2
// 4 1
```