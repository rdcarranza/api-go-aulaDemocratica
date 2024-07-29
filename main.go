package main

import (
	"api-go-aulaDemocratica/src/controladores/env"
	"api-go-aulaDemocratica/src/nucleo/servicios"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {

	//Probando variables de entorno
	env_ := "./.env"
	env_copia := "./src/controladores/env/env.copia"
	if env.VerificarEnv(env_, env_copia) {
		estado_env, err := env.GetEnv("estado_env", env_)
		if err == nil && estado_env == "1" {
			fmt.Println("Archivo env: " + env_ + " - cargado correctamente.")
		} else {
			if err == nil {
				log.Fatal("Verificar configuración de archivo env")
			} else {
				log.Fatal("Verificar configuración de archivo env: " + err.Error())
			}

		}
	}

	u_mdb, _ := env.GetEnv("user_mongodb", env_)
	pw_mdb, _ := env.GetEnv("pw_mongodb", env_)
	h_mdb, _ := env.GetEnv("host_mongodb", env_)
	p_mdb, _ := env.GetEnv("port_mongodb", env_)

	mongo_url := "mongodb://" + u_mdb + ":" + pw_mdb + "@" + h_mdb + ":" + p_mdb + "/"
	fmt.Println("Cadena de conexión mongodb: " + mongo_url)

	/*
		jornada := dominio.JornadaElectoral{
			ID:    uuid.New(),
			Fecha: time.Now(),
		}
		fmt.Print("Jornada creada! - ID: " + jornada.ID.(uuid.UUID).String() + " fecha: " + jornada.Fecha.String())
	*/

	//jornada := servicios.JEServicio(dominio.JornadaElectoral().New())

	//jes := servicios.JEServicio{}
	//jes1 := servicios.servicio_JornadaElectoral()
	//je:=servicios.Servicio_JornadaElectoral()

	jes := servicios.NuevoServicio_JornadaElectora()
	jornada, err := jes.AbrirJE("escuela", 11111111)
	if err == nil {
		//fmt.Println("Jornada creada! - ID: " + jornada.ID.(uuid.UUID).String() + " fecha: " + jornada.Fecha.String() + " - ubicación: " + jornada.Ubicacion + " - DNI del responsable: " + strconv.Itoa(jornada.Dni_resposable))
		fmt.Println("Jornada creada! - ID: " + jornada.GetID().String() + " fecha: " + jornada.GetFecha_apertura().String() + " - ubicación: " + jornada.GetUbicacion() + " - DNI del responsable: " + strconv.Itoa(jornada.GetDni_responsable()))
	}

	jes = servicios.NuevoServicio_JornadaElectora()
	jornada = jes.GetJornada()
	fmt.Println("Jornada creada! - ID: " + jornada.GetID().String() + " fecha: " + jornada.GetFecha_apertura().String() + " - ubicación: " + jornada.GetUbicacion() + " - DNI del responsable: " + strconv.Itoa(jornada.GetDni_responsable()))

	jes = servicios.NuevoServicio_JornadaElectora()
	jornada, err = jes.AbrirJE("escuela2", 11111112)
	if err == nil {
		fmt.Println("Jornada creada! - ID: " + jornada.GetID().String() + " fecha: " + jornada.GetFecha_apertura().String() + " - ubicación: " + jornada.GetUbicacion() + " - DNI del responsable: " + strconv.Itoa(jornada.GetDni_responsable()))
	}

	fmt.Println("Fecha en blanco: " + time.Time{}.String())

}
