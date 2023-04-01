package repositories

import "lavanderia.pro/api/types"

var laundries = []types.Laundry{
	{ID: "1", Name: "Laundry #1", Lat: 0.123, Long: 0.321},
	{ID: "2", Name: "Laundry #2", Lat: 0.123, Long: 0.321},
	{ID: "3", Name: "Laundry #3", Lat: 0.123, Long: 0.321},
}

func FindAllLaundries() []types.Laundry {
	return laundries
}
