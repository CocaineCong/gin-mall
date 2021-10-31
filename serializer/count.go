package serializer

type Count struct {
	FavoriteTotal int `json:"favorite_total"`
	NotPayTotal int `json:"not_pay_total"`
	PayTotal int `json:"pay_total"`
}

func BuildCount(favoriteTotal, notPayTotal, payTotal int) Count {
	return Count{
		FavoriteTotal: favoriteTotal,
		NotPayTotal:   notPayTotal,
		PayTotal:      payTotal,
	}
}
