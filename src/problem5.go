package main

import (
	"fmt"
	"math"
)

func calParkingFee(h int, m int) int {
	total_fee := 0.0
	total_hour := float64(m)/60.0 + float64(h) 			// convet min to hour

	if total_hour < 0.5 {								// less than 30 min. – Free
		total_fee = 0
	}else if total_hour > 6 {							// after 6th hr. calculate with 200 baht/day
		total_fee += 200 * math.Ceil(total_hour/24)		// calculate with day rate
	}else {
		
		total_hour-=0.5 								// decreat 0.5 hr. from free rate
		total_hour = math.Ceil(total_hour ) 			// round up the fraction of hour
		
		if total_hour-3 >= 0 {							// have overtime in case 10 baht/hour
			total_fee += 10 * 3							// calculate with max duration 3 hr.
			total_hour-=3								// decreat 3 hr. for calculate overtime in next case
		}else{											// don't have overtime in case 
			total_fee += 10 * total_hour				// calculate with remaining hours
		}
		
		if total_hour > 0 {								// have overtime in case 20 baht/hour
			total_fee += 20 * total_hour				// calculate with remaining hours.
		}
	}

	return int(total_fee)
}

func main() {
	h := 4  //hour
	m := 30 //minutes

	fmt.Println(calParkingFee(h, m))
}
