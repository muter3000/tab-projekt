const PostgresPool = require('pg').Pool

class Pool {
    constructor() {
        throw new Error('Use Singleton.getInstance()');
    }
    static getInstance() {
        if (!Pool.instance) {
            Pool.instance = new PostgresPool({
                user: "admin",
                password: "admin2137",
                database: "postgres",
                host: "localhost",
                port: 30050
            })
        }
        return Pool.instance;
    }
}

module.exports = Pool