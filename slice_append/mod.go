package sliceappend

type User struct {
	ID string
}

func AppendWithNoAlloc(u []User) []string {
	var res []string
	for _, v := range u {
		res = append(res, v.ID)
	}
	return res
}

func AppendWithPrealloc(u []User) []string {
	res := make([]string, 0, len(u))
	for _, v := range u {
		res = append(res, v.ID)
	}
	return res
}

func CopyWithPreallocAndIndex(u []User) []string {
	res := make([]string, len(u))
	for i, v := range u {
		res[i] = v.ID
	}
	return res
}
