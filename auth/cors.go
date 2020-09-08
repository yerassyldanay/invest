package auth

//func CORSMethodMiddleware(r *mux.Router) mux.MiddlewareFunc {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//
//			//if req.Method != http.MethodOptions {
//				next.ServeHTTP(w, req)
//				return
//			//}
//
//			//w.Header().Set("Access-Control-Allow-Origin", "*")
//			//w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
//			//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
//			//w.Header().Add("Content-Type", "application/json")
//
//		})
//	}
//}
