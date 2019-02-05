
package main

import "os/exec"
import "fmt"

type Notificador struct{
	ExitoUrl string;
}

func (noti Notificador) NotificarExito(horas []string){
	exec.Command("notify-send",fmt.Sprintf("horas disponibles!!! %v",horas)).Run()
	exec.Command("firefox",noti.ExitoUrl).Run();
}
func (noti Notificador) NotificarError(err string){
	exec.Command("notify-send",err).Run()
}