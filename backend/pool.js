const PostgresPool = require('pg').Pool

class Pool {
    constructor() {
        throw new Error('Use Singleton.getInstance()');
    }
    static getInstance() {
        const host = process.env.PSQL_HOST ?? localhost
        const port = process.env.PSQL_PORT ?? 30050
        if (!Pool.instance) {
            Pool.instance = new PostgresPool({
                user: "admin",
                password: "admin2137",
                database: "postgres",
                host: host,
                port: port
            })
        }
        return Pool.instance;
    }
}

module.exports = Pool