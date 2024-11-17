package main

/*
INTENTO DE MANEJO DE BASES DE DATOS CON GOLANG
*/
import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura de usuario
type User struct {
	UserName string
	Password string
}

// Alias para un slice de usuarios
type Datos []User

// Función para agregar un usuario al slice
func addUser(datos *Datos, user User) {
	// Generar hash de la contraseña
	hashPass := sha256.Sum224([]byte(user.Password))
	hashStringPass := hex.EncodeToString(hashPass[:])

	// Actualizar la contraseña con su versión hasheada
	user.Password = hashStringPass

	// Agregar el usuario al slice
	*datos = append(*datos, user)
}

func encriptPass(userModify *User) {
	hashPass := sha256.Sum224([]byte(userModify.Password))

	hashStringPass := hex.EncodeToString(hashPass[:])

	userModify.Password = hashStringPass

}

// Función para obtener conexión a la base de datos
func getDBConnection() (*sql.DB, error) {
	// Configura el Data Source Name (DSN)
	dsn := "root:@tcp(localhost:3307)/users"

	// Abrir conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión: %w", err)
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		db.Close() // Asegurarse de cerrar la conexión si falla el Ping
		return nil, fmt.Errorf("error al verificar la conexión: %w", err)
	}

	return db, nil
}

func mostrar(passw string) string {
	passDecode, _ := hex.DecodeString(passw)
	return string(passDecode)
}

func main() {
	// Conexión a la base de datos
	db, err := getDBConnection()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear un slice de usuarios
	var datosUser Datos
	addUser(&datosUser, User{"Joaquin", "1234"})
	addUser(&datosUser, User{"RAMON", "2314323142"})
	addUser(&datosUser, User{"JULIO", "23143231421"})
	addUser(&datosUser, User{"RICARDO", "23143231422"})
	addUser(&datosUser, User{"EMANUEL", "23143231423"})

	// Insertar usuarios en la base de datos
	if err := insertUsers(db, datosUser); err != nil {
		log.Fatalf("Error al insertar usuarios: %v", err)
	}

	if err := insert(db, User{UserName: "Yamila", Password: "3324234"}); err != nil {
		log.Fatalf("Error al insertar el usuario: %v", err)
	}

	deleteUser(db, "JULIO")

	err = actName(db, "RAMON", "Ramon")

	log.Fatalf("A ver que paso: %v", err)

	passw, _ := mostrarPassword(db, "Joaquin")

	fmt.Printf("La contraseña del usuario : %v es : %v", "Joaquin", mostrar(passw))

	nombres, _ := condicion(db, "R")
	for _, values := range nombres {
		fmt.Println("El nombre es: ", values)
	}

}

// Funcion que devuelve los nombres que empiezen con una letra.
func condicion(db *sql.DB, letra string) ([]string, error) {
	query := "SELECT usuarios.name from usuarios where usuarios.name like ?"
	parametroAdicional := letra + "%"

	stmt, err := db.Prepare(query)

	var nombres []string
	if err != nil {
		return nombres, fmt.Errorf("Erro al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	filas, err := stmt.Query(parametroAdicional)

	if err != nil {
		return nombres, fmt.Errorf("Error en la consulta: %w", err)
	}
	defer filas.Close()

	for filas.Next() {
		var nombre string
		if err := filas.Scan(&nombre); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		nombres = append(nombres, nombre)
	}

	return nombres, nil
}

func mostrarPassword(db *sql.DB, user string) (string, error) {
	query := "SELECT usuarios.passw from usuarios WHERE usuarios.name = (?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return "", fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	var password string

	err = stmt.QueryRow(user).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("No se encontro el nombre : %v", user)
		}
		return "", fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	return password, nil
}

func deleteUser(db *sql.DB, name string) error {
	query := "DELETE usuarios from usuarios WHERE usuarios.name = (?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("Error al preperar la consulta: %w ", err)
	}
	defer stmt.Close()

	_, errorDelete := stmt.Exec(name)
	if errorDelete != nil {
		return fmt.Errorf("Error al borrar el uusario: %s: %w", name, errorDelete)
	}
	fmt.Println("Eliminado")
	return nil
}

func actName(db *sql.DB, nameAct, nameActualizar string) error {
	query := "UPDATE usuarios set name = (?) where name = (?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("Algo malio sal: %w", err)
	}
	defer stmt.Close()

	_, errUpdate := stmt.Exec(nameActualizar, nameAct)
	if errUpdate != nil {
		return fmt.Errorf("Error al actualizar el usuario %v - %w", nameAct, errUpdate)
	}
	fmt.Println("Usuario actualizado")
	return nil
}

func insert(db *sql.DB, usuario User) error {

	encriptPass(&usuario)

	query := "INSERT INTO usuarios (name, passw) VALUE (?, ?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("Error al preperar la consulta: %w ", err)
	}
	defer stmt.Close()

	_, errorAdd := stmt.Exec(usuario.UserName, usuario.Password)
	if errorAdd != nil {
		fmt.Errorf("Error al agregar el uusario: %s: %w", usuario.UserName, errorAdd)
	}
	fmt.Println("Usuario agregado")
	return nil
}

func insertUsers(db *sql.DB, datos Datos) error {
	// Prepara la consulta SQL
	query := "INSERT INTO usuarios (name, passw) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	// Iterar sobre los usuarios y ejecutar la consulta
	for _, user := range datos {
		_, err := stmt.Exec(user.UserName, user.Password)
		if err != nil {
			return fmt.Errorf("error al insertar usuario %s: %w", user.UserName, err)
		}
		fmt.Printf("Usuario %s agregado a la base de datos.\n", user.UserName)
	}

	return nil
}
