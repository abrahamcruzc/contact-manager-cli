# Contact Manager (CLI)

Un gestor de contactos basado en línea de comandos desarrollado con **Go**, **SQLite**, **Gorm** y **Cobra**.

## Características
✅ Agregar contactos
✅ Listar contactos
✅ Buscar contactos por nombre o email
✅ Eliminar contactos
✅ Interfaz de línea de comandos con `Cobra`
✅ Persistencia con `SQLite` usando `Gorm`

## Instalación

### Prerrequisitos
- Tener **Go** instalado ([Descargar Go](https://go.dev/dl/))
- Tener `cobra-cli` instalado:
  ```sh
  go install github.com/spf13/cobra-cli@latest
  ```

### Clonar el repositorio
```sh
git clone https://github.com/tuusuario/contact-manager.git
cd contact-manager
```

### Instalar dependencias
```sh
go mod tidy
```

## Uso

### Agregar un contacto
```sh
contact-manager add --name "Juan Pérez" --email "juan@example.com" --phone "1234567890"
```

### Listar contactos
```sh
contact-manager list
```

### Buscar un contacto
```sh
contact-manager search --name "Juan"
```

### Eliminar un contacto
```sh
contact-manager delete --id 1
```

## Estructura del Proyecto
```
contact-manager/
│── main.go              // Punto de entrada
│── cmd/                 // Comandos de Cobra
│   │── root.go          // Comando principal
│   │── add.go           // Comando para agregar contactos
│   │── list.go          // Comando para listar contactos
│   │── delete.go        // Comando para eliminar contactos
│── database/            // Configuración de SQLite y Gorm
│   │── database.go      // Conexión a la base de datos
│── models/              // Definición de estructuras
│   │── contact.go       // Modelo Contact con Gorm
└── go.mod               // Módulo de Go
```

## Futuras Mejoras 🚀
- Exportar contactos a CSV/JSON
- Soporte para múltiples números de teléfono
- Implementar actualización de contactos

