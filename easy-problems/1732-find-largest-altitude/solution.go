package findlargestalt

func largestAltitude(gain []int) int {
	var trip []int
	// add initial point
	trip = append(trip, 0)
	i := 0
	for _, point := range gain {
		new_altitude := trip[i] + point
		trip = append(trip, new_altitude)
		i++
	}

	max := trip[0]
	// find max
	for _, alt := range trip {
		if alt > max {
			max = alt
		}
	}

	return max
}
