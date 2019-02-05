
package main

import "net/http"
import "encoding/json"
import "strings"
import "errors"
import "time"

type DisponibilidadService struct{

	url string;
	cookie string;

}

type BloqueHora struct{
	Disponible bool;
	Horario string;
}

type target struct{
	Title string;
	Start string;
	End string;
}
func (tt DisponibilidadService)getBloques(itime string,etime string,header string)([]BloqueHora,error){

	resultado:=make([]BloqueHora,0);
	


	t := time.Now()
	//stamp:="1548877878380";
	stamp:=t.Format("20060102150405");
	sdisponible:="Disponible"
	result:=make([]target,0);
		
	url:="https://reservahora.extranjeria.gob.cl/reservarHora/cargaEventosDeReserva.action";
	url=url+"?start="+itime+"&end="+etime+"&_="+stamp;
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
	}
	req.Header.Set("Cookie",header);
	req.Header.Set("Connection","keep-alive");
	response, err :=http.DefaultClient.Do(req);
	if(err!=nil){
		return resultado,err;
	}else{
		decoder:=json.NewDecoder(response.Body);
		decoder.Decode(&result);
		if(len(result)==0){
			return resultado,errors.New("Respuesta Vacia, renovar cookie!");
		}
		
		for _, element := range result {

			aux:=BloqueHora{strings.Contains(strings.ToLower(element.Title),strings.ToLower(sdisponible)),element.Start}
			resultado=append(resultado,aux)
		}
	}

	return resultado,nil;
}