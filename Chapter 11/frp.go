package main

import (
 "fmt"
 "time"

 "github.com/reactivex/rxgo/v2"
)

func main() {
 // Create a channel to simulate incoming temperature data
 temperatureStream := make(chan rxgo.Item)

 // Create an Observable from the channel
 observable := rxgo.FromChannel(temperatureStream)

 // Define the minimum temperature threshold in Celsius
 const minTemperatureCelsius = 10.0

 // Create a filtered and transformed observable
 processedStream := observable.
  Filter(func(item interface{}) bool {
   // Filter out temperatures below the minimum threshold
   temp := item.(float64)
   return temp >= minTemperatureCelsius
  }).
  Map(func(_ context.Context, item interface{}) (interface{}, error) {
   // Convert Celsius to Fahrenheit
   celsius := item.(float64)
   fahrenheit := celsius*9/5 + 32
   return fahrenheit, nil
  }, rxgo.WithBufferedChannel(1))

 // Consume the processed stream
 for item := range processedStream.Observe() {
  if item.Error() {
   fmt.Printf("Received error: %v\n", item.E)
  } else {
   fmt.Printf("Processed temperature: %.2f Fahrenheit\n", item.V)
  }
 }

 // Simulate incoming temperature data
 go func() {
  defer close(temperatureStream)
  for _, temp := range []float64{5, 11, 15, 6, 20, 12} {
   temperatureStream <- rxgo.Of(temp)
   time.Sleep(1 * time.Second)  // Simulate time delay between readings
  }
 }()

 // Keep the main goroutine alive until the observable completes
 time.Sleep(10 * time.Second)
}
