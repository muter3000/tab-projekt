import { Sequelize, Model, DataTypes } from 'sequelize';

class SequelizeConnection {
    constructor() {
        throw new Error('Use Singleton.getInstance()');
    }
    static getInstance() {
        const host = process.env.PSQL_HOST ?? "localhost"
        const port = process.env.PSQL_PORT ?? 30050
        const database = process.env.PSQL_DB ?? "postgres"
        if (!SequelizeConnection.instance) {
            SequelizeConnection.instance = new Sequelize({
                dialect: 'postgres',
                user: "admin",
                password: "admin2137",
                database: database,
                host: host,
                port: port
            })
        }
        return SequelizeConnection.instance;
    }
}

module.exports = SequelizeConnection