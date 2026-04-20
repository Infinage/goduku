package core

type Index struct {
	Row uint8
	Col uint8
}

func Response(data any, err error) map[string]any {
	res := map[string]any{
		"success": err == nil,
		"data":    data,
		"error":   "",
	}
	if err != nil {
		res["error"] = err.Error()
	}
	return res
}
