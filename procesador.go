
package main

type Procesador struct{
	
}
func (Procesador) Procesar( input []BloqueHora) []string{
	result:=make([]string,0)
	
	for _, element := range input {
		if(element.Disponible){
			result=append(result,element.Horario);
		}
	}
	return result;

}