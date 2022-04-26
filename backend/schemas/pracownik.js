import { Sequelize, Model, DataTypes } from 'sequelize';

import sequelizeConnection from '../sequelize'
 
const sequelize = sequelizeConnection.getInstance()

const Pracownik = sequelize.define('Pracownik', {
  id: {type: DataTypes.INTEGER, primaryKey: true},
  PESEL: DataTypes.STRING,
  imie: DataTypes.STRING,
  nazwisko: DataTypes.STRING,
  login: DataTypes.STRING,
  haslo: DataTypes.STRING,
});

export default Pracownik