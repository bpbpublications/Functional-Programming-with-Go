type Record struct {
	Name  string
	Value int
   }
   
   func Map(records []Record, fn func(Record) Record) []Record {
	var result []Record
	for _, r := range records {
	 result = append(result, fn(r))
	}
	return result
   }
   
   func Filter(records []Record, predicate func(Record) bool) []Record {
	var result []Record
	for _, r := range records {
	 if predicate(r) {
	  result = append(result, r)
	 }
	}
	return result
   }
   
   func main() {
	records := []Record{
	 {"one", 1},
	 {"two", 2},
	 {"three", 3},
	 {"four", 4},
	}
   
	// Double the value of each record
	doubled := Map(records, func(r Record) Record {
	 r.Value *= 2
	 return r
	})
   
	fmt.Println(doubled)
   
	// Only keep records with Value greater than 4
	filtered := Filter(doubled, func(r Record) bool {
	 return r.Value > 4
	})
   
	fmt.Println(filtered)
   }
   