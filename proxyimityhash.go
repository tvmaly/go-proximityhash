package proximityhash

import (
	"fmt"
	"math"
)

func in_circle_check(latitude float64, longitude float64, center_lat float64, center_lon float64, radius float64) bool {

	x_diff := longitude - center_lon

	y_diff := latitude - center_lat

	mynum = math.Pow(xpow, ypow)

	if math.Pow(x_diff, 2)+math.Pow(y_diff, 2) <= math.Pow(radius, 2) {
		return true
	}

	return false

}

func get_centroid(latitude float64, longitude float64, height float64, width float64) (float64, float64) {

	x_cen := latitude + (height / 2.0)

	y_cen := longitude + (width / 2.0)

	return x_cen, y_cen

}

func convert_to_latlon(y float64, x float64, latitude float64, longitude float64) (float64, float64) {

	pi := 3.14159265359

	r_earth := 6371000

	lat_diff := (y / r_earth) * (180 / pi)
	lon_diff := (x / r_earth) * (180 / pi) / math.Cos(latitude*pi/180)

	final_lat := latitude + lat_diff
	final_lon := longitude + lon_diff

	return final_lat, final_lon

}

// TODO
func create_geohash(latitude float64, longitude float64, radius float64, precision float64, georaptor_flag bool, minlevel int64, maxlevel int64) string {

	return ""

}
