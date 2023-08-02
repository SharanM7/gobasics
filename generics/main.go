package main

type Value interface {
	string | int
}

type Request[V Value] struct {
	Id   int
	Data Data[V]
}

type Data[V Value] struct {
	Type string
	Val  V
}

func main() {
	// req := Request[int]{
	// 	Id:    10,
	// 	Type:  "pred",
	// 	Value: 10,
	// }

	// fmt.Printf("got val :%+v\n", req)

	// req := Request[int]{
	// Id: 1,
	// Data: []Data[string]{
	// 	// {
	// 	// 	Type: "pred",
	// 	// 	Val:  20,
	// 	// },
	// 	// {
	// 	// 	Type: "unpred",
	// 	// 	Val:  20,
	// 	// },
	// },
	// Data: []Data[Value]{
	// 	Data[int]{},
	// },
	// }

	// val := getVal(req, "unpred")
	// fmt.Println("got val", val)
}

// func getVal[v V](req Request, key string) V {
// 	for _, k := range req.Data {
// 		if k.Type == key {
// 			return k.Val
// 		}
// 	}
// 	return ""
// }
