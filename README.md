# Utilities

A collection of small utilities.

## Documentation and examples

Utilities.GetLocation()

Utilities.GetSunMoonInfo(55.663426, 12.542953)

Simple sample code:

```go
import (
    "fmt"
    "github.com/ClausRN/Utilities"
)

locInfo := Utilities.GetLocation()
fmt.Printf("IPAddress: %s\n", locInfo.IP)
fmt.Printf("Org:  %s\n", locInfo.Organisation)
fmt.Printf("Latitude:  %f\n", locInfo.Latitude)
fmt.Printf("Longitude: %f\n", locInfo.Longitude)

sunmoon := Utilities.GetSunMoonInfo(locInfo.Latitude, locInfo.Longitude)
fmt.Printf("Sunset: %s\n", sunmoon.SunSet)

```
