package helper

func Destructure1[T any](slice []T) T {
	return slice[0]
}

func Destructure2[T any](slice []T) (r1 T, r2 T) {
	return slice[0], slice[1]
}

func Destructure3[T any](slice []T) (r1 T, r2 T, r3 T) {
	return slice[0], slice[1], slice[2]
}

func Destructure4[T any](slice []T) (r1 T, r2 T, r3 T, r4 T) {
	return slice[0], slice[1], slice[2], slice[3]
}

func Destructure5[T any](slice []T) (r1 T, r2 T, r3 T, r4 T, r5 T) {
	return slice[0], slice[1], slice[2], slice[3], slice[4]
}
