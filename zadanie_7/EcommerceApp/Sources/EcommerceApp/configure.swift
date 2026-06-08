import Vapor
import Fluent
import FluentPostgresDriver
import Leaf

// configures your application
public func configure(_ app: Application) async throws {
    // uncomment to serve files from /Public folder
    // app.middleware.use(FileMiddleware(publicDirectory: app.directory.publicDirectory))

    app.databases.use(.postgres(
        configuration: .init(
            hostname: "localhost",
            port: 5432,
            username: "vapor_username",
            password: "vapor_password",
            database: "vapor_database",
            tls: .disable
        )
    ), as: .psql)

    app.views.use(.leaf)

    app.migrations.add(CreateProduct())
    // register routes
    try routes(app)
}
