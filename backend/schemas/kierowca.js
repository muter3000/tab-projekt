import { Sequelize, Model, DataTypes } from 'sequelize';
import sequelizeConnection from '../sequelize.js'
 
const sequelize = sequelizeConnection.getInstance()

const Kierowca = sequelize.define('kierowcy', {
  kierowcy_id: {type: DataTypes.INTEGER, primaryKey: true},
  id: DataTypes.INTEGER,
  pesel: DataTypes.STRING,
  imie: DataTypes.STRING,
  nazwisko: DataTypes.STRING,
  login: DataTypes.STRING,
  haslo: DataTypes.STRING,
},{
  timestamps: false,
  freezeTableName: true,
});

export default Kierowca