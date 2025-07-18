package seeders

import (
    "backend_golang/config"
    "backend_golang/models"
    "time"
)

func SeedConcept() {
    now := time.Now()

    concepts := []models.Concept{
        {
            Name:        "API",
            Description: "Interfaz que permite la comunicación entre aplicaciones.",
            URL:         "https://es.wikipedia.org/wiki/Interfaz_de_programaci%C3%B3n_de_aplicaciones",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Framework",
            Description: "Estructura reutilizable para desarrollo de software.",
            URL:         "https://es.wikipedia.org/wiki/Framework",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "MVC",
            Description: "Patrón de diseño Modelo-Vista-Controlador.",
            URL:         "https://es.wikipedia.org/wiki/Modelo%E2%80%93vista%E2%80%93controlador",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "ORM",
            Description: "Mapeo de objetos a bases de datos relacionales.",
            URL:         "https://es.wikipedia.org/wiki/Mapeo_objeto-relacional",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "REST",
            Description: "Arquitectura para servicios web basada en recursos.",
            URL:         "https://es.wikipedia.org/wiki/Representational_State_Transfer",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "SOAP",
            Description: "Protocolo para intercambio estructurado de información.",
            URL:         "https://es.wikipedia.org/wiki/SOAP",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "JSON",
            Description: "Formato de datos ligero y legible.",
            URL:         "https://es.wikipedia.org/wiki/JSON",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "XML",
            Description: "Lenguaje de marcado extensible para datos estructurados.",
            URL:         "https://es.wikipedia.org/wiki/Extensible_Markup_Language",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Composer",
            Description: "Gestor de dependencias para PHP.",
            URL:         "https://getcomposer.org",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "NPM",
            Description: "Gestor de paquetes para JavaScript.",
            URL:         "https://www.npmjs.com",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Git",
            Description: "Sistema de control de versiones distribuido.",
            URL:         "https://git-scm.com",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "GitHub",
            Description: "Plataforma para alojar código usando Git.",
            URL:         "https://github.com",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "CI/CD",
            Description: "Integración y entrega continua de software.",
            URL:         "https://en.wikipedia.org/wiki/CI/CD",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Unit Test",
            Description: "Pruebas automatizadas a unidades de código.",
            URL:         "https://en.wikipedia.org/wiki/Unit_testing",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Docker",
            Description: "Contenerización de aplicaciones.",
            URL:         "https://www.docker.com",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Kubernetes",
            Description: "Orquestación de contenedores.",
            URL:         "https://kubernetes.io",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "PHP",
            Description: "Lenguaje de programación del lado del servidor.",
            URL:         "https://www.php.net",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Laravel",
            Description: "Framework PHP para desarrollo web moderno.",
            URL:         "https://laravel.com",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Blade",
            Description: "Motor de plantillas de Laravel.",
            URL:         "https://laravel.com/docs/blade",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Eloquent",
            Description: "ORM incluido con Laravel.",
            URL:         "https://laravel.com/docs/eloquent",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Seeder",
            Description: "Clase que permite poblar bases de datos con datos.",
            URL:         "https://laravel.com/docs/seeding",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Migration",
            Description: "Definición de estructura de base de datos en código.",
            URL:         "https://laravel.com/docs/migrations",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Middleware",
            Description: "Filtro que se ejecuta entre la petición y la respuesta.",
            URL:         "https://laravel.com/docs/middleware",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Controller",
            Description: "Clase que maneja la lógica de las peticiones HTTP.",
            URL:         "https://laravel.com/docs/controllers",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
        {
            Name:        "Route",
            Description: "Definición de endpoints de la aplicación.",
            URL:         "",
            CreatedAt:   now,
            UpdatedAt:   now,
        },
    }

    for _, concept := range concepts {
        config.DB.Where(models.Concept{Name: concept.Name}).FirstOrCreate(&concept)
    }
}

