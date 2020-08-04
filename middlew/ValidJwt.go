package middlew

import(
	"net/http"
	"github.com/BonifacioJZ/tiwgo/routers"
)

//Validar el jwt
func ValidJwt(next http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		_,_,_,err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil{
			http.Error(w,"Error en el Token!"+err.Error(),http.StatusBadRequest)
		}
		next.ServeHTTP(w,r)
	}
}